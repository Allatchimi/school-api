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

func (service *Service) Create(
	ctxData *types.ContextData,
	request *data.ExamRequest,
) (result *model.Exam, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Format request
	item := &model.Exam{
		SchoolID:       newRequest.SchoolID,
		YearID:         newRequest.YearID,
		ClassSubjectID: newRequest.ClassSubjectID,
		SequenceID:     newRequest.SequenceID,
		UnitID:         newRequest.UnitID,
		TypeID:         newRequest.TypeID,

		Status:          newRequest.Status,
		Notation:        newRequest.Notation,
		Percentage:      newRequest.Percentage,
		Description:     newRequest.Description,
		LocationType:    newRequest.LocationType,
		LocationDetails: newRequest.LocationDetails,
		Requirements:    newRequest.Requirements,
		AllowedItems:    newRequest.AllowedItems,
		StartDate:       newRequest.StartDate,
		EndDate:         newRequest.EndDate,

		IsRetry: newRequest.IsRetry,
	}
	if newRequest.ClassSubjectID < 1 && newRequest.UnitID < 1 {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessageV2("class subject id or unit id(you should provide one of these fields)")
		return
	}
	if newRequest.ClassSubjectID > 0 && newRequest.SequenceID < 1 {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessageV2("sequence id")
		return
	}

	// Check retry count
	var retryCount int64 = 0
	if newRequest.IsRetry {
		var notRetryCount int64 = 0
		notRetryCount, err = service.Repository.CountAllUniqueIsNotRetry(item)
		if err != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
		if notRetryCount < 1 {
			errCode = http.StatusLocked
			err = constants.Http423LockedErrorMessage()
			return
		}
		retryCount, err = service.Repository.CountAllUniqueIsRetry(item)
		if err != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
		retryCount++
	}
	item.RetryCount = retryCount

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

func (service *Service) CreateType(
	ctxData *types.ContextData,
	request *data.ExamTypeRequest,
) (result *model.ExamType, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Format request
	item := &model.ExamType{
		SchoolID:    newRequest.SchoolID,
		Name:        newRequest.Name,
		Description: newRequest.Description,
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

func (service *Service) Update(
	ctxData *types.ContextData,
	id int64,
	request *data.ExamRequest,
) (result *model.Exam, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Check if the item exists
	var foundItem *model.Exam
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetByIDSchoolID(id, newRequest.SchoolID)
	} else {
		foundItem, err = service.Repository.GetByID(id)
	}
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
	item := &model.Exam{
		SchoolID:       newRequest.SchoolID,
		YearID:         newRequest.YearID,
		ClassSubjectID: newRequest.ClassSubjectID,
		SequenceID:     newRequest.SequenceID,
		UnitID:         newRequest.UnitID,
		TypeID:         newRequest.TypeID,

		Status:          newRequest.Status,
		Notation:        newRequest.Notation,
		Percentage:      newRequest.Percentage,
		Description:     newRequest.Description,
		LocationType:    newRequest.LocationType,
		LocationDetails: newRequest.LocationDetails,
		Requirements:    newRequest.Requirements,
		AllowedItems:    newRequest.AllowedItems,
		StartDate:       newRequest.StartDate,
		EndDate:         newRequest.EndDate,

		IsRetry: newRequest.IsRetry,
	}
	if newRequest.ClassSubjectID < 1 && newRequest.UnitID < 1 {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessageV2("class subject id or unit id(you should provide one of these fields)")
		return
	}
	if newRequest.ClassSubjectID > 0 && newRequest.SequenceID < 1 {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessageV2("sequence id")
		return
	}

	// Check retry count
	var retryCount int64 = 0
	if newRequest.IsRetry {
		var notRetryCount int64 = 0
		notRetryCount, err = service.Repository.CountAllUniqueIsNotRetry(item)
		if err != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
		if notRetryCount < 1 {
			errCode = http.StatusLocked
			err = constants.Http423LockedErrorMessage()
			return
		}
		retryCount, err = service.Repository.CountAllUniqueIsRetry(item)
		if err != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
		retryCount++
	}
	item.RetryCount = retryCount

	// Check unique
	foundUnique, err := service.Repository.GetUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreSameUniqueObjects(foundUnique, item) &&
		!service.Repository.AreSameUniqueObjects(foundUnique, foundItem) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}

	// Update exam
	result, err = service.Repository.UpdateByID(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) UpdateType(
	ctxData *types.ContextData,
	id int64,
	request *data.ExamTypeRequest,
) (result *model.ExamType, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Check if the item exists
	var foundItem *model.ExamType
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetExamTypeByIDSchoolID(id, newRequest.SchoolID)
	} else {
		foundItem, err = service.Repository.GetExamTypeByID(id)
	}
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
	item := &model.ExamType{
		SchoolID:    newRequest.SchoolID,
		Name:        newRequest.Name,
		Description: newRequest.Description,
	}

	// Check unique
	foundUnique, err := service.Repository.GetExamTypeUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreSameExamTypeUniqueObjects(foundUnique, item) &&
		!service.Repository.AreSameExamTypeUniqueObjects(foundUnique, foundItem) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}

	// Update exam
	result, err = service.Repository.UpdateExamTypeByID(id, item)
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
	var foundItem *model.Exam
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetByIDSchoolID(id, ctxData.Jwt.SchoolID)
	} else {
		foundItem, err = service.Repository.GetByID(id)
	}
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

	// Delete
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

func (service *Service) DeleteType(
	ctxData *types.ContextData,
	id int64,
) (affectedRows int64, errCode int, err error) {
	// Check if the item exists
	var foundItem *model.ExamType
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetExamTypeByIDSchoolID(id, ctxData.Jwt.SchoolID)
	} else {
		foundItem, err = service.Repository.GetExamTypeByID(id)
	}
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

	// Delete
	affectedRows, err = service.Repository.DeleteExamTypeByID(id)
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
	ctxData *types.ContextData,
	list []int64,
) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultipleByID(list, ctxData.Jwt.SchoolID)
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

func (service *Service) DeleteMultipleType(
	ctxData *types.ContextData,
	list []int64,
) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultipleExamTypeByID(list, ctxData.Jwt.SchoolID)
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
	ctxData *types.ContextData,
	id int64,
) (result *model.Exam, errCode int, err error) {
	if ctxData.User.Feature != constants.FeatureAdmin {
		result, err = service.Repository.GetByIDSchoolID(id, ctxData.Jwt.SchoolID)
	} else {
		result, err = service.Repository.GetByID(id)
	}
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

func (service *Service) GetType(
	ctxData *types.ContextData,
	id int64,
) (result *model.ExamType, errCode int, err error) {
	if ctxData.User.Feature != constants.FeatureAdmin {
		result, err = service.Repository.GetExamTypeByIDSchoolID(id, ctxData.Jwt.SchoolID)
	} else {
		result, err = service.Repository.GetExamTypeByID(id)
	}
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
	ctxData *types.ContextData,
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllRequest,
) (result []model.Exam, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Get
	result, err = service.Repository.GetAll(filter, pagination, &newRequest)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}

func (service *Service) GetAllExamType(
	ctxData *types.ContextData,
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllExamTypeRequest,
) (result []model.ExamType, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Get
	result, err = service.Repository.GetAllExamType(filter, pagination, &newRequest)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
