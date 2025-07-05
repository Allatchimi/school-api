package level

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/services/school/university/level/data"
	"api/services/school/university/level/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.UniversityLevel) (*model.UniversityLevel, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) CreateLevelDomain(item *model.UniversityLevelDomain) (*model.UniversityLevelDomain, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) Update(id int64, item *model.UniversityLevel) (*model.UniversityLevel, error) {
	result := &model.UniversityLevel{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"school_id":   item.SchoolID,
			"name":        item.Name,
			"description": item.Description,
		},
	).Error
}

func (repository *Repository) UpdateLevelDomain(id int64, item *model.UniversityLevelDomain) (*model.UniversityLevelDomain, error) {
	result := &model.UniversityLevelDomain{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"school_id": item.SchoolID,
			"level_id":  item.LevelID,
			"domain_id": item.DomainID,

			"fees":         item.Fees,
			"program":      item.Program,
			"requirements": item.Requirements,
			"is_valid":     item.IsValid,
			"invalid_date": item.InvalidDate,
		},
	).Error
}

func (repository *Repository) Delete(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.UniversityLevel{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteLevelDomain(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.UniversityLevelDomain{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultiple(list []int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.UniversityLevel{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) DeleteMultipleLevelDomain(list []int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.UniversityLevelDomain{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) GetByID(id int64) (*model.UniversityLevel, error) {
	result := &model.UniversityLevel{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetLevelDomainByID(id int64) (*model.UniversityLevelDomain, error) {
	result := &model.UniversityLevelDomain{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetUniqueObject(item *model.UniversityLevel) (*model.UniversityLevel, error) {
	result := &model.UniversityLevel{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.UniversityLevel{
		SchoolID: item.SchoolID,
		Name:     item.Name,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjects(item1 *model.UniversityLevel, item2 *model.UniversityLevel) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.Name == item2.Name) {
		return true
	}
	return false
}

func (repository *Repository) GetLevelDomainUniqueObject(item *model.UniversityLevelDomain) (*model.UniversityLevelDomain, error) {
	result := &model.UniversityLevelDomain{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.UniversityLevelDomain{
		SchoolID: item.SchoolID,
		LevelID:  item.LevelID,
		DomainID: item.DomainID,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameLevelDomainUniqueObjects(item1 *model.UniversityLevelDomain, item2 *model.UniversityLevelDomain) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.LevelID == item2.LevelID &&
			item1.DomainID == item2.DomainID) {
		return true
	}
	return false
}

func (repository *Repository) GetAll(
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllRequest,
) (result []model.UniversityLevel, err error) {
	result = make([]model.UniversityLevel, 0)
	var where string = ""
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("levels.school_id = %d", request.SchoolID))
		}
	}
	if filter != nil && len(filter.Search) > 0 {
		tempWhere := fmt.Sprintf(
			"(CAST(levels.id AS TEXT) = '%s' OR levels.name ILIKE '%s' OR levels.description ILIKE '%s' OR schools.name ILIKE '%s' OR schools.type ILIKE '%s')",
			filter.Search,
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
		)
		where = helpers.AppendWhereClause(where, tempWhere)
	}
	tmpErr := repository.Db.Preload(clause.Associations).Scopes(
		helpers.PaginationScope(
			repository.Db,
			"SELECT levels.id, levels.name, levels.description, levels.school_id"+
				", levels.created_at, levels.updated_at FROM university_levels levels "+
				"LEFT JOIN schools ON levels.school_id = schools.id",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error

	err = tmpErr
	return
}

func (repository *Repository) GetAllLevelDomain(filter *types.Filter, pagination *types.Pagination, request *data.GetAllLevelDomainRequest) (result []model.UniversityLevelDomain, err error) {
	result = make([]model.UniversityLevelDomain, 0)
	var where string = ""
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("ld.school_id = %d", request.SchoolID))
		}
		if request.LevelID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("ld.level_id = %d", request.LevelID))
		}
		if request.DomainID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("ld.domain_id = %d", request.DomainID))
		}
	}
	if filter != nil && len(filter.Search) > 0 {
		tempWhere := fmt.Sprintf(
			"(CAST(ld.id AS TEXT) = '%s' OR ld.name ILIKE '%s' OR ld.description ILIKE '%s' OR schools.name ILIKE '%s' OR schools.type ILIKE '%s')",
			filter.Search,
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
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
				"SELECT ld.* "+
					"FROM university_level_domains ld "+
					"LEFT JOIN university_levels ON ld.level_id = university_levels.id "+
					"LEFT JOIN university_domains ON ld.domain_id = university_domains.id "+
					"LEFT JOIN schools ON ld.school_id = schools.id ",
				where,
				pagination,
				filter,
			),
		).Find(&result).Error

	err = tmpErr
	return
}
