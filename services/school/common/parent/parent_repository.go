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

func (repository *Repository) CreateParentAssign(item *model.ParentAssign) (*model.ParentAssign, error) {
	result := *item
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) UpdateByID(id int64, item *model.Parent) (*model.Parent, error) {
	result := &model.Parent{}
	return result, repository.Db.Preload(clause.Associations).Model(&model.ParentStudent{}).Where("id = ?", item.ID).Updates(
		map[string]any{
			"user_id": item.UserID,
		},
	).Find(result).Error
}

func (repository *Repository) UpdateParentStudentByID(id int64, item *model.ParentStudent) (*model.ParentStudent, error) {
	result := &model.ParentStudent{}
	return result, repository.Db.Preload(clause.Associations).Model(&model.ParentStudent{}).Where("id = ?", item.ID).Updates(
		map[string]any{
			"parent_id":  item.ParentID,
			"student_id": item.StudentID,
		},
	).Find(result).Error
}

func (repository *Repository) UpdateParentAssignByID(id int64, item *model.ParentAssign) (*model.ParentAssign, error) {
	result := &model.ParentAssign{}
	return result, repository.Db.Preload(clause.Associations).Model(&model.ParentAssign{}).Where("id = ?", item.ID).Updates(
		map[string]any{
			"school_id": item.SchoolID,
			"user_id":   item.UserID,

			"student_list_id": item.StudentListID,

			"status":          item.Status,
			"status_feedback": item.StatusFeedback,

			"message": item.Message,

			"gender":         item.Gender,
			"first_name":     item.FirstName,
			"last_name":      item.LastName,
			"birthday":       item.Birthday,
			"birth_location": item.BirthLocation,

			"document1": item.Document1,
			"document2": item.Document2,
			"document3": item.Document3,
			"document4": item.Document4,
			"document5": item.Document5,
		},
	).Find(result).Error
}

func (repository *Repository) UpdateParentAssignStatusByID(id int64, item *model.ParentAssign) (*model.ParentAssign, error) {
	result := &model.ParentAssign{}
	return result, repository.Db.Preload(clause.Associations).Model(&model.ParentAssign{}).Where("id = ?", item.ID).Updates(
		map[string]any{
			"status":          item.Status,
			"status_feedback": item.StatusFeedback,
		},
	).Find(result).Error
}

func (repository *Repository) DeleteByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.Parent{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteParentStudentByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.ParentStudent{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteParentAssignByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.ParentAssign{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultipleByID(list []int64, schoolID int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	var query *gorm.DB = repository.Db.Where(where)
	if schoolID > 0 {
		query = query.Where("school_id = ?", schoolID)
	}
	query = query.Delete(&model.Parent{})

	result = query.RowsAffected
	err = query.Error
	return
}

func (repository *Repository) DeleteMultipleParentStudentByID(list []int64, schoolID int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	var query *gorm.DB = repository.Db.Where(where)
	if schoolID > 0 {
		query = query.Where("school_id = ?", schoolID)
	}
	query = query.Delete(&model.ParentStudent{})

	result = query.RowsAffected
	err = query.Error
	return
}

func (repository *Repository) DeleteMultipleParentAssignByID(list []int64, schoolID int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	var query *gorm.DB = repository.Db.Where(where)
	if schoolID > 0 {
		query = query.Where("school_id = ?", schoolID)
	}
	query = query.Delete(&model.ParentAssign{})

	result = query.RowsAffected
	err = query.Error
	return
}

func (repository *Repository) GetByID(id int64) (*model.Parent, error) {
	result := &model.Parent{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByUserID(userID int64) (*model.Parent, error) {
	result := &model.Parent{}
	return result, repository.Db.Preload(clause.Associations).
		Where("user_id = ?", userID).Limit(1).Find(result).Error
}

func (repository *Repository) GetParentStudentByID(id int64) (*model.ParentStudent, error) {
	result := &model.ParentStudent{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetParentAssignByID(id int64) (*model.ParentAssign, error) {
	result := &model.ParentAssign{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByIDSchoolID(id int64, schoolID int64) (*model.Parent, error) {
	result := &model.Parent{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetParentStudentByIDSchoolID(id int64, schoolID int64) (*model.ParentStudent, error) {
	result := &model.ParentStudent{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetParentAssignByIDSchoolID(id int64, schoolID int64) (*model.ParentAssign, error) {
	result := &model.ParentAssign{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetUniqueObjectByUserID(
	item *model.Parent,
) (*model.Parent, error) {
	result := &model.Parent{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.Parent{
		SchoolID: item.SchoolID,
		UserID:   item.UserID,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjectsByUserID(
	item1 *model.Parent,
	item2 *model.Parent,
) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.UserID == item2.UserID) {
		return true
	}
	return false
}

func (repository *Repository) GetParentStudentUniqueObjectByUserID(
	item *model.ParentStudent,
) (*model.ParentStudent, error) {
	result := &model.ParentStudent{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.ParentStudent{
		SchoolID:  item.SchoolID,
		ParentID:  item.ParentID,
		StudentID: item.StudentID,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreParentStudentSameUniqueObjectsByUserID(
	item1 *model.ParentStudent,
	item2 *model.ParentStudent,
) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.ParentID == item2.ParentID &&
			item1.StudentID == item2.StudentID) {
		return true
	}
	return false
}

func (repository *Repository) GetParentAssignUniqueObjectByUserID(
	item *model.ParentAssign,
) (*model.ParentAssign, error) {
	result := &model.ParentAssign{}
	return result, nil
}

func (repository *Repository) AreParentAssignSameUniqueObjectsByUserID(
	item1 *model.ParentAssign,
	item2 *model.ParentAssign,
) bool {
	return false
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
			where = helpers.AppendWhereClause(where, "parent_students.school_id = ?")
			args = append(args, request.SchoolID)
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
			CAST(parent_students.id AS TEXT) = ? OR
			schools.name ILIKE ? OR
			schools.type ILIKE ? OR
			students.uid ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like)
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
				LEFT JOIN schools ON parent_students.school_id = schools.id
				LEFT JOIN parents ON parent_students.parent_id = parents.id
				LEFT JOIN students ON parent_students.student_id = students.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}

func (repository *Repository) GetAllParentAssign(
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllParentAssignRequest,
) (result []model.ParentAssign, err error) {
	result = make([]model.ParentAssign, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "parent_assigns.school_id = ?")
			args = append(args, request.SchoolID)
		}
		if request.ParentID > 0 {
			where = helpers.AppendWhereClause(where, "parent_assigns.parent_id = ?")
			args = append(args, request.ParentID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(parent_assigns.id AS TEXT) = ? OR
			schools.name ILIKE ? OR
			schools.type ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Preload("Parent.User").
		Preload("Parent.User.Info").
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT parent_assigns.*
				FROM parent_assigns
				LEFT JOIN schools ON parent_assigns.school_id = schools.id
				LEFT JOIN parents ON parent_assigns.parent_id = parents.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}
