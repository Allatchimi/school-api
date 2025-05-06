package sequence

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/services/school/highschool/sequence/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.HighschoolSequence) (*model.HighschoolSequence, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) Update(id int64, item *model.HighschoolSequence) (*model.HighschoolSequence, error) {
	result := &model.HighschoolSequence{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"school_id":   item.SchoolID,
			"name":        item.Name,
			"description": item.Description,
		},
	).Error
}

func (repository *Repository) Delete(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.HighschoolSequence{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultiple(list []int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.HighschoolSequence{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) GetById(id int64) (*model.HighschoolSequence, error) {
	result := &model.HighschoolSequence{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetUniqueObject(item *model.HighschoolSequence) (*model.HighschoolSequence, error) {
	result := &model.HighschoolSequence{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.HighschoolSequence{
		SchoolID: item.SchoolID,
		Name:     item.Name,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjects(item1 *model.HighschoolSequence, item2 *model.HighschoolSequence) bool {
	if item1 != nil && item2 != nil &&
		!(item1.SchoolID == item2.SchoolID &&
			item1.Name == item2.Name) {
		return true
	}
	return false
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, schoolID int64) (result []model.HighschoolSequence, err error) {
	result = make([]model.HighschoolSequence, 0)
	var where string = ""
	if schoolID > 0 {
		where = fmt.Sprintf("WHERE sequences.school_id = %d", schoolID)
	}
	if filter != nil && len(filter.Search) >= 1 {
		tempWhere := fmt.Sprintf(
			"CAST(sequences.id AS TEXT) = '%s' OR sequences.name ILIKE '%s' OR sequences.description ILIKE '%s' OR schools.name ILIKE '%s' OR schools.type ILIKE '%s'",
			filter.Search,
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
		)
		if strings.HasPrefix(where, "WHERE") {
			where = fmt.Sprintf("%s AND (%s)", where, tempWhere)
		} else {
			where = fmt.Sprintf("WHERE %s", tempWhere)
		}
	}
	tmpErr := repository.Db.Preload(clause.Associations).Scopes(
		helpers.PaginationScope(
			repository.Db,
			"SELECT sequences.id, sequences.name, sequences.description, sequences.school_id"+
				", sequences.created_at, sequences.updated_at FROM highschool_sequences sequences "+
				"LEFT JOIN schools ON sequences.school_id = schools.id",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error

	err = tmpErr
	return
}
