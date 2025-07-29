package exam

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/services/school/common/exam/data"
	"api/services/school/common/exam/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.Exam) (*model.Exam, error) {
	result := *item
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) CreateType(item *model.ExamType) (*model.ExamType, error) {
	result := *item
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) UpdateByID(id int64, item *model.Exam) (*model.Exam, error) {
	result := &model.Exam{}
	fields := map[string]any{
		"school_id": item.SchoolID,
		"year_id":   item.YearID,
		"type_id":   item.TypeID,

		"status":           item.Status,
		"notation":         item.Notation,
		"percentage":       item.Percentage,
		"description":      item.Description,
		"location_type":    item.LocationType,
		"location_details": item.LocationDetails,
		"requirements":     item.Requirements,
		"allowed_items":    item.AllowedItems,
		"start_date":       item.StartDate,
		"end_date":         item.EndDate,
		"is_retry":         item.IsRetry,
		"retry_count":      item.RetryCount,
	}
	if item.ClassSubjectID > 0 {
		fields["class_subject_id"] = item.ClassSubjectID
		if item.SequenceID > 0 {
			fields["sequence_id"] = item.SequenceID
		}
	} else if item.UnitID > 0 {
		fields["unit_id"] = item.UnitID
	}
	return result, repository.Db.
		Preload(clause.Associations).
		Model(&model.Exam{}).
		Where("id = ?", id).
		Updates(
			fields,
		).
		Find(result).Error
}

func (repository *Repository) UpdateExamTypeByID(id int64, data *model.ExamType) (*model.ExamType, error) {
	result := &model.ExamType{}
	fields := map[string]any{
		"school_id": data.SchoolID,

		"name":        data.Name,
		"description": data.Description,
	}
	return result, repository.Db.
		Preload(clause.Associations).
		Model(&model.ExamType{}).
		Where("id = ?", id).
		Updates(
			fields,
		).
		Find(result).Error
}

func (repository *Repository) DeleteByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.Exam{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteExamTypeByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.ExamType{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultipleByID(list []int64, schoolID int64) (result int64, err error) {
	if len(list) < 1 {
		return
	}
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	var query *gorm.DB = repository.Db.Where(where)
	if schoolID > 0 {
		query = query.Where("school_id = ?", schoolID)
	}
	query = query.Delete(&model.Exam{})

	result = query.RowsAffected
	err = query.Error
	return
}

func (repository *Repository) DeleteMultipleExamTypeByID(list []int64, schoolID int64) (result int64, err error) {
	if len(list) < 1 {
		return
	}
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	var query *gorm.DB = repository.Db.Where(where)
	if schoolID > 0 {
		query = query.Where("school_id = ?", schoolID)
	}
	query = query.Delete(&model.ExamType{})

	result = query.RowsAffected
	err = query.Error
	return
}

func (repository *Repository) GetByID(id int64) (*model.Exam, error) {
	result := &model.Exam{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetExamTypeByID(id int64) (*model.ExamType, error) {
	result := &model.ExamType{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByIDSchoolID(id int64, schoolID int64) (*model.Exam, error) {
	result := &model.Exam{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetExamTypeByIDSchoolID(id int64, schoolID int64) (*model.ExamType, error) {
	result := &model.ExamType{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetUniqueObject(item *model.Exam) (*model.Exam, error) {
	result := &model.Exam{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.Exam{
		SchoolID:       item.SchoolID,
		YearID:         item.YearID,
		ClassSubjectID: item.ClassSubjectID,
		SequenceID:     item.SequenceID,
		UnitID:         item.UnitID,
		TypeID:         item.TypeID,
		IsRetry:        item.IsRetry,
		RetryCount:     item.RetryCount,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjects(item1 *model.Exam, item2 *model.Exam) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.YearID == item2.YearID &&
			item1.ClassSubjectID == item2.ClassSubjectID &&
			item1.SequenceID == item2.SequenceID &&
			item1.UnitID == item2.UnitID &&
			item1.TypeID == item2.TypeID &&
			item1.IsRetry == item2.IsRetry &&
			item1.RetryCount == item2.RetryCount) {
		return true
	}
	return false
}

func (repository *Repository) GetExamTypeUniqueObject(item *model.ExamType) (*model.ExamType, error) {
	result := &model.ExamType{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.ExamType{
		SchoolID: item.SchoolID,
		Name:     item.Name,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameExamTypeUniqueObjects(item1 *model.ExamType, item2 *model.ExamType) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.Name == item2.Name) {
		return true
	}
	return false
}

func (repository *Repository) GetAllExamType(
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllExamTypeRequest,
) (result []model.ExamType, err error) {
	result = make([]model.ExamType, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "exam_types.school_id = ?")
			args = append(args, request.SchoolID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(exam_types.id AS TEXT) = ? OR
			exam_types.name ILIKE ? OR
			exam_types.description ILIKE ? OR
			schools.name ILIKE ? OR
			schools.type ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT exam_types.*
				FROM exam_types
				LEFT JOIN schools ON exam_types.school_id = schools.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}

func (repository *Repository) GetAll(
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllRequest,
) (result []model.Exam, err error) {
	result = make([]model.Exam, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "exams.school_id = ?")
			args = append(args, request.SchoolID)
		}
		if request.YearID > 0 {
			where = helpers.AppendWhereClause(where, "exams.year_id = ?")
			args = append(args, request.YearID)
		}
		if request.ClassSubjectID > 0 {
			where = helpers.AppendWhereClause(where, "exams.class_subject_id = ?")
			args = append(args, request.ClassSubjectID)
		}
		if request.SequenceID > 0 {
			where = helpers.AppendWhereClause(where, "exams.sequence_id = ?")
			args = append(args, request.SequenceID)
		}
		if request.UnitID > 0 {
			where = helpers.AppendWhereClause(where, "exams.unit_id = ?")
			args = append(args, request.UnitID)
		}
		if request.TypeID > 0 {
			where = helpers.AppendWhereClause(where, "exams.type_id = ?")
			args = append(args, request.TypeID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(exams.id AS TEXT) = ? OR
			exams.status ILIKE ? OR
			exams.description ILIKE ? OR
			schools.name ILIKE ? OR
			schools.type ILIKE ? OR
			years.name ILIKE ? OR
			highschool_classes.name ILIKE ? OR
			highschool_classes.description ILIKE ? OR
			highschool_subjects.name ILIKE ? OR
			highschool_subjects.description ILIKE ? OR
			highschool_sequences.name ILIKE ? OR
			highschool_sequences.description ILIKE ? OR
			university_units.name ILIKE ? OR
			university_units.description ILIKE ? OR
			exam_types.name ILIKE ? OR
			exam_types.description ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like, like, like, like, like, like, like, like, like, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Preload("ClassSubject.Class").
		Preload("ClassSubject.Subject").
		Preload("Unit.LevelDomain").
		Preload("Unit.LevelDomain.Level").
		Preload("Unit.LevelDomain.Domain").
		Preload("Unit.LevelDomain.Domain.Department").
		Preload("Unit.Semester").
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT exams.*
				FROM exams
				LEFT JOIN schools ON exams.school_id = schools.id
				LEFT JOIN years ON exams.year_id = years.id
				LEFT JOIN exam_types ON exams.type_id = exam_types.id
				LEFT JOIN highschool_class_subjects ON exams.class_subject_id = highschool_class_subjects.id
				LEFT JOIN highschool_sequences ON exams.sequence_id = highschool_sequences.id
				LEFT JOIN university_units ON exams.unit_id = university_units.id
				LEFT JOIN highschool_classes ON highschool_class_subjects.class_id = highschool_classes.id
				LEFT JOIN highschool_subjects ON highschool_class_subjects.subject_id = highschool_subjects.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}

func (repository *Repository) CountAllUniqueIsNotRetry(item *model.Exam) (result int64, err error) {
	var count int64
	countQuery := `
	SELECT COUNT(*) FROM exams
	WHERE exams.school_id = ?
	AND exams.year_id = ?
	AND exams.class_subject_id = ?
	AND exams.sequence_id = ?
	AND exams.unit_id = ?
	AND (exams.is_retry = ? OR exams.is_retry IS NULL)
	`
	args := []any{}
	args = append(args, item.SchoolID, item.YearID, item.ClassSubjectID, item.SequenceID, item.UnitID, false)
	repository.Db.Raw(countQuery, args...).Count(&count)
	result = count
	return
}

func (repository *Repository) CountAllUniqueIsRetry(item *model.Exam) (result int64, err error) {
	query := repository.Db.
		Model(&model.Exam{}).
		Where(&model.Exam{
			SchoolID:       item.SchoolID,
			YearID:         item.YearID,
			ClassSubjectID: item.ClassSubjectID,
			SequenceID:     item.SequenceID,
			UnitID:         item.UnitID,
			TypeID:         item.TypeID,

			IsRetry: true,
		})
	err = query.Count(&result).Error
	return
}
