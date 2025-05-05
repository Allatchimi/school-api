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

func (repository *Repository) Create(teachingUnit *model.UniversityTeachingUnit) (*model.UniversityTeachingUnit, error) {
	result := *teachingUnit
	return &result, repository.Db.Create(&result).Error
}

func (repository *Repository) Update(teachingUnitID int64, userID int64, teachingUnit *model.UniversityTeachingUnit) (*model.UniversityTeachingUnit, error) {
	tempTeachingUnit, err := repository.GetById(teachingUnitID, userID)
	if err != nil || tempTeachingUnit == nil || tempTeachingUnit.ID != teachingUnitID {
		return nil, err
	}

	result := &model.UniversityTeachingUnit{}
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

	result := repository.Db.Where("id = ?", teachingUnitID).Delete(&model.UniversityTeachingUnit{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) GetById(teachingUnitID int64, userID int64) (*model.UniversityTeachingUnit, error) {
	result := &model.UniversityTeachingUnit{}
	return result, repository.Db.Model(&model.UniversityTeachingUnit{}).
		Select("university_teaching_units.*").
		Joins("left join directors on university_teaching_units.school_id = directors.id").
		Where("university_teaching_units.id = ?", teachingUnitID).Where("directors.user_id = ?", userID).Limit(1).Find(result).Error
}

func (repository *Repository) GetByObject(teachingUnit *model.UniversityTeachingUnit) (*model.UniversityTeachingUnit, error) {
	result := &model.UniversityTeachingUnit{}
	return result, repository.Db.Where(teachingUnit).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, userID int64) ([]model.UniversityTeachingUnit, error) {
	var result []model.UniversityTeachingUnit
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
			"university_teaching_units",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error

	// return result, repository.Db.Model(&model.UniversityTeachingUnit{}).
	// 	Select("university_teaching_units.*").
	// 	Joins("left join directors on university_teaching_units.school_id = directors.id").
	// 	Where("directors.user_id = ?", userID).
	// 	Scopes(helpers.PaginationScope(result, pagination, filter, repository.Db)).Find(result).Error
}
