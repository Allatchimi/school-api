package result

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/common/result/data"
	"api/services/school/common/result/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.Result) (result *model.Result, err error) {
	// Create the item
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.Result{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) CreateResultTable(item *model.ResultTable) (result *model.ResultTable, err error) {
	// Create the item
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.ResultTable{}
	err = repository.Db.
		Preload(clause.Associations).
		Preload("School.Info").
		Preload("School.Config").
		Preload("Exam.Type").
		Preload("Exam.ClassSubject").
		Preload("Exam.ClassSubject.Class").
		Preload("Exam.ClassSubject.Class.School").
		Preload("Exam.ClassSubject.Class.Specialty").
		Preload("Exam.ClassSubject.Class.Specialty.Section").
		Preload("Exam.ClassSubject.Subject").
		Preload("Exam.Sequence").
		Preload("Exam.Unit").
		Preload("Exam.Unit.LevelDomain").
		Preload("Exam.Unit.LevelDomain.Level").
		Preload("Exam.Unit.LevelDomain.Domain").
		Preload("Exam.Unit.LevelDomain.Domain.Department").
		Preload("Exam.Unit.LevelDomain.Domain.Department.Faculty").
		Preload("Exam.Unit.Semester").
		First(result, item.ID).Error
	return
}

func (repository *Repository) UpdateByID(id int64, item *model.Result) (result *model.Result, err error) {
	// Update the item
	fields := map[string]any{
		"school_id":  item.SchoolID,
		"student_id": item.StudentID,
		"exam_id":    item.ExamID,

		"score": item.Score,
	}
	err = repository.Db.
		Model(&model.Result{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.Result{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) UpdateResultTableByID(id int64, item *model.ResultTable) (result *model.ResultTable, err error) {
	// Update the item
	fields := map[string]any{
		"school_id": item.SchoolID,
		"exam_id":   item.ExamID,

		"status": item.Status,
	}
	err = repository.Db.
		Model(&model.ResultTable{}).
		Preload("School.Info").
		Preload("School.Config").
		Preload("Exam.Type").
		Preload("Exam.ClassSubject").
		Preload("Exam.ClassSubject.Class").
		Preload("Exam.ClassSubject.Class.School").
		Preload("Exam.ClassSubject.Class.Specialty").
		Preload("Exam.ClassSubject.Class.Specialty.Section").
		Preload("Exam.ClassSubject.Subject").
		Preload("Exam.Sequence").
		Preload("Exam.Unit").
		Preload("Exam.Unit.LevelDomain").
		Preload("Exam.Unit.LevelDomain.Level").
		Preload("Exam.Unit.LevelDomain.Domain").
		Preload("Exam.Unit.LevelDomain.Domain.Department").
		Preload("Exam.Unit.LevelDomain.Domain.Department.Faculty").
		Preload("Exam.Unit.Semester").
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.ResultTable{}
	err = repository.Db.
		Preload(clause.Associations).
		Preload("School.Info").
		Preload("School.Config").
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) DeleteByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.Result{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteResultTableByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.ResultTable{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultipleByID(list []int64, schoolID int64) (result int64, err error) {
	if len(list) < 1 {
		return
	}
	query := repository.Db.Where("id IN ?", list)
	if schoolID > 0 {
		query = query.Where("school_id = ?", schoolID)
	}
	query = query.Delete(&model.Result{})

	result = query.RowsAffected
	err = query.Error
	return
}

func (repository *Repository) DeleteMultipleResultTableByID(list []int64, schoolID int64) (result int64, err error) {
	if len(list) < 1 {
		return
	}
	query := repository.Db.Where("id IN ?", list)
	if schoolID > 0 {
		query = query.Where("school_id = ?", schoolID)
	}
	query = query.Delete(&model.ResultTable{})

	result = query.RowsAffected
	err = query.Error
	return
}

func (repository *Repository) GetByID(id int64) (*model.Result, error) {
	result := &model.Result{}
	return result, repository.Db.
		Preload(clause.Associations).
		Preload("Student.User").
		Preload("Student.User.Info").
		Preload("Exam.School").
		Preload("Exam.Year").
		Preload("Exam.Type").
		Preload("Exam.ClassSubject").
		Preload("Exam.ClassSubject.Class").
		Preload("Exam.ClassSubject.Class.School").
		Preload("Exam.ClassSubject.Class.Specialty").
		Preload("Exam.ClassSubject.Class.Specialty.Section").
		Preload("Exam.ClassSubject.Subject").
		Preload("Exam.Sequence").
		Preload("Exam.Unit").
		Preload("Exam.Unit.LevelDomain").
		Preload("Exam.Unit.LevelDomain.Level").
		Preload("Exam.Unit.LevelDomain.Domain").
		Preload("Exam.Unit.LevelDomain.Domain.Department").
		Preload("Exam.Unit.LevelDomain.Domain.Department.Faculty").
		Preload("Exam.Unit.Semester").
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetResultTableByID(id int64) (*model.ResultTable, error) {
	result := &model.ResultTable{}
	return result, repository.Db.
		Preload(clause.Associations).
		Preload("Exam.School").
		Preload("Exam.Year").
		Preload("Exam.Type").
		Preload("Exam.ClassSubject").
		Preload("Exam.ClassSubject.Class").
		Preload("Exam.ClassSubject.Class.School").
		Preload("Exam.ClassSubject.Class.Specialty").
		Preload("Exam.ClassSubject.Class.Specialty.Section").
		Preload("Exam.ClassSubject.Subject").
		Preload("Exam.Sequence").
		Preload("Exam.Unit").
		Preload("Exam.Unit.LevelDomain").
		Preload("Exam.Unit.LevelDomain.Level").
		Preload("Exam.Unit.LevelDomain.Domain").
		Preload("Exam.Unit.LevelDomain.Domain.Department").
		Preload("Exam.Unit.LevelDomain.Domain.Department.Faculty").
		Preload("Exam.Unit.Semester").
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByIDSchoolID(id int64, schoolID int64) (*model.Result, error) {
	result := &model.Result{}
	return result, repository.Db.
		Preload(clause.Associations).
		Preload("Student.User").
		Preload("Student.User.Info").
		Preload("Exam.School").
		Preload("Exam.Year").
		Preload("Exam.Type").
		Preload("Exam.ClassSubject").
		Preload("Exam.ClassSubject.Class").
		Preload("Exam.ClassSubject.Class.School").
		Preload("Exam.ClassSubject.Class.Specialty").
		Preload("Exam.ClassSubject.Class.Specialty.Section").
		Preload("Exam.ClassSubject.Subject").
		Preload("Exam.Sequence").
		Preload("Exam.Unit").
		Preload("Exam.Unit.LevelDomain").
		Preload("Exam.Unit.LevelDomain.Level").
		Preload("Exam.Unit.LevelDomain.Domain").
		Preload("Exam.Unit.LevelDomain.Domain.Department").
		Preload("Exam.Unit.LevelDomain.Domain.Department.Faculty").
		Preload("Exam.Unit.Semester").
		Where("id = ?", id).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetResultTableByIDSchoolID(id int64, schoolID int64) (*model.ResultTable, error) {
	result := &model.ResultTable{}
	return result, repository.Db.
		Preload(clause.Associations).
		Preload("Exam.School").
		Preload("Exam.Year").
		Preload("Exam.Type").
		Preload("Exam.ClassSubject").
		Preload("Exam.ClassSubject.Class").
		Preload("Exam.ClassSubject.Class.School").
		Preload("Exam.ClassSubject.Class.Specialty").
		Preload("Exam.ClassSubject.Class.Specialty.Section").
		Preload("Exam.ClassSubject.Subject").
		Preload("Exam.Sequence").
		Preload("Exam.Unit").
		Preload("Exam.Unit.LevelDomain").
		Preload("Exam.Unit.LevelDomain.Level").
		Preload("Exam.Unit.LevelDomain.Domain").
		Preload("Exam.Unit.LevelDomain.Domain.Department").
		Preload("Exam.Unit.LevelDomain.Domain.Department.Faculty").
		Preload("Exam.Unit.Semester").
		Where("id = ?", id).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetUniqueObject(item *model.Result) (*model.Result, error) {
	result := &model.Result{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.Result{
		SchoolID:  item.SchoolID,
		StudentID: item.StudentID,
		ExamID:    item.ExamID,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjects(item1 *model.Result, item2 *model.Result) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.StudentID == item2.StudentID &&
			item1.ExamID == item2.ExamID) {
		return true
	}
	return false
}

func (repository *Repository) GetResultTableUniqueObject(item *model.ResultTable) (*model.ResultTable, error) {
	result := &model.ResultTable{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.ResultTable{
		SchoolID: item.SchoolID,
		ExamID:   item.ExamID,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameResultTableUniqueObjects(item1 *model.ResultTable, item2 *model.ResultTable) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
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
		if request.QuarterID > 0 {
			where = helpers.AppendWhereClause(where, "highschool_sequences.quarter_id = ?")
			args = append(args, request.QuarterID)
		}
		if request.UnitID > 0 {
			where = helpers.AppendWhereClause(where, "exams.unit_id = ?")
			args = append(args, request.UnitID)
		}
		if request.SemesterID > 0 {
			where = helpers.AppendWhereClause(where, "university_units.semester_id = ?")
			args = append(args, request.SemesterID)
		}
		if request.ExamID > 0 {
			where = helpers.AppendWhereClause(where, "results.exam_id = ?")
			args = append(args, request.ExamID)
		}
		if request.ExamTypeID > 0 {
			where = helpers.AppendWhereClause(where, "exams.type_id = ?")
			args = append(args, request.ExamTypeID)
		}
		if request.ClassID > 0 {
			where = helpers.AppendWhereClause(where, "highschool_class_subjects.class_id = ?")
			args = append(args, request.ClassID)
		}
		if request.LevelDomainID > 0 {
			where = helpers.AppendWhereClause(where, "university_units.level_domain_id = ?")
			args = append(args, request.LevelDomainID)
		}
		if request.TeacherID > 0 {
			where = helpers.AppendWhereClause(where, "teacher_class_subject_units.teacher_id = ?")
			args = append(args, request.TeacherID)
		}
		if request.StudentID > 0 {
			where = helpers.AppendWhereClause(where, "student_enrolls.student_id = ?")
			args = append(args, request.StudentID)
		}
		if request.ParentID > 0 {
			where = helpers.AppendWhereClause(where, "parent_students.parent_id = ?")
			args = append(args, request.ParentID)
		}
		if len(request.TableStatusList) > 0 {
			placeholders := make([]string, len(request.TableStatusList))
			for i := range request.TableStatusList {
				placeholders[i] = "?"
				args = append(args, request.TableStatusList[i])
			}
			where = helpers.AppendWhereClause(where, fmt.Sprintf("result_tables.status IN (%s)", strings.Join(placeholders, ",")))
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(results.id AS TEXT) = ? OR
			CAST(results.score AS TEXT) = ? OR
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
		Preload("Exam.ClassSubject.Class.School").
		Preload("Exam.ClassSubject.Class.Specialty").
		Preload("Exam.ClassSubject.Class.Specialty.Section").
		Preload("Exam.ClassSubject.Subject").
		Preload("Exam.Sequence").
		Preload("Exam.Unit").
		Preload("Exam.Unit.LevelDomain").
		Preload("Exam.Unit.LevelDomain.Level").
		Preload("Exam.Unit.LevelDomain.Domain").
		Preload("Exam.Unit.LevelDomain.Domain.Department").
		Preload("Exam.Unit.LevelDomain.Domain.Department.Faculty").
		Preload("Exam.Unit.Semester").
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT DISTINCT results.*
				FROM results
				LEFT JOIN schools ON results.school_id = schools.id
				LEFT JOIN exams ON results.exam_id = exams.id
				LEFT JOIN students ON results.student_id = students.id
				LEFT JOIN highschool_class_subjects ON exams.class_subject_id = highschool_class_subjects.id
				LEFT JOIN highschool_sequences ON exams.sequence_id = highschool_sequences.id
				LEFT JOIN university_units ON exams.unit_id = university_units.id
				LEFT JOIN years ON exams.year_id = years.id
				LEFT JOIN exam_types ON exams.type_id = exam_types.id
				LEFT JOIN highschool_classes ON highschool_class_subjects.class_id = highschool_classes.id
				LEFT JOIN highschool_subjects ON highschool_class_subjects.subject_id = highschool_subjects.id
				LEFT JOIN result_tables ON results.school_id = result_tables.school_id AND exams.id = result_tables.exam_id
				
				LEFT JOIN teacher_class_subject_units ON results.school_id = teacher_class_subject_units.school_id
				AND teacher_class_subject_units.year_id = exams.year_id
				AND (
				(teacher_class_subject_units.class_subject_id IS NOT NULL AND exams.class_subject_id = teacher_class_subject_units.class_subject_id)
				OR
				(teacher_class_subject_units.unit_id IS NOT NULL AND exams.unit_id = teacher_class_subject_units.unit_id)
				)
				
				LEFT JOIN student_enrolls ON results.school_id = student_enrolls.school_id
				AND student_enrolls.year_id = exams.year_id
				AND (
				(student_enrolls.class_id IS NOT NULL AND highschool_class_subjects.class_id = student_enrolls.class_id)
				OR
				(student_enrolls.level_domain_id IS NOT NULL AND university_units.level_domain_id = student_enrolls.level_domain_id)
				)
				
				LEFT JOIN parent_students ON student_enrolls.student_id = parent_students.student_id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}

func (repository *Repository) GetAllResultTable(
	filter *types.Filter, pagination *types.Pagination,
	request *data.GetAllResultTableRequest,
) (result []model.ResultTable, err error) {
	result = make([]model.ResultTable, 0)

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
		if request.SemesterID > 0 {
			where = helpers.AppendWhereClause(where, "university_units.semester_id = ?")
			args = append(args, request.SemesterID)
		}
		if request.ExamID > 0 {
			where = helpers.AppendWhereClause(where, "result_tables.exam_id = ?")
			args = append(args, request.ExamID)
		}
		if request.ExamTypeID > 0 {
			where = helpers.AppendWhereClause(where, "exams.type_id = ?")
			args = append(args, request.ExamTypeID)
		}
		if len(request.ExamStatusList) > 0 {
			placeholders := make([]string, len(request.ExamStatusList))
			for i := range request.ExamStatusList {
				placeholders[i] = "?"
				args = append(args, request.ExamStatusList[i])
			}
			where = helpers.AppendWhereClause(where, fmt.Sprintf("result_tables.status IN (%s)", strings.Join(placeholders, ",")))
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(result_tables.id AS TEXT) = ? OR
			result_tables.status = ? OR
			exams.status ILIKE ? OR
			exams.description ILIKE ? OR
			exams.location_type ILIKE ? OR
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
		args = append(args, search, like, like, like, like, like, like, like, like, like, like, like, like, like, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Preload("Exam.School").
		Preload("Exam.Year").
		Preload("Exam.Type").
		Preload("Exam.ClassSubject").
		Preload("Exam.ClassSubject.Class").
		Preload("Exam.ClassSubject.Class.School").
		Preload("Exam.ClassSubject.Class.Specialty").
		Preload("Exam.ClassSubject.Class.Specialty.Section").
		Preload("Exam.ClassSubject.Subject").
		Preload("Exam.Sequence").
		Preload("Exam.Unit").
		Preload("Exam.Unit.LevelDomain").
		Preload("Exam.Unit.LevelDomain.Level").
		Preload("Exam.Unit.LevelDomain.Domain").
		Preload("Exam.Unit.LevelDomain.Domain.Department").
		Preload("Exam.Unit.LevelDomain.Domain.Department.Faculty").
		Preload("Exam.Unit.Semester").
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT DISTINCT result_tables.*
				FROM result_tables
				LEFT JOIN schools ON result_tables.school_id = schools.id
				LEFT JOIN exams ON result_tables.exam_id = exams.id
				LEFT JOIN highschool_class_subjects ON exams.class_subject_id = highschool_class_subjects.id
				LEFT JOIN highschool_sequences ON exams.sequence_id = highschool_sequences.id
				LEFT JOIN university_units ON exams.unit_id = university_units.id
				LEFT JOIN years ON exams.year_id = years.id
				LEFT JOIN exam_types ON exams.type_id = exam_types.id
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
