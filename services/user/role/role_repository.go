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

func (repository *Repository) Create(role *model.Role) (result *model.Role, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = utils.InterfaceToError(r)
		}
	}()

	result = &model.Role{}
	*result = *role
	tmpErr := repository.Db.Preload(clause.Associations).Create(&result).Error

	err = tmpErr
	return
}

func (repository *Repository) UpdateByID(id int64, role *model.Role) (result *model.Role, err error) {
	result = &model.Role{}
	tmpErr := repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"name":        role.Name,
			"feature":     role.Feature,
			"description": role.Description,
		},
	).Error

	err = tmpErr
	return
}

func (repository *Repository) DeleteByID(id int64) (result int64, err error) {
	tmpResult := repository.Db.Where("id = ?", id).Delete(&model.Role{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) DeleteMultipleByID(list []int64) (result int64, err error) {
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
	inputJwtToken *types.JwtToken,
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
			roles.name ILIKE ? OR 
			infos.feature ILIKE ? OR 
			infos.description ILIKE ?
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
				FROM roles `,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}
