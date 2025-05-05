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

func (repository *Repository) CreateTUSubject(item *model.TUSubject) (*model.TUSubject, error) {
	result := *item
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) Update(id int64, item *model.Teacher) (*model.Teacher, error) {
	tempTeacher, err := repository.GetById(id)
	if err != nil || tempTeacher == nil || tempTeacher.ID != id {
		return nil, err
	}

	result := &model.Teacher{}
	return result, repository.Db.Model(result).Where("id = ?", item.ID).Updates(
		map[string]interface{}{
			"school_id": item.SchoolID,
			"user_id":   item.UserID,
			"uid":       item.UID,
		},
	).Error
}

func (repository *Repository) UpdateTUSubject(id int64, item *model.TUSubject) (*model.TUSubject, error) {
	tempTeacher, err := repository.GetById(id)
	if err != nil || tempTeacher == nil || tempTeacher.ID != id {
		return nil, err
	}

	result := &model.TUSubject{}
	return result, repository.Db.Model(result).Where("id = ?", item.ID).Updates(
		map[string]interface{}{
			"teacher_id": item.TeacherID,
			"year_id":    item.YearID,

			"teaching_unit_id": item.TeachingUnitID,
			"subject_id":       item.SubjectID,
		},
	).Error
}

func (repository *Repository) Delete(id int64) (int64, error) {
	foundItem, err := repository.GetById(id)
	if err != nil || foundItem == nil || foundItem.ID != id {
		return -1, err
	}

	result := repository.Db.Where("id = ?", id).Delete(&model.Teacher{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteTUSubject(id int64) (int64, error) {
	foundItem, err := repository.GetTUSubjectById(id)
	if err != nil || foundItem == nil || foundItem.ID != id {
		return -1, err
	}

	result := repository.Db.Where("id = ?", id).Delete(&model.TUSubject{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetById(id int64) (*model.Teacher, error) {
	result := &model.Teacher{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetTUSubjectById(id int64) (*model.TUSubject, error) {
	result := &model.TUSubject{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByObject(item *model.Teacher) (*model.Teacher, error) {
	result := &model.Teacher{}
	return result, repository.Db.Where(item).Limit(1).Find(result).Error
}

func (repository *Repository) GetTUSubjectByObject(item *model.TUSubject) (*model.TUSubject, error) {
	result := &model.TUSubject{}
	return result, repository.Db.Where(item).Limit(1).Find(result).Error
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

func (repository *Repository) GetAllTUSubject(filter *types.Filter, pagination *types.Pagination, teacherID int64) (result []model.TUSubject, err error) {
	result = make([]model.TUSubject, 0)
	var where string = ""
	if teacherID > 0 {
		where = fmt.Sprintf("WHERE teacher_lc.teacher_id = %d", teacherID)
	}
	if filter != nil && len(filter.Search) >= 1 {
		tempWhere := fmt.Sprintf(
			"CAST(teacher_lc.id AS TEXT) = '%s' OR years.name ILIKE '%s'"+
				" OR university_domains.name ILIKE '%s' OR university_domains.description ILIKE '%s'"+
				" OR university_levels.name ILIKE '%s' OR university_levels.description ILIKE '%s'"+
				" OR highschool_classes.name ILIKE '%s' OR highschool_classes.description ILIKE '%s'",
			filter.Search,
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
		)
		where = fmt.Sprintf("WHERE %s", tempWhere)
	}
	tmpErr := repository.Db.Preload(clause.Associations).Scopes(
		helpers.PaginationScope(
			repository.Db,
			"SELECT teacher_lc.id, teacher_lc.teacher_id, teacher_lc.year_id, teacher_lc.domain_id, teacher_lc.level_id, teacher_lc.class_id"+
				", teacher_lc.created_at, teacher_lc.updated_at FROM teacher_level_classes AS teacher_lc "+
				"LEFT JOIN years ON teacher_lc.year_id = years.id "+
				"LEFT JOIN university_domains ON teacher_lc.domain_id = university_domains.id "+
				"LEFT JOIN university_levels ON teacher_lc.level_id = university_levels.id "+
				"LEFT JOIN highschool_classes ON teacher_lc.class_id = highschool_classes.id ",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error

	err = tmpErr
	return
}
