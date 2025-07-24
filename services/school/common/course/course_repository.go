package course

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/services/school/common/course/data"
	"api/services/school/common/course/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.Course) (*model.Course, error) {
	result := *item
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) CreateCourseDocument(item *model.CourseDocument) (*model.CourseDocument, error) {
	result := *item
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) CreateCourseVideo(item *model.CourseVideo) (*model.CourseVideo, error) {
	result := *item
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) CreateCourseComment(item *model.CourseComment) (*model.CourseComment, error) {
	result := *item
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) UpdateByID(id int64, item *model.Course) (*model.Course, error) {
	result := &model.Course{}
	return result, repository.Db.Preload(clause.Associations).Model(&model.Course{}).Where("id = ?", id).Updates(
		map[string]any{
			"school_id":        item.SchoolID,
			"year_id":          item.YearID,
			"class_subject_id": item.ClassSubjectID,
			"unit_id":          item.UnitID,

			"title":       item.Title,
			"description": item.Description,
			"content":     item.Content,
		},
	).Find(result).Error
}

func (repository *Repository) UpdateCourseCommentByID(id int64, item *model.CourseComment) (*model.CourseComment, error) {
	result := &model.CourseComment{}
	return result, repository.Db.Preload(clause.Associations).Model(&model.CourseComment{}).Where("id = ?", id).Updates(
		map[string]any{
			"message":    item.Message,
			"rate":       item.Rate,
			"is_deleted": item.IsDeleted,
		},
	).Find(result).Error
}

func (repository *Repository) DeleteByID(
	id int64,
) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.Course{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteCourseDocumentByCourseID(
	id int64,
) (int64, error) {
	result := repository.Db.Where("course_id = ?", id).Delete(&model.CourseComment{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteCourseVideoByCourseID(
	id int64,
) (int64, error) {
	result := repository.Db.Where("course_id = ?", id).Delete(&model.CourseComment{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultipleByID(list []int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.Course{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) GetByID(
	id int64,
) (*model.Course, error) {
	result := &model.Course{}
	return result, repository.Db.Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetCourseCommentByID(
	id int64,
) (*model.CourseComment, error) {
	result := &model.CourseComment{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
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
		Preload("Unit.Domain").
		Preload("Unit.Level").
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
		if request.CourseID > 0 {
			where = helpers.AppendWhereClause(where, "comments.course_id = ?")
			args = append(args, request.CourseID)
		}
		if request.UserID > 0 {
			where = helpers.AppendWhereClause(where, "comments.user_id = ?")
			args = append(args, request.UserID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(comments.id AS TEXT) = ? OR
			courses.title ILIKE ? OR
			courses.description ILIKE ? OR
			schools.name ILIKE ? OR
			schools.type ILIKE ? OR
			university_units.name ILIKE ? OR
			university_units.description ILIKE ? OR
			highschool_subjects.name ILIKE ? OR
			highschool_subjects.description ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like, like, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Preload("Course.User").
		Preload("Course.User.Info").
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT comments.*
				FROM course_comments comments
				LEFT JOIN courses ON course_comments.course_id = courses.id
				LEFT JOIN schools ON courses.school_id = schools.id
				LEFT JOIN courses ON courses.year_id = courses.id
				LEFT JOIN highschool_class_subjects ON courses.class_subject_id = highschool_class_subjects.id
				LEFT JOIN university_units ON courses.unit_id = university_units.id
				LEFT JOIN highschool_subjects ON highschool_class_subjects.subject_id = highschool_subjects.id
				LEFT JOIN users ON course_comments.user_id = users.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}
