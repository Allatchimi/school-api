package user

import (
	"context"

	"api/common/helpers"
	httpHelper "api/common/helpers/http"
	"api/common/types"
	"api/services/user/user/data"
	"api/services/user/user/model"
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
		Body data.UserRequest
	},
) (result *model.User, errCode int, err error) {

	result, errCode, err = controller.Service.Create(
		httpHelper.GetContextData(ctx),
		&input.Body,
		nil,
	)
	return
}

func (controller *Controller) Update(
	ctx *context.Context,
	input *struct {
		data.UserID
		Body data.UserRequest
	},
) (result *model.User, errCode int, err error) {
	result, errCode, err = controller.Service.Update(
		httpHelper.GetContextData(ctx),
		input.UserID.ID,
		&input.Body,
	)
	return
}

func (controller *Controller) Delete(
	ctx *context.Context,
	input *struct {
		data.UserID
	},
) (result int64, errCode int, err error) {
	result, errCode, err = controller.Service.Delete(
		httpHelper.GetContextData(ctx),
		input.ID,
	)
	return
}

func (controller *Controller) DeleteMultiple(
	ctx *context.Context,
	input *struct {
		Body types.DeleteMultipleRequest
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteMultiple(httpHelper.GetContextData(ctx), input.Body.List)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) Get(
	ctx *context.Context,
	input *struct {
		data.UserID
	},
) (result *model.User, errCode int, err error) {
	result, errCode, err = controller.Service.Get(
		httpHelper.GetContextData(ctx),
		input.ID,
	)
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllRequest
	},
) (result *data.UserResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	userList, errCode, err := controller.Service.GetAll(
		httpHelper.GetContextData(ctx),
		newFilter,
		newPagination,
		&input.GetAllRequest,
	)
	if err != nil {
		return
	}
	result = &data.UserResponseList{
		Data: model.ToResponseList(userList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
