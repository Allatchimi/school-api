package request

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/common/request/data"
	"api/services/school/common/request/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

const MODEL_NAME = "request"
const DEFAULT_ERROR_MESSAGE = "interact with request model"

func (service *Service) Create(inputJwtToken *types.JwtToken, request *data.RequestRequest) (result *model.Request, errCode int, err error) {
	// Format request
	item := &model.Request{
		SchoolID:       request.SchoolID,
		YearID:         request.YearID,
		ClassSubjectID: request.ClassSubjectID,
		SequenceID:     request.SequenceID,
		UnitID:         request.UnitID,
		StudentID:      request.StudentID,

		Status:  request.Status,
		Type:    request.Type,
		Title:   request.Title,
		Message: request.Message,

		Document1: request.Document1,
		Document2: request.Document2,
		Document3: request.Document3,
		Document4: request.Document4,
		Document5: request.Document5,
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

	// Insert request
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) Update(inputJwtToken *types.JwtToken, id int64, request *data.RequestRequest) (result *model.Request, errCode int, err error) {
	// Format request
	item := &model.Request{
		SchoolID:       request.SchoolID,
		YearID:         request.YearID,
		ClassSubjectID: request.ClassSubjectID,
		SequenceID:     request.SequenceID,
		UnitID:         request.UnitID,
		StudentID:      request.StudentID,

		Status:  request.Status,
		Type:    request.Type,
		Title:   request.Title,
		Message: request.Message,

		Document1: request.Document1,
		Document2: request.Document2,
		Document3: request.Document3,
		Document4: request.Document4,
		Document5: request.Document5,
	}

	// Check if request already exists
	foundItem, err := service.Repository.GetByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
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

	// Update request
	result, err = service.Repository.Update(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) Delete(inputJwtToken *types.JwtToken, id int64) (affectedRows int64, errCode int, err error) {
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

func (service *Service) Get(inputJwtToken *types.JwtToken, requestID int64) (result *model.Request, errCode int, err error) {
	result, err = service.Repository.GetByID(requestID)
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

func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, request *data.GetAllRequest) (result []model.Request, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, request)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
