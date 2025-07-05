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

func (repository *Repository) Create(permission *model.Permission) (*model.Permission, error) {
	result := *permission
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) UpdateByID(
	roleID int64,
	tableName string,
	data *model.Permission,
) (result *model.Permission, err error) {
	result = &model.Permission{}
	tmpErr := repository.Db.Preload(clause.Associations).Model(result).Where("role_id = ?", roleID).Where("table_name = ?", tableName).Updates(
		map[string]any{
			"role_id": data.RoleID,

			"table_name": data.TableName,
			"create":     data.Create,
			"read":       data.Read,
			"update":     data.Update,
			"delete":     data.Delete,
		},
	).Error

	err = tmpErr
	return
}

func (repository *Repository) DeleteByID(id int64) (result int64, err error) {
	tmpResult := repository.Db.Where("id = ?", id).Delete(&model.Permission{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) DeleteMultipleByID(list []int64) (result int64, err error) {
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
	var where string = ""
	if request != nil {
		if request.RoleID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("permissions.role_id = %d", request.RoleID))
		}
		if len(request.TableName) > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("permissions.table_name = %s", request.TableName))
		}
	}
	if filter != nil && len(filter.Search) > 0 {
		tempWhere := fmt.Sprintf(
			"(CAST(permissions.id AS TEXT) = '%s' OR CAST(permissions.role_id AS TEXT) ILIKE '%s' OR permissions.table_name ILIKE '%s')",
			filter.Search,
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
		)
		where = helpers.AppendWhereClause(where, tempWhere)
	}
	tmpErr := repository.Db.
		Preload(clause.Associations).
		Scopes(
			helpers.PaginationScope(
				repository.Db,
				"SELECT permissions.* "+
					"FROM permissions ",
				where,
				pagination,
				filter,
			),
		).Find(&result).Error

	err = tmpErr
	return
}
