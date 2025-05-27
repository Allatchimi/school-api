package teacher

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/common/teacher/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.Teacher) (*model.Teacher, error) {
	result := *item
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) CreateTeacherUnitSubject(item *model.TeacherUnitSubject) (*model.TeacherUnitSubject, error) {
	result := *item
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) Update(id int64, item *model.Teacher) (*model.Teacher, error) {
	tempTeacher, err := repository.GetByID(id)
	if err != nil || tempTeacher == nil || tempTeacher.ID != id {
		return nil, err
	}

	result := &model.Teacher{}
	return result, repository.Db.Model(result).Where("id = ?", item.ID).Updates(
		map[string]any{
			"school_id": item.SchoolID,
			"user_id":   item.UserID,
			"uid":       item.UID,
		},
	).Error
}

func (repository *Repository) UpdateTeacherUnitSubject(id int64, item *model.TeacherUnitSubject) (*model.TeacherUnitSubject, error) {
	tempTeacher, err := repository.GetByID(id)
	if err != nil || tempTeacher == nil || tempTeacher.ID != id {
		return nil, err
	}

	result := &model.TeacherUnitSubject{}
	return result, repository.Db.Model(result).Where("id = ?", item.ID).Updates(
		map[string]any{
			"teacher_id": item.TeacherID,
			"year_id":    item.YearID,

			"unit_id":          item.UnitID,
			"class_subject_id": item.ClassSubjectID,
		},
	).Error
}

func (repository *Repository) Delete(id int64) (int64, error) {
	foundItem, err := repository.GetByID(id)
	if err != nil || foundItem == nil || foundItem.ID != id {
		return -1, err
	}

	result := repository.Db.Where("id = ?", id).Delete(&model.Teacher{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteTeacherUnitSubject(id int64) (int64, error) {
	foundItem, err := repository.GetTeacherUnitSubjectByID(id)
	if err != nil || foundItem == nil || foundItem.ID != id {
		return -1, err
	}

	result := repository.Db.Where("id = ?", id).Delete(&model.TeacherUnitSubject{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetByID(id int64) (*model.Teacher, error) {
	result := &model.Teacher{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByUserID(userID int64) (*model.Teacher, error) {
	result := &model.Teacher{}
	return result, repository.Db.Preload(clause.Associations).Where("user_id = ?", userID).Limit(1).Find(result).Error
}

func (repository *Repository) GetTeacherUnitSubjectByID(id int64) (*model.TeacherUnitSubject, error) {
	result := &model.TeacherUnitSubject{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetTeacherUnitSubjectByUserIDUnitID(userID int64, unitID int64) (*model.TeacherUnitSubject, error) {
	result := &model.TeacherUnitSubject{}
	return result, repository.Db.Preload(clause.Associations).
		Where("user_id = ?", userID).Where("unit_id = ?", unitID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetUniqueObjectByUserID(item *model.Teacher) (*model.Teacher, error) {
	result := &model.Teacher{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.Teacher{
		SchoolID: item.SchoolID,
		UserID:   item.UserID,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjectsByUserID(item1 *model.Teacher, item2 *model.Teacher) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.UserID == item2.UserID) {
		return true
	}
	return false
}

func (repository *Repository) GetUniqueObjectByUID(item *model.Teacher) (*model.Teacher, error) {
	result := &model.Teacher{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.Teacher{
		SchoolID: item.SchoolID,
		UID:      item.UID,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjectsByUID(item1 *model.Teacher, item2 *model.Teacher) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.UID == item2.UID) {
		return true
	}
	return false
}

func (repository *Repository) GetUnitSubjectUniqueObject(item *model.TeacherUnitSubject) (*model.TeacherUnitSubject, error) {
	result := &model.TeacherUnitSubject{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.TeacherUnitSubject{
		TeacherID:      item.TeacherID,
		YearID:         item.YearID,
		UnitID:         item.UnitID,
		ClassSubjectID: item.ClassSubjectID,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreUnitSubjectSameUniqueObjects(item1 *model.TeacherUnitSubject, item2 *model.TeacherUnitSubject) bool {
	if item1 != nil && item2 != nil &&
		(item1.TeacherID == item2.TeacherID &&
			item1.YearID == item2.YearID && item1.UnitID == item2.UnitID && item1.ClassSubjectID == item2.ClassSubjectID) {
		return true
	}
	return false
}

func (repository *Repository) GetTeacherUnitSubjectByUserIDSchoolIDYearIDUnitID(userID int64, schoolID int64, yearID int64, classSubjectID int64) (*model.TeacherUnitSubject, error) {
	result := &model.TeacherUnitSubject{}
	where := fmt.Sprintf("WHERE parents.user_id = %d AND students.school_id = %d", userID, schoolID)
	tmpErr := repository.Db.Preload(clause.Associations).Scopes(
		helpers.PaginationScope(
			repository.Db,
			"SELECT parent_students.id, parent_students.parent_id, parent_students.student_id"+
				", parent_students.created_at, parent_students.updated_at FROM parent_students "+
				"LEFT JOIN parents ON parent_students.parent_id = parents.id "+
				"LEFT JOIN students ON parent_students.student_id = students.id",
			where,
			nil,
			nil,
		),
	).Limit(1).Find(&result).Error

	err := tmpErr
	return result, err
}

func (repository *Repository) GetTeacherUnitSubjectByUserIDSchoolIDYearIDClassSubjectID(userID int64, schoolID int64, yearID int64, classSubjectID int64) (*model.TeacherUnitSubject, error) {
	result := &model.TeacherUnitSubject{}
	where := fmt.Sprintf("WHERE parents.user_id = %d AND students.school_id = %d", userID, schoolID)
	tmpErr := repository.Db.Preload(clause.Associations).Scopes(
		helpers.PaginationScope(
			repository.Db,
			"SELECT parent_students.id, parent_students.parent_id, parent_students.student_id"+
				", parent_students.created_at, parent_students.updated_at FROM parent_students "+
				"LEFT JOIN parents ON parent_students.parent_id = parents.id "+
				"LEFT JOIN students ON parent_students.student_id = students.id",
			where,
			nil,
			nil,
		),
	).Limit(1).Find(&result).Error

	err := tmpErr
	return result, err
}

func (repository *Repository) GetTeacherUnitSubjectByUserIDClassSubjectID(userID int64, classSubjectID int64) (*model.TeacherUnitSubject, error) {
	result := &model.TeacherUnitSubject{}
	return result, repository.Db.Preload(clause.Associations).
		Where("user_id = ?", userID).Where("class_subject_id = ?", classSubjectID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, schoolID int64) (result []model.Teacher, err error) {
	result = make([]model.Teacher, 0)
	var where string = ""
	if schoolID > 0 {
		where = fmt.Sprintf("WHERE teachers.school_id = %d", schoolID)
	}
	if filter != nil && len(filter.Search) >= 1 {
		tempWhere := fmt.Sprintf(
			"CAST(teachers.id AS TEXT) = '%s' OR teachers.uid ILIKE '%s' OR schools.name ILIKE '%s' OR schools.type ILIKE '%s' OR users.email ILIKE '%s'",
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
			"SELECT teachers.id, teachers.school_id, teachers.user_id, teachers.uid"+
				", teachers.created_at, teachers.updated_at FROM teachers "+
				"LEFT JOIN schools ON teachers.school_id = schools.id "+
				"LEFT JOIN users ON teachers.user_id = users.id",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error

	err = tmpErr
	return
}

func (repository *Repository) GetAllTeacherUnitSubject(filter *types.Filter, pagination *types.Pagination, schoolID int64, teacherID int64) (result []model.TeacherUnitSubject, err error) {
	result = make([]model.TeacherUnitSubject, 0)
	var where string = ""
	if schoolID > 0 {
		where = fmt.Sprintf("WHERE teachers.school_id = %d", schoolID)
	}
	if teacherID > 0 {
		tempWhere := fmt.Sprintf("tus.teacher_id = %d", teacherID)
		if strings.HasPrefix(where, "WHERE") {
			where = fmt.Sprintf("%s AND %s", where, tempWhere)
		} else {
			where = fmt.Sprintf("WHERE %s", tempWhere)
		}
	}
	if filter != nil && len(filter.Search) >= 1 {
		tempWhere := fmt.Sprintf(
			"CAST(tus.id AS TEXT) = '%s' OR years.name ILIKE '%s'"+
				" OR university_units.name ILIKE '%s' OR university_units.description ILIKE '%s'"+
				filter.Search,
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
		)
		where = fmt.Sprintf("WHERE %s", tempWhere)
	}
	tmpErr := repository.Db.Preload(clause.Associations).
		Preload("Teacher.School").
		Preload("Teacher.User").
		Preload("ClassSubject.Class").
		Preload("ClassSubject.Subject").
		Preload("Unit.Domain").
		Preload("Unit.Level").
		Preload("Unit.Semester").
		Scopes(
			helpers.PaginationScope(
				repository.Db,
				"SELECT tus.id, tus.teacher_id, tus.year_id, tus.unit_id, tus.class_subject_id"+
					", tus.created_at, tus.updated_at FROM teacher_unit_subjects AS tus "+
					"LEFT JOIN teachers ON tus.teacher_id = teachers.id "+
					"LEFT JOIN years ON tus.year_id = years.id "+
					"LEFT JOIN university_units ON tus.unit_id = university_units.id "+
					"LEFT JOIN highschool_class_subjects ON tus.class_subject_id = highschool_class_subjects.id ",
				where,
				pagination,
				filter,
			),
		).Find(&result).Error

	err = tmpErr
	return
}
