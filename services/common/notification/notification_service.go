package notification

import (
	"fmt"
	"net/http"
	"time"

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

func (service *Service) UpdateSeen(
	ctxData *types.ContextData,
	id int64,
	request *data.NotificationSeenRequest,
) (result *model.Notification, errCode int, err error) {
	// Check if exists
	foundItem, err := service.Repository.GetByIDUserID(id, ctxData.Jwt.UserID)
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

	// Format request
	item := &model.Notification{
		UserID:  foundItem.UserID,
		Title:   foundItem.Title,
		Message: foundItem.Message,
		Seen:    request.Seen,
	}

	// Check seen date
	if request.Seen {
		seenAt := new(time.Time)
		*seenAt = time.Now()
		item.SeenAt = seenAt
	} else if item.Seen && !foundItem.Seen {
		item.SeenAt = nil
	}

	// Update
	result, err = service.Repository.UpdateSeenByIDUserID(id, ctxData.Jwt.UserID, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if result == nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) UpdateSeenAll(
	ctxData *types.ContextData,
	request *data.NotificationSeenAllRequest,
) (errCode int, err error) {
	// Check seen date
	var seenAt *time.Time
	if request.Seen {
		seenAt = new(time.Time)
		*seenAt = time.Now()
	}

	// Update
	err = service.Repository.UpdateSeenAllByUserID(ctxData.Jwt.UserID, request.Seen, seenAt)
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
	affectedRows, err = service.Repository.DeleteByIDUserID(id, ctxData.Jwt.UserID)
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

func (service *Service) DeleteAll(
	ctxData *types.ContextData,
) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteAllByUserID(ctxData.Jwt.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) Get(
	ctxData *types.ContextData,
	id int64,
) (result *model.Notification, errCode int, err error) {
	result, err = service.Repository.GetByIDUserID(id, ctxData.Jwt.UserID)
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

func (service *Service) GetNotSeenCount(
	ctxData *types.ContextData,
) (result int64, errCode int, err error) {
	result, err = service.Repository.GetNotSeenCountByUserID(ctxData.Jwt.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) GetAll(
	ctxData *types.ContextData,
	filter *types.Filter, pagination *types.Pagination,
	request *data.GetAllRequest,
) (result []model.Notification, errCode int, err error) {
	result, err = service.Repository.GetAllByUserID(filter, pagination, ctxData.Jwt.UserID)
	fmt.Println(err)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
