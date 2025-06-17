package payment

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/common/payment/data"
	"api/services/school/common/payment/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

const MODEL_NAME = "payment"
const DEFAULT_ERROR_MESSAGE = "interact with payment model"

func (service *Service) Create(
	inputJwtToken *types.JwtToken,
	request *data.PaymentRequest,
) (result *model.Payment, errCode int, err error) {
	// Format request
	item := &model.Payment{
		SchoolID:        request.SchoolID,
		StudentEnrollID: request.StudentEnrollID,

		Amount:        request.Amount,
		Currency:      request.Currency,
		PaymentDate:   request.PaymentDate,
		PaymentMethod: request.PaymentMethod,
		PaymentStatus: request.PaymentStatus,
		PaymentNote:   request.PaymentNote,
	}

	// Create
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) Update(
	inputJwtToken *types.JwtToken,
	id int64,
	request *data.PaymentRequest,
) (result *model.Payment, errCode int, err error) {
	// Check if payment exists
	foundItem, err := service.Repository.GetByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil || foundItem.ID < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Format request
	item := &model.Payment{
		SchoolID:        request.SchoolID,
		StudentEnrollID: request.StudentEnrollID,

		Amount:        request.Amount,
		Currency:      request.Currency,
		PaymentDate:   request.PaymentDate,
		PaymentMethod: request.PaymentMethod,
		PaymentStatus: request.PaymentStatus,
		PaymentNote:   request.PaymentNote,
	}

	// Update payment
	result, err = service.Repository.Update(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) Delete(
	inputJwtToken *types.JwtToken,
	id int64,
) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	return
}

func (service *Service) DeleteMultiple(
	inputJwtToken *types.JwtToken,
	list []int64,
) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultiple(list)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	return
}

func (service *Service) Get(
	inputJwtToken *types.JwtToken,
	id int64,
) (result *model.Payment, errCode int, err error) {
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

func (service *Service) GetAll(
	inputJwtToken *types.JwtToken,
	filter *types.Filter,
	pagination *types.Pagination,
	schoolID int64,
) (result []model.Payment, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, &data.GetAllRequest{SchoolID: schoolID})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
