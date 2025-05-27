package school

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/services/school/common/school/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.School) (*model.School, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) CreateInfo(item *model.SchoolInfo) (*model.SchoolInfo, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) CreateConfig(item *model.SchoolConfig) (*model.SchoolConfig, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) UpdateByID(id int64, item *model.School) (*model.School, error) {
	result := &model.School{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"name": item.Name,
			"type": item.Type,
			"logo": item.Logo,
		},
	).Error
}

func (repository *Repository) UpdateConfigInfoByID(id int64, configID int64, infoID int64) (*model.School, error) {
	result := &model.School{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"config_id": configID,
			"info_id":   infoID,
		},
	).Error
}

func (repository *Repository) UpdateInfoByID(id int64, item *model.SchoolInfo) (*model.SchoolInfo, error) {
	result := &model.SchoolInfo{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"full_name":   item.FullName,
			"description": item.Description,
			"slogan":      item.Slogan,

			"phone_number1": item.PhoneNumber1,
			"phone_number2": item.PhoneNumber2,
			"phone_number3": item.PhoneNumber3,

			"email1": item.Email1,
			"email2": item.Email2,
			"email3": item.Email3,

			"founder":    item.Founder,
			"founded_at": item.FoundedAt,

			"address":            item.Address,
			"location_longitude": item.LocationLongitude,
			"location_latitude":  item.LocationLatitude,

			"image1": item.Image1,
			"image2": item.Image2,
			"image3": item.Image3,
			"image4": item.Image4,
		},
	).Error
}

func (repository *Repository) UpdateConfigByID(id int64, item *model.SchoolConfig) (*model.SchoolConfig, error) {
	result := &model.SchoolConfig{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"email_domain":  item.EmailDomain,
			"color_primary": item.ColorPrimary,
		},
	).Error
}

func (repository *Repository) DeleteByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.School{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultipleByID(list []int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.School{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) GetByID(id int64) (*model.School, error) {
	result := &model.School{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetUniqueObject(item *model.School) (*model.School, error) {
	result := &model.School{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.School{
		Name: item.Name,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjects(item1 *model.School, item2 *model.School) bool {
	if item1 != nil && item2 != nil &&
		(item1.Name == item2.Name) {
		return true
	}
	return false
}

func (repository *Repository) GetInfoByID(id int64) (*model.SchoolInfo, error) {
	result := &model.SchoolInfo{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetConfigByID(id int64) (*model.SchoolConfig, error) {
	result := &model.SchoolConfig{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, typeName string) (result []model.School, err error) {
	result = make([]model.School, 0)
	var where string = ""
	if len(typeName) > 0 {
		where = fmt.Sprintf("WHERE schools.type = '%s'", typeName)
	}
	if filter != nil && len(filter.Search) >= 1 {
		tempWhere := fmt.Sprintf(
			"CAST(schools.id AS TEXT) = '%s' OR schools.name ILIKE '%s' OR CAST(schools.type AS TEXT) ILIKE '%s' OR infos.full_name ILIKE '%s' OR infos.slogan ILIKE '%s' OR infos.founder ILIKE '%s'",
			filter.Search,
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
		)
		if strings.HasPrefix(where, "WHERE") {
			where = fmt.Sprintf("%s AND (%s)", where, tempWhere)
		} else {
			where = fmt.Sprintf("WHERE %s", tempWhere)
		}
	}
	newFilter := filter
	newFilter.OrderBy = "schools." + newFilter.OrderBy
	tmpErr := repository.Db.Preload(clause.Associations).Scopes(
		helpers.PaginationScope(
			repository.Db,
			"SELECT schools.id, schools.name, schools.type, schools.logo, schools.currency, schools.school_config_id, schools.school_info_id"+
				", schools.created_at, schools.updated_at FROM schools "+
				"LEFT JOIN school_infos as infos ON schools.school_info_id = infos.id",
			where,
			pagination,
			newFilter,
		),
	).Find(&result).Error

	err = tmpErr
	return
}
