package parent

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/services/school/common/parent/data"
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

func (repository *Repository) UpdateByID(id int64, item *model.Parent) (*model.Parent, error) {
	tempParent, err := repository.GetByID(id)
	if err != nil || tempParent == nil || tempParent.ID != id {
		return nil, err
	}

	result := &model.Parent{}
	return result, repository.Db.Model(&model.ParentStudent{}).Where("id = ?", item.ID).Updates(
		map[string]any{
			"user_id": item.UserID,
		},
	).Find(result).Error
}

func (repository *Repository) UpdateParentStudent(id int64, item *model.ParentStudent) (*model.ParentStudent, error) {
	tempParent, err := repository.GetByID(id)
	if err != nil || tempParent == nil || tempParent.ID != id {
		return nil, err
	}

	result := &model.ParentStudent{}
	return result, repository.Db.Model(&model.ParentStudent{}).Where("id = ?", item.ID).Updates(
		map[string]any{
			"parent_id":  item.ParentID,
			"student_id": item.StudentID,
		},
	).Find(result).Error
}

func (repository *Repository) DeleteByID(id int64) (int64, error) {
	foundItem, err := repository.GetByID(id)
	if err != nil || foundItem == nil || foundItem.ID != id {
		return -1, err
	}

	result := repository.Db.Where("id = ?", id).Delete(&model.Parent{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteParentStudentByID(id int64) (int64, error) {
	foundItem, err := repository.GetParentStudentByID(id)
	if err != nil || foundItem == nil || foundItem.ID != id {
		return -1, err
	}

	result := repository.Db.Where("id = ?", id).Delete(&model.ParentStudent{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultipleByID(list []int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.Parent{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) CountAll(schoolID int64) (result int64, err error) {
	if schoolID <= 1 {
		err = repository.Db.Model(&model.Parent{}).Count(&result).Error
		return
	}
	err = repository.Db.Model(&model.Parent{}).Where("school_id = ?", schoolID).Count(&result).Error
	return
}

func (repository *Repository) GetByID(id int64) (*model.Parent, error) {
	result := &model.Parent{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByUserID(userID int64) (*model.Parent, error) {
	result := &model.Parent{}
	return result, repository.Db.
		Preload(clause.Associations).
		Preload("User.Info").
		Where("user_id = ?", userID).Limit(1).Find(result).Error
}

func (repository *Repository) GetParentStudentByID(id int64) (*model.ParentStudent, error) {
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

func (repository *Repository) GetAll(
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllRequest,
) (result []model.Parent, err error) {
	result = make([]model.Parent, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "parents.school_id = ?")
			args = append(args, request.SchoolID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(parents.id AS TEXT) = ? OR
			schools.name ILIKE ? OR
			schools.type ILIKE ? OR
			users.email ILIKE ? OR
			CAST(users.phone_number AS TEXT) ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Preload("User.Info").
		Preload("User.Role").
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT parents.*
				FROM parents
				LEFT JOIN schools ON parents.school_id = schools.id
				LEFT JOIN users ON parents.user_id = users.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}

func (repository *Repository) GetAllParentStudent(
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllParentStudentRequest,
) (result []model.ParentStudent, err error) {
	result = make([]model.ParentStudent, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "students.school_id = ?")
			args = append(args, request.SchoolID)
		}
		if request.YearID > 0 {
			where = helpers.AppendWhereClause(where, "student_enrolls.year_id = ?")
			args = append(args, request.YearID)
		}
		if request.ClassSubjectID > 0 {
			where = helpers.AppendWhereClause(where, "student_enrolls.class_subject_id = ?")
			args = append(args, request.ClassSubjectID)
		}
		if request.UnitID > 0 {
			where = helpers.AppendWhereClause(where, "student_enrolls.unit_id = ?")
			args = append(args, request.UnitID)
		}
		if request.ParentID > 0 {
			where = helpers.AppendWhereClause(where, "parent_students.parent_id = ?")
			args = append(args, request.ParentID)
		}
		if request.StudentID > 0 {
			where = helpers.AppendWhereClause(where, "parent_students.student_id = ?")
			args = append(args, request.StudentID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(directors.id AS TEXT) = ? OR
			students.uid ILIKE ? OR
			schools.name ILIKE ? OR
			schools.type ILIKE ? OR
			years.name ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Preload("Parent.User").
		Preload("Parent.User.Info").
		Preload("Student.User").
		Preload("Student.User.Info").
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT parent_students.*
				FROM parent_students
				LEFT JOIN parents ON parent_students.parent_id = parents.id
				LEFT JOIN students ON parent_students.student_id = students.id
				LEFT JOIN schools ON students.school_id = schools.id
				LEFT JOIN student_enrolls ON students.id = student_enrolls.student_id
				LEFT JOIN years ON student_enrolls.year_id = years.id
				LEFT JOIN highschool_class_subjects ON student_enrolls.class_subject_id = highschool_class_subjects.id
				LEFT JOIN university_units ON student_enrolls.unit_id = university_units.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}
