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

func (repository *Repository) Update(id int64, item *model.Quiz) (*model.Quiz, error) {
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

func (repository *Repository) Delete(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.Quiz{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteQuizQuestion(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.QuizQuestion{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteQuizQuestionOption(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.QuizQuestionOption{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteQuizAttempt(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.QuizAttempt{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteQuizAttemptQuestion(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.QuizAttemptQuestion{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultiple(list []int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.Quiz{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) GetById(id int64) (*model.Quiz, error) {
	result := &model.Quiz{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination) (result []model.Quiz, err error) {
	result = make([]model.Quiz, 0)
	var where string = ""
	if filter != nil && len(filter.Search) >= 1 {
		where = fmt.Sprintf(
			"WHERE CAST(id AS TEXT) = '%s' OR title ILIKE '%s' OR description ILIKE '%s' OR start_date ILIKE '%s' OR end_date ILIKE '%s'",
			filter.Search,
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
		)
	}
	tmpErr := repository.Db.Preload(clause.Associations).Scopes(
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

func (repository *Repository) GetAllQuizAttempt(filter *types.Filter, pagination *types.Pagination) (result []model.QuizAttempt, err error) {
	result = make([]model.QuizAttempt, 0)
	var where string = ""
	if filter != nil && len(filter.Search) >= 1 {
		where = fmt.Sprintf(
			"WHERE CAST(id AS TEXT) = '%s' OR name ILIKE '%s' OR start_date ILIKE '%s' OR end_date ILIKE '%s'",
			filter.Search,
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
		)
	}
	tmpErr := repository.Db.Preload(clause.Associations).Scopes(
		helpers.PaginationScope(
			repository.Db,
			"SELECT * FROM quiz_attempts",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error

	err = tmpErr
	return
}
