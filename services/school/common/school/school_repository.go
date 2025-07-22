package school

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/services/school/common/school/data"
	"api/services/school/common/school/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.School) (*model.School, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) CreateSchoolInfo(item *model.SchoolInfo) (*model.SchoolInfo, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) CreateSchoolConfig(item *model.SchoolConfig) (*model.SchoolConfig, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) UpdateByID(id int64, item *model.School) (*model.School, error) {
	result := &model.School{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"name":   item.Name,
			"type":   item.Type,
			"status": item.Status,

			"deployment_request":  item.DeploymentRequest,
			"deployment_status":   item.DeploymentStatus,
			"deployment_feedback": item.DeploymentFeedback,
			"deployment_count":    item.DeploymentCount,

			"favicon":    item.Favicon,
			"logo":       item.Logo,
			"logo_white": item.LogoWhite,

			"currency":      item.Currency,
			"payment_count": item.PaymentCount,
		},
	).Error
}

func (repository *Repository) UpdateDeploymentStatusByID(id int64, request *data.SchoolDeploymentStatusRequest) (*model.School, error) {
	result := &model.School{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"deployment_status":   request.Status,
			"deployment_feedback": request.Feedback,
		},
	).Error
}

func (repository *Repository) UpdateConfigInfoByID(id int64, configID int64, infoID int64) (*model.School, error) {
	result := &model.School{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"config_id": configID,
			"info_id":   infoID,
		},
	).Error
}

func (repository *Repository) UpdateSchoolInfoByID(id int64, item *model.SchoolInfo) (*model.SchoolInfo, error) {
	result := &model.SchoolInfo{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"full_name":   item.FullName,
			"description": item.Description,
			"motto":       item.Motto,

			"phone_number1": item.PhoneNumber1,
			"phone_number2": item.PhoneNumber2,
			"phone_number3": item.PhoneNumber3,

			"email1": item.Email1,
			"email2": item.Email2,
			"email3": item.Email3,

			"founder":    item.Founder,
			"founded_at": item.FoundedAt,

			"address":            item.Address,
			"po_box":             item.PoBox,
			"location_longitude": item.LocationLongitude,
			"location_latitude":  item.LocationLatitude,

			"social_media_telegram": item.SocialMediaTelegram,
			"social_media_whasapp":  item.SocialMediaWhasapp,
			"social_media_youtube":  item.SocialMediaYoutube,
			"social_media_twitter":  item.SocialMediaTwitter,
			"social_media_facebook": item.SocialMediaFacebook,

			"image1": item.Image1,
			"image2": item.Image2,
			"image3": item.Image3,
			"image4": item.Image4,
			"image5": item.Image5,
		},
	).Error
}

func (repository *Repository) UpdateSchoolConfigByID(id int64, item *model.SchoolConfig) (*model.SchoolConfig, error) {
	result := &model.SchoolConfig{}
	return result, repository.Db.Preload(clause.Associations).Model(result).Where("id = ?", id).Updates(
		map[string]any{
			"website_domain_name":    item.WebsiteDomainName,
			"user_email_domain_name": item.UserEmailDomainName,
			"support_email":          item.SupportEmail,

			"google_workspace_credentials":       item.GoogleWorkspaceCredentials,
			"google_workspace_user_email_domain": item.GoogleWorkspaceUserEmailDomain,

			"sms_user_id":        item.SmsUserID,
			"whatsapp_token":     item.WhatsappToken,
			"whatsapp_phone_id":  item.WhatsappPhoneID,
			"telegram_bot_token": item.TelegramBotToken,

			"website_title":       item.WebsiteTitle,
			"website_description": item.WebsiteDescription,

			"color_primary":          item.ColorPrimary,
			"color_primary_bg":       item.ColorPrimaryBg,
			"color_primary_bg_hover": item.ColorPrimaryBgHover,
		},
	).Error
}

func (repository *Repository) DeleteByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.School{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultipleByID(list []int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.School{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) CountAll() (result int64, err error) {
	err = repository.Db.Model(&model.School{}).Count(&result).Error
	return
}

func (repository *Repository) GetByID(id int64) (*model.School, error) {
	result := &model.School{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetUniqueObject(item *model.School) (*model.School, error) {
	result := &model.School{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.School{
		Name: item.Name,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjects(item1 *model.School, item2 *model.School) bool {
	if item1 != nil && item2 != nil &&
		(item1.Name == item2.Name) {
		return true
	}
	return false
}

func (repository *Repository) GetSchoolInfoByID(id int64) (*model.SchoolInfo, error) {
	result := &model.SchoolInfo{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetSchoolConfigByID(id int64) (*model.SchoolConfig, error) {
	result := &model.SchoolConfig{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination, request *data.GetAllRequest) (result []model.School, err error) {
	result = make([]model.School, 0)
	var where string = ""
	if request != nil {
		if len(request.Type) > 0 {
			where = helpers.AppendWhereClause(where, fmt.Sprintf("schools.type = '%s'", request.Type))
		}
	}
	if filter != nil && len(filter.Search) > 0 {
		tempWhere := fmt.Sprintf(
			"(CAST(schools.id AS TEXT) = '%s' OR schools.name ILIKE '%s' OR CAST(schools.type AS TEXT) ILIKE '%s' OR infos.full_name ILIKE '%s' OR infos.slogan ILIKE '%s' OR infos.founder ILIKE '%s')",
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
	newFilter.OrderBy = "schools." + newFilter.OrderBy
	tmpErr := repository.Db.
		Preload(clause.Associations).
		Scopes(
			helpers.PaginationScope(
				repository.Db,
				"SELECT schools.* "+
					"FROM schools "+
					"LEFT JOIN school_infos as infos ON schools.info_id = infos.id "+
					"LEFT JOIN school_configs as configs ON schools.config_id = configs.id ",
				where,
				pagination,
				newFilter,
			),
		).Find(&result).Error

	err = tmpErr
	return
}
