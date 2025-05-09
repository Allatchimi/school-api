package student

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/common/student/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.Student) (*model.Student, error) {
	result := *item
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) CreateLevelClass(item *model.StudentLevelClass) (*model.StudentLevelClass, error) {
	result := *item
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) Update(id int64, item *model.Student) (*model.Student, error) {
	tempStudent, err := repository.GetById(id)
	if err != nil || tempStudent == nil || tempStudent.ID != id {
		return nil, err
	}

	result := &model.Student{}
	return result, repository.Db.Model(result).Where("id = ?", item.ID).Updates(
		map[string]any{
			"school_id": item.SchoolID,
			"user_id":   item.UserID,
			"uid":       item.UID,
		},
	).Error
}

func (repository *Repository) UpdateLevelClass(id int64, item *model.StudentLevelClass) (*model.StudentLevelClass, error) {
	tempStudent, err := repository.GetById(id)
	if err != nil || tempStudent == nil || tempStudent.ID != id {
		return nil, err
	}

	result := &model.StudentLevelClass{}
	return result, repository.Db.Model(result).Where("id = ?", item.ID).Updates(
		map[string]any{
			"student_id": item.StudentID,
			"year_id":    item.YearID,

			"domain_id": item.DomainID,
			"level_id":  item.LevelID,

			"class_id": item.ClassID,
		},
	).Error
}

func (repository *Repository) Delete(id int64) (int64, error) {
	foundItem, err := repository.GetById(id)
	if err != nil || foundItem == nil || foundItem.ID != id {
		return -1, err
	}

	result := repository.Db.Where("id = ?", id).Delete(&model.Student{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteLevelClass(id int64) (int64, error) {
	foundItem, err := repository.GetLevelClassById(id)
	if err != nil || foundItem == nil || foundItem.ID != id {
		return -1, err
	}

	result := repository.Db.Where("id = ?", id).Delete(&model.StudentLevelClass{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetById(id int64) (*model.Student, error) {
	result := &model.Student{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetLevelClassById(id int64) (*model.StudentLevelClass, error) {
	result := &model.StudentLevelClass{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByObject(item *model.Student) (*model.Student, error) {
	result := &model.Student{}
	return result, repository.Db.Where(item).Limit(1).Find(result).Error
}

func (repository *Repository) GetLevelClassByObject(item *model.StudentLevelClass) (*model.StudentLevelClass, error) {
	result := &model.StudentLevelClass{}
	return result, repository.Db.Where(item).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, schoolID int64) (result []model.Student, err error) {
	result = make([]model.Student, 0)
	var where string = ""
	if schoolID > 0 {
		where = fmt.Sprintf("WHERE students.school_id = %d", schoolID)
	}
	if filter != nil && len(filter.Search) >= 1 {
		tempWhere := fmt.Sprintf(
			"CAST(students.id AS TEXT) = '%s' OR students.uid ILIKE '%s' OR schools.name ILIKE '%s' OR schools.type ILIKE '%s' OR users.email ILIKE '%s' OR users.phone_number ILIKE '%s'",
			filter.Search,
			"%"+filter.Search+"%",
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
			"SELECT students.id, students.school_id, students.user_id, students.uid"+
				", students.created_at, students.updated_at FROM students "+
				"LEFT JOIN schools ON students.school_id = schools.id "+
				"LEFT JOIN users ON students.user_id = users.id",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error

	err = tmpErr
	return
}

func (repository *Repository) GetAllLevelClass(filter *types.Filter, pagination *types.Pagination, studentID int64) (result []model.StudentLevelClass, err error) {
	result = make([]model.StudentLevelClass, 0)
	var where string = ""
	if studentID > 0 {
		where = fmt.Sprintf("WHERE student_lc.student_id = %d", studentID)
	}
	if filter != nil && len(filter.Search) >= 1 {
		tempWhere := fmt.Sprintf(
			"CAST(student_lc.id AS TEXT) = '%s' OR years.name ILIKE '%s'"+
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
			"SELECT student_lc.id, student_lc.student_id, student_lc.year_id, student_lc.level_id"+
				", student_lc.created_at, student_lc.updated_at FROM student_level_classes AS student_lc "+
				"LEFT JOIN years ON student_lc.year_id = years.id "+
				"LEFT JOIN university_domains ON student_lc.domain_id = university_domains.id "+
				"LEFT JOIN university_levels ON student_lc.level_id = university_levels.id "+
				"LEFT JOIN highschool_classes ON student_lc.class_id = highschool_classes.id ",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error

	err = tmpErr
	return
}
