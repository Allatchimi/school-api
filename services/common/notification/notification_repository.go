package notification

import (
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/services/common/notification/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.Notification) (*model.Notification, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) UpdateSeenByIDUserID(id int64, userID int64, item *model.Notification) (*model.Notification, error) {
	result := &model.Notification{}
	return result, repository.Db.Preload(clause.Associations).
		Model(result).
		Where("id = ?", id).
		Where("user_id = ?", userID).
		Updates(
			map[string]any{
				"seen":    item.Seen,
				"seen_at": item.SeenAt,
			},
		).Error
}

func (repository *Repository) UpdateSeenAllByUserID(userID int64, seen bool, seenAt *time.Time) error {
	return repository.Db.Preload(clause.Associations).
		Model(&model.Notification{}).
		Where("user_id = ?", userID).
		Where("seen = ?", !seen).
		Updates(
			map[string]any{
				"seen":    seen,
				"seen_at": seenAt,
			},
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

func (repository *Repository) GetNotSeenCount(userID int64) (result int64, err error) {
	return result, repository.Db.
		Model(&model.Notification{}).
		Where("user_id = ?", userID).
		Where("seen = ?", false).
		Count(&result).
		Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, userID int64) (result []model.Notification, err error) {
	result = make([]model.Notification, 0)
	var where string = ""
	if userID > 0 {
		where = helpers.AppendWhereClause(where, fmt.Sprintf("notifications.user_id = %d", userID))
	}
	if filter != nil && len(filter.Search) > 0 {
		tempWhere := fmt.Sprintf(
			"(notifications.title ILIKE '%s')",
			"%"+filter.Search+"%",
		)
		where = helpers.AppendWhereClause(where, tempWhere)
	}
	tmpErr := repository.Db.
		Preload(clause.Associations).
		Scopes(
			helpers.PaginationScope(
				repository.Db,
				"SELECT notifications.* "+
					"FROM notifications "+
					"LEFT JOIN users ON notifications.user_id = users.id ",
				where,
				pagination,
				filter,
			),
		).Find(&result).Error

	err = tmpErr
	return
}
