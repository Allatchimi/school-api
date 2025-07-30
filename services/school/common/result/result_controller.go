package result

import (
	"context"

	"api/common/helpers"
	httpHelper "api/common/helpers/http"
	"api/common/types"
	"api/services/school/common/result/data"
	"api/services/school/common/result/model"
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
		Body data.ResultRequest
	},
) (result *model.Result, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		httpHelper.GetContextData(ctx),
		&input.Body,
	)
	return
}

func (controller *Controller) CreateTable(
	ctx *context.Context,
	input *struct {
		Body data.ResultTableRequest
	},
) (result *model.ResultTable, errCode int, err error) {
	result, errCode, err = controller.Service.CreateTable(
		httpHelper.GetContextData(ctx),
		&input.Body,
	)
	return
}

func (controller *Controller) Update(
	ctx *context.Context,
	input *struct {
		data.ResultID
		Body data.ResultRequest
	},
) (result *model.Result, errCode int, err error) {
	result, errCode, err = controller.Service.Update(
		httpHelper.GetContextData(ctx),
		input.ID,
		&input.Body,
	)
	return
}

func (controller *Controller) UpdateTable(
	ctx *context.Context,
	input *struct {
		data.ResultTableID
		Body data.ResultTableRequest
	},
) (result *model.ResultTable, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateTable(
		httpHelper.GetContextData(ctx),
		input.ID,
		&input.Body,
	)
	return
}

func (controller *Controller) DeleteTable(
	ctx *context.Context,
	input *struct {
		data.ResultTableID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteTable(httpHelper.GetContextData(ctx), input.ID)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) Delete(
	ctx *context.Context,
	input *struct {
		data.ResultID
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

func (controller *Controller) DeleteMultipleTable(
	ctx *context.Context,
	input *struct {
		Body types.DeleteMultipleRequest
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteMultipleTable(httpHelper.GetContextData(ctx), input.Body.List)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) GetTable(
	ctx *context.Context,
	input *struct {
		data.ResultTableID
	},
) (result *model.ResultTable, errCode int, err error) {
	result, errCode, err = controller.Service.GetTable(httpHelper.GetContextData(ctx), input.ID)
	return
}

func (controller *Controller) Get(
	ctx *context.Context,
	input *struct {
		data.ResultID
	},
) (result *model.Result, errCode int, err error) {
	result, errCode, err = controller.Service.Get(httpHelper.GetContextData(ctx), input.ID)
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllRequest
	},
) (result *data.ResultResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	resultList, errCode, err := controller.Service.GetAll(httpHelper.GetContextData(ctx), newFilter, newPagination, &input.GetAllRequest)
	if err != nil {
		return
	}
	result = &data.ResultResponseList{
		Data: model.ToResultResponseList(resultList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}

func (controller *Controller) GetAllTable(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllResultTableRequest
	},
) (result *data.ResultTableResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	resultList, errCode, err := controller.Service.GetAllTable(httpHelper.GetContextData(ctx), newFilter, newPagination, &input.GetAllResultTableRequest)
	if err != nil {
		return
	}
	model.ToResultTableResponseList(resultList)
	result = &data.ResultTableResponseList{
		Data: model.ToResultTableResponseList(resultList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
