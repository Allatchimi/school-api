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

func (repository *Repository) Create(item *model.Quiz) (*model.Quiz, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) CreateQuizQuestion(item *model.QuizQuestion) (*model.QuizQuestion, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) CreateQuizQuestionOption(item *model.QuizQuestionOption) (*model.QuizQuestionOption, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) CreateQuizAnswer(item *model.QuizAnswer) (*model.QuizAnswer, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) UpdateByID(id int64, item *model.Quiz) (*model.Quiz, error) {
	result := &model.Quiz{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"school_id":        item.SchoolID,
			"year_id":          item.YearID,
			"class_subject_id": item.ClassSubjectID,
			"unit_id":          item.UnitID,

			"title":       item.Title,
			"description": item.Description,
			"status":      item.Status,
		},
	).Error
}

func (repository *Repository) UpdateQuizQuestionSolutionByID(id int64, item *model.QuizQuestion) (*model.QuizQuestion, error) {
	result := &model.QuizQuestion{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"solution_id": item.SolutionID,
		},
	).Error
}

func (repository *Repository) DeleteByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.Quiz{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteQuizQuestionByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.QuizQuestion{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteQuizQuestionOptionByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.QuizQuestionOption{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultipleByID(list []int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.Quiz{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) DeleteMultipleQuizQuestionByID(list []int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.QuizQuestion{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) DeleteMultipleQuizQuestionOptionByID(list []int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.QuizQuestionOption{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) GetByID(id int64) (*model.Quiz, error) {
	result := &model.Quiz{}
	return result, repository.Db.Preload(clause.Associations).
		Preload("Questions.Options").
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetAllQuizAnswerByStudentIDQuizQuestionIDs(studentID int64, list []int64) (*model.QuizAnswer, error) {
	result := &model.QuizAnswer{}
	where := fmt.Sprintf("quiz_question_id IN (%s)", utils.ListIntToString(list))
	return result, repository.Db.Preload(clause.Associations).
		Where("student_id = ?", studentID).Where(where).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(
	filter *types.Filter,
	pagination *types.Pagination,
	schoolID int64,
	yearID int64,
	classSubjectID int64,
	unitID int64,
) (result []model.Quiz, err error) {
	result = make([]model.Quiz, 0)
	var where string = ""
	if schoolID > 0 {
		where = helpers.AppendWhereClause(where, fmt.Sprintf("quizzes.school_id = %d", schoolID))
	}
	if yearID > 0 {
		where = helpers.AppendWhereClause(where, fmt.Sprintf("quizzes.year_id = %d", yearID))
	}
	if classSubjectID > 0 {
		where = helpers.AppendWhereClause(where, fmt.Sprintf("quizzes.class_subject_id = %d", classSubjectID))
	}
	if unitID > 0 {
		where = helpers.AppendWhereClause(where, fmt.Sprintf("quizzes.unit_id = %d", unitID))
	}
	if filter != nil && len(filter.Search) >= 1 {
		tempWhere := fmt.Sprintf(
			"WHERE CAST(quizzes.id AS TEXT) = '%s' OR quizzes.title ILIKE '%s' OR quizzes.description ILIKE '%s' OR schools.name ILIKE '%s' OR years.name ILIKE '%s' OR highschool_class_subjects.name ILIKE '%s' OR university_units.name ILIKE '%s'",
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
	tmpErr := repository.Db.Preload(clause.Associations).
		Preload("Questions.Options").
		Preload("ClassSubject.Class").
		Preload("ClassSubject.Subject").
		Preload("Unit.Domain").
		Preload("Unit.Level").
		Preload("Unit.Semester").
		Scopes(
			helpers.PaginationScope(
				repository.Db,
				"SELECT quizzes.* "+
					"FROM quizzes "+
					"LEFT JOIN schools ON quizzes.school_id = schools.id "+
					"LEFT JOIN years ON quizzes.year_id = years.id "+
					"LEFT JOIN highschool_class_subjects ON quizzes.class_subject_id = highschool_class_subjects.id "+
					"LEFT JOIN university_units ON quizzes.unit_id = university_units.id ",
				where,
				pagination,
				filter,
			),
		).Find(&result).Error

	err = tmpErr
	return
}

func (repository *Repository) GetAllQuizQuestionOption(
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllQuizQuestionOptionRequest,
) (result []model.QuizQuestionOption, err error) {
	result = make([]model.QuizQuestionOption, 0)
	err = repository.Db.Preload(clause.Associations).
		Where("quiz_question_id = ?", request.QuizQuestionID.ID).Limit(1).Find(result).Error
	return
}

func (repository *Repository) GetAllQuizAnswer(
	filter *types.Filter,
	pagination *types.Pagination,
	quizID int64,
	quizQuestionID int64,
	studentID int64,
) (result []model.QuizAnswer, err error) {
	result = make([]model.QuizAnswer, 0)
	var where string = ""
	if quizID > 0 {
		where = helpers.AppendWhereClause(where, fmt.Sprintf("quiz_questions.quiz_id = %d", quizID))
	}
	if quizQuestionID > 0 {
		where = helpers.AppendWhereClause(where, fmt.Sprintf("quiz_answers.quiz_question_id = %d", quizQuestionID))
	}
	if studentID > 0 {
		where = helpers.AppendWhereClause(where, fmt.Sprintf("quiz_answers.student_id = %d", studentID))
	}
	if filter != nil && len(filter.Search) >= 1 {
		tempWhere := fmt.Sprintf(
			"CAST(quiz_answers.id AS TEXT) = '%s' OR quizzes.title ILIKE '%s' OR quizzes.description ILIKE '%s' OR students.uid ILIKE '%s' OR years.name ILIKE '%s' OR highschool_class_subjects.name ILIKE '%s' OR university_units.name ILIKE '%s'",
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
		Preload("Student.User").
		Preload("Student.User.Info").
		Scopes(
			helpers.PaginationScope(
				repository.Db,
				"SELECT quiz_answers.* "+
					"FROM quiz_answers "+
					"LEFT JOIN students ON quiz_answers.student_id = students.id "+
					"LEFT JOIN quiz_questions ON quiz_answers.quiz_question_id = quiz_questions.id "+
					"LEFT JOIN quiz_question_options ON quiz_answers.quiz_question_option_id = quiz_question_options.id "+
					"LEFT JOIN quizzes ON quiz_questions.quiz_id = quizzes.id "+
					"LEFT JOIN schools ON quizzes.school_id = schools.id "+
					"LEFT JOIN years ON quizzes.year_id = years.id "+
					"LEFT JOIN highschool_class_subjects ON quizzes.class_subject_id = highschool_class_subjects.id "+
					"LEFT JOIN university_units ON quizzes.unit_id = university_units.id ",
				where,
				pagination,
				filter,
			),
		).Find(&result).Error

	err = tmpErr
	return
}
