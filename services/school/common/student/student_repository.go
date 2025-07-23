package student

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/services/school/common/student/data"
	"api/services/school/common/student/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.Student) (*model.Student, error) {
	result := *item
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) CreateStudentEnroll(item *model.StudentEnroll) (*model.StudentEnroll, error) {
	result := *item
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) UpdateByID(id int64, item *model.Student) (*model.Student, error) {
	foundItem, err := repository.GetByID(id)
	if err != nil || foundItem == nil || foundItem.ID != id {
		return nil, err
	}
	result := &model.Student{}
	return result, repository.Db.Model(&model.Student{}).Where("id = ?", item.ID).Updates(
		map[string]any{
			"school_id": item.SchoolID,
			"user_id":   item.UserID,
			"uid":       item.UID,
		},
	).Find(result).Error
}

func (repository *Repository) UpdateStudentEnrollByID(id int64, item *model.StudentEnroll) (*model.StudentEnroll, error) {
	foundItem, err := repository.GetByID(id)
	if err != nil || foundItem == nil || foundItem.ID != id {
		return nil, err
	}
	result := &model.StudentEnroll{}
	return result, repository.Db.Model(&model.StudentEnroll{}).Where("id = ?", item.ID).Updates(
		map[string]any{
			"school_id":       item.SchoolID,
			"year_id":         item.YearID,
			"class_id":        item.ClassID,
			"level_domain_id": item.LevelDomainID,
			"student_id":      item.StudentID,

			"email":        item.Email,
			"phone_number": item.PhoneNumber,

			"origin":          item.Origin,
			"status":          item.Status,
			"status_feedback": item.StatusFeedback,

			"message": item.Message,

			"gender":     item.Gender,
			"first_name": item.FirstName,
			"last_name":  item.LastName,
			"birthday":   item.Birthday,

			"document1": item.Document1,
			"document2": item.Document2,
			"document3": item.Document3,
			"document4": item.Document4,
			"document5": item.Document5,
		},
	).Find(result).Error
}

func (repository *Repository) DeleteByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.Student{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteStudentEnrollByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.StudentEnroll{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultipleByID(list []int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.Student{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) CountAll(schoolID int64) (result int64, err error) {
	if schoolID <= 1 {
		err = repository.Db.Model(&model.Student{}).Count(&result).Error
		return
	}
	err = repository.Db.Model(&model.Student{}).Where("school_id = ?", schoolID).Count(&result).Error
	return
}

func (repository *Repository) GetByID(id int64) (*model.Student, error) {
	result := &model.Student{}
	return result, repository.Db.
		Preload(clause.Associations).
		Preload("User.Info").
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByUserIDSchoolID(userID int64, schoolID int64) (*model.Student, error) {
	result := &model.Student{}
	return result, repository.Db.
		Preload(clause.Associations).
		Preload("User.Info").
		Where("user_id = ?", userID).Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetStudentEnrollByID(id int64) (*model.StudentEnroll, error) {
	result := &model.StudentEnroll{}
	return result, repository.Db.
		Preload(clause.Associations).
		Preload("Student.User").
		Preload("Student.User.Info").
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetUniqueObjectByUserID(item *model.Student) (*model.Student, error) {
	result := &model.Student{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.Student{
		SchoolID: item.SchoolID,
		UserID:   item.UserID,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjectsByUserID(item1 *model.Student, item2 *model.Student) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.UserID == item2.UserID) {
		return true
	}
	return false
}

func (repository *Repository) GetUniqueObjectByUID(item *model.Student) (*model.Student, error) {
	result := &model.Student{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.Student{
		SchoolID: item.SchoolID,
		UID:      item.UID,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjectsByUID(item1 *model.Student, item2 *model.Student) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.UID == item2.UID) {
		return true
	}
	return false
}

func (repository *Repository) GetStudentEnrollUniqueObject(item *model.StudentEnroll) (*model.StudentEnroll, error) {
	result := &model.StudentEnroll{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.StudentEnroll{
		StudentID:     item.StudentID,
		YearID:        item.YearID,
		LevelDomainID: item.LevelDomainID,
		ClassID:       item.ClassID,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreStudentEnrollSameUniqueObjects(item1 *model.StudentEnroll, item2 *model.StudentEnroll) bool {
	if item1 != nil && item2 != nil &&
		(item1.StudentID == item2.StudentID &&
			item1.YearID == item2.YearID && item1.LevelDomainID == item2.LevelDomainID && item1.ClassID == item2.ClassID) {
		return true
	}
	return false
}

func (repository *Repository) GetAll(
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllRequest,
) (result []model.Student, err error) {
	result = make([]model.Student, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "students.school_id = ?")
			args = append(args, request.SchoolID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(students.id AS TEXT) = ? OR 
			students.uid ILIKE ? OR 
			schools.name ILIKE ? OR
			schools.type ILIKE ? OR
			users.email ILIKE ? OR 
			CAST(users.phone_number AS TEXT) ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Preload("User.Info").
		Preload("User.Role").
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT students.* 
				FROM students 
				LEFT JOIN schools ON students.school_id = schools.id 
				LEFT JOIN users ON students.user_id = users.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}

func (repository *Repository) GetAllStudentEnroll(
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllStudentEnrollRequest,
) (result []model.StudentEnroll, err error) {
	result = make([]model.StudentEnroll, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "student_enrolls.school_id = ?")
			args = append(args, request.SchoolID)
		}
		if request.YearID > 0 {
			where = helpers.AppendWhereClause(where, "student_enrolls.year_id = ?")
			args = append(args, request.YearID)
		}
		if request.ClassID > 0 {
			where = helpers.AppendWhereClause(where, "student_enrolls.class_id = ?")
			args = append(args, request.ClassID)
		}
		if request.LevelDomainID > 0 {
			where = helpers.AppendWhereClause(where, "student_enrolls.level_domain_id = ?")
			args = append(args, request.LevelDomainID)
		}
		if request.StudentID > 0 {
			where = helpers.AppendWhereClause(where, "student_enrolls.student_id = ?")
			args = append(args, request.StudentID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(student_enrolls.id AS TEXT) = ? OR 
			schools.name ILIKE ? OR
			schools.type ILIKE ? OR
			years.name ILIKE ? OR 
			highschool_classes.name ILIKE ? OR 
			students.uid ILIKE ? 
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Preload("Class.Specialty").
		Preload("LevelDomain.Level").
		Preload("LevelDomain.Domain").
		Preload("LevelDomain.Domain.Department").
		Preload("Student.User").
		Preload("Student.User.Info").
		Preload("Student.User.Role").
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT student_enrolls.* 
				FROM student_enrolls 
				LEFT JOIN schools ON student_enrolls.school_id = schools.id 
				LEFT JOIN years ON student_enrolls.year_id = years.id 
				LEFT JOIN highschool_classes ON student_enrolls.class_id = highschool_classes.id
				LEFT JOIN students ON student_enrolls.student_id = students.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}
