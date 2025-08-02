package notification

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/services/others/notification/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.Notification) (result *model.Notification, err error) {
	item.Seen = false
	item.SeenAt = nil
	// Create the item
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.Notification{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) UpdateSeenByIDUserID(id int64, userID int64, item *model.Notification) (*model.Notification, error) {
	result := &model.Notification{}
	fields := map[string]any{
		"seen":    item.Seen,
		"seen_at": item.SeenAt,
	}
	return result, repository.Db.
		Preload(clause.Associations).
		Model(&model.Notification{}).
		Where("id = ?", id).
		Where("user_id = ?", userID).
		Updates(
			fields,
		).
		Find(result).Error
}

func (repository *Repository) UpdateSeenAllByUserID(userID int64, seen bool, seenAt *time.Time) error {
	fields := map[string]any{
		"seen":    seen,
		"seen_at": seenAt,
	}
	return repository.Db.
		Preload(clause.Associations).
		Model(&model.Notification{}).
		Where("user_id = ?", userID).
		Updates(
			fields,
		).Error
}

func (repository *Repository) DeleteByIDUserID(id int64, userID int64) (int64, error) {
	tempUnit, err := repository.GetByIDUserID(id, userID)
	if err != nil || tempUnit == nil || tempUnit.ID != id {
		return -1, err
	}

	result := repository.Db.Where("id = ?", id).Where("user_id = ?", userID).Delete(&model.Notification{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteAllByUserID(userID int64) (int64, error) {
	result := repository.Db.Where("user_id = ?", userID).Delete(&model.Notification{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetByID(id int64) (*model.Notification, error) {
	result := &model.Notification{}
	return result, repository.Db.Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByIDUserID(id int64, userID int64) (*model.Notification, error) {
	result := &model.Notification{}
	return result, repository.Db.Where("id = ?", id).Where("user_id = ?", userID).Limit(1).Find(result).Error
}

func (repository *Repository) GetNotSeenCountByUserID(userID int64) (result int64, err error) {
	var count int64
	countQuery := `
	SELECT COUNT(*) FROM notifications
	WHERE notifications.user_id = ?
	AND (notifications.seen = ? OR notifications.seen IS NULL)
	`
	args := []any{}
	args = append(args, userID, false)
	repository.Db.Raw(countQuery, args...).Count(&count)
	result = count
	return
}

func (repository *Repository) GetAllByUserID(filter *types.Filter, pagination *types.Pagination, userID int64) (result []model.Notification, err error) {
	result = make([]model.Notification, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if userID > 0 {
		where = helpers.AppendWhereClause(where, "notifications.user_id = ?")
		args = append(args, userID)
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(notifications.id AS TEXT) = ? OR
			notifications.title ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT DISTINCT notifications.*
				FROM notifications
				LEFT JOIN users ON notifications.user_id = users.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}
