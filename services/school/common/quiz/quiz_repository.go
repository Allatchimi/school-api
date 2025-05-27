package quiz

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
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

func (repository *Repository) CreateQuizAttempt(item *model.QuizAttempt) (*model.QuizAttempt, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) CreateQuizAttemptQuestion(item *model.QuizAttemptQuestion) (*model.QuizAttemptQuestion, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) UpdateByID(id int64, item *model.Quiz) (*model.Quiz, error) {
	result := &model.Quiz{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"title":            item.Title,
			"description":      item.Description,
			"start_date":       item.StartDate,
			"end_date":         item.EndDate,
			"school_id":        item.SchoolID,
			"year_id":          item.YearID,
			"unit_id":          item.UnitID,
			"class_subject_id": item.ClassSubjectID,
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

func (repository *Repository) DeleteQuizAttemptByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.QuizAttempt{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteQuizAttemptQuestionByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.QuizAttemptQuestion{})
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
			"WHERE CAST(quizzes.id AS TEXT) = '%s' OR quizzes.title ILIKE '%s' OR quizzes.description ILIKE '%s' OR quizzes.start_date ILIKE '%s' OR quizzes.end_date ILIKE '%s'",
			filter.Search,
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
		)
		where = helpers.AppendWhereClause(where, tempWhere)
	}
	tmpErr := repository.Db.Preload(clause.Associations).
		Preload("Questions.Options").
		Scopes(
			helpers.PaginationScope(
				repository.Db,
				"SELECT * FROM quizzes",
				where,
				pagination,
				filter,
			),
		).Find(&result).Error

	err = tmpErr
	return
}

func (repository *Repository) GetAllQuizAttempt(
	filter *types.Filter,
	pagination *types.Pagination,
	quizID int64,
	schoolID int64,
	yearID int64,
	classSubjectID int64,
	unitID int64,
) (result []model.QuizAttempt, err error) {
	result = make([]model.QuizAttempt, 0)
	var where string = ""
	if quizID > 0 {
		where = helpers.AppendWhereClause(where, fmt.Sprintf("quiz_attempts.quiz_id = %d", quizID))
	}
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
			"CAST(quiz_attemps.id AS TEXT) = '%s' OR quizzes.title ILIKE '%s' OR quizzes.description ILIKE '%s' OR students.uid ILIKE '%s'",
			filter.Search,
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
		)
		where = helpers.AppendWhereClause(where, tempWhere)
	}
	tmpErr := repository.Db.Preload(clause.Associations).Scopes(
		helpers.PaginationScope(
			repository.Db,
			"SELECT quiz_attemps.id, quiz_attemps.quiz_id, quiz_attemps.student_id"+
				", quiz_attemps.created_at, quiz_attemps.updated_at FROM quiz_attemps "+
				"LEFT JOIN quizzes ON quiz_attemps.quiz_id = quizzes.id "+
				"LEFT JOIN students ON quiz_attemps.student_id = students.id",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error

	err = tmpErr
	return
}
