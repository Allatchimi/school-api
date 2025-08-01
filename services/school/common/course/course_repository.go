package course

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/common/course/data"
	"api/services/school/common/course/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.Course) (result *model.Course, err error) {
	// Create the item
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.Course{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) CreateCourseDocument(item *model.CourseDocument) (result *model.CourseDocument, err error) {
	// Create the item
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.CourseDocument{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) CreateCourseVideo(item *model.CourseVideo) (result *model.CourseVideo, err error) {
	// Create the item
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.CourseVideo{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) CreateCourseComment(item *model.CourseComment) (result *model.CourseComment, err error) {
	// Create the item
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.CourseComment{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) UpdateByID(id int64, item *model.Course) (result *model.Course, err error) {
	// Update the item
	fields := map[string]any{
		"school_id":        item.SchoolID,
		"year_id":          item.YearID,
		"class_subject_id": nil,
		"unit_id":          nil,

		"title":       item.Title,
		"description": item.Description,
		"content":     item.Content,
	}
	if item.ClassSubjectID > 0 {
		fields["class_subject_id"] = item.ClassSubjectID
	}
	if item.UnitID > 0 {
		fields["unit_id"] = item.UnitID
	}
	err = repository.Db.
		Model(&model.Course{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.Course{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) UpdateCourseCommentByID(id int64, item *model.CourseComment) (result *model.CourseComment, err error) {
	// Update the item
	fields := map[string]any{
		"message":    item.Message,
		"rate":       item.Rate,
		"is_deleted": item.IsDeleted,
	}
	err = repository.Db.
		Model(&model.CourseComment{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.CourseComment{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) DeleteByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.Course{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteCourseCommentByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.CourseComment{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteCourseDocumentByCourseID(
	id int64,
) (int64, error) {
	result := repository.Db.Where("course_id = ?", id).Delete(&model.CourseDocument{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteCourseVideoByCourseID(
	id int64,
) (int64, error) {
	result := repository.Db.Where("course_id = ?", id).Delete(&model.CourseVideo{})
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
	query = query.Delete(&model.Course{})

	result = query.RowsAffected
	err = query.Error
	return
}

func (repository *Repository) GetByID(id int64) (*model.Course, error) {
	result := &model.Course{}
	return result, repository.Db.
		Preload(clause.Associations).
		Preload("ClassSubject.Class").
		Preload("ClassSubject.Subject").
		Preload("Unit.LevelDomain").
		Preload("Unit.LevelDomain.Level").
		Preload("Unit.LevelDomain.Domain").
		Preload("Unit.LevelDomain.Domain.Department").
		Preload("Unit.Semester").
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetCourseCommentByID(id int64) (*model.CourseComment, error) {
	result := &model.CourseComment{}
	return result, repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByIDSchoolID(id int64, schoolID int64) (*model.Course, error) {
	result := &model.Course{}
	return result, repository.Db.
		Preload(clause.Associations).
		Preload("ClassSubject.Class").
		Preload("ClassSubject.Subject").
		Preload("Unit.LevelDomain").
		Preload("Unit.LevelDomain.Level").
		Preload("Unit.LevelDomain.Domain").
		Preload("Unit.LevelDomain.Domain.Department").
		Preload("Unit.Semester").
		Where("id = ?", id).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetCourseCommentByIDSchoolID(id int64, schoolID int64) (*model.CourseComment, error) {
	result := &model.CourseComment{}
	return result, repository.Db.
		Preload(clause.Associations).
		Joins("LEFT JOIN courses ON course_comments.course_id = courses.id").
		Where("id = ?", id).
		Where("courses.school_id = ?", schoolID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllRequest,
) (result []model.Course, err error) {
	result = make([]model.Course, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "courses.school_id = ?")
			args = append(args, request.SchoolID)
		}
		if request.YearID > 0 {
			where = helpers.AppendWhereClause(where, "courses.year_id = ?")
			args = append(args, request.YearID)
		}
		if request.ClassSubjectID > 0 {
			where = helpers.AppendWhereClause(where, "courses.class_subject_id = ?")
			args = append(args, request.ClassSubjectID)
		}
		if request.UnitID > 0 {
			where = helpers.AppendWhereClause(where, "courses.unit_id = ?")
			args = append(args, request.UnitID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(courses.id AS TEXT) = ? OR
			courses.title ILIKE ? OR
			courses.description ILIKE ? OR
			schools.name ILIKE ? OR
			schools.type ILIKE ? OR
			years.name ILIKE ? OR
			highschool_subjects.name ILIKE ? OR
			highschool_subjects.description ILIKE ? OR
			university_units.name ILIKE ? OR
			university_units.description ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like, like, like, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Preload("ClassSubject.Class").
		Preload("ClassSubject.Subject").
		Preload("Unit.LevelDomain").
		Preload("Unit.LevelDomain.Level").
		Preload("Unit.LevelDomain.Domain").
		Preload("Unit.LevelDomain.Domain.Department").
		Preload("Unit.Semester").
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT courses.*
				FROM courses
				LEFT JOIN schools ON courses.school_id = schools.id
				LEFT JOIN years ON courses.year_id = years.id
				LEFT JOIN highschool_class_subjects ON courses.class_subject_id = highschool_class_subjects.id
				LEFT JOIN university_units ON courses.unit_id = university_units.id
				LEFT JOIN highschool_subjects ON highschool_class_subjects.subject_id = highschool_subjects.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}

func (repository *Repository) GetAllCourseComment(
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllCourseCommentRequest,
) (result []model.CourseComment, err error) {
	result = make([]model.CourseComment, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "courses.school_id = ?")
			args = append(args, request.SchoolID)
		}
		if request.CourseID > 0 {
			where = helpers.AppendWhereClause(where, "comments.course_id = ?")
			args = append(args, request.CourseID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(comments.id AS TEXT) = ? OR
			schools.name ILIKE ? OR
			schools.type ILIKE ? OR
			university_units.name ILIKE ? OR
			university_units.description ILIKE ? OR
			highschool_subjects.name ILIKE ? OR
			highschool_subjects.description ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Preload("User.Info").
		Preload("User.Role").
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT comments.*
				FROM course_comments comments
				LEFT JOIN courses ON comments.course_id = courses.id
				LEFT JOIN schools ON courses.school_id = schools.id
				LEFT JOIN years ON courses.year_id = years.id
				LEFT JOIN highschool_class_subjects ON courses.class_subject_id = highschool_class_subjects.id
				LEFT JOIN university_units ON courses.unit_id = university_units.id
				LEFT JOIN highschool_subjects ON highschool_class_subjects.subject_id = highschool_subjects.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}
