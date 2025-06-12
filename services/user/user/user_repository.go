package user

import (
	"api/common/constants"
	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/services/user/user/data"
	"api/services/user/user/model"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.User) (*model.User, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) CreateUserInfo(item *model.UserInfo) (*model.UserInfo, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) CreateUserMfa(item *model.UserMfa) (*model.UserMfa, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) UpdateByID(id int64, item *model.User) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"email":        item.Email,
			"phone_number": item.PhoneNumber,
			"role_id":      item.RoleID,
			"is_activated": item.IsActivated,
		},
	).Error
}

func (repository *Repository) UpdateUserInfoByID(id int64, item *model.UserInfo) (*model.UserInfo, error) {
	result := &model.UserInfo{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"username":   item.Username,
			"first_name": item.FirstName,
			"last_name":  item.LastName,
			"address":    item.Address,
			"image":      item.Image,
			"language":   item.Language,
		},
	).Error
}

func (repository *Repository) DeleteByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.User{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultipleByID(list []int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.User{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) GetByID(id int64) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByEmail(email string) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Preload(clause.Associations).
		Where(
			"login_method = ?", constants.AuthLoginMethodDefault,
		).Where(
		"email = ?", email,
	).Limit(1).Find(result).Error
}

func (repository *Repository) GetByPhoneNumber(phoneNumber uint64) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Preload(clause.Associations).
		Where(
			"login_method = ?", constants.AuthLoginMethodDefault,
		).Where(
		"phone_number = ?", phoneNumber,
	).Limit(1).Find(result).Error
}

func (repository *Repository) GetByProvider(provider string, providerUserID string) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Preload(clause.Associations).
		Where(
			"login_method = ?", constants.AuthLoginMethodProvider,
		).Where(
		"provider = ?", provider,
	).Where(
		"provider_user_id = ?", providerUserID,
	).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, request *data.GetAllRequest) (result []model.User, err error) {
	result = make([]model.User, 0)
	var where string = ""
	if request != nil {
		if len(request.RoleName) > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("roles.name = %s", request.RoleName))
		}
	}
	if filter != nil && len(filter.Search) >= 1 {
		tempWhere := fmt.Sprintf(
			"(CAST(users.id AS TEXT) = '%s' OR users.email ILIKE '%s' OR CAST(users.phone_number AS TEXT) ILIKE '%s' OR infos.first_name ILIKE '%s' OR infos.last_name ILIKE '%s' OR infos.username ILIKE '%s')",
			filter.Search,
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
		)
		where = helpers.AppendWhereClause(where, tempWhere)
	}
	newFilter := filter
	newFilter.OrderBy = "users." + newFilter.OrderBy
	tmpErr := repository.Db.
		Preload(clause.Associations).
		Scopes(
			helpers.PaginationScope(
				repository.Db,
				"SELECT users.* "+
					"LEFT JOIN user_infos AS infos ON users.user_info_id = infos.id "+
					"LEFT JOIN roles ON users.role_id = roles.id",
				where,
				pagination,
				newFilter,
			),
		).Find(&result).Error

	err = tmpErr
	return
}

// ----------------- Authentication -----------------
func (repository *Repository) UpdateUserPassword(id int64, password string) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Update("password", password).Error
}

func (repository *Repository) UpdateUserActivation(id int64, item *model.User) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"is_activated": item.IsActivated,
			"activated_at": item.ActivatedAt,
			"user_info_id": item.UserInfoID,
			"user_mfa_id":  item.UserMfaID,
		},
	).Error
}

// ----------------- Profile -----------------

func (repository *Repository) UpdateUserEmail(id int64, email string) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"email": email,
		},
	).Error
}
func (repository *Repository) UpdateUserPhoneNumber(id int64, phoneNumber uint64) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"phone_number": phoneNumber,
		},
	).Error
}

func (repository *Repository) UpdateProfileInfo(id int64, item *model.UserInfo) (*model.UserInfo, error) {
	result := &model.UserInfo{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"username":   item.Username,
			"first_name": item.FirstName,
			"last_name":  item.LastName,
			"address":    item.Address,
			"image":      item.Image,
			"language":   item.Language,
		},
	).Error
}

func (repository *Repository) UpdateProfileMfa(id int64, column string, value bool) (*model.UserMfa, error) {
	result := &model.UserMfa{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"" + column: value,
		},
	).Error
}
