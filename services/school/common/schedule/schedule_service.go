package schedule

import (
	"net/http"

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

func (service *Service) Create(inputJwtToken *types.JwtToken, request *data.ScheduleRequest) (schedule *model.Schedule, errCode int, err error) {
	// Format request
	item := &model.Schedule{
		SchoolID:       request.SchoolID,
		YearID:         request.YearID,
		ClassSubjectID: request.ClassSubjectID,
		UnitID:         request.UnitID,
		Type:           request.Type,
		DayOfTheWeek:   request.DayOfTheWeek,
		RepeatCount:    request.RepeatCount,
		RepeatType:     request.RepeatType,
		StartTime:      request.StartTime,
		EndTime:        request.EndTime,
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

	// Insert schedule
	schedule, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) Update(inputJwtToken *types.JwtToken, id int64, request *data.ScheduleRequest) (schedule *model.Schedule, errCode int, err error) {
	// Format request
	item := &model.Schedule{
		SchoolID:       request.SchoolID,
		YearID:         request.YearID,
		ClassSubjectID: request.ClassSubjectID,
		UnitID:         request.UnitID,
		Type:           request.Type,
		DayOfTheWeek:   request.DayOfTheWeek,
		RepeatCount:    request.RepeatCount,
		RepeatType:     request.RepeatType,
		StartTime:      request.StartTime,
		EndTime:        request.EndTime,
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

	// Update schedule
	schedule, err = service.Repository.Update(id, item)
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

func (service *Service) Get(inputJwtToken *types.JwtToken, scheduleID int64) (schedule *model.Schedule, errCode int, err error) {
	schedule, err = service.Repository.GetByID(scheduleID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if schedule == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	return
}

func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, request *data.GetAllRequest) (schedule []model.Schedule, errCode int, err error) {
	schedule, err = service.Repository.GetAll(filter, pagination, request)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
