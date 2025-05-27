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

func (repository *Repository) DeleteByID(
	id int64,
) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.Document{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetByID(
	id int64,
) (*model.Document, error) {
	result := &model.Document{}
	return result, repository.Db.Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(
	filter *types.Filter,
	pagination *types.Pagination,
	schoolID int64,
	yearID int64,
	classSubjectID int64,
	unitID int64,
) ([]model.Document, error) {
	var result []model.Document
	var where string = ""
	if schoolID > 0 {
		where = helpers.AppendWhereClause(where, fmt.Sprintf("documents.school_id = %d", schoolID))
	}
	if yearID > 0 {
		where = helpers.AppendWhereClause(where, fmt.Sprintf("documents.year_id = %d", yearID))
	}
	if classSubjectID > 0 {
		where = helpers.AppendWhereClause(where, fmt.Sprintf("documents.class_subject_id = %d", classSubjectID))
	}
	if unitID > 0 {
		where = helpers.AppendWhereClause(where, fmt.Sprintf("documents.unit_id = %d", unitID))
	}
	if filter != nil && len(filter.Search) >= 1 {
		tempWhere := fmt.Sprintf(
			"(type ILIKE %s OR WHERE name ILIKE %s)",
			filter.Search,
			filter.Search,
		)
		where = helpers.AppendWhereClause(where, tempWhere)
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
