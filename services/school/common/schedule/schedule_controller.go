package schedule

import (
	"context"

	"api/common/helpers"
	httpHelper "api/common/helpers/http"
	"api/common/types"
	"api/services/school/common/schedule/data"
	"api/services/school/common/schedule/model"
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
		Body data.ScheduleRequest
	},
) (result *model.Schedule, errCode int, err error) {
	if input.Body.IsCommon {
		result, errCode, err = controller.Service.CreateCommon(
			httpHelper.GetContextData(ctx),
			&input.Body,
		)
		return
	}
	result, errCode, err = controller.Service.Create(
		httpHelper.GetContextData(ctx),
		&input.Body,
	)
	return
}

func (controller *Controller) Update(
	ctx *context.Context,
	input *struct {
		data.ScheduleID
		Body data.ScheduleRequest
	},
) (result *model.Schedule, errCode int, err error) {
	if input.Body.IsCommon {
		result, errCode, err = controller.Service.UpdateCommon(
			httpHelper.GetContextData(ctx),
			input.ID,
			&input.Body,
		)
		return
	}
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
		data.ScheduleID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.Delete(httpHelper.GetContextData(ctx), input.ID)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) DeleteCommon(
	ctx *context.Context,
	input *struct {
		data.ScheduleID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteCommon(httpHelper.GetContextData(ctx), input.ID)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) Get(
	ctx *context.Context,
	input *struct {
		data.ScheduleID
	},
) (result *model.Schedule, errCode int, err error) {
	result, errCode, err = controller.Service.Get(httpHelper.GetContextData(ctx), input.ID)
	return
}

func (controller *Controller) GetCommon(
	ctx *context.Context,
	input *struct {
		data.ScheduleID
	},
) (result *model.Schedule, errCode int, err error) {
	result, errCode, err = controller.Service.GetCommon(httpHelper.GetContextData(ctx), input.ID)
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllRequest
	},
) (result *data.ScheduleResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	scheduleList, errCode, err := controller.Service.GetAll(httpHelper.GetContextData(ctx), newFilter, newPagination, &input.GetAllRequest)
	if err != nil {
		return
	}
	result = &data.ScheduleResponseList{
		Data: model.ToScheduleResponseList(scheduleList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}

func (controller *Controller) GetAllWeeklyView(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllRequest
	},
) (result *data.ScheduleWeeklyViewResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	scheduleList, errCode, err := controller.Service.GetAll(httpHelper.GetContextData(ctx), newFilter, newPagination, &input.GetAllRequest)
	if err != nil {
		return
	}
	result = &data.ScheduleWeeklyViewResponseList{
		Data: model.ToScheduleWeeklyViewResponseList(scheduleList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
