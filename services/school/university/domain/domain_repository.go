package domain

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/services/school/university/domain/data"
	"api/services/school/university/domain/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.UniversityDomain) (*model.UniversityDomain, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) Update(id int64, item *model.UniversityDomain) (*model.UniversityDomain, error) {
	result := &model.UniversityDomain{}
	fmt.Println(item)
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"school_id":     item.SchoolID,
			"department_id": item.DepartmentID,
			"name":          item.Name,
			"description":   item.Description,
		},
	).Error
}

func (repository *Repository) Delete(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.UniversityDomain{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultiple(list []int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.UniversityDomain{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) GetByID(id int64) (*model.UniversityDomain, error) {
	result := &model.UniversityDomain{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetUniqueObject(item *model.UniversityDomain) (*model.UniversityDomain, error) {
	result := &model.UniversityDomain{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.UniversityDomain{
		SchoolID: item.SchoolID,
		Name:     item.Name,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjects(item1 *model.UniversityDomain, item2 *model.UniversityDomain) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.Name == item2.Name) {
		return true
	}
	return false
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, request *data.GetAllRequest) (result []model.UniversityDomain, err error) {
	result = make([]model.UniversityDomain, 0)
	var where string = ""
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("domains.school_id = %d", request.SchoolID))
		}
	}
	if filter != nil && len(filter.Search) >= 1 {
		tempWhere := fmt.Sprintf(
			"(CAST(domains.id AS TEXT) = '%s' OR domains.name ILIKE '%s' OR domains.description ILIKE '%s' OR schools.name ILIKE '%s' OR schools.type ILIKE '%s' OR departments.name ILIKE '%s' OR departments.description ILIKE '%s')",
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
		Preload("Department.Faculty").
		Scopes(
			helpers.PaginationScope(
				repository.Db,
				"SELECT domains.* "+
					"FROM university_domains domains "+
					"LEFT JOIN schools ON domains.school_id = schools.id "+
					"LEFT JOIN university_departments AS departments ON domains.department_id = departments.id",
				where,
				pagination,
				filter,
			),
		).Find(&result).Error

	err = tmpErr
	return
}
