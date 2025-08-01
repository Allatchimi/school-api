package result

import (
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
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.ResultTable{}
	err = repository.Db.
		Preload(clause.Associations).
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
		Preload("Exam.ClassSubject.Subject").
		Preload("Exam.Sequence").
		Preload("Exam.Unit").
		Preload("Exam.Unit.LevelDomain").
		Preload("Exam.Unit.LevelDomain.Level").
		Preload("Exam.Unit.LevelDomain.Domain").
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
		Preload("Exam.ClassSubject.Subject").
		Preload("Exam.Sequence").
		Preload("Exam.Unit").
		Preload("Exam.Unit.LevelDomain").
		Preload("Exam.Unit.LevelDomain.Level").
		Preload("Exam.Unit.LevelDomain.Domain").
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
		Preload("Exam.ClassSubject.Subject").
		Preload("Exam.Sequence").
		Preload("Exam.Unit").
		Preload("Exam.Unit.LevelDomain").
		Preload("Exam.Unit.LevelDomain.Level").
		Preload("Exam.Unit.LevelDomain.Domain").
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
		Preload("Exam.ClassSubject.Subject").
		Preload("Exam.Sequence").
		Preload("Exam.Unit").
		Preload("Exam.Unit.LevelDomain").
		Preload("Exam.Unit.LevelDomain.Level").
		Preload("Exam.Unit.LevelDomain.Domain").
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
		if request.StudentID > 0 {
			where = helpers.AppendWhereClause(where, "results.student_id = ?")
			args = append(args, request.StudentID)
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
		Preload("Exam.ClassSubject.Subject").
		Preload("Exam.Sequence").
		Preload("Exam.Unit").
		Preload("Exam.Unit.LevelDomain").
		Preload("Exam.Unit.LevelDomain.Level").
		Preload("Exam.Unit.LevelDomain.Domain").
		Preload("Exam.Unit.Semester").
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT results.*
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
				LEFT JOIN highschool_subjects ON highschool_class_subjects.subject_id = highschool_subjects.id`,
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
		Preload("Exam.ClassSubject.Subject").
		Preload("Exam.Sequence").
		Preload("Exam.Unit").
		Preload("Exam.Unit.LevelDomain").
		Preload("Exam.Unit.LevelDomain.Level").
		Preload("Exam.Unit.LevelDomain.Domain").
		Preload("Exam.Unit.Semester").
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT result_tables.*
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
