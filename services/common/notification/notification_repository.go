package notification

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/services/common/notification/data"
	"api/services/common/notification/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) GetByID(notificationID int64) (*model.Notification, error) {
	result := &model.Notification{}
	return result, repository.Db.Where("id = ?", notificationID).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, request *data.GetAllRequest) (result []model.Notification, err error) {
	result = make([]model.Notification, 0)
	var where string = ""
	if filter != nil && len(filter.Search) >= 1 {
		tempWhere := fmt.Sprintf(
			"(title ILIKE '%s' OR message ILIKE '%s')",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
		)
		where = helpers.AppendWhereClause(where, tempWhere)
	}
	tmpErr := repository.Db.Preload(clause.Associations).Scopes(
		helpers.PaginationScope(
			repository.Db,
			"SELECT * FROM notifications",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error

	err = tmpErr
	return
}
