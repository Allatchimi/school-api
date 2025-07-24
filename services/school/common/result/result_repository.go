package result

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/services/school/common/result/data"
	"api/services/school/common/result/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(data *model.Result) (*model.Result, error) {
	result := *data
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) Update(id int64, data *model.Result) (*model.Result, error) {
	foundItem, err := repository.GetByID(id)
	if err != nil || foundItem == nil || foundItem.ID != id {
		return nil, err
	}

	result := &model.Result{}
	return result, repository.Db.Model(&model.Result{}).Where("id = ?", id).Updates(
		map[string]any{
			"student_id": data.StudentID,
			"exam_id":    data.ExamID,

			"value": data.Value,
		},
	).Find(result).Error
}

func (repository *Repository) DeleteByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.Result{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultipleByID(list []int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.Result{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) GetByID(id int64) (*model.Result, error) {
	result := &model.Result{}
	return result, repository.Db.Model(&model.Result{}).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetUniqueObject(item *model.Result) (*model.Result, error) {
	result := &model.Result{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.Result{
		StudentID: item.StudentID,
		ExamID:    item.ExamID,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjects(item1 *model.Result, item2 *model.Result) bool {
	if item1 != nil && item2 != nil &&
		(item1.StudentID == item2.StudentID &&
			item1.ExamID == item2.ExamID) {
		return true
	}
	return false
}

func (repository *Repository) GetAll(
	filter *types.Filter, pagination *types.Pagination,
	request *data.GetAllRequest,
) (result []model.Result, err error) {
	result = make([]model.Result, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.StudentID > 0 {
			where = helpers.AppendWhereClause(where, "results.student_id = ?")
			args = append(args, request.StudentID)
		}
		if request.ExamID > 0 {
			where = helpers.AppendWhereClause(where, "results.exam_id = ?")
			args = append(args, request.ExamID)
		}
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
			CAST(results.id AS TEXT) = ? OR
			CAST(results.value AS TEXT) = ? OR
			exams.status ILIKE ? OR
			exams.description ILIKE ? OR
			exams.location_type ILIKE ? OR
			students.uid ILIKE ? OR
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
		args = append(args, search, search, like, like, like, like, like, like, like, like, like, like, like, like, like, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Preload("Student.User").
		Preload("Student.User.Info").
		Preload("Exam.School").
		Preload("Exam.Year").
		Preload("Exam.Type").
		Preload("Exam.ClassSubject").
		Preload("Exam.ClassSubject.Class").
		Preload("Exam.ClassSubject.Subject").
		Preload("Exam.Sequence").
		Preload("Exam.Unit").
		Preload("Exam.Unit.Level").
		Preload("Exam.Unit.Domain").
		Preload("Exam.Unit.Semester").
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT results.*
				FROM results
				LEFT JOIN exams ON results.exam_id = exams.id
				LEFT JOIN students ON results.student_id = students.id
				LEFT JOIN schools ON exams.school_id = schools.id
				LEFT JOIN years ON exams.year_id = years.id
				LEFT JOIN highschool_class_subjects ON exams.class_subject_id = highschool_class_subjects.id
				LEFT JOIN highschool_sequences ON exams.sequence_id = highschool_sequences.id
				LEFT JOIN university_units ON exams.unit_id = university_units.id
				LEFT JOIN highschool_classes ON highschool_class_subjects.class_id = highschool_classes.id
				LEFT JOIN highschool_subjects ON highschool_class_subjects.subject_id = highschool_subjects.id
				LEFT JOIN exam_types ON exams.type_id = exam_types.id `,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}
