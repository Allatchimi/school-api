package payment

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/services/school/common/payment/data"
	"api/services/school/common/payment/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.Payment) (*model.Payment, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) Update(id int64, item *model.Payment) (*model.Payment, error) {
	result := &model.Payment{}
	return result, repository.Db.Preload(clause.Associations).Model(&model.Payment{}).Where("id = ?", id).Updates(
		map[string]any{
			"school_id":         item.SchoolID,
			"student_enroll_id": item.StudentEnrollID,

			"amount":   item.Amount,
			"currency": item.Currency,
			"date":     item.Date,
			"method":   item.Method,
			"status":   item.Status,
			"message":  item.Message,
		},
	).Find(result).Error
}

func (repository *Repository) DeleteByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.Payment{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultipleByID(list []int64, schoolID int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	var query *gorm.DB = repository.Db.Where(where)
	if schoolID > 0 {
		query = query.Where("school_id = ?", schoolID)
	}
	query = query.Delete(&model.Payment{})

	result = query.RowsAffected
	err = query.Error
	return
}

func (repository *Repository) GetByID(id int64) (*model.Payment, error) {
	result := &model.Payment{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByIDSchoolID(id int64, schoolID int64) (*model.Payment, error) {
	result := &model.Payment{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllRequest,
) (result []model.Payment, err error) {
	result = make([]model.Payment, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "payments.school_id = ?")
			args = append(args, request.SchoolID)
		}
		if request.StudentEnrollID > 0 {
			where = helpers.AppendWhereClause(where, "payments.student_enroll_id = ?")
			args = append(args, request.StudentEnrollID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(payments.id AS TEXT) = ? OR
			payments.method ILIKE ? OR
			payments.status ILIKE ? OR
			schools.name ILIKE ? OR
			schools.type ILIKE ? OR
			students.uid ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Preload("School").
		Preload("StudentEnroll.Student").
		Preload("StudentEnroll.Student.User").
		Preload("StudentEnroll.Student.User.Info").
		Preload("StudentEnroll.Year").
		Preload("StudentEnroll.Class").
		Preload("StudentEnroll.LevelDomain").
		Preload("StudentEnroll.LevelDomain.Level").
		Preload("StudentEnroll.LevelDomain.Domain").
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT payments.*
				FROM payments
				LEFT JOIN schools ON payments.school_id = schools.id
				LEFT JOIN student_enrolls ON payments.student_enroll_id = student_enrolls.id
				LEFT JOIN students ON student_enrolls.student_id = students.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}
