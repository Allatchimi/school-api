package specialty

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/services/school/highschool/specialty/data"
	"api/services/school/highschool/specialty/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.HighschoolSpecialty) (*model.HighschoolSpecialty, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) Update(id int64, item *model.HighschoolSpecialty) (*model.HighschoolSpecialty, error) {
	result := &model.HighschoolSpecialty{}
	fmt.Println(item)
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"school_id":   item.SchoolID,
			"section_id":  item.SectionID,
			"name":        item.Name,
			"description": item.Description,
		},
	).Error
}

func (repository *Repository) Delete(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.HighschoolSpecialty{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultiple(list []int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.HighschoolSpecialty{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) GetByID(id int64) (*model.HighschoolSpecialty, error) {
	result := &model.HighschoolSpecialty{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetUniqueObject(item *model.HighschoolSpecialty) (*model.HighschoolSpecialty, error) {
	result := &model.HighschoolSpecialty{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.HighschoolSpecialty{
		SchoolID: item.SchoolID,
		Name:     item.Name,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjects(item1 *model.HighschoolSpecialty, item2 *model.HighschoolSpecialty) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.Name == item2.Name) {
		return true
	}
	return false
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, request *data.GetAllRequest) (result []model.HighschoolSpecialty, err error) {
	result = make([]model.HighschoolSpecialty, 0)
	var where string = ""
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("specialties.school_id = %d", request.SchoolID))
		}
	}
	if filter != nil && len(filter.Search) >= 1 {
		tempWhere := fmt.Sprintf(
			"(CAST(specialties.id AS TEXT) = '%s' OR specialties.name ILIKE '%s' OR specialties.description ILIKE '%s' OR schools.name ILIKE '%s' OR schools.type ILIKE '%s' OR sections.name ILIKE '%s' OR sections.description ILIKE '%s')",
			filter.Search,
			"%"+filter.Search+"%",
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
				"SELECT specialties.* "+
					"FROM highschool_specialties specialties "+
					"LEFT JOIN schools ON specialties.school_id = schools.id "+
					"LEFT JOIN highschool_sections AS sections ON specialties.section_id = sections.id",
				where,
				pagination,
				filter,
			),
		).Find(&result).Error

	err = tmpErr
	return
}
