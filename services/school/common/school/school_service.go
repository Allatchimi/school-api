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

func (service *Service) Create(
	ctxData *types.ContextData,
	request *data.SchoolRequest,
) (result *model.School, errCode int, err error) {
	// Format request
	item := &model.School{
		Name:               request.Name,
		Type:               request.Type,
		Status:             request.Status,
		Favicon:            request.Favicon,
		Logo:               request.Logo,
		LogoWhite:          request.LogoWhite,
		Currency:           request.Currency,
		PaymentCount:       request.PaymentCount,
		DeploymentRequest:  constants.SCHOOL_DEPLOYMENT_REQUEST_CREATE,
		DeploymentStatus:   constants.SCHOOL_DEPLOYMENT_STATUS_INITIATED,
		DeploymentFeedback: "",
		DeploymentCount:    1,
		Info:               model.FromInfoRequest(request.Info),
		Config:             model.FromConfigRequest(request.Config),
	}

	// Check unique
	foundUnique, err := service.Repository.GetUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreSameUniqueObjects(foundUnique, item) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}

	// Check unique config
	foundConfigUnique, err := service.Repository.GetSchoolConfigUniqueObject(item.Config)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreSchoolConfigSameUniqueObjects(foundConfigUnique, item.Config) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}

	// Create info
	newInfo, err := service.Repository.CreateSchoolInfo(item.Info)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Create config
	newConfig, err := service.Repository.CreateSchoolConfig(item.Config)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Create school
	item.InfoID = newInfo.ID
	item.ConfigID = newConfig.ID
	result, err = service.Repository.Create(item)
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

func (service *Service) Update(
	ctxData *types.ContextData,
	id int64,
	request *data.SchoolRequest,
) (result *model.School, errCode int, err error) {
	// Check if the item exists
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
	if (foundItem.DeploymentRequest == constants.SCHOOL_DEPLOYMENT_REQUEST_CREATE || foundItem.DeploymentRequest == constants.SCHOOL_DEPLOYMENT_REQUEST_UPDATE) &&
		(foundItem.DeploymentStatus == constants.SCHOOL_DEPLOYMENT_STATUS_INITIATED ||
			foundItem.DeploymentStatus == constants.SCHOOL_DEPLOYMENT_STATUS_PENDING) {
		errCode = http.StatusLocked
		err = constants.Http423LockedErrorMessage()
		return
	}

	// Format request
	item := &model.School{
		Name:               request.Name,
		Type:               request.Type,
		Status:             request.Status,
		DeploymentRequest:  constants.SCHOOL_DEPLOYMENT_REQUEST_UPDATE,
		DeploymentStatus:   constants.SCHOOL_DEPLOYMENT_STATUS_INITIATED,
		DeploymentFeedback: "",
		DeploymentCount:    foundItem.DeploymentCount + 1,
		Favicon:            request.Favicon,
		Logo:               request.Logo,
		LogoWhite:          request.LogoWhite,
		Currency:           request.Currency,
		PaymentCount:       request.PaymentCount,
		Info:               model.FromInfoRequest(request.Info),
		Config:             model.FromConfigRequest(request.Config),
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

	// Check unique config
	foundConfigUnique, err := service.Repository.GetSchoolConfigUniqueObject(item.Config)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreSchoolConfigSameUniqueObjects(foundConfigUnique, item.Config) && !service.Repository.AreSchoolConfigSameUniqueObjects(foundConfigUnique, foundItem.Config) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}

	// Update info
	newInfo, errCode, err := service.UpdateInfo(ctxData, foundItem.InfoID, item.Info)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Update config
	newConfig, _, err := service.UpdateConfig(ctxData, foundItem.ConfigID, item.Config)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Update school
	if foundItem.IsSameDeploymentAsRequest(request) && foundItem.Config.IsSameDeploymentAsRequest(request.Config) {
		item.DeploymentStatus = constants.SCHOOL_DEPLOYMENT_STATUS_DONE_NO_CHANGES
	}
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
	if result == nil || result.ID < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	result.Info = newInfo
	result.Config = newConfig

	// Deploy school
	go func() {
		if !(foundItem.IsSameDeploymentAsRequest(request) && foundItem.Config.IsSameDeploymentAsRequest(request.Config)) {
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

func (service *Service) UpdateDeploymentStatus(
	ctxData *types.ContextData,
	id int64,
	request *data.SchoolDeploymentStatusRequest,
) (errCode int, err error) {
	// Check if the item exists
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
		_, err = service.Repository.DeleteByID(foundItem.ID)
		_, err = service.Repository.DeleteSchoolInfoByID(foundItem.InfoID)
		_, err = service.Repository.DeleteSchoolInfoByID(foundItem.ConfigID)
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

func (service *Service) UpdateInfo(
	ctxData *types.ContextData,
	id int64,
	item *model.SchoolInfo,
) (result *model.SchoolInfo, errCode int, err error) {
	// Check if the item exists
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

func (service *Service) UpdateConfig(
	ctxData *types.ContextData,
	id int64,
	item *model.SchoolConfig,
) (result *model.SchoolConfig, errCode int, err error) {
	// Check if the item exists
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

func (service *Service) Delete(
	ctxData *types.ContextData,
	id int64,
) (affectedRows int64, errCode int, err error) {
	// Check if the item exists
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
	if foundItem.DeploymentRequest == constants.SCHOOL_DEPLOYMENT_REQUEST_DELETE && (foundItem.DeploymentStatus == constants.SCHOOL_DEPLOYMENT_STATUS_INITIATED ||
		foundItem.DeploymentStatus == constants.SCHOOL_DEPLOYMENT_STATUS_PENDING) {
		errCode = http.StatusLocked
		err = constants.Http423LockedErrorMessage()
		return
	}

	// Update deployment status
	foundItem.DeploymentStatus = constants.SCHOOL_DEPLOYMENT_STATUS_PENDING
	foundItem.DeploymentRequest = constants.SCHOOL_DEPLOYMENT_REQUEST_DELETE
	foundItem.DeploymentFeedback = ""
	service.Repository.UpdateByID(id, foundItem)

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

func (service *Service) DeleteMultiple(
	ctxData *types.ContextData,
	list []int64,
) (affectedRows int64, errCode int, err error) {
	for _, id := range list {
		total, _, err := service.Delete(ctxData, id)
		if err != nil && total > 0 {
			affectedRows++
		}
	}
	return
}

func (service *Service) Get(
	ctxData *types.ContextData,
	id int64,
) (result *model.School, errCode int, err error) {
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

func (service *Service) GetPublic(
	ctxData *types.ContextData,
) (result *model.School, errCode int, err error) {
	result, err = service.Repository.GetByID(ctxData.Jwt.SchoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if result == nil || result.ID < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	return
}

func (service *Service) GetAll(
	ctxData *types.ContextData,
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllRequest,
) (result []model.School, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, request)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
