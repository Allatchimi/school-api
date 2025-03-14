package subject

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/highschool/subject/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(subject *model.Subject) (*model.Subject, error) {
	result := *subject
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) AddProfessor(professor *model.SubjectProfessor) (*model.SubjectProfessor, error) {
	result := *professor
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) Update(subjectID int64, userID int64, subject *model.Subject) (*model.Subject, error) {
	tempSubject, err := repository.GetById(subjectID, userID)
	if err != nil || tempSubject == nil || tempSubject.ID != subjectID {
		return nil, err
	}

	result := &model.Subject{}
	return result, repository.Db.Model(result).Where("id = ?", subjectID).Updates(
		map[string]interface{}{
			"name":         subject.Name,
			"description":  subject.Description,
			"coefficient":  subject.Coefficient,
			"program":      subject.Program,
			"requirements": subject.Requirements,
		},
	).Error
}

func (repository *Repository) Delete(subjectID int64, userID int64) (int64, error) {
	tempSubject, err := repository.GetById(subjectID, userID)
	if err != nil || tempSubject == nil || tempSubject.ID != subjectID {
		return -1, err
	}

	result := repository.Db.Where("id = ?", subjectID).Delete(&model.Subject{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteProfessor(subjectID int64, userID int64) (int64, error) {
	professor, err := repository.GetProfessorById(subjectID, userID)
	if err != nil || professor == nil || professor.ID != subjectID {
		return -1, err
	}

	result := repository.Db.Where("id = ?", subjectID).Where("user_id = ?", userID).Delete(&model.SubjectProfessor{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetById(subjectID int64, userID int64) (*model.Subject, error) {
	result := &model.Subject{}
	return result, repository.Db.Model(&model.Subject{}).
		Select("subjects.*").
		Joins("left join school_directors on subjects.school_id = school_directors.id").
		Where("subjects.id = ?", subjectID).Where("school_directors.user_id = ?", userID).Limit(1).Find(result).Error
}

func (repository *Repository) GetProfessorById(subjectProfessorID int64, userID int64) (*model.SubjectProfessor, error) {
	result := &model.SubjectProfessor{}
	return result, repository.Db.Where("id = ?", subjectProfessorID).Where("user_id = ?", userID).Limit(1).Find(result).Error
}

func (repository *Repository) GetByObject(subject *model.Subject) (*model.Subject, error) {
	result := &model.Subject{}
	return result, repository.Db.Where(subject).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, userID int64, schoolID int64) (result []model.Subject, err error) {
	result = make([]model.Subject, 0)
	var where string = ""
	if schoolID > 0 {
		where = fmt.Sprintf("WHERE subjects.school_id = %d", schoolID)
	}
	if filter != nil && len(filter.Search) >= 1 {
		tempWhere := fmt.Sprintf(
			"CAST(subjects.id AS TEXT) = '%s' OR subjects.name ILIKE '%s' OR subjects.description ILIKE '%s' OR schools.name ILIKE '%s' OR schools.type ILIKE '%s' OR classes.name ILIKE '%s' OR classes.description ILIKE '%s'",
			filter.Search,
			"%"+filter.Search+"%",
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
			"SELECT subjects.id, subjects.name, subjects.description, subjects.school_id, subjects.section_id"+
				", subjects.created_at, subjects.updated_at FROM highschool_specialties subjects "+
				"LEFT JOIN schools ON subjects.school_id = schools.id "+
				"LEFT JOIN highschool_sections AS classes ON subjects.section_id = classes.id",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error

	err = tmpErr
	return
}
