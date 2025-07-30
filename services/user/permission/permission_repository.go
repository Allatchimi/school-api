package permission

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/services/user/permission/data"
	"api/services/user/permission/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.Permission) (result *model.Permission, err error) {
	// Create the item
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.Permission{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) UpdateByID(
	roleID int64,
	tableName string,
	item *model.Permission,
) (result *model.Permission, err error) {
	// Update the item
	fields := map[string]any{
		"role_id": item.RoleID,

		"table_name": item.TableName,
		"create":     item.Create,
		"read":       item.Read,
		"update":     item.Update,
		"delete":     item.Delete,
	}
	err = repository.Db.
		Model(&model.Permission{}).
		Where("role_id = ?", roleID).
		Where("table_name = ?", tableName).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.Permission{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", roleID).
		First(result).Error
	return
}

func (repository *Repository) DeleteByID(id int64) (result int64, err error) {
	tmpResult := repository.Db.Where("id = ?", id).Delete(&model.Permission{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) DeleteMultipleByID(list []int64) (result int64, err error) {
	if len(list) < 1 {
		return
	}
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.Permission{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) GetByRoleIDTableName(
	roleID int64,
	tableName string,
) (*model.Permission, error) {
	result := &model.Permission{}
	return result, repository.Db.Preload(clause.Associations).Where("role_id = ?", roleID).Where("table_name = ?", tableName).Limit(1).Find(result).Error
}

func (repository *Repository) GetByRoleIDTableNameMultiple(
	roleID int64,
	tableName1 string,
	tableName2 string,
) (*model.Permission, error) {
	result := &model.Permission{}
	return result, repository.Db.Preload(clause.Associations).Where("role_id = ?", roleID).Where("table_name = ? or table_name = ?", tableName1, tableName2).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllRequest,
) (result []model.Permission, err error) {
	result = make([]model.Permission, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.RoleID > 0 {
			where = helpers.AppendWhereClause(where, "permissions.role_id = ?")
			args = append(args, request.RoleID)
		}
		if len(request.TableName) > 0 {
			where = helpers.AppendWhereClause(where, "permissions.table_name = ?")
			args = append(args, request.TableName)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(permissions.id AS TEXT) = ? OR
			permissions.table_name ILIKE ? OR
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
				`SELECT permissions.*
				FROM permissions
				LEFT JOIN roles ON permissions.role_id = roles.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}
