package director

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/services/school/common/director/data"
	"api/services/school/common/director/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(
	item *model.Director,
) (result *model.Director, err error) {
	// Create the item
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.Director{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) UpdateByID(
	id int64,
	item *model.Director,
) (result *model.Director, err error) {
	// Update the item
	fields := map[string]any{
		"school_id": item.SchoolID,
		"user_id":   item.UserID,
	}
	err = repository.Db.
		Model(&model.Director{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.Director{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) DeleteByID(
	id int64,
) (int64, error) {
	foundItem, err := repository.GetByID(id)
	if err != nil || foundItem == nil || foundItem.ID != id {
		return -1, err
	}
	result := repository.Db.Where("id = ?", id).Delete(&model.Director{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultipleByID(list []int64) (result int64, err error) {
	if len(list) < 1 {
		return
	}
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.Director{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) GetByID(
	id int64,
) (*model.Director, error) {
	result := &model.Director{}
	return result, repository.Db.
		Preload(clause.Associations).
		Preload("User.Info").
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByUserID(
	userID int64,
) (*model.Director, error) {
	result := &model.Director{}
	return result, repository.Db.
		Preload(clause.Associations).
		Preload("User.Info").
		Where("user_id = ?", userID).Limit(1).Find(result).Error
}

func (repository *Repository) GetUniqueObjectByUserID(
	item *model.Director,
) (*model.Director, error) {
	result := &model.Director{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.Director{
		SchoolID: item.SchoolID,
		UserID:   item.UserID,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjectsByUserID(
	item1 *model.Director,
	item2 *model.Director,
) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.UserID == item2.UserID) {
		return true
	}
	return false
}

func (repository *Repository) GetUniqueObjectByUID(
	item *model.Director,
) (*model.Director, error) {
	result := &model.Director{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.Director{
		SchoolID: item.SchoolID,
		UID:      item.UID,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjectsByUID(
	item1 *model.Director,
	item2 *model.Director,
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
) (result []model.Director, err error) {
	result = make([]model.Director, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "directors.school_id = ?")
			args = append(args, request.SchoolID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(directors.id AS TEXT) = ? OR
			directors.uid ILIKE ? OR
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
				`SELECT directors.*
				FROM directors
				LEFT JOIN schools ON directors.school_id = schools.id
				LEFT JOIN users ON directors.user_id = users.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}
