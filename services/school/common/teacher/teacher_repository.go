package teacher

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
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
) (result *model.Teacher, err error) {
	// Create the item
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.Teacher{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) CreateTeacherClassSubjectUnit(
	item *model.TeacherClassSubjectUnit,
) (result *model.TeacherClassSubjectUnit, err error) {
	// Create the item
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.TeacherClassSubjectUnit{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) UpdateByID(
	id int64,
	item *model.Teacher,
) (result *model.Teacher, err error) {
	// Update the item
	fields := map[string]any{
		"school_id": item.SchoolID,
		"user_id":   item.UserID,

		"uid": item.UID,
	}
	err = repository.Db.
		Model(&model.Teacher{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.Teacher{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) UpdateTeacherClassSubjectUnitByID(
	id int64,
	item *model.TeacherClassSubjectUnit,
) (result *model.TeacherClassSubjectUnit, err error) {
	// Update the item
	fields := map[string]any{
		"school_id":        item.SchoolID,
		"year_id":          item.YearID,
		"class_subject_id": nil,
		"unit_id":          nil,
		"teacher_id":       item.TeacherID,
	}
	if item.ClassSubjectID > 0 {
		fields["class_subject_id"] = item.ClassSubjectID
	}
	if item.UnitID > 0 {
		fields["unit_id"] = item.UnitID
	}
	err = repository.Db.
		Model(&model.TeacherClassSubjectUnit{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.TeacherClassSubjectUnit{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) DeleteByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.Teacher{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteTeacherClassSubjectUnitByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.TeacherClassSubjectUnit{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultipleByID(list []int64, schoolID int64) (result int64, err error) {
	if len(list) < 1 {
		return
	}
	query := repository.Db.Where("id IN ?", list)
	if schoolID > 0 {
		query = query.Where("school_id = ?", schoolID)
	}
	query = query.Delete(&model.Teacher{})

	result = query.RowsAffected
	err = query.Error
	return
}

func (repository *Repository) DeleteMultipleTeacherClassSubjectUnitByID(list []int64, schoolID int64) (result int64, err error) {
	if len(list) < 1 {
		return
	}
	query := repository.Db.Where("id IN ?", list)
	if schoolID > 0 {
		query = query.Where("school_id = ?", schoolID)
	}
	query = query.Delete(&model.TeacherClassSubjectUnit{})

	result = query.RowsAffected
	err = query.Error
	return
}

func (repository *Repository) GetByID(id int64) (*model.Teacher, error) {
	result := &model.Teacher{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByUserID(userID int64) (*model.Teacher, error) {
	result := &model.Teacher{}
	return result, repository.Db.Preload(clause.Associations).
		Where("user_id = ?", userID).Limit(1).Find(result).Error
}

func (repository *Repository) GetTeacherClassSubjectUnitByID(id int64) (*model.TeacherClassSubjectUnit, error) {
	result := &model.TeacherClassSubjectUnit{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByIDSchoolID(id int64, schoolID int64) (*model.Teacher, error) {
	result := &model.Teacher{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetTeacherClassSubjectUnitByIDSchoolID(id int64, schoolID int64) (*model.TeacherClassSubjectUnit, error) {
	result := &model.TeacherClassSubjectUnit{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
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
		SchoolID:       item.SchoolID,
		YearID:         item.YearID,
		ClassSubjectID: item.ClassSubjectID,
		UnitID:         item.UnitID,
		TeacherID:      item.TeacherID,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreTeacherClassSubjectUnitSameUniqueObjects(
	item1 *model.TeacherClassSubjectUnit,
	item2 *model.TeacherClassSubjectUnit,
) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.YearID == item2.YearID &&
			item1.ClassSubjectID == item2.ClassSubjectID &&
			item1.UnitID == item2.UnitID &&
			item1.TeacherID == item2.TeacherID) {
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
				`SELECT DISTINCT teachers.*
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
			where = helpers.AppendWhereClause(where, "tcsu.school_id = ?")
			args = append(args, request.SchoolID)
		}
		if request.YearID > 0 {
			where = helpers.AppendWhereClause(where, "tcsu.year_id = ?")
			args = append(args, request.YearID)
		}
		if request.ClassSubjectID > 0 {
			where = helpers.AppendWhereClause(where, "tcsu.class_subject_id = ?")
			args = append(args, request.ClassSubjectID)
		}
		if request.UnitID > 0 {
			where = helpers.AppendWhereClause(where, "tcsu.unit_id = ?")
			args = append(args, request.UnitID)
		}
		if request.ClassID > 0 {
			where = helpers.AppendWhereClause(where, "highschool_class_subjects.class_id = ?")
			args = append(args, request.ClassID)
		}
		if request.LevelDomainID > 0 {
			where = helpers.AppendWhereClause(where, "university_units.level_domain_id = ?")
			args = append(args, request.LevelDomainID)
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
			schools.name ILIKE ? OR
			schools.type ILIKE ? OR
			years.name ILIKE ? OR
			highschool_classes.name ILIKE ? OR
			highschool_classes.description ILIKE ? OR
			highschool_subjects.name ILIKE ? OR
			highschool_subjects.description ILIKE ? OR
			university_units.name ILIKE ? OR
			university_units.description ILIKE ? OR
			teachers.uid ILIKE ? OR
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like, like, like, like, like, like, like, like)
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
				`SELECT DISTINCT tcsu.*
				FROM teacher_class_subject_units AS tcsu
				LEFT JOIN schools ON tcsu.school_id = schools.id
				LEFT JOIN years ON tcsu.year_id = years.id
				LEFT JOIN highschool_class_subjects ON tcsu.class_subject_id = highschool_class_subjects.id
				LEFT JOIN university_units ON tcsu.unit_id = university_units.id
				LEFT JOIN teachers ON tcsu.teacher_id = teachers.id
				LEFT JOIN highschool_classes ON highschool_class_subjects.class_id = highschool_classes.id
				LEFT JOIN highschool_subjects ON highschool_class_subjects.subject_id = highschool_subjects.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}
