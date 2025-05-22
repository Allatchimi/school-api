package health

import (
	"api/services/history/model"

	"gorm.io/gorm"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) GetHistory() (*model.History, error) {
	result := &model.History{}
	return result, repository.Db.Limit(1).Find(result).Error
}
