package exam

import (
	"fmt"

	"gorm.io/gorm"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/common/exam/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) CreateType(data *model.ExamType) (*model.ExamType, error) {
	result := *data
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) Create(data *model.Exam) (*model.Exam, error) {
	result := *data
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) UpdateType(id int64, data *model.ExamType) (*model.ExamType, error) {
	tempExam, err := repository.GetTypeById(id, data.SchoolID)
	if err != nil || tempExam == nil || tempExam.ID != id {
		return nil, err
	}

	result := &model.ExamType{}
	return result, repository.Db.Model(result).Where("id = ?", id).Where("school_id = ?", data.SchoolID).Updates(
		map[string]any{
			"name":        data.Name,
			"description": data.Description,
		},
	).Error
}

func (repository *Repository) Update(id int64, data *model.Exam) (*model.Exam, error) {
	tempExam, err := repository.GetById(id, data.SchoolID)
	if err != nil || tempExam == nil || tempExam.ID != id {
		return nil, err
	}

	result := &model.Exam{}
	return result, repository.Db.Model(result).Where("id = ?", id).Where("school_id = ?", data.SchoolID).Updates(
		map[string]any{
			"percentage":       data.Percentage,
			"description":      data.Description,
			"type_id":          data.TypeID,
			"teaching_unit_id": data.TeachingUnitID,
			"subject_id":       data.SubjectID,
		},
	).Error
}

func (repository *Repository) DeleteType(id int64, schoolID int64) (int64, error) {
	tempExam, err := repository.GetTypeById(id, schoolID)
	if err != nil || tempExam == nil || tempExam.ID != id {
		return -1, err
	}

	result := repository.Db.Where("id = ?", id).Where("school_id = ?", schoolID).Delete(&model.ExamType{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) Delete(id int64, schoolID int64) (int64, error) {
	tempExam, err := repository.GetTypeById(id, schoolID)
	if err != nil || tempExam == nil || tempExam.ID != id {
		return -1, err
	}

	result := repository.Db.Where("id = ?", id).Where("school_id = ?", schoolID).Delete(&model.Exam{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetTypeById(id int64, schoolID int64) (*model.ExamType, error) {
	result := &model.ExamType{}
	return result, repository.Db.Model(&model.ExamType{}).Where("id = ?", id).Where("school_id = ?", schoolID).Limit(1).Find(result).Error
}

func (repository *Repository) GetById(id int64, schoolID int64) (*model.Exam, error) {
	result := &model.Exam{}
	return result, repository.Db.Model(&model.Exam{}).Where("id = ?", id).Where("school_id = ?", schoolID).Limit(1).Find(result).Error
}

func (repository *Repository) GetByObject(data *model.Exam) (*model.Exam, error) {
	result := &model.Exam{}
	return result, repository.Db.Model(&model.Exam{}).Where("id = ?", data.ID).Where("school_id = ?", data.SchoolID).Limit(1).Find(result).Error
}

func (repository *Repository) GetTypeByObject(data *model.ExamType) (*model.ExamType, error) {
	result := &model.ExamType{}
	return result, repository.Db.Model(&model.ExamType{}).Where("id = ?", data.ID).Where("school_id = ?", data.SchoolID).Limit(1).Find(result).Error
}

func (repository *Repository) GetAllType(filter *types.Filter, pagination *types.Pagination, schoolID int64) ([]model.ExamType, error) {
	var result []model.ExamType
	var where string = ""
	if filter != nil && len(filter.Search) >= 1 {
		where = fmt.Sprintf(
			"WHERE school_id = '%d' AND (type ILIKE %s OR WHERE name ILIKE %s OR WHERE description ILIKE %s)",
			schoolID,
			filter.Search,
			filter.Search,
			filter.Search,
		)
	}
	return result, repository.Db.Scopes(
		helpers.PaginationScope(
			repository.Db,
			"exam_types",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, schoolID int64) ([]model.Exam, error) {
	var result []model.Exam
	var where string = ""
	if filter != nil && len(filter.Search) >= 1 {
		where = fmt.Sprintf(
			"WHERE school_id = '%d' AND (type ILIKE %s OR WHERE description ILIKE %s)",
			schoolID,
			filter.Search,
			filter.Search,
		)
	}
	return result, repository.Db.Scopes(
		helpers.PaginationScope(
			repository.Db,
			"exams",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error
}
