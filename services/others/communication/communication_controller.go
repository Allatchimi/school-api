package communication

import (
	"context"

	"api/common/helpers"
	httpHelper "api/common/helpers/http"
	"api/common/types"
	"api/services/others/communication/data"
	"api/services/others/communication/model"
)

type Controller struct {
	Service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{Service: service}
}

func (controller *Controller) Create(
	ctx *context.Context,
	input *struct {
		Body data.CommunicationRequest
	},
) (result *model.Communication, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		httpHelper.GetContextData(ctx),
		&input.Body,
	)
	return
}

func (controller *Controller) Get(
	ctx *context.Context,
	input *struct {
		data.CommunicationID
	},
) (result *model.Communication, errCode int, err error) {
	communication, errCode, err := controller.Service.Get(httpHelper.GetContextData(ctx), input.ID)
	if err != nil {
		return
	}
	result = communication
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllRequest
	},
) (result *data.CommunicationResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	communicationList, errCode, err := controller.Service.GetAll(httpHelper.GetContextData(ctx), newFilter, newPagination, &input.GetAllRequest)
	if err != nil {
		return
	}
	result = &data.CommunicationResponseList{
		Data: model.ToResponseList(communicationList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
