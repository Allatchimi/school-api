package health

import (
	"api/services/user/role/model"

	"gorm.io/gorm"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) GetRole() (*model.Role, error) {
	result := &model.Role{
		Name: "default",
	}
	return result, repository.Db.Limit(1).Find(result).Error
}
