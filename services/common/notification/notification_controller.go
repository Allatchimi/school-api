package notification

import (
	"context"

	"api/common/helpers"
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
	notification, errCode, err := controller.Service.Get(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = notification
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
	notificationList, errCode, err := controller.Service.GetAll(helpers.GetJwtContext(ctx), newFilter, newPagination, &input.GetAllRequest)
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
