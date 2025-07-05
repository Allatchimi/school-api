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
) (*model.Director, error) {
	result := *item
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) UpdateByID(
	id int64,
	item *model.Director,
) (*model.Director, error) {
	tempDirector, err := repository.GetByID(id)
	if err != nil || tempDirector == nil || tempDirector.ID != id {
		return nil, err
	}
	result := &model.Director{}
	return result, repository.Db.Model(result).Where("id = ?", item.ID).Updates(
		map[string]any{
			"school_id": item.SchoolID,
			"user_id":   item.UserID,
			"uid":       item.UID,
		},
	).Error
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
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.Director{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) CountAll(schoolID int64) (result int64, err error) {
	if schoolID <= 1 {
		err = repository.Db.Model(&model.Director{}).Count(&result).Error
		return
	}
	err = repository.Db.Model(&model.Director{}).Where("school_id = ?", schoolID).Count(&result).Error
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

func (repository *Repository) GetByUserIDSchoolID(
	userID int64,
	schoolID int64,
) (*model.Director, error) {
	result := &model.Director{}
	return result, repository.Db.
		Preload(clause.Associations).
		Preload("User.Info").
		Where("user_id = ?", userID).Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllRequest,
) (result []model.Director, err error) {
	result = make([]model.Director, 0)
	var where string = ""
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("directors.school_id = %d", request.SchoolID))
		}
	}
	if filter != nil && len(filter.Search) > 0 {
		tempWhere := fmt.Sprintf(
			"(CAST(directors.id AS TEXT) = '%s' OR directors.uid ILIKE '%s' OR schools.name ILIKE '%s' OR schools.type ILIKE '%s' OR users.email ILIKE '%s')",
			filter.Search,
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
		)
		where = helpers.AppendWhereClause(where, tempWhere)
	}
	tmpErr := repository.Db.
		Preload(clause.Associations).
		Scopes(
			helpers.PaginationScope(
				repository.Db,
				"SELECT directors.* "+
					"FROM directors "+
					"LEFT JOIN schools ON directors.school_id = schools.id "+
					"LEFT JOIN users ON directors.user_id = users.id ",
				where,
				pagination,
				filter,
			),
		).Find(&result).Error

	err = tmpErr
	return
}
