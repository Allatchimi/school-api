package semester

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/services/school/university/semester/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.UniversitySemester) (*model.UniversitySemester, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) Update(id int64, item *model.UniversitySemester) (*model.UniversitySemester, error) {
	result := &model.UniversitySemester{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]interface{}{
			"school_id":   item.SchoolID,
			"name":        item.Name,
			"description": item.Description,
		},
	).Error
}

func (repository *Repository) Delete(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.UniversitySemester{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultiple(list []int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.UniversitySemester{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) GetById(id int64) (*model.UniversitySemester, error) {
	result := &model.UniversitySemester{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetBySchoolIDName(schoolID int64, name string) (*model.UniversitySemester, error) {
	result := &model.UniversitySemester{}
	return result, repository.Db.Preload(clause.Associations).Where("school_id = ?", schoolID).Where("name = ?", name).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, schoolID int64) (result []model.UniversitySemester, err error) {
	result = make([]model.UniversitySemester, 0)
	var where string = ""
	if schoolID > 0 {
		where = fmt.Sprintf("WHERE semesters.school_id = %d", schoolID)
	}
	if filter != nil && len(filter.Search) >= 1 {
		tempWhere := fmt.Sprintf(
			"CAST(semesters.id AS TEXT) = '%s' OR semesters.name ILIKE '%s' OR semesters.description ILIKE '%s' OR schools.name ILIKE '%s' OR schools.type ILIKE '%s'",
			filter.Search,
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
			"SELECT semesters.id, semesters.name, semesters.description, semesters.school_id"+
				", semesters.created_at, semesters.updated_at FROM university_semesters semesters "+
				"LEFT JOIN schools ON semesters.school_id = schools.id",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error

	err = tmpErr
	return
}
