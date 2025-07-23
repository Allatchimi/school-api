package department

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/services/school/university/department/data"
	"api/services/school/university/department/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.UniversityDepartment) (*model.UniversityDepartment, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) Update(id int64, item *model.UniversityDepartment) (*model.UniversityDepartment, error) {
	result := &model.UniversityDepartment{}
	return result, repository.Db.Preload(clause.Associations).Model(&model.UniversityDepartment{}).Where("id = ?", id).Updates(
		map[string]any{
			"school_id":   item.SchoolID,
			"faculty_id":  item.FacultyID,
			"name":        item.Name,
			"description": item.Description,
		},
	).Find(result).Error
}

func (repository *Repository) Delete(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.UniversityDepartment{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultiple(list []int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.UniversityDepartment{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) GetByID(id int64) (*model.UniversityDepartment, error) {
	result := &model.UniversityDepartment{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetUniqueObject(item *model.UniversityDepartment) (*model.UniversityDepartment, error) {
	result := &model.UniversityDepartment{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.UniversityDepartment{
		SchoolID: item.SchoolID,
		Name:     item.Name,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjects(item1 *model.UniversityDepartment, item2 *model.UniversityDepartment) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.Name == item2.Name) {
		return true
	}
	return false
}

func (repository *Repository) GetAll(
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllRequest,
) (result []model.UniversityDepartment, err error) {
	result = make([]model.UniversityDepartment, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "departments.school_id = ?")
			args = append(args, request.SchoolID)
		}
		if request.FacultyID > 0 {
			where = helpers.AppendWhereClause(where, "departments.faculty_id = ?")
			args = append(args, request.FacultyID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(departments.id AS TEXT) = ? OR 
			departments.name ILIKE ? OR 
			departments.description ILIKE ? OR 
			schools.name ILIKE ? OR 
			schools.type ILIKE ? OR 
			university_faculties.name ILIKE ? OR 
			university_faculties.description ILIKE ? 
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
				`SELECT departments.* 
				FROM university_departments departments 
				LEFT JOIN schools ON departments.school_id = schools.id 
				LEFT JOIN university_faculties ON departments.faculty_id = university_faculties.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}
