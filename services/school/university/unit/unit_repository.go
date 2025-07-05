package unit

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/university/unit/data"
	"api/services/school/university/unit/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.UniversityUnit) (*model.UniversityUnit, error) {
	result := *item
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) Update(id int64, item *model.UniversityUnit) (*model.UniversityUnit, error) {
	tempUnit, err := repository.GetByID(id)
	if err != nil || tempUnit == nil || tempUnit.ID != id {
		return nil, err
	}

	result := &model.UniversityUnit{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"school_id":       item.SchoolID,
			"level_domain_id": item.LevelDomainID,
			"semester_id":     item.SemesterID,

			"name":         item.Name,
			"description":  item.Description,
			"credit":       item.Credit,
			"program":      item.Program,
			"requirements": item.Requirements,

			"is_valid":     item.IsValid,
			"invalid_date": item.InvalidDate,
		},
	).Error
}

func (repository *Repository) Delete(id int64) (int64, error) {
	tempUnit, err := repository.GetByID(id)
	if err != nil || tempUnit == nil || tempUnit.ID != id {
		return -1, err
	}

	result := repository.Db.Where("id = ?", id).Delete(&model.UniversityUnit{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetByID(id int64) (*model.UniversityUnit, error) {
	result := &model.UniversityUnit{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetUniqueObject(item *model.UniversityUnit) (*model.UniversityUnit, error) {
	result := &model.UniversityUnit{}
	return result, repository.Db.Where(&model.UniversityUnit{
		SchoolID:      item.SchoolID,
		LevelDomainID: item.LevelDomainID,
		SemesterID:    item.SemesterID,

		Name: item.Name,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjects(item1 *model.UniversityUnit, item2 *model.UniversityUnit) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.LevelDomainID == item2.LevelDomainID &&
			item1.SemesterID == item2.SemesterID &&
			item1.Name == item2.Name) {
		return true
	}
	return false
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, request *data.GetAllRequest) (result []model.UniversityUnit, err error) {
	result = make([]model.UniversityUnit, 0)
	var where string = ""
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("units.school_id = %d", request.SchoolID))
		}
		if request.LevelDomainID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("units.level_domain_id = %d", request.LevelDomainID))
		}
		if request.SemesterID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("units.semester_id = %d", request.SemesterID))
		}
	}
	if filter != nil && len(filter.Search) > 0 {
		tempWhere := fmt.Sprintf(
			"(CAST(units.id AS TEXT) = '%s' OR units.name ILIKE '%s' OR units.description ILIKE '%s' OR schools.name ILIKE '%s' OR schools.type ILIKE '%s' OR university_semesters.name ILIKE '%s')",
			filter.Search,
			"%"+filter.Search+"%",
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
				"SELECT units.* "+
					"FROM university_units units "+
					"LEFT JOIN schools ON units.school_id = schools.id "+
					"LEFT JOIN university_level_domains ON units.level_domain_id = university_level_domains.id "+
					"LEFT JOIN university_semesters ON units.semester_id = university_semesters.id ",
				where,
				pagination,
				filter,
			),
		).Find(&result).Error

	err = tmpErr
	return
}
