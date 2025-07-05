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
	return result, repository.Db.Model(result).Where("id = ?", item.ID).Updates(
		map[string]any{
			"school_id": item.SchoolID,
			"user_id":   item.UserID,
			"uid":       item.UID,
		},
	).Error
}

func (repository *Repository) UpdateStudentEnrollByID(id int64, item *model.StudentEnroll) (*model.StudentEnroll, error) {
	foundItem, err := repository.GetByID(id)
	if err != nil || foundItem == nil || foundItem.ID != id {
		return nil, err
	}
	result := &model.StudentEnroll{}
	return result, repository.Db.Model(result).Where("id = ?", item.ID).Updates(
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
	).Error
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

func (repository *Repository) GetStudentEnrollByUserIDSchoolIDYearIDClassSubjectID(
	userID int64, schoolID int64, yearID int64, classSubjectID int64,
) (*model.StudentEnroll, error) {
	result := &model.StudentEnroll{}

	where := helpers.AppendWhereClause("", fmt.Sprintf("(students.user_id = %d AND students.school_id = %d AND student_enrolls.year_id = %d AND highschool_class_subjects.id = %d)", userID, schoolID, yearID, classSubjectID))

	tmpErr := repository.Db.
		Preload(clause.Associations).
		Scopes(
			helpers.PaginationScope(
				repository.Db,
				"SELECT student_enrolls.* "+
					"FROM student_enrolls "+
					"LEFT JOIN students ON student_enrolls.student_id = students.id "+
					"LEFT JOIN highschool_class_subjects ON student_enrolls.class_id = highschool_class_subjects.class_id ",
				where,
				nil,
				nil,
			),
		).Limit(1).Find(&result).Error

	err := tmpErr
	return result, err
}

func (repository *Repository) GetStudentEnrollByUserIDSchoolIDYearIDUnitID(userID int64, schoolID int64, yearID int64, unitID int64) (*model.StudentEnroll, error) {
	result := &model.StudentEnroll{}

	where := helpers.AppendWhereClause("", fmt.Sprintf("(students.user_id = %d AND students.school_id = %d AND student_enrolls.year_id = %d AND university_units.id = %d)", userID, schoolID, yearID, unitID))

	tmpErr := repository.Db.
		Preload(clause.Associations).
		Scopes(
			helpers.PaginationScope(
				repository.Db,
				"SELECT student_enrolls.* "+
					"FROM student_enrolls "+
					"LEFT JOIN students ON student_enrolls.student_id = students.id "+
					"LEFT JOIN university_units ON student_enrolls.level_domain_id = university_units.level_domain_id ",
				where,
				nil,
				nil,
			),
		).Limit(1).Find(&result).Error

	err := tmpErr
	return result, err
}

func (repository *Repository) GetStudentEnrollByUserIDSchoolIDYearIDClassID(
	userID int64, schoolID int64, yearID int64, classID int64,
) (*model.StudentEnroll, error) {
	result := &model.StudentEnroll{}

	where := helpers.AppendWhereClause("", fmt.Sprintf("(students.user_id = %d AND students.school_id = %d AND student_enrolls.year_id = %d AND student_enrolls.class_id = %d)", userID, schoolID, yearID, classID))

	tmpErr := repository.Db.
		Preload(clause.Associations).
		Scopes(
			helpers.PaginationScope(
				repository.Db,
				"SELECT student_enrolls.* "+
					"FROM student_enrolls "+
					"LEFT JOIN students ON student_enrolls.student_id = students.id ",
				where,
				nil,
				nil,
			),
		).Limit(1).Find(&result).Error

	err := tmpErr
	return result, err
}

func (repository *Repository) GetStudentEnrollByUserIDSchoolIDYearIDLevelDomainID(userID int64, schoolID int64, yearID int64, levelDomainID int64) (*model.StudentEnroll, error) {
	result := &model.StudentEnroll{}

	where := helpers.AppendWhereClause("", fmt.Sprintf("(students.user_id = %d AND students.school_id = %d AND student_enrolls.year_id = %d AND student_enrolls.level_domain_id = %d)", userID, schoolID, yearID, levelDomainID))

	tmpErr := repository.Db.
		Preload(clause.Associations).
		Scopes(
			helpers.PaginationScope(
				repository.Db,
				"SELECT student_enrolls.* "+
					"FROM student_enrolls "+
					"LEFT JOIN students ON student_enrolls.student_id = students.id ",
				where,
				nil,
				nil,
			),
		).Limit(1).Find(&result).Error

	err := tmpErr
	return result, err
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, request *data.GetAllRequest) (result []model.Student, err error) {
	result = make([]model.Student, 0)
	var where string = ""
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("students.school_id = %d", request.SchoolID))
		}
	}
	if filter != nil && len(filter.Search) > 0 {
		tempWhere := fmt.Sprintf(
			"(CAST(students.id AS TEXT) = '%s' OR students.uid ILIKE '%s' OR schools.name ILIKE '%s' OR schools.type ILIKE '%s' OR users.email ILIKE '%s')",
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
		Preload("User.Info").
		Scopes(
			helpers.PaginationScope(
				repository.Db,
				"SELECT students.* "+
					"FROM students "+
					"LEFT JOIN schools ON students.school_id = schools.id "+
					"LEFT JOIN users ON students.user_id = users.id ",
				where,
				pagination,
				filter,
			),
		).Find(&result).Error

	err = tmpErr
	return
}

func (repository *Repository) GetAllStudentEnroll(filter *types.Filter, pagination *types.Pagination, request *data.GetAllStudentEnrollRequest) (result []model.StudentEnroll, err error) {
	result = make([]model.StudentEnroll, 0)
	var where string = ""
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("student_enrolls.school_id = %d", request.SchoolID))
		}
		if request.YearID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("student_enrolls.year_id = %d", request.YearID))
		}
		if request.ClassSubjectID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("student_enrolls.class_id = %d", request.ClassSubjectID))
		}
		if request.UnitID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("student_enrolls.level_domain_id = %d", request.UnitID))
		}
		if request.StudentID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("student_enrolls.student_id = %d", request.StudentID))
		}
	}
	if filter != nil && len(filter.Search) > 0 {
		tempWhere := fmt.Sprintf(
			"(CAST(student_enrolls.id AS TEXT) = '%s' OR years.name ILIKE '%s' OR highschool_classes.name ILIKE '%s' OR highschool_classes.description ILIKE '%s' OR university_level_domains.program ILIKE '%s' OR university_level_domains.requirements ILIKE '%s')",
			filter.Search,
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
		)
		where = helpers.AppendWhereClause(where, tempWhere)
	}
	tmpErr := repository.Db.
		Preload(clause.Associations).
		Preload("Student.School").
		Preload("Student.User").
		Preload("Student.User.Info").
		Preload("Class.Specialty").
		Preload("LevelDomain.Domain").
		Preload("LevelDomain.Level").
		Scopes(
			helpers.PaginationScope(
				repository.Db,
				"SELECT student_enrolls.* "+
					"FROM student_enrolls "+
					"LEFT JOIN students ON student_enrolls.student_id = students.id "+
					"LEFT JOIN years ON student_enrolls.year_id = years.id "+
					"LEFT JOIN highschool_classes ON student_enrolls.class_id = highschool_classes.id "+
					"LEFT JOIN university_level_domains ON student_enrolls.level_domain_id = university_level_domains.id ",
				where,
				pagination,
				filter,
			),
		).Find(&result).Error

	err = tmpErr
	return
}
