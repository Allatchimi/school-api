package sequence

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/services/school/highschool/sequence/data"
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
	return result, repository.Db.Preload(clause.Associations).Model(&model.HighschoolSequence{}).Where("id = ?", id).Updates(
		map[string]any{
			"school_id":  item.SchoolID,
			"quarter_id": item.QuarterID,

			"name":        item.Name,
			"description": item.Description,
		},
	).Find(result).Error
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

func (repository *Repository) GetByID(id int64) (*model.HighschoolSequence, error) {
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
		(item1.SchoolID == item2.SchoolID &&
			item1.Name == item2.Name) {
		return true
	}
	return false
}

func (repository *Repository) GetAll(
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllRequest,
) (result []model.HighschoolSequence, err error) {
	result = make([]model.HighschoolSequence, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "sequences.school_id = ?")
			args = append(args, request.SchoolID)
		}
		if request.QuarterID > 0 {
			where = helpers.AppendWhereClause(where, "sequences.quarter_id = ?")
			args = append(args, request.QuarterID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(sequences.id AS TEXT) = ? OR 
			sequences.name ILIKE ? OR 
			sequences.description ILIKE ? OR 
			schools.name ILIKE ? OR 
			schools.type ILIKE ? OR 
			highschool_quarters.name ILIKE ? OR 
			highschool_quarters.description ILIKE ? OR 
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT sequences.* 
				FROM highschool_sequences sequences 
				LEFT JOIN schools ON sequences.school_id = schools.id 
				LEFT JOIN highschool_quarters ON sequences.quarter_id = highschool_quarters.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}
