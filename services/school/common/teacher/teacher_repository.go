package teacher

import (
	"fmt"

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

func (repository *Repository) CreateTeacherClassSubjectUnit(item *model.TeacherClassSubjectUnit) (*model.TeacherClassSubjectUnit, error) {
	result := *item
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) UpdateByID(id int64, item *model.Teacher) (*model.Teacher, error) {
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

func (repository *Repository) UpdateTeacherClassSubjectUnitByID(id int64, item *model.TeacherClassSubjectUnit) (*model.TeacherClassSubjectUnit, error) {
	tempTeacher, err := repository.GetByID(id)
	if err != nil || tempTeacher == nil || tempTeacher.ID != id {
		return nil, err
	}
	result := &model.TeacherClassSubjectUnit{}
	return result, repository.Db.Model(result).Where("id = ?", item.ID).Updates(
		map[string]any{
			"teacher_id":       item.TeacherID,
			"year_id":          item.YearID,
			"class_subject_id": item.ClassSubjectID,
			"unit_id":          item.UnitID,
		},
	).Error
}

func (repository *Repository) DeleteByID(id int64) (int64, error) {
	foundItem, err := repository.GetByID(id)
	if err != nil || foundItem == nil || foundItem.ID != id {
		return -1, err
	}
	result := repository.Db.Where("id = ?", id).Delete(&model.Teacher{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteTeacherClassSubjectUnitByID(id int64) (int64, error) {
	foundItem, err := repository.GetTeacherClassSubjectUnitByID(id)
	if err != nil || foundItem == nil || foundItem.ID != id {
		return -1, err
	}
	result := repository.Db.Where("id = ?", id).Delete(&model.TeacherClassSubjectUnit{})
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

func (repository *Repository) GetTeacherClassSubjectUnitByID(id int64) (*model.TeacherClassSubjectUnit, error) {
	result := &model.TeacherClassSubjectUnit{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
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

func (repository *Repository) GetTeacherClassSubjectUnitUniqueObject(item *model.TeacherClassSubjectUnit) (*model.TeacherClassSubjectUnit, error) {
	result := &model.TeacherClassSubjectUnit{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.TeacherClassSubjectUnit{
		TeacherID:      item.TeacherID,
		YearID:         item.YearID,
		UnitID:         item.UnitID,
		ClassSubjectID: item.ClassSubjectID,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreTeacherClassSubjectUnitSameUniqueObjects(item1 *model.TeacherClassSubjectUnit, item2 *model.TeacherClassSubjectUnit) bool {
	if item1 != nil && item2 != nil &&
		(item1.TeacherID == item2.TeacherID &&
			item1.YearID == item2.YearID && item1.UnitID == item2.UnitID && item1.ClassSubjectID == item2.ClassSubjectID) {
		return true
	}
	return false
}

func (repository *Repository) GetByUserIDSchoolID(userID int64, schoolID int64) (*model.Teacher, error) {
	result := &model.Teacher{}
	return result, repository.Db.Preload(clause.Associations).
		Where("user_id = ?", userID).Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetTeacherClassSubjectUnitByUserIDClassSubjectID(userID int64, classSubjectID int64) (*model.TeacherClassSubjectUnit, error) {
	result := &model.TeacherClassSubjectUnit{}
	return result, repository.Db.Preload(clause.Associations).
		Where("user_id = ?", userID).Where("class_subject_id = ?", classSubjectID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetTeacherClassSubjectUnitByUserIDUnitID(userID int64, unitID int64) (*model.TeacherClassSubjectUnit, error) {
	result := &model.TeacherClassSubjectUnit{}
	return result, repository.Db.Preload(clause.Associations).
		Where("user_id = ?", userID).Where("unit_id = ?", unitID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetTeacherClassSubjectUnitByUserIDSchoolIDYearIDClassSubjectID(userID int64, schoolID int64, yearID int64, classSubjectID int64) (*model.TeacherClassSubjectUnit, error) {
	result := &model.TeacherClassSubjectUnit{}
	where := fmt.Sprintf("WHERE teachers.user_id = %d AND teachers.school_id = %d AND teacher_unit_subjects.year_id = %d AND teacher_unit_subjects.class_subject_id = %d", userID, schoolID, yearID, classSubjectID)
	tmpErr := repository.Db.Preload(clause.Associations).
		Scopes(
			helpers.PaginationScope(
				repository.Db,
				"SELECT teacher_unit_subjects.id, teacher_unit_subjects.teacher_id, teacher_unit_subjects.year_id, teacher_unit_subjects.class_subject_id, teacher_unit_subjects.unit_id"+
					", teacher_unit_subjects.created_at, teacher_unit_subjects.updated_at FROM teacher_unit_subjects "+
					"LEFT JOIN teachers ON teacher_unit_subjects.teacher_id = teachers.id ",
				where,
				nil,
				nil,
			),
		).Limit(1).Find(&result).Error

	err := tmpErr
	return result, err
}

func (repository *Repository) GetTeacherClassSubjectUnitByUserIDSchoolIDYearIDUnitID(userID int64, schoolID int64, yearID int64, unitID int64) (*model.TeacherClassSubjectUnit, error) {
	result := &model.TeacherClassSubjectUnit{}
	where := fmt.Sprintf("WHERE teachers.user_id = %d AND teachers.school_id = %d AND teacher_unit_subjects.year_id = %d AND teacher_unit_subjects.unit_id = %d", userID, schoolID, yearID, unitID)
	tmpErr := repository.Db.Preload(clause.Associations).
		Scopes(
			helpers.PaginationScope(
				repository.Db,
				"SELECT teacher_unit_subjects.id, teacher_unit_subjects.teacher_id, teacher_unit_subjects.year_id, teacher_unit_subjects.class_subject_id, teacher_unit_subjects.unit_id"+
					", teacher_unit_subjects.created_at, teacher_unit_subjects.updated_at FROM teacher_unit_subjects "+
					"LEFT JOIN teachers ON teacher_unit_subjects.teacher_id = teachers.id ",
				where,
				nil,
				nil,
			),
		).Limit(1).Find(&result).Error

	err := tmpErr
	return result, err
}

func (repository *Repository) GetTeacherClassSubjectUnitByUserIDSchoolIDYearIDClassID(userID int64, schoolID int64, yearID int64, classID int64) (*model.TeacherClassSubjectUnit, error) {
	result := &model.TeacherClassSubjectUnit{}
	where := fmt.Sprintf("WHERE teachers.user_id = %d AND teachers.school_id = %d AND teacher_unit_subjects.year_id = %d AND highschool_class_subjects.class_id = %d", userID, schoolID, yearID, classID)
	tmpErr := repository.Db.Preload(clause.Associations).
		Preload("Unit.LevelDomain").
		Scopes(
			helpers.PaginationScope(
				repository.Db,
				"SELECT teacher_unit_subjects.id, teacher_unit_subjects.teacher_id, teacher_unit_subjects.year_id, teacher_unit_subjects.class_subject_id, teacher_unit_subjects.unit_id"+
					", teacher_unit_subjects.created_at, teacher_unit_subjects.updated_at FROM teacher_unit_subjects "+
					"LEFT JOIN teachers ON teacher_unit_subjects.teacher_id = teachers.id "+
					"LEFT JOIN highschool_class_subjects ON teacher_unit_subjects.class_subject_id = highschool_class_subjects.id",
				where,
				nil,
				nil,
			),
		).Limit(1).Find(&result).Error

	err := tmpErr
	return result, err
}

func (repository *Repository) GetTeacherClassSubjectUnitByUserIDSchoolIDYearIDLevelDomainID(userID int64, schoolID int64, yearID int64, levelDomainID int64) (*model.TeacherClassSubjectUnit, error) {
	result := &model.TeacherClassSubjectUnit{}
	where := fmt.Sprintf("WHERE teachers.user_id = %d AND teachers.school_id = %d AND teacher_unit_subjects.year_id = %d AND university_units.level_domain_id = %d", userID, schoolID, yearID, levelDomainID)
	tmpErr := repository.Db.Preload(clause.Associations).
		Preload("Unit.LevelDomain").
		Scopes(
			helpers.PaginationScope(
				repository.Db,
				"SELECT teacher_unit_subjects.id, teacher_unit_subjects.teacher_id, teacher_unit_subjects.year_id, teacher_unit_subjects.class_subject_id, teacher_unit_subjects.unit_id"+
					", teacher_unit_subjects.created_at, teacher_unit_subjects.updated_at FROM teacher_unit_subjects "+
					"LEFT JOIN teachers ON teacher_unit_subjects.teacher_id = teachers.id "+
					"LEFT JOIN university_units ON teacher_unit_subjects.unit_id = university_units.id",
				where,
				nil,
				nil,
			),
		).Limit(1).Find(&result).Error

	err := tmpErr
	return result, err
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, schoolID int64) (result []model.Teacher, err error) {
	result = make([]model.Teacher, 0)
	var where string = ""
	if schoolID > 0 {
		where = helpers.AppendWhereClause(where, fmt.Sprintf("WHERE teachers.school_id = %d", schoolID))
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
		where = helpers.AppendWhereClause(where, tempWhere)
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

func (repository *Repository) GetAllTeacherClassSubjectUnit(filter *types.Filter, pagination *types.Pagination, schoolID int64, teacherID int64) (result []model.TeacherClassSubjectUnit, err error) {
	result = make([]model.TeacherClassSubjectUnit, 0)
	var where string = ""
	if schoolID > 0 {
		where = helpers.AppendWhereClause(where, fmt.Sprintf("WHERE teachers.school_id = %d", schoolID))
	}
	if teacherID > 0 {
		where = helpers.AppendWhereClause(where, fmt.Sprintf("tcsu.teacher_id = %d", teacherID))
	}
	if filter != nil && len(filter.Search) >= 1 {
		tempWhere := fmt.Sprintf(
			"CAST(tcsu.id AS TEXT) = '%s' OR years.name ILIKE '%s'"+
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
				"SELECT tcsu.id, tcsu.teacher_id, tcsu.year_id, tcsu.class_subject_id, tcsu.unit_id"+
					", tcsu.created_at, tcsu.updated_at FROM teacher_class_subject_units AS tcsu "+
					"LEFT JOIN teachers ON tcsu.teacher_id = teachers.id "+
					"LEFT JOIN years ON tcsu.year_id = years.id "+
					"LEFT JOIN highschool_class_subjects ON tcsu.class_subject_id = highschool_class_subjects.id "+
					"LEFT JOIN university_units ON tcsu.unit_id = university_units.id",
				where,
				pagination,
				filter,
			),
		).Find(&result).Error

	err = tmpErr
	return
}
