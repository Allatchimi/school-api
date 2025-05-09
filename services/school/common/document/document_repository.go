package document

import (
	"fmt"

	"gorm.io/gorm"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/common/document/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.Document) (*model.Document, error) {
	result := *item
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) Delete(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.Document{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetById(id int64) (*model.Document, error) {
	result := &model.Document{}
	return result, repository.Db.Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByObject(item *model.Document) (*model.Document, error) {
	result := &model.Document{}
	return result, repository.Db.Where(item).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, schoolID int64) ([]model.Document, error) {
	var result []model.Document
	var where string = ""
	if filter != nil && len(filter.Search) >= 1 {
		where = fmt.Sprintf(
			"WHERE school_id = '%d' AND (type ILIKE %s OR WHERE name ILIKE %s)",
			schoolID,
			filter.Search,
			filter.Search,
		)
	}
	return result, repository.Db.Scopes(
		helpers.PaginationScope(
			repository.Db,
			"documents",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error
}
