package notification

import (
	"context"

	"api/common/helpers"
	httpHelper "api/common/helpers/http"
	"api/common/types"
	"api/services/common/notification/data"
	"api/services/common/notification/model"
)

type Controller struct {
	Service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{Service: service}
}

func (controller *Controller) Get(
	ctx *context.Context,
	input *struct {
		data.NotificationID
	},
) (result *model.Notification, errCode int, err error) {
	notification, errCode, err := controller.Service.Get(httpHelper.GetContextData(ctx), input.ID)
	if err != nil {
		return
	}
	result = notification
	return
}

func (controller *Controller) GetNotSeenCount(
	ctx *context.Context,
	input *struct {
		data.NotificationNotSeenCountRequest
	},
) (result int64, errCode int, err error) {
	count, errCode, err := controller.Service.GetNotSeenCount(httpHelper.GetContextData(ctx))
	if err != nil {
		return
	}
	result = count
	return
}

func (controller *Controller) UpdateSeen(
	ctx *context.Context,
	input *struct {
		data.NotificationID
		Body data.NotificationSeenRequest
	},
) (result *model.Notification, errCode int, err error) {
	notification, errCode, err := controller.Service.UpdateSeen(httpHelper.GetContextData(ctx), input.ID, &input.Body)
	if err != nil {
		return
	}
	result = notification
	return
}

func (controller *Controller) UpdateSeenAll(
	ctx *context.Context,
	input *struct {
		Body data.NotificationSeenAllRequest
	},
) (errCode int, err error) {
	errCode, err = controller.Service.UpdateSeenAll(httpHelper.GetContextData(ctx), &input.Body)
	if err != nil {
		return
	}
	return
}

func (controller *Controller) Delete(
	ctx *context.Context,
	input *struct {
		data.NotificationID
	},
) (result int64, errCode int, err error) {
	result, errCode, err = controller.Service.Delete(httpHelper.GetContextData(ctx), input.ID)
	if err != nil {
		return
	}
	return
}

func (controller *Controller) DeleteAll(
	ctx *context.Context,
	input *struct{},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteAll(httpHelper.GetContextData(ctx))
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllRequest
	},
) (result *data.NotificationResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	notificationList, errCode, err := controller.Service.GetAll(httpHelper.GetContextData(ctx), newFilter, newPagination, &input.GetAllRequest)
	if err != nil {
		return
	}
	result = &data.NotificationResponseList{
		Data: model.ToResponseList(notificationList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
