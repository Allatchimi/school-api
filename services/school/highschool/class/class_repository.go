package class

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/services/school/highschool/class/data"
	"api/services/school/highschool/class/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.HighschoolClass) (*model.HighschoolClass, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) CreateClassSubject(item *model.HighschoolClassSubject) (*model.HighschoolClassSubject, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) Update(id int64, item *model.HighschoolClass) (*model.HighschoolClass, error) {
	result := &model.HighschoolClass{}
	fmt.Println(item)
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"school_id":    item.SchoolID,
			"specialty_id": item.SpecialtyID,
			"name":         item.Name,
			"description":  item.Description,
		},
	).Error
}

func (repository *Repository) UpdateClassSubject(id int64, item *model.HighschoolClassSubject) (*model.HighschoolClassSubject, error) {
	result := &model.HighschoolClassSubject{}
	fmt.Println(item)
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"school_id":  item.SchoolID,
			"subject_id": item.SubjectID,
			"class_id":   item.ClassID,

			"coefficient":  item.Coefficient,
			"program":      item.Program,
			"requirements": item.Requirements,
			"is_valid":     item.IsValid,
			"invalid_date": item.InvalidDate,
		},
	).Error
}

func (repository *Repository) Delete(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.HighschoolClass{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteClassSubject(id int64) (int64, error) {
	foundItem, err := repository.GetClassSubjectByID(id)
	if err != nil || foundItem == nil || foundItem.ID != id {
		return -1, err
	}

	result := repository.Db.Where("id = ?", id).Delete(&model.HighschoolClassSubject{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultiple(list []int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.HighschoolClass{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) DeleteMultipleClassSubject(list []int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.HighschoolClassSubject{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) GetByID(id int64) (*model.HighschoolClass, error) {
	result := &model.HighschoolClass{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetClassSubjectByID(id int64) (*model.HighschoolClassSubject, error) {
	result := &model.HighschoolClassSubject{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
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

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, request *data.GetAllRequest) (result []model.HighschoolClass, err error) {
	result = make([]model.HighschoolClass, 0)
	var where string = ""
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("classes.school_id = %d", request.SchoolID))
		}
		if request.SpecialtyID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("classes.specialty_id = %d", request.SpecialtyID))
		}
	}
	if filter != nil && len(filter.Search) > 0 {
		tempWhere := fmt.Sprintf(
			"(CAST(classes.id AS TEXT) = '%s' OR classes.name ILIKE '%s' OR classes.description ILIKE '%s' OR schools.name ILIKE '%s' OR schools.type ILIKE '%s' OR specialties.name ILIKE '%s' OR specialties.description ILIKE '%s')",
			filter.Search,
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
		)
		where = helpers.AppendWhereClause(where, tempWhere)
	}
	tmpErr := repository.Db.Preload(clause.Associations).Scopes(
		helpers.PaginationScope(
			repository.Db,
			"SELECT classes.* "+
				"FROM highschool_classes classes "+
				"LEFT JOIN schools ON classes.school_id = schools.id "+
				"LEFT JOIN highschool_specialties AS specialties ON classes.specialty_id = specialties.id",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error

	err = tmpErr
	return
}

func (repository *Repository) GetAllClassSubject(filter *types.Filter, pagination *types.Pagination, request *data.GetAllClassSubjectRequest) (result []model.HighschoolClassSubject, err error) {
	result = make([]model.HighschoolClassSubject, 0)
	var where string = ""
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("cs.school_id = %d", request.SchoolID))
		}
		if request.ClassID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("cs.class_id = %d", request.ClassID))
		}
		if request.SubjectID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("cs.subject_id = %d", request.SubjectID))
		}
	}
	if filter != nil && len(filter.Search) > 0 {
		tempWhere := fmt.Sprintf(
			"(CAST(cs.id AS TEXT) = '%s' OR highschool_classes.name ILIKE '%s' OR highschool_classes.description ILIKE '%s' OR highschool_subjects.name ILIKE '%s' OR highschool_subjects.description ILIKE '%s' OR schools.name ILIKE '%s' OR schools.type ILIKE '%s')",
			filter.Search,
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
		)
		where = helpers.AppendWhereClause(where, tempWhere)
	}
	tmpErr := repository.Db.Preload(clause.Associations).Scopes(
		helpers.PaginationScope(
			repository.Db,
			"SELECT cs.* "+
				"FROM highschool_class_subjects AS cs "+
				"LEFT JOIN highschool_classes ON cs.class_id = highschool_classes.id "+
				"LEFT JOIN highschool_subjects ON cs.subject_id = highschool_subjects.id "+
				"LEFT JOIN schools ON cs.school_id = schools.id ",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error

	err = tmpErr
	return
}
