package quarter

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/highschool/quarter/data"
	"api/services/school/highschool/quarter/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.HighschoolQuarter) (result *model.HighschoolQuarter, err error) {
	// Create the item
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.HighschoolQuarter{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) UpdateByID(id int64, item *model.HighschoolQuarter) (result *model.HighschoolQuarter, err error) {
	// Update the item
	fields := map[string]any{
		"school_id": item.SchoolID,

		"name":        item.Name,
		"description": item.Description,
	}
	err = repository.Db.
		Model(&model.HighschoolQuarter{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.HighschoolQuarter{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) DeleteByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.HighschoolQuarter{})
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
	query = query.Delete(&model.HighschoolQuarter{})

	result = query.RowsAffected
	err = query.Error
	return
}

func (repository *Repository) GetByID(id int64) (*model.HighschoolQuarter, error) {
	result := &model.HighschoolQuarter{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByIDSchoolID(id int64, schoolID int64) (*model.HighschoolQuarter, error) {
	result := &model.HighschoolQuarter{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetUniqueObject(item *model.HighschoolQuarter) (*model.HighschoolQuarter, error) {
	result := &model.HighschoolQuarter{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.HighschoolQuarter{
		SchoolID: item.SchoolID,
		Name:     item.Name,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjects(item1 *model.HighschoolQuarter, item2 *model.HighschoolQuarter) bool {
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
) (result []model.HighschoolQuarter, err error) {
	result = make([]model.HighschoolQuarter, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "quarters.school_id = ?")
			args = append(args, request.SchoolID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(quarters.id AS TEXT) = ? OR
			quarters.name ILIKE ? OR
			quarters.description ILIKE ? OR
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
				`SELECT quarters.*
				FROM highschool_quarters quarters
				LEFT JOIN schools ON quarters.school_id = schools.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}
