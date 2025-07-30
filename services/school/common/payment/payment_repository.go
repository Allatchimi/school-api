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

func (repository *Repository) Create(item *model.Payment) (result *model.Payment, err error) {
	// Create the item
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.Payment{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) UpdateByID(id int64, item *model.Payment) (result *model.Payment, err error) {
	// Update the item
	fields := map[string]any{
		"school_id":         item.SchoolID,
		"student_enroll_id": item.StudentEnrollID,

		"amount":   item.Amount,
		"currency": item.Currency,
		"date":     item.Date,
		"method":   item.Method,
		"status":   item.Status,
		"message":  item.Message,
	}
	err = repository.Db.
		Model(&model.Payment{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.Payment{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) DeleteByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.Payment{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultipleByID(list []int64, schoolID int64) (result int64, err error) {
	if len(list) < 1 {
		return
	}
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
			CAST(payments.amount AS TEXT) = ? OR
			payments.method ILIKE ? OR
			payments.status ILIKE ? OR
			schools.name ILIKE ? OR
			schools.type ILIKE ? OR
			students.uid ILIKE ? OR
			highschool_classes.name ? OR
			highschool_classes.description ? OR
			university_levels.name ? OR
			university_levels.description ? OR
			highschool_domains.name ? OR
			highschool_domains.description ? OR
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, search, like, like, like, like, like, like, like, like, like, like, like)
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
				LEFT JOIN students ON student_enrolls.student_id = students.id
				LEFT JOIN highschool_classes ON student_enrolls.class_id = highschool_classes.id
				LEFT JOIN university_level_domains ON student_enrolls.level_domain_id = university_level_domains.id
				LEFT JOIN university_levels ON university_level_domains.level_id = university_levels.id
				LEFT JOIN university_domains ON university_level_domains.domain_id = university_domains.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}
