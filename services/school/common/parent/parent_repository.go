package parent

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/common/parent/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.Parent) (*model.Parent, error) {
	result := *item
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) CreateParentStudent(item *model.ParentStudent) (*model.ParentStudent, error) {
	result := *item
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) Update(id int64, item *model.Parent) (*model.Parent, error) {
	tempParent, err := repository.GetById(id)
	if err != nil || tempParent == nil || tempParent.ID != id {
		return nil, err
	}

	result := &model.Parent{}
	return result, repository.Db.Model(result).Where("id = ?", item.ID).Updates(
		map[string]any{
			"user_id": item.UserID,
			"uid":     item.UID,
		},
	).Error
}

func (repository *Repository) UpdateParentStudent(id int64, item *model.ParentStudent) (*model.ParentStudent, error) {
	tempParent, err := repository.GetById(id)
	if err != nil || tempParent == nil || tempParent.ID != id {
		return nil, err
	}

	result := &model.ParentStudent{}
	return result, repository.Db.Model(result).Where("id = ?", item.ID).Updates(
		map[string]any{
			"parent_id":  item.ParentID,
			"student_id": item.StudentID,
		},
	).Error
}

func (repository *Repository) Delete(id int64) (int64, error) {
	foundItem, err := repository.GetById(id)
	if err != nil || foundItem == nil || foundItem.ID != id {
		return -1, err
	}

	result := repository.Db.Where("id = ?", id).Delete(&model.Parent{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteParentStudent(id int64) (int64, error) {
	foundItem, err := repository.GetParentStudentById(id)
	if err != nil || foundItem == nil || foundItem.ID != id {
		return -1, err
	}

	result := repository.Db.Where("id = ?", id).Delete(&model.ParentStudent{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetById(id int64) (*model.Parent, error) {
	result := &model.Parent{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetParentStudentById(id int64) (*model.ParentStudent, error) {
	result := &model.ParentStudent{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByObject(item *model.Parent) (*model.Parent, error) {
	result := &model.Parent{}
	return result, repository.Db.Where(item).Limit(1).Find(result).Error
}

func (repository *Repository) GetParentStudentByObject(item *model.ParentStudent) (*model.ParentStudent, error) {
	result := &model.ParentStudent{}
	return result, repository.Db.Where(item).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, schoolID int64) (result []model.Parent, err error) {
	result = make([]model.Parent, 0)
	var where string = ""
	if schoolID > 0 {
		where = fmt.Sprintf("WHERE parents.school_id = %d", schoolID)
	}
	if filter != nil && len(filter.Search) >= 1 {
		tempWhere := fmt.Sprintf(
			"CAST(parents.id AS TEXT) = '%s' OR parents.uid ILIKE '%s' OR schools.name ILIKE '%s' OR schools.type ILIKE '%s' OR users.email ILIKE '%s' OR users.phone_number ILIKE '%s'",
			filter.Search,
			"%"+filter.Search+"%",
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
			"SELECT parents.id, parents.school_id, parents.user_id, parents.uid"+
				", parents.created_at, parents.updated_at FROM parents "+
				"LEFT JOIN users ON parents.user_id = users.id",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error

	err = tmpErr
	return
}

func (repository *Repository) GetAllParentStudent(filter *types.Filter, pagination *types.Pagination, parentID int64) (result []model.ParentStudent, err error) {
	result = make([]model.ParentStudent, 0)
	var where string = ""
	if parentID > 0 {
		where = fmt.Sprintf("WHERE parent_lc.parent_id = %d", parentID)
	}
	if filter != nil && len(filter.Search) >= 1 {
		tempWhere := fmt.Sprintf(
			"CAST(parent_lc.id AS TEXT) = '%s' OR students.uid ILIKE '%s'",
			filter.Search,
			"%"+filter.Search+"%",
		)
		where = fmt.Sprintf("WHERE %s", tempWhere)
	}
	tmpErr := repository.Db.Preload(clause.Associations).Scopes(
		helpers.PaginationScope(
			repository.Db,
			"SELECT parent_lc.id, parent_lc.parent_id, parent_lc.student_id"+
				", parent_lc.created_at, parent_lc.updated_at FROM parent_level_classes AS parent_lc "+
				"LEFT JOIN students ON parent_lc.student_id = students.id ",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error

	err = tmpErr
	return
}
