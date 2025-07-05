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
	return result, repository.Db.Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"student_id": data.StudentID,
			"exam_id":    data.ExamID,

			"value": data.Value,
		},
	).Error
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
	var where string = ""
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("exams.school_id = %d", request.SchoolID))
		}
		if request.YearID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("exams.year_id = %d", request.YearID))
		}
		if request.TypeID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("exams.type_id = %d", request.TypeID))
		}
		if request.UnitID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("exams.unit_id = %d", request.UnitID))
		}
		if request.ClassSubjectID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("exams.class_subject_id = %d", request.ClassSubjectID))
		}
		if request.SequenceID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("exams.sequence_id = %d", request.SequenceID))
		}
		if request.StudentID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("results.student_id = %d", request.StudentID))
		}
		if request.ExamID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("results.exam_id = %d", request.ExamID))
		}
	}
	if filter != nil && len(filter.Search) > 0 {
		tempWhere := fmt.Sprintf(
			"(CAST(results.id AS TEXT) = '%s' OR exams.description ILIKE '%s')",
			filter.Search,
			"%"+filter.Search+"%",
		)
		where = helpers.AppendWhereClause(where, tempWhere)
	}
	tmpErr := repository.Db.
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
			helpers.PaginationScope(
				repository.Db,
				"SELECT results.* "+
					"FROM results "+
					"LEFT JOIN students ON results.student_id = students.id "+
					"LEFT JOIN exams ON results.exam_id = exams.id ",
				where,
				pagination,
				filter,
			),
		).Find(&result).Error

	err = tmpErr
	return
}
