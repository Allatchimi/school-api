package school

import (
	"fmt"
	"net/http"

	"api/common/constants"
	deploymentHelper "api/common/helpers/deployment"
	"api/common/types"
	"api/common/utils"
	"api/services/school/common/school/data"
	"api/services/school/common/school/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

const MODEL_NAME = "school"
const DEFAULT_ERROR_MESSAGE = "interact with school model"

func (service *Service) Create(inputJwtToken *types.JwtToken, request *data.SchoolRequest) (result *model.School, errCode int, err error) {
	// Format request
	item := &model.School{
		Name:   request.Name,
		Type:   request.Type,
		Status: request.Status,

		Favicon:   request.Favicon,
		Logo:      request.Logo,
		LogoWhite: request.LogoWhite,

		Currency:     request.Currency,
		PaymentCount: request.PaymentCount,

		Info:   model.FromInfoRequest(request.Info),
		Config: model.FromConfigRequest(request.Config),
	}

	// Create info
	newInfo, err := service.Repository.CreateSchoolInfo(&model.SchoolInfo{
		FullName:    item.Info.FullName,
		Description: item.Info.Description,
		Motto:       item.Info.Motto,

		PhoneNumber1: item.Info.PhoneNumber1,
		PhoneNumber2: item.Info.PhoneNumber2,
		PhoneNumber3: item.Info.PhoneNumber3,

		Email1: item.Info.Email1,
		Email2: item.Info.Email2,
		Email3: item.Info.Email3,

		Founder:   item.Info.Founder,
		FoundedAt: item.Info.FoundedAt,

		Address:           item.Info.Address,
		LocationLongitude: item.Info.LocationLongitude,
		LocationLatitude:  item.Info.LocationLatitude,

		SocialMediaTelegram: item.Info.SocialMediaTelegram,
		SocialMediaWhasapp:  item.Info.SocialMediaWhasapp,
		SocialMediaYoutube:  item.Info.SocialMediaYoutube,
		SocialMediaTwitter:  item.Info.SocialMediaTwitter,
		SocialMediaFacebook: item.Info.SocialMediaFacebook,

		Image1: item.Info.Image1,
		Image2: item.Info.Image2,
		Image3: item.Info.Image3,
		Image4: item.Info.Image4,
		Image5: item.Info.Image5,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Create config
	newConfig, err := service.Repository.CreateSchoolConfig(&model.SchoolConfig{
		WebsiteDomainName:   item.Config.WebsiteDomainName,
		UserEmailDomainName: item.Config.UserEmailDomainName,
		SupportEmail:        item.Config.SupportEmail,

		GoogleWorkspaceCredentials:     item.Config.GoogleWorkspaceCredentials,
		GoogleWorkspaceUserEmailDomain: item.Config.GoogleWorkspaceUserEmailDomain,

		SmsUserID:        item.Config.SmsUserID,
		WhatsappToken:    item.Config.WhatsappToken,
		WhatsappPhoneID:  item.Config.WhatsappPhoneID,
		TelegramBotToken: item.Config.TelegramBotToken,

		WebsiteTitle:       item.Config.WebsiteTitle,
		WebsiteDescription: item.Config.WebsiteDescription,

		ColorPrimary:        item.Config.ColorPrimary,
		ColorPrimaryBg:      item.Config.ColorPrimaryBg,
		ColorPrimaryBgHover: item.Config.ColorPrimaryBgHover,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Create school
	result, err = service.Repository.Create(&model.School{
		Name:   item.Name,
		Type:   item.Type,
		Status: item.Status,

		DeploymentRequest: constants.SCHOOL_DEPLOYMENT_REQUEST_CREATE,
		DeploymentStatus:  constants.SCHOOL_DEPLOYMENT_STATUS_INITIATED,
		DeploymentCount:   1,

		Favicon:   item.Favicon,
		Logo:      item.Logo,
		LogoWhite: item.LogoWhite,

		Currency:     item.Currency,
		PaymentCount: item.PaymentCount,

		ConfigID: newConfig.ID,
		InfoID:   newInfo.ID,
	})
	if err != nil {
		pgState, errPgState := utils.ExtractSQLState(err.Error())
		if errPgState == nil {
			if pgState == constants.PG_ERROR_UNIQUE_COLUMN {
				errCode = http.StatusFound
				err = constants.Http302ErrorMessage(MODEL_NAME)
				return
			}
		}
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Deploy school
	go func() {
		err := deploymentHelper.DeploySchool(result)
		if err != nil {
			service.Repository.UpdateDeploymentStatusByID(result.ID, &data.SchoolDeploymentStatusRequest{
				Status:   constants.SCHOOL_DEPLOYMENT_STATUS_FAILED,
				Feedback: fmt.Sprintf("Failed to deploy school! Error: %s", err.Error()),
			})
			return
		}
		service.Repository.UpdateDeploymentStatusByID(result.ID, &data.SchoolDeploymentStatusRequest{
			Status:   constants.SCHOOL_DEPLOYMENT_STATUS_PENDING,
			Feedback: "School deployment pushed to GitHub! Now waiting for deployment to complete.",
		})
	}()
	return
}

func (service *Service) Update(inputJwtToken *types.JwtToken, id int64, request *data.SchoolRequest) (result *model.School, errCode int, err error) {
	// Check if school exists
	foundItem, err := service.Repository.GetByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil || foundItem.ID != id {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Check if the deployment status is pending
	if foundItem.DeploymentStatus == constants.SCHOOL_DEPLOYMENT_STATUS_INITIATED ||
		foundItem.DeploymentStatus == constants.SCHOOL_DEPLOYMENT_STATUS_PENDING {
		errCode = http.StatusLocked
		err = constants.Http423LockedErrorMessage()
		return
	}

	// Format request
	item := &model.School{
		Name:   request.Name,
		Type:   request.Type,
		Status: request.Status,

		DeploymentRequest: constants.SCHOOL_DEPLOYMENT_REQUEST_UPDATE,
		DeploymentStatus:  constants.SCHOOL_DEPLOYMENT_STATUS_INITIATED,
		DeploymentCount:   foundItem.DeploymentCount + 1,

		Favicon:   request.Favicon,
		Logo:      request.Logo,
		LogoWhite: request.LogoWhite,

		Currency:     request.Currency,
		PaymentCount: request.PaymentCount,

		Info:   model.FromInfoRequest(request.Info),
		Config: model.FromConfigRequest(request.Config),
	}

	// Check unique
	foundUnique, err := service.Repository.GetUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreSameUniqueObjects(foundUnique, item) && !service.Repository.AreSameUniqueObjects(foundUnique, foundItem) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}

	// Update school
	result, err = service.Repository.UpdateByID(id, item)
	if err != nil {
		pgState, errPgState := utils.ExtractSQLState(err.Error())
		if errPgState == nil {
			if pgState == constants.PG_ERROR_UNIQUE_COLUMN {
				errCode = http.StatusFound
				err = constants.Http302ErrorMessage(MODEL_NAME)
				return
			}
		}
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if result == nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Update config
	newConfig, _, err := service.UpdateConfig(inputJwtToken, foundItem.ConfigID, item.Config)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	result.Config = newConfig

	// Update info
	newInfo, errCode, err := service.UpdateInfo(inputJwtToken, foundItem.InfoID, item.Info)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	result.Info = newInfo

	// Deploy school
	go func() {
		if !(result.IsSameDeploymentAsRequest(request) && newConfig.IsSameDeploymentAsRequest(request.Config)) {
			err := deploymentHelper.DeploySchool(result)
			if err != nil {
				service.Repository.UpdateDeploymentStatusByID(result.ID, &data.SchoolDeploymentStatusRequest{
					Status:   constants.SCHOOL_DEPLOYMENT_STATUS_FAILED,
					Feedback: fmt.Sprintf("Failed to deploy school! Error: %s", err.Error()),
				})
				return
			}
			service.Repository.UpdateDeploymentStatusByID(result.ID, &data.SchoolDeploymentStatusRequest{
				Status:   constants.SCHOOL_DEPLOYMENT_STATUS_PENDING,
				Feedback: "School deployment pushed to GitHub! Now waiting for deployment to complete.",
			})
		}
	}()
	return
}

func (service *Service) UpdateDeploymentStatus(inputJwtToken *types.JwtToken, id int64, request *data.SchoolDeploymentStatusRequest) (errCode int, err error) {
	// Check if school exists
	foundItem, err := service.Repository.GetByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil || foundItem.ID != id {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Delete school if deployment request is delete
	if foundItem.DeploymentRequest == constants.SCHOOL_DEPLOYMENT_REQUEST_DELETE &&
		foundItem.DeploymentStatus == constants.SCHOOL_DEPLOYMENT_STATUS_PENDING &&
		(request.Status == constants.SCHOOL_DEPLOYMENT_STATUS_DONE ||
			request.Status == constants.SCHOOL_DEPLOYMENT_STATUS_DONE_NO_CHANGES) {
		_, err = service.Repository.DeleteByID(id)
		if err != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
		return
	}

	// Update status
	_, err = service.Repository.UpdateDeploymentStatusByID(id, request)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) UpdateInfo(inputJwtToken *types.JwtToken, id int64, item *model.SchoolInfo) (result *model.SchoolInfo, errCode int, err error) {
	// Check if school info already exists
	foundItem, err := service.Repository.GetSchoolInfoByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil || foundItem.ID != id {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Update
	result, err = service.Repository.UpdateSchoolInfoByID(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) UpdateConfig(inputJwtToken *types.JwtToken, id int64, item *model.SchoolConfig) (result *model.SchoolConfig, errCode int, err error) {
	// Check if school config already exists
	foundItem, err := service.Repository.GetSchoolConfigByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil || foundItem.ID != id {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Update
	result, err = service.Repository.UpdateSchoolConfigByID(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) Delete(inputJwtToken *types.JwtToken, id int64) (affectedRows int64, errCode int, err error) {
	// Check if school exists
	foundItem, err := service.Repository.GetByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil || foundItem.ID < 1 || foundItem.ID != id {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Check if the deployment status is pending
	if foundItem.DeploymentStatus == constants.SCHOOL_DEPLOYMENT_STATUS_INITIATED ||
		foundItem.DeploymentStatus == constants.SCHOOL_DEPLOYMENT_STATUS_PENDING {
		errCode = http.StatusLocked
		err = constants.Http423LockedErrorMessage()
		return
	}

	// Delete school deployment
	go func() {
		err := deploymentHelper.DeleteSchoolDeployment(id)
		if err != nil {
			service.Repository.UpdateDeploymentStatusByID(id, &data.SchoolDeploymentStatusRequest{
				Status:   constants.SCHOOL_DEPLOYMENT_STATUS_FAILED,
				Feedback: fmt.Sprintf("Failed to delete school! %s", err.Error()),
			})
			return
		}
		service.Repository.UpdateDeploymentStatusByID(id, &data.SchoolDeploymentStatusRequest{
			Status:   constants.SCHOOL_DEPLOYMENT_STATUS_PENDING,
			Feedback: "Deleted school deployment pushed to GitHub! Now waiting for deletion to complete.",
		})
	}()
	return
}

func (service *Service) DeleteMultiple(inputJwtToken *types.JwtToken, list []int64) (affectedRows int64, errCode int, err error) {
	for _, id := range list {
		total, _, err := service.Delete(inputJwtToken, id)
		if err != nil && total > 0 {
			affectedRows++
		}
	}
	return
}

func (service *Service) Get(inputJwtToken *types.JwtToken, id int64) (result *model.School, errCode int, err error) {
	result, err = service.Repository.GetByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	return
}

func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, request *data.GetAllRequest) (result []model.School, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, request)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
