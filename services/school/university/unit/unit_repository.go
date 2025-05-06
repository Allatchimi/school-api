package unit

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
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
	tempUnit, err := repository.GetById(id)
	if err != nil || tempUnit == nil || tempUnit.ID != id {
		return nil, err
	}

	result := &model.UniversityUnit{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"school_id":   item.SchoolID,
			"domain_id":   item.DomainID,
			"level_id":    item.LevelID,
			"semester_id": item.SemesterID,

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
	tempUnit, err := repository.GetById(id)
	if err != nil || tempUnit == nil || tempUnit.ID != id {
		return -1, err
	}

	result := repository.Db.Where("id = ?", id).Delete(&model.UniversityUnit{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetById(id int64) (*model.UniversityUnit, error) {
	result := &model.UniversityUnit{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetUniqueObject(item *model.UniversityUnit) (*model.UniversityUnit, error) {
	result := &model.UniversityUnit{}
	return result, repository.Db.Where(&model.UniversityUnit{
		SchoolID:   item.SchoolID,
		DomainID:   item.DomainID,
		LevelID:    item.LevelID,
		SemesterID: item.SemesterID,

		Name: item.Name,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjects(item1 *model.UniversityUnit, item2 *model.UniversityUnit) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.DomainID == item2.DomainID &&
			item1.LevelID == item2.LevelID &&
			item1.SemesterID == item2.SemesterID &&
			item1.Name == item2.Name) {
		return true
	}
	return false
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, schoolID int64) (result []model.UniversityUnit, err error) {
	result = make([]model.UniversityUnit, 0)
	var where string = ""
	if schoolID > 0 {
		where = fmt.Sprintf("WHERE tus.school_id = %d", schoolID)
	}
	if filter != nil && len(filter.Search) >= 1 {
		tempWhere := fmt.Sprintf(
			"CAST(tus.id AS TEXT) = '%s' OR tus.name ILIKE '%s' OR tus.description ILIKE '%s' OR schools.name ILIKE '%s' OR schools.type ILIKE '%s' OR university_domains.name ILIKE '%s' OR university_levels.name ILIKE '%s' OR university_semesters.name ILIKE '%s'",
			filter.Search,
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
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
	tmpErr := repository.Db.Preload(clause.Associations).Scopes(
		helpers.PaginationScope(
			repository.Db,
			"SELECT tus.id, tus.name, tus.description, tus.school_id, tus.domain_id, tus.level_id, tus.semester_id"+
				", tus.created_at, tus.updated_at FROM university_units tus "+
				"LEFT JOIN schools ON tus.school_id = schools.id "+
				"LEFT JOIN university_domains ON tus.domain_id = university_domains.id "+
				"LEFT JOIN university_levels ON tus.level_id = university_levels.id "+
				"LEFT JOIN university_semesters ON tus.semester_id = university_semesters.id ",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error

	err = tmpErr
	return
}
