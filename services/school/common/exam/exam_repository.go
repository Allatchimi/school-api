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

func (repository *Repository) Create(data *model.Exam) (*model.Exam, error) {
	result := *data
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) CreateType(data *model.ExamType) (*model.ExamType, error) {
	result := *data
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) Update(id int64, data *model.Exam) (*model.Exam, error) {
	tempExam, err := repository.GetByID(id)
	if err != nil || tempExam == nil || tempExam.ID != id {
		return nil, err
	}

	result := &model.Exam{}
	return result, repository.Db.Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"school_id":        data.SchoolID,
			"year_id":          data.YearID,
			"type_id":          data.TypeID,
			"unit_id":          data.UnitID,
			"class_subject_id": data.ClassSubjectID,
			"sequence_id":      data.SequenceID,

			"status":           data.Status,
			"percentage":       data.Percentage,
			"description":      data.Description,
			"location_type":    data.LocationType,
			"location_details": data.LocationDetails,
			"requirements":     data.Requirements,
			"allowed_items":    data.AllowedItems,
			"start_date":       data.StartDate,
			"end_date":         data.EndDate,
		},
	).Find(result).Error
}

func (repository *Repository) UpdateType(id int64, data *model.ExamType) (*model.ExamType, error) {
	tempExam, err := repository.GetTypeByID(id)
	if err != nil || tempExam == nil || tempExam.ID != id {
		return nil, err
	}

	result := &model.ExamType{}
	return result, repository.Db.Model(&model.ExamType{}).Where("id = ?", id).Updates(
		map[string]any{
			"school_id":   data.SchoolID,
			"name":        data.Name,
			"description": data.Description,
		},
	).Find(result).Error
}

func (repository *Repository) DeleteByID(id int64) (int64, error) {
	tempExam, err := repository.GetTypeByID(id)
	if err != nil || tempExam == nil || tempExam.ID != id {
		return -1, err
	}

	result := repository.Db.Where("id = ?", id).Delete(&model.Exam{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteTypeByID(id int64) (int64, error) {
	tempExam, err := repository.GetTypeByID(id)
	if err != nil || tempExam == nil || tempExam.ID != id {
		return -1, err
	}

	result := repository.Db.Where("id = ?", id).Delete(&model.ExamType{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultipleByID(list []int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.Exam{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) GetByID(id int64) (*model.Exam, error) {
	result := &model.Exam{}
	return result, repository.Db.Model(&model.Exam{}).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetTypeByID(id int64) (*model.ExamType, error) {
	result := &model.ExamType{}
	return result, repository.Db.Model(&model.ExamType{}).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetUniqueObject(item *model.Exam) (*model.Exam, error) {
	result := &model.Exam{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.Exam{
		SchoolID:       item.SchoolID,
		YearID:         item.YearID,
		TypeID:         item.TypeID,
		UnitID:         item.UnitID,
		ClassSubjectID: item.ClassSubjectID,
		SequenceID:     item.SequenceID,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjects(item1 *model.Exam, item2 *model.Exam) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.YearID == item2.YearID &&
			item1.TypeID == item2.TypeID &&
			item1.UnitID == item2.UnitID &&
			item1.ClassSubjectID == item2.ClassSubjectID &&
			item1.SequenceID == item2.SequenceID) {
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
		Preload("Unit.Domain").
		Preload("Unit.Domain.Department").
		Preload("Unit.Level").
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
