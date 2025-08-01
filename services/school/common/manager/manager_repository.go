package manager

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/common/manager/data"
	"api/services/school/common/manager/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(
	item *model.Manager,
) (result *model.Manager, err error) {
	// Create the item
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.Manager{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) UpdateByID(
	id int64,
	item *model.Manager,
) (result *model.Manager, err error) {
	// Update the item
	fields := map[string]any{
		"school_id": item.SchoolID,
		"user_id":   item.UserID,

		"uid": item.UID,
	}
	err = repository.Db.
		Model(&model.Manager{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.Manager{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) DeleteByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.Manager{})
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
	query = query.Delete(&model.Manager{})

	result = query.RowsAffected
	err = query.Error
	return
}

func (repository *Repository) GetByID(id int64) (*model.Manager, error) {
	result := &model.Manager{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByUserID(userID int64) (*model.Manager, error) {
	result := &model.Manager{}
	return result, repository.Db.Preload(clause.Associations).
		Where("user_id = ?", userID).Limit(1).Find(result).Error
}

func (repository *Repository) GetByIDSchoolID(id int64, schoolID int64) (*model.Manager, error) {
	result := &model.Manager{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetUniqueObjectByUserID(
	item *model.Manager,
) (*model.Manager, error) {
	result := &model.Manager{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.Manager{
		SchoolID: item.SchoolID,
		UserID:   item.UserID,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjectsByUserID(
	item1 *model.Manager,
	item2 *model.Manager,
) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.UserID == item2.UserID) {
		return true
	}
	return false
}

func (repository *Repository) GetUniqueObjectByUID(
	item *model.Manager,
) (*model.Manager, error) {
	result := &model.Manager{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.Manager{
		SchoolID: item.SchoolID,
		UID:      item.UID,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjectsByUID(
	item1 *model.Manager,
	item2 *model.Manager,
) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.UID == item2.UID) {
		return true
	}
	return false
}

func (repository *Repository) GetAll(
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllRequest,
) (result []model.Manager, err error) {
	result = make([]model.Manager, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "managers.school_id = ?")
			args = append(args, request.SchoolID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(managers.id AS TEXT) = ? OR
			managers.uid ILIKE ? OR
			schools.name ILIKE ? OR
			schools.type ILIKE ? OR
			users.email ILIKE ? OR
			CAST(users.phone_number AS TEXT) ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Preload("User.Info").
		Preload("User.Role").
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT managers.*
				FROM managers
				LEFT JOIN schools ON managers.school_id = schools.id
				LEFT JOIN users ON managers.user_id = users.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}
