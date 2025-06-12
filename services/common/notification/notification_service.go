package notification

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/common/notification/data"
	"api/services/common/notification/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

const MODEL_NAME = "notification"
const DEFAULT_ERROR_MESSAGE = "interact with notification model"

func (service *Service) Get(inputJwtToken *types.JwtToken, id int64) (result *model.Notification, errCode int, err error) {
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

func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, request *data.GetAllRequest) (result []model.Notification, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, request)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
