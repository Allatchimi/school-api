package result

import (
	"context"

	"api/common/helpers"
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
		helpers.GetJwtContext(ctx),
		&model.Result{
			StudentID: input.Body.StudentID,
			ExamID:    input.Body.ExamID,

			Value: input.Body.Value,
		},
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
		helpers.GetJwtContext(ctx), input.ID,
		&model.Result{
			StudentID: input.Body.StudentID,
			ExamID:    input.Body.ExamID,

			Value: input.Body.Value,
		},
	)
	return
}

func (controller *Controller) Delete(
	ctx *context.Context,
	input *struct {
		data.ResultID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.Delete(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) Get(
	ctx *context.Context,
	input *struct {
		data.ResultID
	},
) (result *model.Result, errCode int, err error) {
	result, errCode, err = controller.Service.Get(helpers.GetJwtContext(ctx), input.ID)
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
	resultList, errCode, err := controller.Service.GetAll(helpers.GetJwtContext(ctx), newFilter, newPagination, &input.GetAllRequest)
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
