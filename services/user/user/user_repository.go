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

func (repository *Repository) Create(item *model.User) (result *model.User, err error) {
	// Create the item
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.User{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) CreateUserInfo(item *model.UserInfo) (result *model.UserInfo, err error) {
	// Create the item
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.UserInfo{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) CreateUserConfig(item *model.UserConfig) (result *model.UserConfig, err error) {
	// Create the item
	err = repository.Db.Create(item).Error
	if err != nil {
		return nil, err
	}

	// Find the created item
	result = &model.UserConfig{}
	err = repository.Db.
		Preload(clause.Associations).
		First(result, item.ID).Error
	return
}

func (repository *Repository) UpdateByID(id int64, item *model.User) (result *model.User, err error) {
	// Update the item
	fields := map[string]any{
		"school_id": item.SchoolID,
		"role_id":   item.RoleID,

		"email":        item.Email,
		"phone_number": item.PhoneNumber,
		"status":       item.Status,
		"is_activated": item.IsActivated,
		"activated_at": item.ActivatedAt,
	}
	err = repository.Db.
		Model(&model.User{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.User{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) UpdateEmailByID(id int64, email string) (result *model.User, err error) {
	// Update the item
	fields := map[string]any{
		"email": email,
	}
	err = repository.Db.
		Model(&model.User{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.User{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) UpdatePhoneNumberByID(id int64, phoneNumber uint64) (result *model.User, err error) {
	// Update the item
	fields := map[string]any{
		"phone_number": phoneNumber,
	}
	err = repository.Db.
		Model(&model.User{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.User{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) UpdatePasswordByID(id int64, password string) (result *model.User, err error) {
	// Update the item
	fields := map[string]any{
		"password": password,
	}
	err = repository.Db.
		Model(&model.User{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.User{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) UpdateActivationByID(id int64, item *model.User) (result *model.User, err error) {
	// Update the item
	fields := map[string]any{
		"is_activated": item.IsActivated,
		"activated_at": item.ActivatedAt,
		"info_id":      item.InfoID,
		"config_id":    item.ConfigID,
	}
	err = repository.Db.
		Model(&model.User{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.User{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) UpdateUserConfigAllowNotificationByID(id int64, enabled bool) (result *model.UserConfig, err error) {
	// Update the item
	fields := map[string]any{
		"allow_notifications": enabled,
	}
	err = repository.Db.
		Model(&model.UserConfig{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.UserConfig{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) UpdateUserConfigFieldByID(id int64, column string, value bool) (result *model.UserConfig, err error) {
	// Update the item
	fields := map[string]any{}
	if len(column) > 0 {
		fields[column] = value
	}
	err = repository.Db.
		Model(&model.UserConfig{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.UserConfig{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) UpdateUserInfoByID(id int64, item *model.UserInfo) (result *model.UserInfo, err error) {
	// Update the item
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
	err = repository.Db.
		Model(&model.UserInfo{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.UserInfo{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) UpdateUserConfigByID(id int64, item *model.UserConfig) (result *model.UserConfig, err error) {
	// Update the item
	fields := map[string]any{
		"whatsapp_phone_number": item.WhatsappPhoneNumber,
		"telegram_chat_id":      item.TelegramChatID,
	}
	err = repository.Db.
		Model(&model.UserConfig{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.UserConfig{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(result).Error
	return
}

func (repository *Repository) UpdateUserConfigWebPushSubscriptionByID(id int64, endpoint string, KeyP256dh string, keyAuth string) (result *model.UserConfig, err error) {
	// Update the item
	fields := map[string]any{
		"web_push_subscription_endpoint":   endpoint,
		"web_push_subscription_key_p256dh": KeyP256dh,
		"web_push_subscription_key_auth":   keyAuth,
	}
	err = repository.Db.
		Model(&model.UserConfig{}).
		Where("id = ?", id).
		Updates(fields).Error
	if err != nil {
		return
	}

	// Find the updated item
	result = &model.UserConfig{}
	err = repository.Db.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(result).Error
	return
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
