package class

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/highschool/class/data"
	"api/services/school/highschool/class/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.HighschoolClass) (result *model.HighschoolClass, err error) {
	// Create the item
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.HighschoolClass{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) CreateClassSubject(item *model.HighschoolClassSubject) (result *model.HighschoolClassSubject, err error) {
	// Create the item
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.HighschoolClassSubject{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) UpdateByID(id int64, item *model.HighschoolClass) (result *model.HighschoolClass, err error) {
	// Update the item
	fields := map[string]any{
		"school_id":    item.SchoolID,
		"specialty_id": item.SpecialtyID,

		"name":        item.Name,
		"description": item.Description,
	}
	err = repository.Db.
		Model(&model.HighschoolClass{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.HighschoolClass{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) UpdateClassSubjectByID(id int64, item *model.HighschoolClassSubject) (result *model.HighschoolClassSubject, err error) {
	// Update the item
	fields := map[string]any{
		"school_id":  item.SchoolID,
		"subject_id": item.SubjectID,
		"class_id":   item.ClassID,

		"coefficient":  item.Coefficient,
		"program":      item.Program,
		"requirements": item.Requirements,
		"is_valid":     item.IsValid,
		"invalid_date": item.InvalidDate,
	}
	err = repository.Db.
		Model(&model.HighschoolClassSubject{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.HighschoolClassSubject{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) DeleteByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.HighschoolClass{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteClassSubjectByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.HighschoolClassSubject{})
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
	query = query.Delete(&model.HighschoolClass{})

	result = query.RowsAffected
	err = query.Error
	return
}

func (repository *Repository) DeleteMultipleClassSubjectByID(list []int64, schoolID int64) (result int64, err error) {
	if len(list) < 1 {
		return
	}
	query := repository.Db.Where("id IN ?", list)
	if schoolID > 0 {
		query = query.Where("school_id = ?", schoolID)
	}
	query = query.Delete(&model.HighschoolClassSubject{})

	result = query.RowsAffected
	err = query.Error
	return
}

func (repository *Repository) GetByID(id int64) (*model.HighschoolClass, error) {
	result := &model.HighschoolClass{}
	return result, repository.Db.
		Preload(clause.Associations).
		Preload("Specialty.Section").
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetClassSubjectByID(id int64) (*model.HighschoolClassSubject, error) {
	result := &model.HighschoolClassSubject{}
	return result, repository.Db.
		Preload(clause.Associations).
		Preload("Class.Specialty").
		Preload("Class.Specialty.Section").
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByIDSchoolID(id int64, schoolID int64) (*model.HighschoolClass, error) {
	result := &model.HighschoolClass{}
	return result, repository.Db.
		Preload(clause.Associations).
		Preload("Specialty.Section").
		Where("id = ?", id).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetClassSubjectByIDSchoolID(id int64, schoolID int64) (*model.HighschoolClassSubject, error) {
	result := &model.HighschoolClassSubject{}
	return result, repository.Db.
		Preload(clause.Associations).
		Preload("Class.Specialty").
		Preload("Class.Specialty.Section").
		Where("id = ?", id).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetUniqueObject(item *model.HighschoolClass) (*model.HighschoolClass, error) {
	result := &model.HighschoolClass{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.HighschoolClass{
		SchoolID: item.SchoolID,
		Name:     item.Name,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjects(item1 *model.HighschoolClass, item2 *model.HighschoolClass) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.Name == item2.Name) {
		return true
	}
	return false
}

func (repository *Repository) GetClassSubjectUniqueObject(item *model.HighschoolClassSubject) (*model.HighschoolClassSubject, error) {
	result := &model.HighschoolClassSubject{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.HighschoolClassSubject{
		SchoolID:  item.SchoolID,
		ClassID:   item.ClassID,
		SubjectID: item.SubjectID,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameClassSubjectUniqueObjects(item1 *model.HighschoolClassSubject, item2 *model.HighschoolClassSubject) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.ClassID == item2.ClassID &&
			item1.SubjectID == item2.SubjectID) {
		return true
	}
	return false
}

func (repository *Repository) GetAll(
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllRequest,
) (result []model.HighschoolClass, err error) {
	result = make([]model.HighschoolClass, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "classes.school_id = ?")
			args = append(args, request.SchoolID)
		}
		if request.SpecialtyID > 0 {
			where = helpers.AppendWhereClause(where, "classes.specialty_id = ?")
			args = append(args, request.SpecialtyID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(classes.id AS TEXT) = ? OR
			classes.name ILIKE ? OR
			classes.description ILIKE ? OR
			schools.name ILIKE ? OR
			schools.type ILIKE ? OR
			highschool_specialties.name ILIKE ? OR
			highschool_specialties.description ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Preload("Specialty.Section").
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT classes.*
				FROM highschool_classes classes
				LEFT JOIN schools ON classes.school_id = schools.id
				LEFT JOIN highschool_specialties ON classes.specialty_id = highschool_specialties.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}

func (repository *Repository) GetAllClassSubject(
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllClassSubjectRequest,
) (result []model.HighschoolClassSubject, err error) {
	result = make([]model.HighschoolClassSubject, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "cs.school_id = ?")
			args = append(args, request.SchoolID)
		}
		if request.ClassID > 0 {
			where = helpers.AppendWhereClause(where, "cs.class_id = ?")
			args = append(args, request.ClassID)
		}
		if request.SubjectID > 0 {
			where = helpers.AppendWhereClause(where, "cs.subkect_id = ?")
			args = append(args, request.SubjectID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(cs.id AS TEXT) = ? OR
			highschool_classes.name ILIKE ? OR
			highschool_classes.description ILIKE ? OR
			highschool_subjects.name ILIKE ? OR
			highschool_subjects.description ILIKE ? OR
			schools.name ILIKE ? OR
			schools.type ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Preload("Class.Specialty").
		Preload("Class.Specialty.Section").
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT cs.*
				FROM highschool_class_subjects AS cs
				LEFT JOIN highschool_classes ON cs.class_id = highschool_classes.id
				LEFT JOIN highschool_subjects ON cs.subject_id = highschool_subjects.id
				LEFT JOIN schools ON cs.school_id = schools.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}
