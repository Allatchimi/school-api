package quiz

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/services/school/common/quiz/data"
	"api/services/school/common/quiz/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.Quiz) (result *model.Quiz, err error) {
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.Quiz{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) CreateQuizQuestion(item *model.QuizQuestion) (result *model.QuizQuestion, err error) {
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.QuizQuestion{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) CreateQuizQuestionOption(item *model.QuizQuestionOption) (result *model.QuizQuestionOption, err error) {
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.QuizQuestionOption{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) CreateQuizAnswer(item *model.QuizAnswer) (result *model.QuizAnswer, err error) {
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.QuizAnswer{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) UpdateByID(id int64, item *model.Quiz) (result *model.Quiz, err error) {
	// Update the item
	fields := map[string]any{
		"school_id":        item.SchoolID,
		"year_id":          item.YearID,
		"class_subject_id": nil,
		"unit_id":          nil,

		"title":       item.Title,
		"description": item.Description,
		"status":      item.Status,
	}
	if item.ClassSubjectID > 0 {
		fields["class_subject_id"] = item.ClassSubjectID
	}
	if item.UnitID > 0 {
		fields["unit_id"] = item.UnitID
	}
	err = repository.Db.
		Model(&model.Quiz{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.Quiz{}
	err = repository.Db.
		Preload(clause.Associations).
		Preload("ClassSubject.Class").
		Preload("ClassSubject.Class.School").
		Preload("ClassSubject.Class.Specialty").
		Preload("ClassSubject.Class.Specialty.Section").
		Preload("ClassSubject.Subject").
		Preload("Unit.LevelDomain").
		Preload("Unit.LevelDomain.School").
		Preload("Unit.LevelDomain.Level").
		Preload("Unit.LevelDomain.Domain").
		Preload("Unit.LevelDomain.Domain.Department").
		Preload("Unit.LevelDomain.Domain.Department.Faculty").
		Preload("Unit.Semester").
		Preload("Questions.Options").
		Preload("Questions.Solution").
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) UpdateQuizQuestionSolutionByID(id int64, item *model.QuizQuestion) (result *model.QuizQuestion, err error) {
	// Update the item
	fields := map[string]any{
		"solution_id": item.SolutionID,
	}
	err = repository.Db.
		Model(&model.QuizQuestion{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.QuizQuestion{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) DeleteByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.Quiz{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteQuizQuestionByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.QuizQuestion{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteQuizQuestionByQuizID(quizID int64) (int64, error) {
	result := repository.Db.
		Where("quiz_id = ?", quizID).
		Delete(&model.QuizQuestion{})
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
	query = query.Delete(&model.Quiz{})

	result = query.RowsAffected
	err = query.Error
	return
}

func (repository *Repository) GetByID(id int64) (*model.Quiz, error) {
	result := &model.Quiz{}
	return result, repository.Db.
		Preload(clause.Associations).
		Preload("ClassSubject.Class").
		Preload("ClassSubject.Class.School").
		Preload("ClassSubject.Class.Specialty").
		Preload("ClassSubject.Class.Specialty.Section").
		Preload("ClassSubject.Subject").
		Preload("Unit.LevelDomain").
		Preload("Unit.LevelDomain.School").
		Preload("Unit.LevelDomain.Level").
		Preload("Unit.LevelDomain.Domain").
		Preload("Unit.LevelDomain.Domain.Department").
		Preload("Unit.LevelDomain.Domain.Department.Faculty").
		Preload("Unit.Semester").
		Preload("Questions.Options").
		Preload("Questions.Solution").
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByIDSchoolID(id int64, schoolID int64) (*model.Quiz, error) {
	result := &model.Quiz{}
	return result, repository.Db.
		Preload(clause.Associations).
		Preload("ClassSubject.Class").
		Preload("ClassSubject.Class.School").
		Preload("ClassSubject.Class.Specialty").
		Preload("ClassSubject.Class.Specialty.Section").
		Preload("ClassSubject.Subject").
		Preload("Unit.LevelDomain").
		Preload("Unit.LevelDomain.School").
		Preload("Unit.LevelDomain.Level").
		Preload("Unit.LevelDomain.Domain").
		Preload("Unit.LevelDomain.Domain.Department").
		Preload("Unit.LevelDomain.Domain.Department.Faculty").
		Preload("Unit.Semester").
		Preload("Questions.Options").
		Preload("Questions.Solution").
		Where("id = ?", id).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetAllQuizAnswerByStudentIDQuizQuestionIDs(studentID int64, list []int64) (*model.QuizAnswer, error) {
	result := &model.QuizAnswer{}
	where := fmt.Sprintf("quiz_question_id IN (%s)", utils.ListIntToString(list))
	return result, repository.Db.
		Preload(clause.Associations).
		Where("student_id = ?", studentID).
		Where(where).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllRequest,
) (result []model.Quiz, err error) {
	result = make([]model.Quiz, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "quizzes.school_id = ?")
			args = append(args, request.SchoolID)
		}
		if request.YearID > 0 {
			where = helpers.AppendWhereClause(where, "quizzes.year_id = ?")
			args = append(args, request.YearID)
		}
		if request.ClassSubjectID > 0 {
			where = helpers.AppendWhereClause(where, "quizzes.class_subject_id = ?")
			args = append(args, request.ClassSubjectID)
		}
		if request.UnitID > 0 {
			where = helpers.AppendWhereClause(where, "quizzes.unit_id = ?")
			args = append(args, request.UnitID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(quizzes.id AS TEXT) = ? OR
			quizzes.title ILIKE ? OR
			quizzes.description ILIKE ? OR
			quizzes.status ILIKE ? OR
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
		args = append(args, search, like, like, like, like, like, like, like, like, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Preload("ClassSubject.Class").
		Preload("ClassSubject.Class.School").
		Preload("ClassSubject.Class.Specialty").
		Preload("ClassSubject.Class.Specialty.Section").
		Preload("ClassSubject.Subject").
		Preload("Unit.LevelDomain").
		Preload("Unit.LevelDomain.School").
		Preload("Unit.LevelDomain.Level").
		Preload("Unit.LevelDomain.Domain").
		Preload("Unit.LevelDomain.Domain.Department").
		Preload("Unit.LevelDomain.Domain.Department.Faculty").
		Preload("Unit.Semester").
		Preload("Questions.Options").
		Preload("Questions.Solution").
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT DISTINCT quizzes.*
				FROM quizzes
				LEFT JOIN schools ON quizzes.school_id = schools.id
				LEFT JOIN years ON quizzes.year_id = years.id
				LEFT JOIN highschool_class_subjects ON quizzes.class_subject_id = highschool_class_subjects.id
				LEFT JOIN university_units ON quizzes.unit_id = university_units.id
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

func (repository *Repository) GetAllQuizAnswer(
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllQuizAnswerRequest,
) (result []model.QuizAnswer, err error) {
	result = make([]model.QuizAnswer, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "students.school_id = ?")
			args = append(args, request.SchoolID)
		}
		if request.StudentID > 0 {
			where = helpers.AppendWhereClause(where, "quiz_answers.student_id = ?")
			args = append(args, request.StudentID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(quiz_answers.id AS TEXT) = ? OR
			quiz_questions.title ILIKE ? OR
			quiz_questions.description ILIKE ? OR
			students.uid ILIKE ? OR
			schools.name ILIKE ? OR
			schools.type ILIKE ? OR
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Preload("Student.User").
		Preload("Student.User.Info").
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT DISTINCT quiz_answers.*
				FROM quiz_answers
				LEFT JOIN quiz_questions ON quiz_answers.quiz_question_id = quiz_questions.id
				LEFT JOIN students ON quiz_answers.student_id = students.id
				LEFT JOIN schools ON students.school_id = schools.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}
