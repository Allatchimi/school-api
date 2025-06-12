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
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"school_id":        item.SchoolID,
			"year_id":          item.YearID,
			"class_subject_id": item.ClassSubjectID,
			"unit_id":          item.UnitID,

			"title":       item.Title,
			"description": item.Description,
			"content":     item.Content,
		},
	).Error
}

func (repository *Repository) UpdateCourseCommentByID(id int64, item *model.CourseComment) (*model.CourseComment, error) {
	result := &model.CourseComment{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"message":    item.Message,
			"rate":       item.Rate,
			"is_deleted": item.IsDeleted,
		},
	).Error
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
) ([]model.Course, error) {
	var result []model.Course
	var where string = ""
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("courses.school_id = %d", request.SchoolID))
		}
		if request.YearID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("courses.year_id = %d", request.YearID))
		}
		if request.ClassSubjectID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("courses.class_subject_id = %d", request.ClassSubjectID))
		}
		if request.UnitID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("courses.unit_id = %d", request.UnitID))
		}
	}
	if filter != nil && len(filter.Search) >= 1 {
		tempWhere := fmt.Sprintf(
			"(title ILIKE %s OR WHERE description ILIKE %s)",
			filter.Search,
			filter.Search,
		)
		where = helpers.AppendWhereClause(where, tempWhere)
	}
	return result, repository.Db.Scopes(
		helpers.PaginationScope(
			repository.Db,
			"courses",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error
}

func (repository *Repository) GetAllCourseComment(
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllCourseCommentRequest,
) ([]model.CourseComment, error) {
	var result []model.CourseComment
	var where string = ""
	if request != nil {
		if request.CourseID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("courses.course_id = %d", request.CourseID))
		}
	}
	if filter != nil && len(filter.Search) >= 1 {
		tempWhere := fmt.Sprintf(
			"(message ILIKE %s)",
			filter.Search,
		)
		where = helpers.AppendWhereClause(where, tempWhere)
	}
	return result, repository.Db.Scopes(
		helpers.PaginationScope(
			repository.Db,
			"course_comments",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error
}
