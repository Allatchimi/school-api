package manager

import (
	"context"

	"api/common/helpers"
	httpHelper "api/common/helpers/http"
	"api/common/types"
	"api/services/school/common/manager/data"
	"api/services/school/common/manager/model"
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
		Body data.ManagerRequest
	},
) (result *model.Manager, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		httpHelper.GetContextData(ctx),
		&input.Body,
	)
	return
}

func (controller *Controller) Update(
	ctx *context.Context,
	input *struct {
		data.ManagerID
		Body data.ManagerRequest
	},
) (result *model.Manager, errCode int, err error) {
	result, errCode, err = controller.Service.Update(
		httpHelper.GetContextData(ctx),
		input.ID,
		&input.Body,
	)
	return
}

func (controller *Controller) Delete(
	ctx *context.Context,
	input *struct {
		data.ManagerID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.Delete(httpHelper.GetContextData(ctx), input.ID)
	if err != nil {
		return
	}
	result = affectedRows
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
		data.ManagerID
	},
) (result *model.Manager, errCode int, err error) {
	manager, errCode, err := controller.Service.Get(httpHelper.GetContextData(ctx), input.ID)
	if err != nil {
		return
	}
	result = manager
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllRequest
	},
) (result *data.ManagerResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	managerList, errCode, err := controller.Service.GetAll(
		httpHelper.GetContextData(ctx),
		newFilter,
		newPagination,
		&input.GetAllRequest,
	)
	if err != nil {
		return
	}
	result = &data.ManagerResponseList{
		Data: model.ToManagerResponseList(managerList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
