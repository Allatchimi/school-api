package domain

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/services/school/university/domain/data"
	"api/services/school/university/domain/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.UniversityDomain) (*model.UniversityDomain, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) UpdateByID(id int64, item *model.UniversityDomain) (*model.UniversityDomain, error) {
	result := &model.UniversityDomain{}
	return result, repository.Db.Preload(clause.Associations).Model(&model.UniversityDomain{}).Where("id = ?", id).Updates(
		map[string]any{
			"school_id":     item.SchoolID,
			"department_id": item.DepartmentID,

			"name":        item.Name,
			"description": item.Description,
		},
	).Find(result).Error
}

func (repository *Repository) DeleteByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.UniversityDomain{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultipleByID(list []int64, schoolID int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	var query *gorm.DB = repository.Db.Where(where)
	if schoolID > 0 {
		query = query.Where("school_id = ?", schoolID)
	}
	query = query.Delete(&model.UniversityDomain{})

	result = query.RowsAffected
	err = query.Error
	return
}

func (repository *Repository) GetByID(id int64) (*model.UniversityDomain, error) {
	result := &model.UniversityDomain{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByIDSchoolID(id int64, schoolID int64) (*model.UniversityDomain, error) {
	result := &model.UniversityDomain{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetUniqueObject(item *model.UniversityDomain) (*model.UniversityDomain, error) {
	result := &model.UniversityDomain{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.UniversityDomain{
		SchoolID: item.SchoolID,
		Name:     item.Name,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjects(item1 *model.UniversityDomain, item2 *model.UniversityDomain) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.Name == item2.Name) {
		return true
	}
	return false
}

func (repository *Repository) GetAll(
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllRequest,
) (result []model.UniversityDomain, err error) {
	result = make([]model.UniversityDomain, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "domains.school_id = ?")
			args = append(args, request.SchoolID)
		}
		if request.DepartmentID > 0 {
			where = helpers.AppendWhereClause(where, "domains.department_id = ?")
			args = append(args, request.DepartmentID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(domains.id AS TEXT) = ? OR
			domains.name ILIKE ? OR
			domains.description ILIKE ? OR
			schools.name ILIKE ? OR
			schools.type ILIKE ? OR
			university_departments.name ILIKE ? OR
			university_departments.description ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT domains.*
				FROM university_domains domains
				LEFT JOIN schools ON domains.school_id = schools.id
				LEFT JOIN university_departments ON domains.department_id = university_departments.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}
