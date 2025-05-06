package role

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
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

func (repository *Repository) Update(roleID int64, role *model.Role) (result *model.Role, err error) {
	result = &model.Role{}
	tmpErr := repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", roleID).Updates(
		map[string]any{
			"name":        role.Name,
			"feature":     role.Feature,
			"description": role.Description,
		},
	).Error

	err = tmpErr
	return
}

func (repository *Repository) Delete(roleID int64) (result int64, err error) {
	tmpResult := repository.Db.Where("id = ?", roleID).Delete(&model.Role{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) DeleteMultiple(list []int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.Role{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) GetByID(roleID int64) (result *model.Role, err error) {
	result = &model.Role{}
	tmpErr := repository.Db.Preload(clause.Associations).Where("id = ?", roleID).Limit(1).Find(result).Error

	err = tmpErr
	return
}

func (repository *Repository) GetByName(name string) (result *model.Role, err error) {
	result = &model.Role{}
	tmpErr := repository.Db.Preload(clause.Associations).Where("name = ?", name).Limit(1).Find(result).Error

	err = tmpErr
	return
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination) (result []model.Role, err error) {
	result = make([]model.Role, 0)
	var where string = ""
	if filter != nil && len(filter.Search) >= 1 {
		where = fmt.Sprintf(
			"WHERE CAST(id AS TEXT) = '%s' OR name ILIKE '%s' OR feature ILIKE '%s' OR description ILIKE '%s'",
			filter.Search,
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
		)
	}
	tmpErr := repository.Db.Preload(clause.Associations).Scopes(
		helpers.PaginationScope(
			repository.Db,
			"SELECT * FROM roles",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error

	err = tmpErr
	return
}
