package exam

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/common/exam/data"
	"api/services/school/common/exam/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

const MODEL_NAME = "exam/type"
const DEFAULT_ERROR_MESSAGE = "interact with exam/type model"

func (service *Service) Create(ctxData *types.ContextData, request *data.ExamRequest) (result *model.Exam, errCode int, err error) {
	// Format request
	item := &model.Exam{
		SchoolID:       request.SchoolID,
		YearID:         request.YearID,
		TypeID:         request.TypeID,
		UnitID:         request.UnitID,
		ClassSubjectID: request.ClassSubjectID,
		SequenceID:     request.SequenceID,

		Status:          request.Status,
		Percentage:      request.Percentage,
		Description:     request.Description,
		LocationType:    request.LocationType,
		LocationDetails: request.LocationDetails,
		Requirements:    request.Requirements,
		AllowedItems:    request.AllowedItems,
		StartDate:       request.StartDate,
		EndDate:         request.EndDate,
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

	// Insert exam
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) CreateType(ctxData *types.ContextData, request *data.ExamTypeRequest) (result *model.ExamType, errCode int, err error) {
	// Format request
	item := &model.ExamType{
		SchoolID:    request.SchoolID,
		Name:        request.Name,
		Description: request.Description,
	}

	// Check unique
	foundUnique, err := service.Repository.GetExamTypeUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreSameExamTypeUniqueObjects(foundUnique, item) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}

	// Insert exam
	result, err = service.Repository.CreateType(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) Update(ctxData *types.ContextData, id int64, request *data.ExamRequest) (result *model.Exam, errCode int, err error) {
	// Check if exam already exists
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

	// Format request
	item := &model.Exam{
		SchoolID:       request.SchoolID,
		YearID:         request.YearID,
		TypeID:         request.TypeID,
		UnitID:         request.UnitID,
		ClassSubjectID: request.ClassSubjectID,
		SequenceID:     request.SequenceID,

		Status:          request.Status,
		Percentage:      request.Percentage,
		Description:     request.Description,
		LocationType:    request.LocationType,
		LocationDetails: request.LocationDetails,
		Requirements:    request.Requirements,
		AllowedItems:    request.AllowedItems,
		StartDate:       request.StartDate,
		EndDate:         request.EndDate,
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

	// Update exam
	result, err = service.Repository.Update(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) UpdateType(ctxData *types.ContextData, id int64, request *data.ExamTypeRequest) (result *model.ExamType, errCode int, err error) {
	// Check if exam already exists
	foundItem, err := service.Repository.GetTypeByID(id)
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

	// Format request
	item := &model.ExamType{
		SchoolID:    request.SchoolID,
		Name:        request.Name,
		Description: request.Description,
	}

	// Check unique
	foundUnique, err := service.Repository.GetExamTypeUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreSameExamTypeUniqueObjects(foundUnique, item) && !service.Repository.AreSameExamTypeUniqueObjects(foundUnique, foundItem) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}

	// Update exam
	result, err = service.Repository.UpdateType(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) Delete(ctxData *types.ContextData, id int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteByID(id)
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

func (service *Service) DeleteType(ctxData *types.ContextData, id int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteTypeByID(id)
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

func (service *Service) DeleteMultiple(ctxData *types.ContextData, list []int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultipleByID(list)
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

func (service *Service) Get(ctxData *types.ContextData, examID int64) (result *model.Exam, errCode int, err error) {
	result, err = service.Repository.GetByID(examID)
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

func (service *Service) GetType(ctxData *types.ContextData, examID int64) (result *model.ExamType, errCode int, err error) {
	result, err = service.Repository.GetTypeByID(examID)
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

func (service *Service) GetAll(ctxData *types.ContextData, filter *types.Filter, pagination *types.Pagination, request *data.GetAllRequest) (result []model.Exam, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, request)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}

func (service *Service) GetAllExamType(ctxData *types.ContextData, filter *types.Filter, pagination *types.Pagination, request *data.GetAllExamTypeRequest) (result []model.ExamType, errCode int, err error) {
	result, err = service.Repository.GetAllExamType(filter, pagination, request)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
