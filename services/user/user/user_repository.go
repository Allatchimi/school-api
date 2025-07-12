package user

import (
	"api/common/constants"
	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/services/user/user/data"
	"api/services/user/user/model"

	dataMonitoring "api/services/common/monitoring/data"
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

func (repository *Repository) CreateUserConfig(item *model.UserConfig) (*model.UserConfig, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) UpdateByID(id int64, item *model.User) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"email":        item.Email,
			"phone_number": item.PhoneNumber,
			"status":       item.Status,
			"school_id":    item.SchoolID,
			"role_id":      item.RoleID,
			"is_activated": item.IsActivated,
		},
	).Error
}

func (repository *Repository) UpdateEmailByID(id int64, email string) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"email": email,
		},
	).Error
}

func (repository *Repository) UpdatePhoneNumberByID(id int64, phoneNumber uint64) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"phone_number": phoneNumber,
		},
	).Error
}

func (repository *Repository) UpdatePasswordByID(id int64, password string) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Update("password", password).Error
}

func (repository *Repository) UpdateActivationByID(id int64, item *model.User) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"is_activated":   item.IsActivated,
			"activated_at":   item.ActivatedAt,
			"user_info_id":   item.UserInfoID,
			"user_config_id": item.UserConfigID,
		},
	).Error
}

func (repository *Repository) UpdateUserConfigAllowNotificationByID(id int64, enabled bool) (*model.UserConfig, error) {
	result := &model.UserConfig{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"allow_notifications": enabled,
		},
	).Error
}

func (repository *Repository) UpdateUserConfigFieldByID(id int64, column string, value bool) (*model.UserConfig, error) {
	result := &model.UserConfig{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"" + column: value,
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

			"gender":         item.Gender,
			"birthday":       item.Birthday,
			"birth_location": item.BirthLocation,
			"address":        item.Address,
			"language":       item.Language,
			"image":          item.Image,
		},
	).Error
}
func (repository *Repository) UpdateUserConfigByID(id int64, item *model.UserConfig) (*model.UserConfig, error) {
	result := &model.UserConfig{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"whatsapp_phone_number": item.WhatsappPhoneNumber,
			"telegram_chat_id":      item.TelegramChatID,
		},
	).Error
}

func (repository *Repository) UpdateUserConfigWebPushSubscriptionByID(userID int64, endpoint string, KeyP256dh string, keyAuth string) (*model.UserConfig, error) {
	result := &model.UserConfig{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", userID).Updates(
		map[string]any{
			"web_push_subscription_endpoint":   endpoint,
			"web_push_subscription_key_p256dh": KeyP256dh,
			"web_push_subscription_key_auth":   keyAuth,
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

func (repository *Repository) CountAllGroupByYear(result *[]dataMonitoring.UsersByYearResponse) (err error) {
	err = repository.Db.
		Model(&model.User{}).
		Select("EXTRACT(YEAR FROM created_at) AS year, COUNT(*) AS count").
		Group("year").
		Order("year").
		Find(&result).Error
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

func (repository *Repository) GetByEmailSchoolID(email string, schoolID int64) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Preload(clause.Associations).
		Where(
			"login_method = ?", constants.AuthLoginMethodDefault,
		).Where(
		"email = ?", email,
	).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
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

func (repository *Repository) GetByPhoneNumberSchoolID(phoneNumber uint64, schoolID int64) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Preload(clause.Associations).
		Where(
			"login_method = ?", constants.AuthLoginMethodDefault,
		).Where(
		"phone_number = ?", phoneNumber,
	).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
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

func (repository *Repository) GetByProviderSchoolID(provider string, providerUserID string, schoolID int64) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Preload(clause.Associations).
		Where(
			"login_method = ?", constants.AuthLoginMethodProvider,
		).Where(
		"provider = ?", provider,
	).Where(
		"provider_user_id = ?", providerUserID,
	).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, request *data.GetAllRequest) (result []model.User, err error) {
	result = make([]model.User, 0)
	var where string = ""
	if request != nil {
		if len(request.RoleName) > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("roles.name = %s", request.RoleName))
		}
	}
	if filter != nil && len(filter.Search) > 0 {
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
					"FROM users "+
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
