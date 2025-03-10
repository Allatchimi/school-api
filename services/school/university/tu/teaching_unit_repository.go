package tu

import (
	"fmt"

	"gorm.io/gorm"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/university/tu/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(teachingUnit *model.TeachingUnit) (*model.TeachingUnit, error) {
	result := *teachingUnit
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) AddProfessor(professor *model.TeachingUnitProfessor) (*model.TeachingUnitProfessor, error) {
	result := *professor
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) Update(teachingUnitID int64, userID int64, teachingUnit *model.TeachingUnit) (*model.TeachingUnit, error) {
	tempTeachingUnit, err := repository.GetById(teachingUnitID, userID)
	if err != nil || tempTeachingUnit == nil || tempTeachingUnit.ID != teachingUnitID {
		return nil, err
	}

	result := &model.TeachingUnit{}
	return result, repository.Db.Model(result).Where("id = ?", teachingUnitID).Updates(
		map[string]interface{}{
			"name":         teachingUnit.Name,
			"description":  teachingUnit.Description,
			"credit":       teachingUnit.Credit,
			"program":      teachingUnit.Program,
			"requirements": teachingUnit.Requirements,

			"is_valid":     teachingUnit.IsValid,
			"invalid_date": teachingUnit.InvalidDate,
		},
	).Error
}

func (repository *Repository) Delete(teachingUnitID int64, userID int64) (int64, error) {
	tempTeachingUnit, err := repository.GetById(teachingUnitID, userID)
	if err != nil || tempTeachingUnit == nil || tempTeachingUnit.ID != teachingUnitID {
		return -1, err
	}

	result := repository.Db.Where("id = ?", teachingUnitID).Delete(&model.TeachingUnit{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteProfessor(teachingUnitID int64, userID int64) (int64, error) {
	professor, err := repository.GetProfessorById(teachingUnitID, userID)
	if err != nil || professor == nil || professor.ID != teachingUnitID {
		return -1, err
	}

	result := repository.Db.Where("id = ?", teachingUnitID).Where("user_id = ?", userID).Delete(&model.TeachingUnitProfessor{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetById(teachingUnitID int64, userID int64) (*model.TeachingUnit, error) {
	result := &model.TeachingUnit{}
	return result, repository.Db.Model(&model.TeachingUnit{}).
		Select("teaching_units.*").
		Joins("left join school_directors on teaching_units.school_id = school_directors.id").
		Where("teaching_units.id = ?", teachingUnitID).Where("school_directors.user_id = ?", userID).Limit(1).Find(result).Error
}

func (repository *Repository) GetProfessorById(teachingUnitProfessorID int64, userID int64) (*model.TeachingUnitProfessor, error) {
	result := &model.TeachingUnitProfessor{}
	return result, repository.Db.Where("id = ?", teachingUnitProfessorID).Where("user_id = ?", userID).Limit(1).Find(result).Error
}

func (repository *Repository) GetByObject(teachingUnit *model.TeachingUnit) (*model.TeachingUnit, error) {
	result := &model.TeachingUnit{}
	return result, repository.Db.Where(teachingUnit).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, userID int64) ([]model.TeachingUnit, error) {
	var result []model.TeachingUnit
	var where string = ""
	if filter != nil && len(filter.Search) >= 1 {
		where = fmt.Sprintf(
			"WHERE name ILIKE %s OR WHERE description ILIKE %s OR WHERE program ILIKE %s OR WHERE requirements ILIKE %s",
			filter.Search,
			filter.Search,
			filter.Search,
			filter.Search,
		)
	}
	return result, repository.Db.Scopes(
		helpers.PaginationScope(
			repository.Db,
			"teaching_units",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error

	// return result, repository.Db.Model(&model.TeachingUnit{}).
	// 	Select("teaching_units.*").
	// 	Joins("left join school_directors on teaching_units.school_id = school_directors.id").
	// 	Where("school_directors.user_id = ?", userID).
	// 	Scopes(helpers.PaginationScope(result, pagination, filter, repository.Db)).Find(result).Error
}
