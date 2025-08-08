package unit

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/university/unit/data"
	"api/services/school/university/unit/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.UniversityUnit) (result *model.UniversityUnit, err error) {
	// Create the item
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.UniversityUnit{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) UpdateByID(id int64, item *model.UniversityUnit) (result *model.UniversityUnit, err error) {
	// Update the item
	fields := map[string]any{
		"school_id":       item.SchoolID,
		"level_domain_id": item.LevelDomainID,
		"semester_id":     item.SemesterID,

		"name":         item.Name,
		"description":  item.Description,
		"credit":       item.Credit,
		"program":      item.Program,
		"requirements": item.Requirements,
		"is_valid":     item.IsValid,
		"invalid_date": item.InvalidDate,
	}
	err = repository.Db.
		Model(&model.UniversityUnit{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.UniversityUnit{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) DeleteByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.UniversityUnit{})
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
	query = query.Delete(&model.UniversityUnit{})

	result = query.RowsAffected
	err = query.Error
	return
}

func (repository *Repository) GetByID(id int64) (*model.UniversityUnit, error) {
	result := &model.UniversityUnit{}
	return result, repository.Db.
		Preload(clause.Associations).
		Preload("LevelDomain.Level").
		Preload("LevelDomain.Domain.Department").
		Preload("LevelDomain.Domain.Department.Faculty").
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByIDSchoolID(id int64, schoolID int64) (*model.UniversityUnit, error) {
	result := &model.UniversityUnit{}
	return result, repository.Db.
		Preload(clause.Associations).
		Preload("LevelDomain.Level").
		Preload("LevelDomain.Domain.Department").
		Preload("LevelDomain.Domain.Department.Faculty").
		Where("id = ?", id).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetUniqueObject(item *model.UniversityUnit) (*model.UniversityUnit, error) {
	result := &model.UniversityUnit{}
	return result, repository.Db.Where(&model.UniversityUnit{
		SchoolID:      item.SchoolID,
		LevelDomainID: item.LevelDomainID,
		SemesterID:    item.SemesterID,

		Name: item.Name,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjects(item1 *model.UniversityUnit, item2 *model.UniversityUnit) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.LevelDomainID == item2.LevelDomainID &&
			item1.SemesterID == item2.SemesterID &&
			item1.Name == item2.Name) {
		return true
	}
	return false
}

func (repository *Repository) GetAll(
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllRequest,
) (result []model.UniversityUnit, err error) {
	result = make([]model.UniversityUnit, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "units.school_id = ?")
			args = append(args, request.SchoolID)
		}
		if request.LevelDomainID > 0 {
			where = helpers.AppendWhereClause(where, "units.level_domain_id = ?")
			args = append(args, request.LevelDomainID)
		}
		if request.SemesterID > 0 {
			where = helpers.AppendWhereClause(where, "units.semester_id = ?")
			args = append(args, request.SemesterID)
		}
		if request.OnlyValid {
			where = helpers.AppendWhereClause(where, "units.is_valid = ?")
			args = append(args, true)
		}
		if request.TeacherID > 0 {
			where = helpers.AppendWhereClause(where, "teacher_class_subject_units.teacher_id = ?")
			args = append(args, request.TeacherID)
		}
		if request.StudentID > 0 {
			where = helpers.AppendWhereClause(where, "student_enrolls.student_id = ?")
			args = append(args, request.StudentID)
		}
		if request.ParentID > 0 {
			where = helpers.AppendWhereClause(where, "parent_students.parent_id = ?")
			args = append(args, request.ParentID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(units.id AS TEXT) = ? OR
			units.name ILIKE ? OR
			units.description ILIKE ? OR
			schools.name ILIKE ? OR
			schools.type ILIKE ? OR
			university_semesters.name ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Preload("LevelDomain.Level").
		Preload("LevelDomain.Domain.Department").
		Preload("LevelDomain.Domain.Department.Faculty").
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT DISTINCT units.*
				FROM university_units units
				LEFT JOIN schools ON units.school_id = schools.id
				LEFT JOIN university_level_domains ON units.level_domain_id = university_level_domains.id
				LEFT JOIN university_semesters ON units.semester_id = university_semesters.id

				LEFT JOIN teacher_class_subject_units ON units.school_id = teacher_class_subject_units.school_id
				AND (
				teacher_class_subject_units.unit_id IS NOT NULL AND units.id = teacher_class_subject_units.unit_id
				)
				LEFT JOIN student_enrolls ON units.school_id = student_enrolls.school_id
				AND (
				student_enrolls.level_domain_id IS NOT NULL AND units.level_domain_id = student_enrolls.level_domain_id
				)
				LEFT JOIN parent_students ON student_enrolls.student_id = parent_students.student_id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}
