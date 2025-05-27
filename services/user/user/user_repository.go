package user

import (
	"api/common/constants"
	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/services/user/user/model"
	"fmt"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(user *model.User) (*model.User, error) {
	result := *user
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) AssignRole(userID int64, roleID int64) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", userID).Updates(
		map[string]any{
			"role_id": roleID,
		},
	).Error
}

func (repository *Repository) UpdateByID(userID int64, user *model.User) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", userID).Updates(
		map[string]any{
			"email":        user.Email,
			"phone_number": user.PhoneNumber,
			"role_id":      user.RoleID,
			"is_activated": user.IsActivated,
		},
	).Error
}

func (repository *Repository) DeleteByID(userID int64) (int64, error) {
	result := repository.Db.Where("id = ?", userID).Delete(&model.User{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteRoleByUserIDRoleID(userID int64, roleID int64) (int64, error) {
	result := repository.Db.Model(&model.User{}).Where("id = ?", userID).Where("role_id = ?", roleID).Updates(
		map[string]any{
			"role_id": nil,
		},
	)
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultipleByID(list []int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.User{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) GetByID(userID int64) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", userID).Limit(1).Find(result).Error
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

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, roleName string) (result []model.User, err error) {
	result = make([]model.User, 0)
	var where string = ""
	if len(roleName) > 0 {
		where = fmt.Sprintf("WHERE roles.name = '%s'", roleName)
	}
	if filter != nil && len(filter.Search) >= 1 {
		tempWhere := fmt.Sprintf(
			"CAST(users.id AS TEXT) = '%s' OR users.email ILIKE '%s' OR CAST(users.phone_number AS TEXT) ILIKE '%s' OR infos.first_name ILIKE '%s' OR infos.last_name ILIKE '%s' OR infos.username ILIKE '%s'",
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
	newFilter := filter
	newFilter.OrderBy = "users." + newFilter.OrderBy
	tmpErr := repository.Db.Preload(clause.Associations).Scopes(
		helpers.PaginationScope(
			repository.Db,
			"SELECT users.id, users.email, users.phone_number, users.login_method, users.provider, users.provider_user_id"+
				", users.is_activated, users.activated_at, users.role_id, users.user_info_id, users.user_mfa_id"+
				", users.created_at, users.updated_at FROM users "+
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

func (repository *Repository) CreateUserInfo(userInfo *model.UserInfo) (*model.UserInfo, error) {
	result := *userInfo
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}
func (repository *Repository) CreateUserMfa(userMfa *model.UserMfa) (*model.UserMfa, error) {
	result := *userMfa
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}
func (repository *Repository) UpdateUserPassword(userID int64, password string) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", userID).Update("password", password).Error
}

func (repository *Repository) UpdateUserActivation(userID int64, user *model.User) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", userID).Updates(
		map[string]any{
			"is_activated": user.IsActivated,
			"activated_at": user.ActivatedAt,
			"user_info_id": user.UserInfoID,
			"user_mfa_id":  user.UserMfaID,
		},
	).Error
}

// ----------------- Profile -----------------

func (repository *Repository) UpdateEmail(userID int64, email string) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", userID).Updates(
		map[string]any{
			"email": email,
		},
	).Error
}
func (repository *Repository) UpdatePhoneNumber(userID int64, phoneNumber uint64) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", userID).Updates(
		map[string]any{
			"phone_number": phoneNumber,
		},
	).Error
}
func (repository *Repository) UpdatePassword(userID int64, password string) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", userID).Updates(
		map[string]any{
			"password": password,
		},
	).Error
}

func (repository *Repository) UpdateProfileInfo(userInfoID int64, userInfo *model.UserInfo) (*model.UserInfo, error) {
	result := &model.UserInfo{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", userInfoID).Updates(
		map[string]any{
			"username":   userInfo.Username,
			"first_name": userInfo.FirstName,
			"last_name":  userInfo.LastName,
			"address":    userInfo.Address,
			"image":      userInfo.Image,
			"language":   userInfo.Language,
		},
	).Error
}

func (repository *Repository) UpdateProfileMfa(userMfaID int64, column string, value bool) (*model.UserMfa, error) {
	result := &model.UserMfa{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", userMfaID).Updates(
		map[string]any{
			"" + column: value,
		},
	).Error
}
