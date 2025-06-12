package schedule

import (
	"fmt"

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
	schedule := *data
	return &schedule, repository.Db.Create(&schedule).Error
}

func (repository *Repository) Update(id int64, data *model.Schedule) (*model.Schedule, error) {
	tempSchedule, err := repository.GetByID(id)
	if err != nil || tempSchedule == nil || tempSchedule.ID != id {
		return nil, err
	}

	schedule := &model.Schedule{}
	return schedule, repository.Db.Model(schedule).Where("id = ?", id).Updates(
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
		},
	).Error
}

func (repository *Repository) Delete(id int64) (int64, error) {
	tempSchedule, err := repository.GetByID(id)
	if err != nil || tempSchedule == nil || tempSchedule.ID != id {
		return -1, err
	}

	schedule := repository.Db.Where("id = ?", id).Delete(&model.Schedule{})
	return schedule.RowsAffected, schedule.Error
}

func (repository *Repository) GetByID(id int64) (*model.Schedule, error) {
	schedule := &model.Schedule{}
	return schedule, repository.Db.Model(&model.Schedule{}).Where("id = ?", id).Limit(1).Find(schedule).Error
}

func (repository *Repository) GetUniqueObject(item *model.Schedule) (*model.Schedule, error) {
	schedule := &model.Schedule{}
	return schedule, repository.Db.Preload(clause.Associations).Where(&model.Schedule{
		SchoolID:       item.SchoolID,
		YearID:         item.YearID,
		ClassSubjectID: item.ClassSubjectID,
		UnitID:         item.UnitID,
	}).Limit(1).Find(schedule).Error
}

func (repository *Repository) AreSameUniqueObjects(item1 *model.Schedule, item2 *model.Schedule) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.YearID == item2.YearID &&
			item1.ClassSubjectID == item2.ClassSubjectID &&
			item1.UnitID == item2.UnitID) {
		return true
	}
	return false
}

func (repository *Repository) GetAll(
	filter *types.Filter, pagination *types.Pagination,
	request *data.GetAllRequest,
) (schedule []model.Schedule, err error) {
	schedule = make([]model.Schedule, 0)
	var where string = ""
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("schedules.school_id = %d", request.SchoolID))
		}
		if request.YearID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("schedules.year_id = %d", request.YearID))
		}
		if request.ClassSubjectID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("schedules.class_subject_id = %d", request.ClassSubjectID))
		}
		if request.UnitID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("schedules.unit_id = %d", request.UnitID))
		}
	}
	if filter != nil && len(filter.Search) >= 1 {
		tempWhere := fmt.Sprintf(
			"(CAST(schedules.id AS TEXT) = '%s' OR schedules.type ILIKE '%s' OR schedules.day_of_the_week ILIKE '%s' OR schedules.repeat_count ILIKE '%s' OR schedules.repeat_type ILIKE '%s' OR schedules.start_time ILIKE '%s' OR schedules.end_time ILIKE '%s')",
			filter.Search,
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
		)
		where = helpers.AppendWhereClause(where, tempWhere)
	}
	tmpErr := repository.Db.
		Preload(clause.Associations).
		Preload("ClassSubject.Class").
		Preload("ClassSubject.Subject").
		Preload("Unit.Domain").
		Preload("Unit.Level").
		Preload("Unit.Semester").
		Scopes(
			helpers.PaginationScope(
				repository.Db,
				"SELECT schedules.* "+
					"FROM schedules "+
					"LEFT JOIN schools ON schedules.school_id = schools.id "+
					"LEFT JOIN years ON schedules.year_id = years.id "+
					"LEFT JOIN highschool_class_subjects ON schedules.class_subject_id = highschool_class_subjects.id "+
					"LEFT JOIN university_units ON schedules.unit_id = university_units.id ",
				where,
				pagination,
				filter,
			),
		).Find(&schedule).Error

	err = tmpErr
	return
}
