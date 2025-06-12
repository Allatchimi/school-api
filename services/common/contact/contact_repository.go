package contact

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/services/common/contact/data"
	"api/services/common/contact/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(contact *model.Contact) (*model.Contact, error) {
	result := *contact
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) Delete(roleID int64) (result int64, err error) {
	tmpResult := repository.Db.Where("id = ?", roleID).Delete(&model.Contact{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) DeleteMultiple(list []int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.Contact{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) GetByID(contactID int64) (*model.Contact, error) {
	result := &model.Contact{}
	return result, repository.Db.Where("id = ?", contactID).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, request *data.GetAllRequest) (result []model.Contact, err error) {
	result = make([]model.Contact, 0)
	var where string = ""
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("contacts.school_id = %d", request.SchoolID))
		}
	}
	if filter != nil && len(filter.Search) >= 1 {
		tempWhere := fmt.Sprintf(
			"(subject ILIKE '%s' OR email ILIKE '%s' OR message ILIKE '%s')",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
		)
		where = helpers.AppendWhereClause(where, tempWhere)
	}
	tmpErr := repository.Db.Preload(clause.Associations).Scopes(
		helpers.PaginationScope(
			repository.Db,
			"SELECT * FROM contacts",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error

	err = tmpErr
	return
}
