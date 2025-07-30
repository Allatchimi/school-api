package role

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/services/user/role/data"
	"api/services/user/role/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.Role) (result *model.Role, err error) {
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.Role{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) UpdateByID(id int64, item *model.Role) (result *model.Role, err error) {
	// Update the item
	fields := map[string]any{
		"feature":     item.Feature,
		"name":        item.Name,
		"description": item.Description,
	}
	err = repository.Db.
		Model(&model.Role{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.Role{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) DeleteByID(id int64) (result int64, err error) {
	tmpResult := repository.Db.Where("id = ?", id).Delete(&model.Role{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) DeleteMultipleByID(list []int64) (result int64, err error) {
	if len(list) < 1 {
		return
	}
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.Role{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) GetByID(id int64) (result *model.Role, err error) {
	result = &model.Role{}
	tmpErr := repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error

	err = tmpErr
	return
}

func (repository *Repository) GetByName(name string) (result *model.Role, err error) {
	result = &model.Role{}
	tmpErr := repository.Db.Preload(clause.Associations).Where("name = ?", name).Limit(1).Find(result).Error

	err = tmpErr
	return
}

func (repository *Repository) GetAll(
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllRequest,
) (result []model.Role, err error) {
	result = make([]model.Role, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if len(request.Feature) > 0 {
			where = helpers.AppendWhereClause(where, "roles.feature = ?")
			args = append(args, request.Feature)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(roles.id AS TEXT) = ? OR
			roles.feature ILIKE ? OR
			roles.name ILIKE ? OR
			roles.description ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT roles.*
				FROM roles`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}
