package user

import (
	"api/common/constants"
	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/services/user/user/data"
	"api/services/user/user/model"

	dataMonitoring "api/services/others/monitoring/data"
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
	fields := map[string]any{
		"email":        item.Email,
		"phone_number": item.PhoneNumber,
		"status":       item.Status,
		"school_id":    nil,
		"role_id":      item.RoleID,
		"is_activated": item.IsActivated,
	}
	if item.SchoolID > 0 {
		fields["school_id"] = item.SchoolID
	}
	return result, repository.Db.
		Preload(clause.Associations).
		Model(&model.User{}).
		Where("id = ?", id).
		Updates(
			fields,
		).
		Find(result).Error
}

func (repository *Repository) UpdateEmailByID(id int64, email string) (*model.User, error) {
	result := &model.User{}
	fields := map[string]any{
		"email": email,
	}
	return result, repository.Db.
		Preload(clause.Associations).
		Model(&model.User{}).
		Where("id = ?", id).
		Updates(
			fields,
		).
		Find(result).Error
}

func (repository *Repository) UpdatePhoneNumberByID(id int64, phoneNumber uint64) (*model.User, error) {
	result := &model.User{}
	fields := map[string]any{
		"phone_number": phoneNumber,
	}
	return result, repository.Db.
		Preload(clause.Associations).
		Model(&model.User{}).
		Where("id = ?", id).
		Updates(
			fields,
		).
		Find(result).Error
}

func (repository *Repository) UpdatePasswordByID(id int64, password string) (*model.User, error) {
	result := &model.User{}
	fields := map[string]any{
		"password": password,
	}
	return result, repository.Db.
		Preload(clause.Associations).
		Model(&model.User{}).
		Where("id = ?", id).
		Updates(
			fields,
		).
		Find(result).Error
}

func (repository *Repository) UpdateActivationByID(id int64, item *model.User) (*model.User, error) {
	result := &model.User{}
	fields := map[string]any{
		"is_activated": item.IsActivated,
		"activated_at": item.ActivatedAt,
		"info_id":      item.InfoID,
		"config_id":    item.ConfigID,
	}
	return result, repository.Db.
		Preload(clause.Associations).
		Model(&model.User{}).
		Where("id = ?", id).
		Updates(
			fields,
		).
		Find(result).Error
}

func (repository *Repository) UpdateUserConfigAllowNotificationByID(id int64, enabled bool) (*model.UserConfig, error) {
	result := &model.UserConfig{}
	fields := map[string]any{
		"allow_notifications": enabled,
	}
	return result, repository.Db.
		Preload(clause.Associations).
		Model(&model.UserConfig{}).
		Where("id = ?", id).
		Updates(
			fields,
		).
		Find(result).Error
}

func (repository *Repository) UpdateUserConfigFieldByID(id int64, column string, value bool) (*model.UserConfig, error) {
	result := &model.UserConfig{}
	fields := map[string]any{}
	if len(column) > 0 {
		fields[column] = value
	}
	return result, repository.Db.
		Preload(clause.Associations).
		Model(&model.UserConfig{}).
		Where("id = ?", id).
		Updates(
			fields,
		).
		Find(result).Error
}
func (repository *Repository) UpdateUserInfoByID(id int64, item *model.UserInfo) (*model.UserInfo, error) {
	result := &model.UserInfo{}
	fields := map[string]any{
		"username":       item.Username,
		"first_name":     item.FirstName,
		"last_name":      item.LastName,
		"gender":         item.Gender,
		"birthday":       item.Birthday,
		"birth_location": item.BirthLocation,
		"address":        item.Address,
		"language":       item.Language,
		"image":          item.Image,
	}
	return result, repository.Db.
		Preload(clause.Associations).
		Model(&model.UserInfo{}).
		Where("id = ?", id).
		Updates(
			fields,
		).
		Find(result).Error
}
func (repository *Repository) UpdateUserConfigByID(id int64, item *model.UserConfig) (*model.UserConfig, error) {
	result := &model.UserConfig{}
	fields := map[string]any{
		"whatsapp_phone_number": item.WhatsappPhoneNumber,
		"telegram_chat_id":      item.TelegramChatID,
	}
	return result, repository.Db.
		Preload(clause.Associations).
		Model(&model.UserConfig{}).
		Where("id = ?", id).
		Updates(
			fields,
		).
		Find(result).Error
}

func (repository *Repository) UpdateUserConfigWebPushSubscriptionByID(id int64, endpoint string, KeyP256dh string, keyAuth string) (*model.UserConfig, error) {
	result := &model.UserConfig{}
	fields := map[string]any{
		"web_push_subscription_endpoint":   endpoint,
		"web_push_subscription_key_p256dh": KeyP256dh,
		"web_push_subscription_key_auth":   keyAuth,
	}
	return result, repository.Db.
		Preload(clause.Associations).
		Model(&model.UserConfig{}).
		Where("id = ?", id).
		Updates(
			fields,
		).
		Find(result).Error
}

func (repository *Repository) DeleteByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.User{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultipleByID(list []int64, schoolID int64) (result int64, err error) {
	if len(list) < 1 {
		return
	}
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	var query *gorm.DB = repository.Db.Where(where)
	if schoolID > 0 {
		query = query.Where("school_id = ?", schoolID)
	}
	query = query.Delete(&model.User{})

	result = query.RowsAffected
	err = query.Error
	return
}

func (repository *Repository) GetByID(id int64) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByIDSchoolID(id int64, schoolID int64) (*model.User, error) {
	result := &model.User{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetByEmailSchoolID(email string, schoolID int64) (*model.User, error) {
	result := &model.User{}
	if schoolID < 1 {
		return result, repository.Db.Preload(clause.Associations).
			Where(
				"login_method = ?", constants.AuthLoginMethodDefault,
			).Where(
			"email = ?", email,
		).
			Or("school_id < ?", 1).Or("school_id = ?", nil).
			Limit(1).Find(result).Error
	}
	return result, repository.Db.Preload(clause.Associations).
		Where(
			"login_method = ?", constants.AuthLoginMethodDefault,
		).Where(
		"email = ?", email,
	).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetByPhoneNumberSchoolID(phoneNumber uint64, schoolID int64) (*model.User, error) {
	result := &model.User{}
	if schoolID < 1 {
		return result, repository.Db.Preload(clause.Associations).
			Where(
				"login_method = ?", constants.AuthLoginMethodDefault,
			).Where(
			"phone_number = ?", phoneNumber,
		).
			Or("school_id < ?", 1).Or("school_id = ?", nil).
			Limit(1).Find(result).Error
	}
	return result, repository.Db.Preload(clause.Associations).
		Where(
			"login_method = ?", constants.AuthLoginMethodDefault,
		).Where(
		"phone_number = ?", phoneNumber,
	).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
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

func (repository *Repository) GetAll(
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllRequest,
) (result []model.User, err error) {
	result = make([]model.User, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "users.school_id = ?")
			args = append(args, request.SchoolID)
		}
		if request.RoleID > 0 {
			where = helpers.AppendWhereClause(where, "users.role_id = ?")
			args = append(args, request.RoleID)
		}
		if len(request.RoleName) > 0 {
			where = helpers.AppendWhereClause(where, "roles.name = ?")
			args = append(args, request.RoleName)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(users.id AS TEXT) = ? OR
			users.email ILIKE ? OR
			CAST(users.phone_number AS TEXT) ILIKE ? OR
			infos.first_name ILIKE ? OR
			infos.last_name ILIKE ? OR
			infos.username ILIKE ? OR
			roles.name ILIKE ? OR
			schools.name ILIKE ? OR
			schools.type ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like, like, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT users.*
				FROM users
				LEFT JOIN user_infos AS infos ON users.info_id = infos.id
				LEFT JOIN roles ON users.role_id = roles.id
				LEFT JOIN schools ON users.school_id = schools.id`,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

	return
}

func (repository *Repository) CountAllByFeature(schoolID int64, feature string) (result int64, err error) {
	query := repository.Db.
		Model(&model.User{}).
		Joins("LEFT JOIN roles ON users.role_id = roles.id").
		Where("roles.feature = ?", feature)

	if schoolID > 0 {
		query = query.Where("users.school_id = ?", schoolID)
	}
	err = query.Count(&result).Error
	return
}

func (repository *Repository) CountAllGroupByFeature(schoolID int64, result *[]dataMonitoring.UsersByFeatureResponse) (err error) {
	query := repository.Db.
		Model(&model.User{}).
		Joins("LEFT JOIN roles ON users.role_id = roles.id")

	if schoolID > 0 {
		query = query.Where("users.school_id = ?", schoolID)
	}
	err = query.
		Select("roles.feature AS feature, COUNT(*) AS count").
		Group("feature").
		Order("feature").
		Find(&result).Error
	return
}

func (repository *Repository) CountAllGroupByGender(schoolID int64, result *[]dataMonitoring.UsersByGenderResponse) (err error) {
	query := repository.Db.
		Model(&model.User{}).
		Joins("LEFT JOIN user_infos ON users.info_id = user_infos.id")

	if schoolID > 0 {
		query = query.Where("users.school_id = ?", schoolID)
	}
	err = query.
		Select("user_infos.gender AS gender, COUNT(*) AS count").
		Group("gender").
		Order("gender").
		Find(&result).Error
	return
}

func (repository *Repository) CountAllGroupByMonth(schoolID int64, result *[]dataMonitoring.UsersByMonthResponse) (err error) {
	query := repository.Db.Model(&model.User{})

	if schoolID > 0 {
		query = query.Where("users.school_id = ?", schoolID)
	}
	err = query.
		Select("EXTRACT(MONTH FROM created_at) AS month, COUNT(*) AS count").
		Group("month").
		Order("month").
		Find(&result).Error
	return
}

func (repository *Repository) CountAllGroupByYear(schoolID int64, result *[]dataMonitoring.UsersByYearResponse) (err error) {
	query := repository.Db.Model(&model.User{})

	if schoolID > 0 {
		query = query.Where("users.school_id = ?", schoolID)
	}
	err = query.
		Select("EXTRACT(YEAR FROM created_at) AS year, COUNT(*) AS count").
		Group("year").
		Order("year").
		Find(&result).Error
	return
}
