package teacher

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/services/school/common/teacher/data"
	"api/services/school/common/teacher/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(
	item *model.Teacher,
) (*model.Teacher, error) {
	result := *item
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) CreateTeacherClassSubjectUnit(
	item *model.TeacherClassSubjectUnit,
) (*model.TeacherClassSubjectUnit, error) {
	result := *item
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) UpdateByID(
	id int64,
	item *model.Teacher,
) (*model.Teacher, error) {
	foundItem, err := repository.GetByID(id)
	if err != nil || foundItem == nil || foundItem.ID != id {
		return nil, err
	}
	result := &model.Teacher{}
	return result, repository.Db.Model(&model.Teacher{}).Where("id = ?", item.ID).Updates(
		map[string]any{
			"school_id": item.SchoolID,
			"user_id":   item.UserID,
			"uid":       item.UID,
		},
	).Find(result).Error
}

func (repository *Repository) UpdateTeacherClassSubjectUnitByID(
	id int64,
	item *model.TeacherClassSubjectUnit,
) (*model.TeacherClassSubjectUnit, error) {
	foundItem, err := repository.GetByID(id)
	if err != nil || foundItem == nil || foundItem.ID != id {
		return nil, err
	}
	result := &model.TeacherClassSubjectUnit{}
	return result, repository.Db.Model(&model.TeacherClassSubjectUnit{}).Where("id = ?", item.ID).Updates(
		map[string]any{
			"teacher_id":       item.TeacherID,
			"year_id":          item.YearID,
			"class_subject_id": item.ClassSubjectID,
			"unit_id":          item.UnitID,
		},
	).Find(result).Error
}

func (repository *Repository) DeleteByID(
	id int64,
) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.Teacher{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteTeacherClassSubjectUnitByID(
	id int64,
) (int64, error) {
	foundItem, err := repository.GetTeacherClassSubjectUnitByID(id)
	if err != nil || foundItem == nil || foundItem.ID != id {
		return -1, err
	}
	result := repository.Db.Where("id = ?", id).Delete(&model.TeacherClassSubjectUnit{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultipleByID(list []int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.Teacher{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) CountAll(schoolID int64) (result int64, err error) {
	if schoolID <= 1 {
		err = repository.Db.Model(&model.Teacher{}).Count(&result).Error
		return
	}
	err = repository.Db.Model(&model.Teacher{}).Where("school_id = ?", schoolID).Count(&result).Error
	return
}

func (repository *Repository) GetByID(
	id int64,
) (*model.Teacher, error) {
	result := &model.Teacher{}
	return result, repository.Db.
		Preload(clause.Associations).
		Preload("User.Info").
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByUserID(
	userID int64,
) (*model.Teacher, error) {
	result := &model.Teacher{}
	return result, repository.Db.
		Preload(clause.Associations).
		Preload("User.Info").
		Where("user_id = ?", userID).Limit(1).Find(result).Error
}

func (repository *Repository) GetTeacherClassSubjectUnitByID(
	id int64,
) (*model.TeacherClassSubjectUnit, error) {
	result := &model.TeacherClassSubjectUnit{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetUniqueObjectByUserID(
	item *model.Teacher,
) (*model.Teacher, error) {
	result := &model.Teacher{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.Teacher{
		SchoolID: item.SchoolID,
		UserID:   item.UserID,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjectsByUserID(
	item1 *model.Teacher,
	item2 *model.Teacher,
) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.UserID == item2.UserID) {
		return true
	}
	return false
}

func (repository *Repository) GetUniqueObjectByUID(
	item *model.Teacher,
) (*model.Teacher, error) {
	result := &model.Teacher{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.Teacher{
		SchoolID: item.SchoolID,
		UID:      item.UID,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjectsByUID(
	item1 *model.Teacher,
	item2 *model.Teacher,
) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.UID == item2.UID) {
		return true
	}
	return false
}

func (repository *Repository) GetTeacherClassSubjectUnitUniqueObject(
	item *model.TeacherClassSubjectUnit,
) (*model.TeacherClassSubjectUnit, error) {
	result := &model.TeacherClassSubjectUnit{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.TeacherClassSubjectUnit{
		TeacherID:      item.TeacherID,
		YearID:         item.YearID,
		UnitID:         item.UnitID,
		ClassSubjectID: item.ClassSubjectID,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreTeacherClassSubjectUnitSameUniqueObjects(
	item1 *model.TeacherClassSubjectUnit,
	item2 *model.TeacherClassSubjectUnit,
) bool {
	if item1 != nil && item2 != nil &&
		(item1.TeacherID == item2.TeacherID &&
			item1.YearID == item2.YearID && item1.UnitID == item2.UnitID && item1.ClassSubjectID == item2.ClassSubjectID) {
		return true
	}
	return false
}

func (repository *Repository) GetTeacherClassSubjectUnitByUserIDClassSubjectID(
	userID int64,
	classSubjectID int64,
) (*model.TeacherClassSubjectUnit, error) {
	result := &model.TeacherClassSubjectUnit{}
	return result, repository.Db.Preload(clause.Associations).
		Where("user_id = ?", userID).Where("class_subject_id = ?", classSubjectID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetTeacherClassSubjectUnitByUserIDUnitID(
	userID int64,
	unitID int64,
) (*model.TeacherClassSubjectUnit, error) {
	result := &model.TeacherClassSubjectUnit{}
	return result, repository.Db.Preload(clause.Associations).
		Where("user_id = ?", userID).Where("unit_id = ?", unitID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllRequest,
) (result []model.Teacher, err error) {
	result = make([]model.Teacher, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "teachers.school_id = ?")
			args = append(args, request.SchoolID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(teachers.id AS TEXT) = ? OR
			teachers.uid ILIKE ? OR
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
				`SELECT teachers.*
				FROM teachers
				LEFT JOIN schools ON teachers.school_id = schools.id
				LEFT JOIN users ON teachers.user_id = users.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}

func (repository *Repository) GetAllTeacherClassSubjectUnit(
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllTeacherClassSubjectUnitRequest,
) (result []model.TeacherClassSubjectUnit, err error) {
	result = make([]model.TeacherClassSubjectUnit, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "teachers.school_id = ?")
			args = append(args, request.SchoolID)
		}
		if request.TeacherID > 0 {
			where = helpers.AppendWhereClause(where, "tcsu.teacher_id = ?")
			args = append(args, request.TeacherID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(tcsu.id AS TEXT) = ? OR
			teachers.uid ILIKE ? OR
			years.name ILIKE ? OR
			university_units.name ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Preload("Teacher.User").
		Preload("Teacher.User.Info").
		Preload("Teacher.User.Role").
		Preload("Unit.LevelDomain").
		Preload("Unit.LevelDomain.Level").
		Preload("Unit.LevelDomain.Domain").
		Preload("Unit.LevelDomain.Domain.Department").
		Preload("Unit.Semester").
		Preload("ClassSubject.Class").
		Preload("ClassSubject.Class.Specialty").
		Preload("ClassSubject.Class.Specialty.Section").
		Preload("ClassSubject.Subject").
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT tcsu.*
				FROM teacher_class_subject_units AS tcsu
				LEFT JOIN teachers ON tcsu.teacher_id = teachers.id
				LEFT JOIN years ON tcsu.year_id = years.id
				LEFT JOIN highschool_class_subjects ON tcsu.class_subject_id = highschool_class_subjects.id
				LEFT JOIN university_units ON tcsu.unit_id = university_units.id `,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}
