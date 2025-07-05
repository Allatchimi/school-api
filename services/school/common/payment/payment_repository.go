package payment

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/services/school/common/payment/data"
	"api/services/school/common/payment/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.Payment) (*model.Payment, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) Update(id int64, item *model.Payment) (*model.Payment, error) {
	result := &model.Payment{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"school_id":         item.SchoolID,
			"student_enroll_id": item.StudentEnrollID,

			"amount":         item.Amount,
			"currency":       item.Currency,
			"payment_date":   item.PaymentDate,
			"payment_method": item.PaymentMethod,
			"payment_status": item.PaymentStatus,
			"payment_note":   item.PaymentNote,
		},
	).Error
}

func (repository *Repository) Delete(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.Payment{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultiple(list []int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.Payment{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) GetByID(id int64) (*model.Payment, error) {
	result := &model.Payment{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllRequest,
) (result []model.Payment, err error) {
	result = make([]model.Payment, 0)
	var where string = ""
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("payments.school_id = %d", request.SchoolID))
		}
		if request.StudentEnrollID > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("payments.student_enroll_id = %d", request.StudentEnrollID))
		}
	}
	if filter != nil && len(filter.Search) > 0 {
		tempWhere := fmt.Sprintf(
			"CAST(payments.id AS TEXT) = '%s' OR payments.currency ILIKE '%s' OR payments.payment_method ILIKE '%s' OR payments.payment_status ILIKE '%s' OR payments.payment_note ILIKE '%s' OR schools.name ILIKE '%s'",
			filter.Search,
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
		)
		where = helpers.AppendWhereClause(where, tempWhere)
	}
	tmpErr := repository.Db.
		Preload(clause.Associations).
		Preload("School").
		Preload("StudentEnroll.Student").
		Preload("StudentEnroll.Student.User").
		Preload("StudentEnroll.Student.User.Info").
		Preload("StudentEnroll.Year").
		Preload("StudentEnroll.Class").
		Preload("StudentEnroll.LevelDomain").
		Preload("StudentEnroll.LevelDomain.Level").
		Preload("StudentEnroll.LevelDomain.Domain").
		Scopes(
			helpers.PaginationScope(
				repository.Db,
				"SELECT payments.* "+
					"FROM payments "+
					"LEFT JOIN schools ON payments.school_id = schools.id "+
					"LEFT JOIN student_enrolls ON payments.student_enroll_id = student_enrolls.id ",
				where,
				pagination,
				filter,
			),
		).Find(&result).Error

	err = tmpErr
	return
}
