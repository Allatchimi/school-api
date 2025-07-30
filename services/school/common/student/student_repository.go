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

func (repository *Repository) Create(item *model.Student) (result *model.Student, err error) {
	// Create the item
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.Student{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) CreateStudentPreEnroll(item *model.StudentPreEnroll) (result *model.StudentPreEnroll, err error) {
	// Create the item
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.StudentPreEnroll{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) CreateStudentEnroll(item *model.StudentEnroll) (result *model.StudentEnroll, err error) {
	// Create the item
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.StudentEnroll{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) UpdateByID(id int64, item *model.Student) (result *model.Student, err error) {
	// Update the item
	fields := map[string]any{
		"school_id": item.SchoolID,
		"user_id":   item.UserID,

		"uid": item.UID,
	}
	err = repository.Db.
		Model(&model.Student{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.Student{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) UpdateStudentEnrollByID(id int64, item *model.StudentEnroll) (result *model.StudentEnroll, err error) {
	// Update the item
	fields := map[string]any{
		"school_id":       item.SchoolID,
		"year_id":         item.YearID,
		"class_id":        nil,
		"level_domain_id": nil,
		"student_id":      item.StudentID,

		"origin":          item.Origin,
		"origin_feedback": item.OriginFeedback,
	}
	if item.LevelDomainID > 0 {
		fields["level_domain_id"] = item.LevelDomainID
	}
	if item.ClassID > 0 {
		fields["class_id"] = item.ClassID
	}
	err = repository.Db.
		Model(&model.StudentEnroll{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.StudentEnroll{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) UpdateStudentPreEnrollByID(id int64, item *model.StudentPreEnroll) (result *model.StudentPreEnroll, err error) {
	// Update the item
	fields := map[string]any{
		"school_id": item.SchoolID,
		"year_id":   item.YearID,

		"message":    item.Message,
		"gender":     item.Gender,
		"first_name": item.FirstName,
		"last_name":  item.LastName,
		"birthday":   item.Birthday,
		"document1":  item.Document1,
		"document2":  item.Document2,
		"document3":  item.Document3,
		"document4":  item.Document4,
		"document5":  item.Document5,
	}
	err = repository.Db.
		Model(&model.StudentPreEnroll{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.StudentPreEnroll{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) UpdateStudentPreEnrollStatusByID(id int64, item *model.StudentPreEnroll) (result *model.StudentPreEnroll, err error) {
	// Update the item
	fields := map[string]any{
		"status":          item.Status,
		"status_feedback": item.StatusFeedback,
	}
	err = repository.Db.
		Model(&model.StudentPreEnroll{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.StudentPreEnroll{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) DeleteByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.Student{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteStudentEnrollByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.StudentEnroll{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteStudentPreEnrollByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.StudentPreEnroll{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultipleByID(list []int64, schoolID int64) (result int64, err error) {
	if len(list) < 1 {
		return
	}
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	var query *gorm.DB = repository.Db.Where(where)
	if schoolID > 0 {
		query = query.Where("school_id = ?", schoolID)
	}
	query = query.Delete(&model.Student{})

	result = query.RowsAffected
	err = query.Error
	return
}

func (repository *Repository) DeleteMultipleStudentEnrollByID(list []int64, schoolID int64) (result int64, err error) {
	if len(list) < 1 {
		return
	}
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	var query *gorm.DB = repository.Db.Where(where)
	if schoolID > 0 {
		query = query.Where("school_id = ?", schoolID)
	}
	query = query.Delete(&model.StudentEnroll{})

	result = query.RowsAffected
	err = query.Error
	return
}

func (repository *Repository) DeleteMultipleStudentPreEnrollByID(list []int64, schoolID int64) (result int64, err error) {
	if len(list) < 1 {
		return
	}
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	var query *gorm.DB = repository.Db.Where(where)
	if schoolID > 0 {
		query = query.Where("school_id = ?", schoolID)
	}
	query = query.Delete(&model.StudentPreEnroll{})

	result = query.RowsAffected
	err = query.Error
	return
}

func (repository *Repository) GetByID(id int64) (*model.Student, error) {
	result := &model.Student{}
	return result, repository.Db.Preload(clause.Associations).
		Preload("School.Info").
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByUserID(userID int64) (*model.Student, error) {
	result := &model.Student{}
	return result, repository.Db.Preload(clause.Associations).
		Where("user_id = ?", userID).Limit(1).Find(result).Error
}

func (repository *Repository) GetByUserIDSchoolID(userID int64, schoolID int64) (*model.Student, error) {
	result := &model.Student{}
	return result, repository.Db.Preload(clause.Associations).
		Where("user_id = ?", userID).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetStudentEnrollByID(id int64) (*model.StudentEnroll, error) {
	result := &model.StudentEnroll{}
	return result, repository.Db.Preload(clause.Associations).
		Preload("Class.Specialty").
		Preload("Class.Specialty.Section").
		Preload("LevelDomain.Level").
		Preload("LevelDomain.Domain").
		Preload("LevelDomain.Domain.Department").
		Preload("LevelDomain.Domain.Department.Faculty").
		Preload("Student.User").
		Preload("Student.User.Info").
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetStudentPreEnrollByID(id int64) (*model.StudentPreEnroll, error) {
	result := &model.StudentPreEnroll{}
	return result, repository.Db.Preload(clause.Associations).
		Preload("Class.Specialty").
		Preload("Class.Specialty.Section").
		Preload("LevelDomain.Level").
		Preload("LevelDomain.Domain").
		Preload("LevelDomain.Domain.Department").
		Preload("LevelDomain.Domain.Department.Faculty").
		Preload("User.Info").
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByIDSchoolID(id int64, schoolID int64) (*model.Student, error) {
	result := &model.Student{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetStudentEnrollByIDSchoolID(id int64, schoolID int64) (*model.StudentEnroll, error) {
	result := &model.StudentEnroll{}
	return result, repository.Db.Preload(clause.Associations).
		Preload("Class.Specialty").
		Preload("Class.Specialty.Section").
		Preload("LevelDomain.Level").
		Preload("LevelDomain.Domain").
		Preload("LevelDomain.Domain.Department").
		Preload("LevelDomain.Domain.Department.Faculty").
		Preload("Student.User").
		Preload("Student.User.Info").
		Where("id = ?", id).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetStudentPreEnrollByIDSchoolID(id int64, schoolID int64) (*model.StudentPreEnroll, error) {
	result := &model.StudentPreEnroll{}
	return result, repository.Db.Preload(clause.Associations).
		Preload("School.Info").
		Preload("Class.Specialty").
		Preload("Class.Specialty.Section").
		Preload("LevelDomain.Level").
		Preload("LevelDomain.Domain").
		Preload("LevelDomain.Domain.Department").
		Preload("LevelDomain.Domain.Department.Faculty").
		Preload("User.Info").
		Where("id = ?", id).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
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

func (repository *Repository) GetStudentPreEnrollUniqueObject(item *model.StudentPreEnroll) (*model.StudentPreEnroll, error) {
	result := &model.StudentPreEnroll{}
	return result, nil
}

func (repository *Repository) AreStudentPreEnrollSameUniqueObjects(item1 *model.StudentPreEnroll, item2 *model.StudentPreEnroll) bool {
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
			where = helpers.AppendWhereClause(where, "enrolls.school_id = ?")
			args = append(args, request.SchoolID)
		}
		if request.YearID > 0 {
			where = helpers.AppendWhereClause(where, "enrolls.year_id = ?")
			args = append(args, request.YearID)
		}
		if request.ClassID > 0 {
			where = helpers.AppendWhereClause(where, "enrolls.class_id = ?")
			args = append(args, request.ClassID)
		}
		if request.LevelDomainID > 0 {
			where = helpers.AppendWhereClause(where, "enrolls.level_domain_id = ?")
			args = append(args, request.LevelDomainID)
		}
		if request.StudentID > 0 {
			where = helpers.AppendWhereClause(where, "enrolls.student_id = ?")
			args = append(args, request.StudentID)
		}
		if request.TeacherID > 0 {
			where = helpers.AppendWhereClause(where, "enrolls.student_id = ?")
			args = append(args, request.TeacherID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(enrolls.id AS TEXT) = ? OR
			schools.name ILIKE ? OR
			schools.type ILIKE ? OR
			years.name ILIKE ? OR
			highschool_classes.name ILIKE ? OR
			highschool_classes.description ILIKE ? OR
			university_levels.name ILIKE ? OR
			university_levels.description ILIKE ? OR
			university_domains.name ILIKE ? OR
			university_domains.description ILIKE ? OR
			students.uid ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like, like, like, like, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Preload("Class.Specialty").
		Preload("Class.Specialty.Section").
		Preload("LevelDomain.Level").
		Preload("LevelDomain.Domain").
		Preload("LevelDomain.Domain.Department").
		Preload("LevelDomain.Domain.Department.Faculty").
		Preload("Student.User").
		Preload("Student.User.Info").
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT enrolls.*
				FROM student_enrolls enrolls
				LEFT JOIN schools ON enrolls.school_id = schools.id
				LEFT JOIN years ON enrolls.year_id = years.id
				LEFT JOIN highschool_classes ON enrolls.class_id = highschool_classes.id
				LEFT JOIN university_level_domains ON enrolls.level_domain_id = university_level_domains.id
				LEFT JOIN students ON enrolls.student_id = students.id
				LEFT JOIN university_levels ON university_level_domains.level_id = university_levels.id
				LEFT JOIN university_domains ON university_level_domains.domain_id = university_domains.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}

func (repository *Repository) GetAllStudentPreEnroll(
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllStudentPreEnrollRequest,
) (result []model.StudentPreEnroll, err error) {
	result = make([]model.StudentPreEnroll, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "pre_enrolls.school_id = ?")
			args = append(args, request.SchoolID)
		}
		if request.YearID > 0 {
			where = helpers.AppendWhereClause(where, "pre_enrolls.year_id = ?")
			args = append(args, request.YearID)
		}
		if request.ClassID > 0 {
			where = helpers.AppendWhereClause(where, "pre_enrolls.class_id = ?")
			args = append(args, request.ClassID)
		}
		if request.LevelDomainID > 0 {
			where = helpers.AppendWhereClause(where, "pre_enrolls.level_domain_id = ?")
			args = append(args, request.LevelDomainID)
		}
		if request.UserID > 0 {
			where = helpers.AppendWhereClause(where, "pre_enrolls.user_id = ?")
			args = append(args, request.UserID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(pre_enrolls.id AS TEXT) = ? OR
			CAST(pre_enrolls.user_id AS TEXT) = ? OR
			pre_enrolls.gender ILIKE ? OR
			pre_enrolls.first_name ILIKE ? OR
			pre_enrolls.last_name ILIKE ? OR
			pre_enrolls.status ILIKE ? OR
			schools.name ILIKE ? OR
			schools.type ILIKE ? OR
			years.name ILIKE ? OR
			highschool_classes.name ILIKE ? OR
			highschool_classes.description ILIKE ? OR
			university_levels.name ILIKE ? OR
			university_levels.description ILIKE ? OR
			university_domains.name ILIKE ? OR
			users.email ILIKE ? OR
			CAST(users.phone_number AS TEXT) ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, search, like, like, like, like, like, like, like, like, like, like, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Preload("Class.Specialty").
		Preload("Class.Specialty.Section").
		Preload("LevelDomain.Level").
		Preload("LevelDomain.Domain").
		Preload("LevelDomain.Domain.Department").
		Preload("LevelDomain.Domain.Department.Faculty").
		Preload("User").
		Preload("User.Info").
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT pre_enrolls.*
				FROM student_pre_enrolls pre_enrolls
				LEFT JOIN schools ON pre_enrolls.school_id = schools.id
				LEFT JOIN years ON pre_enrolls.year_id = years.id
				LEFT JOIN highschool_classes ON pre_enrolls.class_id = highschool_classes.id
				LEFT JOIN university_level_domains ON pre_enrolls.level_domain_id = university_level_domains.id
				LEFT JOIN users ON pre_enrolls.user_id = users.id
				LEFT JOIN university_levels ON university_level_domains.level_id = university_levels.id
				LEFT JOIN university_domains ON university_level_domains.domain_id = university_domains.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}
