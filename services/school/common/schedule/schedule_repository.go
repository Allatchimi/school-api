package schedule

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/common/schedule/data"
	"api/services/school/common/schedule/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(data *model.Schedule) (*model.Schedule, error) {
	result := *data
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) CreateScheduleCommon(data *model.ScheduleCommon) (*model.ScheduleCommon, error) {
	result := *data
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) Update(id int64, data *model.Schedule) (*model.Schedule, error) {
	tempSchedule, err := repository.GetByID(id)
	if err != nil || tempSchedule == nil || tempSchedule.ID != id {
		return nil, err
	}

	result := &model.Schedule{}
	return result, repository.Db.Model(&model.Schedule{}).Where("id = ?", id).Updates(
		map[string]any{
			"school_id":        data.SchoolID,
			"year_id":          data.YearID,
			"class_subject_id": data.ClassSubjectID,
			"unit_id":          data.UnitID,

			"type":            data.Type,
			"day_of_the_week": data.DayOfTheWeek,
			"repeat_count":    data.RepeatCount,
			"repeat_type":     data.RepeatType,
			"start_time":      data.StartTime,
			"end_time":        data.EndTime,
			"is_valid":        data.IsValid,
			"invalid_date":    data.InvalidDate,
		},
	).Find(result).Error
}

func (repository *Repository) UpdateScheduleCommon(id int64, data *model.ScheduleCommon) (*model.ScheduleCommon, error) {
	tempSchedule, err := repository.GetScheduleCommonByID(id)
	if err != nil || tempSchedule == nil || tempSchedule.ID != id {
		return nil, err
	}

	result := &model.ScheduleCommon{}
	return result, repository.Db.Model(&model.ScheduleCommon{}).Where("id = ?", id).Updates(
		map[string]any{
			"school_id": data.SchoolID,
			"year_id":   data.YearID,

			"type":            data.Type,
			"day_of_the_week": data.DayOfTheWeek,
			"repeat_count":    data.RepeatCount,
			"repeat_type":     data.RepeatType,
			"start_time":      data.StartTime,
			"end_time":        data.EndTime,
			"is_valid":        data.IsValid,
			"invalid_date":    data.InvalidDate,
		},
	).Find(result).Error
}

func (repository *Repository) DeleteByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.Schedule{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteScheduleCommonByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.ScheduleCommon{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetByID(id int64) (*model.Schedule, error) {
	result := &model.Schedule{}
	return result, repository.Db.Model(&model.Schedule{}).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetScheduleCommonByID(id int64) (*model.ScheduleCommon, error) {
	result := &model.ScheduleCommon{}
	return result, repository.Db.Model(&model.ScheduleCommon{}).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetUniqueObject(item *model.Schedule) (*model.Schedule, error) {
	result := &model.Schedule{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.Schedule{
		SchoolID:       item.SchoolID,
		YearID:         item.YearID,
		ClassSubjectID: item.ClassSubjectID,
		UnitID:         item.UnitID,
		DayOfTheWeek:   item.DayOfTheWeek,
		RepeatCount:    item.RepeatCount,
		RepeatType:     item.RepeatType,
		StartTime:      item.StartTime,
		EndTime:        item.EndTime,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjects(item1 *model.Schedule, item2 *model.Schedule) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.YearID == item2.YearID &&
			item1.ClassSubjectID == item2.ClassSubjectID &&
			item1.UnitID == item2.UnitID &&
			item1.DayOfTheWeek == item2.DayOfTheWeek &&
			item1.RepeatCount == item2.RepeatCount &&
			item1.RepeatType == item2.RepeatType &&
			item1.StartTime == item2.StartTime &&
			item1.EndTime == item2.EndTime) {
		return true
	}
	return false
}

func (repository *Repository) GetScheduleCommonUniqueObject(item *model.ScheduleCommon) (*model.ScheduleCommon, error) {
	result := &model.ScheduleCommon{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.ScheduleCommon{
		SchoolID:     item.SchoolID,
		YearID:       item.YearID,
		DayOfTheWeek: item.DayOfTheWeek,
		RepeatCount:  item.RepeatCount,
		RepeatType:   item.RepeatType,
		StartTime:    item.StartTime,
		EndTime:      item.EndTime,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreScheduleCommonSameUniqueObjects(item1 *model.ScheduleCommon, item2 *model.ScheduleCommon) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.YearID == item2.YearID &&
			item1.DayOfTheWeek == item2.DayOfTheWeek &&
			item1.RepeatCount == item2.RepeatCount &&
			item1.RepeatType == item2.RepeatType &&
			item1.StartTime == item2.StartTime &&
			item1.EndTime == item2.EndTime) {
		return true
	}
	return false
}

func (repository *Repository) GetAll(
	filter *types.Filter, pagination *types.Pagination,
	request *data.GetAllRequest,
) (result []model.Schedule, err error) {
	result = make([]model.Schedule, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "schedules.school_id = ?")
			args = append(args, request.SchoolID)
		}
		if request.YearID > 0 {
			where = helpers.AppendWhereClause(where, "schedules.year_id = ?")
			args = append(args, request.YearID)
		}
		if request.ClassSubjectID > 0 {
			where = helpers.AppendWhereClause(where, "schedules.class_subject_id = ?")
			args = append(args, request.ClassSubjectID)
		}
		if request.UnitID > 0 {
			where = helpers.AppendWhereClause(where, "schedules.unit_id = ?")
			args = append(args, request.UnitID)
		}
		if len(request.Type) > 0 {
			where = helpers.AppendWhereClause(where, "schedules.type = ?")
			args = append(args, request.Type)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(schedules.id AS TEXT) = ? OR
			schedules.type ILIKE ? OR
			schedules.day_of_the_week ILIKE ? OR
			schedules.repeat_type ILIKE ? OR
			schedules.start_time ILIKE ? OR
			schedules.end_time ILIKE ? OR
			schools.name ILIKE ? OR
			schools.type ILIKE ? OR
			years.name ILIKE ? OR
			highschool_classes.name ILIKE ? OR
			highschool_classes.description ILIKE ? OR
			highschool_subjects.name ILIKE ? OR
			highschool_subjects.description ILIKE ? OR
			university_units.name ILIKE ? OR
			university_units.description ILIKE ? 
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like, like, like, like, like, like, like, like, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Preload("ClassSubject.Class").
		Preload("ClassSubject.Subject").
		Preload("Unit.LevelDomain").
		Preload("Unit.LevelDomain.Domain").
		Preload("Unit.LevelDomain.Level").
		Preload("Unit.Semester").
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT schedules.*
				FROM schedules
				LEFT JOIN schools ON schedules.school_id = schools.id
				LEFT JOIN years ON schedules.year_id = years.id
				LEFT JOIN highschool_class_subjects ON schedules.class_subject_id = highschool_class_subjects.id
				LEFT JOIN university_units ON schedules.unit_id = university_units.id
				LEFT JOIN highschool_classes ON highschool_class_subjects.class_id = highschool_classes.id
				LEFT JOIN highschool_subjects ON highschool_class_subjects.subject_id = highschool_subjects.id `,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}

func (repository *Repository) GetAllScheduleCommon(
	filter *types.Filter, pagination *types.Pagination,
	request *data.GetAllRequest,
) (result []model.ScheduleCommon, err error) {
	result = make([]model.ScheduleCommon, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "schedules.school_id = ?")
			args = append(args, request.SchoolID)
		}
		if request.YearID > 0 {
			where = helpers.AppendWhereClause(where, "schedules.year_id = ?")
			args = append(args, request.YearID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(schedules.id AS TEXT) = ? OR
			schedules.type ILIKE ? OR
			schedules.day_of_the_week ILIKE ? OR
			schedules.repeat_type ILIKE ? OR
			schedules.start_time ILIKE ? OR
			schedules.end_time ILIKE ? OR
			schools.name ILIKE ? OR
			schools.type ILIKE ? OR
			years.name ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like, like, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT schedules.*
				FROM schedule_commons schedules
				LEFT JOIN schools ON schedules.school_id = schools.id
				LEFT JOIN years ON schedules.year_id = years.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}
