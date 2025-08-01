package communication

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/services/others/communication/data"
	"api/services/others/communication/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.Communication) (result *model.Communication, err error) {
	// Create the item
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.Communication{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) Delete(id int64) (result int64, err error) {
	tmpResult := repository.Db.Where("id = ?", id).Delete(&model.Communication{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) DeleteMultipleByID(list []int64) (result int64, err error) {
	if len(list) < 1 {
		return
	}
	tmpResult := repository.Db.
		Where("id IN (?)", list).
		Delete(&model.Communication{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) GetByID(id int64) (*model.Communication, error) {
	result := &model.Communication{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByIDSchoolID(id int64, schoolID int64) (*model.Communication, error) {
	result := &model.Communication{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Where("school_id = ?", schoolID).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, request *data.GetAllRequest) (result []model.Communication, err error) {
	result = make([]model.Communication, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "communications.school_id = ?")
			args = append(args, request.SchoolID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(communications.id AS TEXT) = ? OR
			communications.subject ILIKE ? OR
			communications.message ILIKE ? OR
			schools.name ILIKE ? OR
			schools.type ILIKE ? OR
			roles.name ILIKE ? OR
			roles.feature ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT communications.*
				FROM communications
				LEFT JOIN schools ON communications.school_id = schools.id
				LEFT JOIN roles ON communications.role_id = roles.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}
