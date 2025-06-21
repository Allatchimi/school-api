package schedule

import (
	"net/http"
	"time"

	"api/common/constants"
	"api/common/types"
	"api/services/school/common/schedule/data"
	"api/services/school/common/schedule/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

const MODEL_NAME = "schedule"
const DEFAULT_ERROR_MESSAGE = "interact with schedule model"

func (service *Service) Create(inputJwtToken *types.JwtToken, request *data.ScheduleRequest) (result *model.Schedule, errCode int, err error) {
	// Format request
	item := &model.Schedule{
		SchoolID:       request.SchoolID,
		YearID:         request.YearID,
		ClassSubjectID: request.ClassSubjectID,
		UnitID:         request.UnitID,

		Type:         request.Type,
		DayOfTheWeek: request.DayOfTheWeek,
		RepeatCount:  request.RepeatCount,
		RepeatType:   request.RepeatType,
		StartTime:    request.StartTime,
		EndTime:      request.EndTime,
		IsValid:      request.IsValid,
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

	// Check invalid date
	if !item.IsValid {
		invalidDate := new(time.Time)
		*invalidDate = time.Now()
		item.InvalidDate = invalidDate
	}

	// Insert schedule
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) CreateGeneric(inputJwtToken *types.JwtToken, request *data.ScheduleRequest) (result *model.Schedule, errCode int, err error) {
	// Format request
	item := &model.ScheduleGeneric{
		SchoolID: request.SchoolID,
		YearID:   request.YearID,

		Type:         request.Type,
		DayOfTheWeek: request.DayOfTheWeek,
		RepeatCount:  request.RepeatCount,
		RepeatType:   request.RepeatType,
		StartTime:    request.StartTime,
		EndTime:      request.EndTime,
		IsValid:      request.IsValid,
	}

	// Check unique
	foundUnique, err := service.Repository.GetScheduleGenericUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreScheduleGenericSameUniqueObjects(foundUnique, item) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}

	// Check invalid date
	if !item.IsValid {
		invalidDate := new(time.Time)
		*invalidDate = time.Now()
		item.InvalidDate = invalidDate
	}

	// Insert schedule
	resultGeneric, err := service.Repository.CreateScheduleGeneric(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Cast to Schedule
	result = &model.Schedule{
		SchoolID: resultGeneric.SchoolID,
		School:   resultGeneric.School,
		YearID:   resultGeneric.YearID,
		Year:     resultGeneric.Year,

		Type:         resultGeneric.Type,
		DayOfTheWeek: resultGeneric.DayOfTheWeek,
		RepeatCount:  resultGeneric.RepeatCount,
		RepeatType:   resultGeneric.RepeatType,
		StartTime:    resultGeneric.StartTime,
		EndTime:      resultGeneric.EndTime,
		IsValid:      resultGeneric.IsValid,
		InvalidDate:  resultGeneric.InvalidDate,
	}
	result.ID = resultGeneric.ID
	result.CreatedAt = resultGeneric.CreatedAt
	result.UpdatedAt = resultGeneric.UpdatedAt
	return
}

func (service *Service) Update(inputJwtToken *types.JwtToken, id int64, request *data.ScheduleRequest) (result *model.Schedule, errCode int, err error) {
	// Format request
	item := &model.Schedule{
		SchoolID:       request.SchoolID,
		YearID:         request.YearID,
		ClassSubjectID: request.ClassSubjectID,
		UnitID:         request.UnitID,

		Type:         request.Type,
		DayOfTheWeek: request.DayOfTheWeek,
		RepeatCount:  request.RepeatCount,
		RepeatType:   request.RepeatType,
		StartTime:    request.StartTime,
		EndTime:      request.EndTime,
		IsValid:      request.IsValid,
	}

	// Check if schedule already exists
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

	// Check invalid date
	if !item.IsValid && foundItem.IsValid {
		invalidDate := new(time.Time)
		*invalidDate = time.Now()
		item.InvalidDate = invalidDate
	} else if item.IsValid && !foundItem.IsValid {
		item.InvalidDate = nil
	}

	// Update schedule
	result, err = service.Repository.Update(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) UpdateGeneric(inputJwtToken *types.JwtToken, id int64, request *data.ScheduleRequest) (result *model.Schedule, errCode int, err error) {
	// Format request
	item := &model.ScheduleGeneric{
		SchoolID: request.SchoolID,
		YearID:   request.YearID,

		Type:         request.Type,
		DayOfTheWeek: request.DayOfTheWeek,
		RepeatCount:  request.RepeatCount,
		RepeatType:   request.RepeatType,
		StartTime:    request.StartTime,
		EndTime:      request.EndTime,
		IsValid:      request.IsValid,
	}

	// Check if schedule already exists
	foundItem, err := service.Repository.GetScheduleGenericByID(id)
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
	foundUnique, err := service.Repository.GetScheduleGenericUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreScheduleGenericSameUniqueObjects(foundUnique, item) && !service.Repository.AreScheduleGenericSameUniqueObjects(foundUnique, foundItem) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}

	// Check invalid date
	if !item.IsValid && foundItem.IsValid {
		invalidDate := new(time.Time)
		*invalidDate = time.Now()
		item.InvalidDate = invalidDate
	} else if item.IsValid && !foundItem.IsValid {
		item.InvalidDate = nil
	}

	// Update schedule
	resultGeneric, err := service.Repository.UpdateScheduleGeneric(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Cast to Schedule
	result = &model.Schedule{
		SchoolID: resultGeneric.SchoolID,
		School:   resultGeneric.School,
		YearID:   resultGeneric.YearID,
		Year:     resultGeneric.Year,

		Type:         resultGeneric.Type,
		DayOfTheWeek: resultGeneric.DayOfTheWeek,
		RepeatCount:  resultGeneric.RepeatCount,
		RepeatType:   resultGeneric.RepeatType,
		StartTime:    resultGeneric.StartTime,
		EndTime:      resultGeneric.EndTime,
		IsValid:      resultGeneric.IsValid,
		InvalidDate:  resultGeneric.InvalidDate,
	}
	result.ID = resultGeneric.ID
	result.CreatedAt = resultGeneric.CreatedAt
	result.UpdatedAt = resultGeneric.UpdatedAt
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

func (service *Service) DeleteGeneric(inputJwtToken *types.JwtToken, id int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteScheduleGeneric(id)
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

func (service *Service) Get(inputJwtToken *types.JwtToken, id int64) (result *model.Schedule, errCode int, err error) {
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

func (service *Service) GetGeneric(inputJwtToken *types.JwtToken, id int64) (result *model.Schedule, errCode int, err error) {
	resultGeneric, err := service.Repository.GetScheduleGenericByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if resultGeneric == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	result = model.ConvertScheduleGenericToSchedule(resultGeneric)
	return
}

func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, request *data.GetAllRequest) (result []model.Schedule, errCode int, err error) {
	// Get default schedules
	if request.Type == "default" {
		resultDefault, errDefault := service.Repository.GetAll(filter, pagination, request)
		if errDefault != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
		result = resultDefault
		return
	}
	// Get generic schedules
	if request.Type == "generic" {
		resultGeneric, errGeneric := service.Repository.GetAllScheduleGeneric(filter, pagination, request)
		if errGeneric != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
		resultDefault := make([]model.Schedule, 0)
		result = model.ListAppendGenericSchedules(resultDefault, resultGeneric)
		return
	}

	// Get all schedules
	resultDefault, err := service.Repository.GetAll(filter, pagination, request)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	resultGeneric, err := service.Repository.GetAllScheduleGeneric(filter, pagination, request)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	result = model.ListAppendGenericSchedules(resultDefault, resultGeneric)
	return
}
