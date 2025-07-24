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

func (service *Service) Create(ctxData *types.ContextData, request *data.ScheduleRequest) (result *model.Schedule, errCode int, err error) {
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

func (service *Service) CreateCommon(ctxData *types.ContextData, request *data.ScheduleRequest) (result *model.Schedule, errCode int, err error) {
	// Format request
	item := &model.ScheduleCommon{
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
	foundUnique, err := service.Repository.GetScheduleCommonUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreScheduleCommonSameUniqueObjects(foundUnique, item) {
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
	resultCommon, err := service.Repository.CreateScheduleCommon(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Cast to Schedule
	result = &model.Schedule{
		SchoolID: resultCommon.SchoolID,
		School:   resultCommon.School,
		YearID:   resultCommon.YearID,
		Year:     resultCommon.Year,

		Type:         resultCommon.Type,
		DayOfTheWeek: resultCommon.DayOfTheWeek,
		RepeatCount:  resultCommon.RepeatCount,
		RepeatType:   resultCommon.RepeatType,
		StartTime:    resultCommon.StartTime,
		EndTime:      resultCommon.EndTime,
		IsValid:      resultCommon.IsValid,
		InvalidDate:  resultCommon.InvalidDate,
	}
	result.ID = resultCommon.ID
	result.CreatedAt = resultCommon.CreatedAt
	result.UpdatedAt = resultCommon.UpdatedAt
	return
}

func (service *Service) Update(ctxData *types.ContextData, id int64, request *data.ScheduleRequest) (result *model.Schedule, errCode int, err error) {
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

func (service *Service) UpdateCommon(ctxData *types.ContextData, id int64, request *data.ScheduleRequest) (result *model.Schedule, errCode int, err error) {
	// Format request
	item := &model.ScheduleCommon{
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
	foundItem, err := service.Repository.GetScheduleCommonByID(id)
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
	foundUnique, err := service.Repository.GetScheduleCommonUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreScheduleCommonSameUniqueObjects(foundUnique, item) && !service.Repository.AreScheduleCommonSameUniqueObjects(foundUnique, foundItem) {
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
	resultCommon, err := service.Repository.UpdateScheduleCommon(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Cast to Schedule
	result = &model.Schedule{
		SchoolID: resultCommon.SchoolID,
		School:   resultCommon.School,
		YearID:   resultCommon.YearID,
		Year:     resultCommon.Year,

		Type:         resultCommon.Type,
		DayOfTheWeek: resultCommon.DayOfTheWeek,
		RepeatCount:  resultCommon.RepeatCount,
		RepeatType:   resultCommon.RepeatType,
		StartTime:    resultCommon.StartTime,
		EndTime:      resultCommon.EndTime,
		IsValid:      resultCommon.IsValid,
		InvalidDate:  resultCommon.InvalidDate,
	}
	result.ID = resultCommon.ID
	result.CreatedAt = resultCommon.CreatedAt
	result.UpdatedAt = resultCommon.UpdatedAt
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

func (service *Service) DeleteCommon(ctxData *types.ContextData, id int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteScheduleCommonByID(id)
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

func (service *Service) Get(ctxData *types.ContextData, id int64) (result *model.Schedule, errCode int, err error) {
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

func (service *Service) GetCommon(ctxData *types.ContextData, id int64) (result *model.Schedule, errCode int, err error) {
	resultCommon, err := service.Repository.GetScheduleCommonByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if resultCommon == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	result = model.ConvertScheduleCommonToSchedule(resultCommon)
	return
}

func (service *Service) GetAll(ctxData *types.ContextData, filter *types.Filter, pagination *types.Pagination, request *data.GetAllRequest) (result []model.Schedule, errCode int, err error) {
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
	// Get common schedules
	if request.Type == "common" {
		resultCommon, errCommon := service.Repository.GetAllScheduleCommon(filter, pagination, request)
		if errCommon != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
		resultDefault := make([]model.Schedule, 0)
		result = model.ListAppendCommonSchedules(resultDefault, resultCommon)
		return
	}

	// Get all schedules
	resultDefault, err := service.Repository.GetAll(filter, pagination, request)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	resultCommon, err := service.Repository.GetAllScheduleCommon(filter, pagination, request)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	result = model.ListAppendCommonSchedules(resultDefault, resultCommon)
	return
}
