package class

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
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

func (repository *Repository) Delete(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.HighschoolClass{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteClassSubject(id int64) (int64, error) {
	foundItem, err := repository.GetClassSubjectById(id)
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

func (repository *Repository) GetById(id int64) (*model.HighschoolClass, error) {
	result := &model.HighschoolClass{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetClassSubjectById(id int64) (*model.HighschoolClassSubject, error) {
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
		ClassID:   item.ClassID,
		SubjectID: item.SubjectID,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameClassSubjectUniqueObjects(item1 *model.HighschoolClassSubject, item2 *model.HighschoolClassSubject) bool {
	if item1 != nil && item2 != nil &&
		(item1.ClassID == item2.ClassID &&
			item1.SubjectID == item2.SubjectID) {
		return true
	}
	return false
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, schoolID int64) (result []model.HighschoolClass, err error) {
	result = make([]model.HighschoolClass, 0)
	var where string = ""
	if schoolID > 0 {
		where = fmt.Sprintf("WHERE classes.school_id = %d", schoolID)
	}
	if filter != nil && len(filter.Search) >= 1 {
		tempWhere := fmt.Sprintf(
			"CAST(classes.id AS TEXT) = '%s' OR classes.name ILIKE '%s' OR classes.description ILIKE '%s' OR schools.name ILIKE '%s' OR schools.type ILIKE '%s' OR specialties.name ILIKE '%s' OR specialties.description ILIKE '%s'",
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
			"SELECT classes.id, classes.name, classes.description, classes.school_id, classes.specialty_id"+
				", classes.created_at, classes.updated_at FROM highschool_classes classes "+
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

func (repository *Repository) GetAllClassSubject(filter *types.Filter, pagination *types.Pagination, schoolID int64, classID int64) (result []model.HighschoolClassSubject, err error) {
	result = make([]model.HighschoolClassSubject, 0)
	var where string = ""
	if schoolID > 0 {
		where = fmt.Sprintf("WHERE highschool_classes.school_id = %d", schoolID)
	}
	if classID > 0 {
		tempWhere := fmt.Sprintf("cs.class_id = %d", classID)
		if strings.HasPrefix(where, "WHERE") {
			where = fmt.Sprintf("%s AND %s", where, tempWhere)
		} else {
			where = fmt.Sprintf("WHERE %s", tempWhere)
		}
	}
	if filter != nil && len(filter.Search) >= 1 {
		tempWhere := fmt.Sprintf(
			"CAST(cs.id AS TEXT) = '%s' OR highschool_classes.name ILIKE '%s' OR highschool_classes.description ILIKE '%s'"+
				" OR highschool_subjects.name ILIKE '%s' OR highschool_subjects.description ILIKE '%s'",
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
			"SELECT cs.id, cs.class_id, cs.subject_id, cs.coefficient, cs.program, cs.requirements, cs.is_valid, cs.invalid_date"+
				", cs.created_at, cs.updated_at FROM highschool_class_subjects AS cs "+
				"LEFT JOIN highschool_classes ON cs.class_id = highschool_classes.id "+
				"LEFT JOIN highschool_subjects ON cs.subject_id = highschool_subjects.id ",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error

	err = tmpErr
	return
}
