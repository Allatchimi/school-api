package level

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/university/level/data"
	"api/services/school/university/level/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.UniversityLevel) (result *model.UniversityLevel, err error) {
	// Create the item
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.UniversityLevel{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) CreateLevelDomain(item *model.UniversityLevelDomain) (result *model.UniversityLevelDomain, err error) {
	// Create the item
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.UniversityLevelDomain{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) UpdateByID(id int64, item *model.UniversityLevel) (result *model.UniversityLevel, err error) {
	// Update the item
	fields := map[string]any{
		"school_id": item.SchoolID,

		"name":        item.Name,
		"description": item.Description,
	}
	err = repository.Db.
		Model(&model.UniversityLevel{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.UniversityLevel{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) UpdateLevelDomainByID(id int64, item *model.UniversityLevelDomain) (result *model.UniversityLevelDomain, err error) {
	// Update the item
	fields := map[string]any{
		"school_id": item.SchoolID,
		"level_id":  item.LevelID,
		"domain_id": item.DomainID,

		"fees":         item.Fees,
		"program":      item.Program,
		"requirements": item.Requirements,
		"is_valid":     item.IsValid,
		"invalid_date": item.InvalidDate,
	}
	err = repository.Db.
		Model(&model.UniversityLevelDomain{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.UniversityLevelDomain{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) DeleteByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.UniversityLevel{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteLevelDomainByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.UniversityLevelDomain{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultipleByID(list []int64, schoolID int64) (result int64, err error) {
	if len(list) < 1 {
		return
	}
	query := repository.Db.Where("id IN ?", list)
	if schoolID > 0 {
		query = query.Where("school_id = ?", schoolID)
	}
	query = query.Delete(&model.UniversityLevel{})

	result = query.RowsAffected
	err = query.Error
	return
}

func (repository *Repository) DeleteMultipleLevelDomainByID(list []int64, schoolID int64) (result int64, err error) {
	if len(list) < 1 {
		return
	}
	query := repository.Db.Where("id IN ?", list)
	if schoolID > 0 {
		query = query.Where("school_id = ?", schoolID)
	}
	query = query.Delete(&model.UniversityLevelDomain{})

	result = query.RowsAffected
	err = query.Error
	return
}

func (repository *Repository) GetByID(id int64) (*model.UniversityLevel, error) {
	result := &model.UniversityLevel{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetLevelDomainByID(id int64) (*model.UniversityLevelDomain, error) {
	result := &model.UniversityLevelDomain{}
	return result, repository.Db.
		Preload(clause.Associations).
		Preload("Domain.Department").
		Preload("Domain.Department.Faculty").
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByIDSchoolID(id int64, schoolID int64) (*model.UniversityLevel, error) {
	result := &model.UniversityLevel{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetLevelDomainByIDSchoolID(id int64, schoolID int64) (*model.UniversityLevelDomain, error) {
	result := &model.UniversityLevelDomain{}
	return result, repository.Db.
		Preload(clause.Associations).
		Preload("Domain.Department").
		Preload("Domain.Department.Faculty").
		Where("id = ?", id).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
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

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "levels.school_id = ?")
			args = append(args, request.SchoolID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(levels.id AS TEXT) = ? OR
			levels.name ILIKE ? OR
			levels.description ILIKE ? OR
			schools.name ILIKE ? OR
			schools.type ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT DISTINCT levels.*
				FROM university_levels levels
				LEFT JOIN schools ON levels.school_id = schools.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}

func (repository *Repository) GetAllLevelDomain(
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllLevelDomainRequest,
) (result []model.UniversityLevelDomain, err error) {
	result = make([]model.UniversityLevelDomain, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "ld.school_id = ?")
			args = append(args, request.SchoolID)
		}
		if request.LevelID > 0 {
			where = helpers.AppendWhereClause(where, "ld.level_id = ?")
			args = append(args, request.LevelID)
		}
		if request.DomainID > 0 {
			where = helpers.AppendWhereClause(where, "ld.domain_id = ?")
			args = append(args, request.DomainID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(ld.id AS TEXT) = ? OR
			university_levels.name ILIKE ? OR
			university_levels.description ILIKE ? OR
			university_domains.name ILIKE ? OR
			university_domains.description ILIKE ? OR
			schools.name ILIKE ? OR
			schools.type ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Preload("Domain.Department").
		Preload("Domain.Department.Faculty").
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT DISTINCT ld.*
				FROM university_level_domains ld
				LEFT JOIN university_levels ON ld.level_id = university_levels.id
				LEFT JOIN university_domains ON ld.domain_id = university_domains.id
				LEFT JOIN schools ON ld.school_id = schools.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}
