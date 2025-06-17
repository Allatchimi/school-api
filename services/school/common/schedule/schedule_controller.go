package schedule

import (
	"context"

	"api/common/helpers"
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
	if input.Body.IsGeneric {
		result, errCode, err = controller.Service.CreateGeneric(
			helpers.GetJwtContext(ctx),
			&input.Body,
		)
		return
	}
	result, errCode, err = controller.Service.Create(
		helpers.GetJwtContext(ctx),
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
	if input.Body.IsGeneric {
		result, errCode, err = controller.Service.UpdateGeneric(
			helpers.GetJwtContext(ctx),
			input.ID,
			&input.Body,
		)
		return
	}
	result, errCode, err = controller.Service.Update(
		helpers.GetJwtContext(ctx),
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
	affectedRows, errCode, err := controller.Service.Delete(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) DeleteGeneric(
	ctx *context.Context,
	input *struct {
		data.ScheduleID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteGeneric(helpers.GetJwtContext(ctx), input.ID)
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
	result, errCode, err = controller.Service.Get(helpers.GetJwtContext(ctx), input.ID)
	return
}

func (controller *Controller) GetGeneric(
	ctx *context.Context,
	input *struct {
		data.ScheduleID
	},
) (result *model.Schedule, errCode int, err error) {
	result, errCode, err = controller.Service.GetGeneric(helpers.GetJwtContext(ctx), input.ID)
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
	scheduleList, errCode, err := controller.Service.GetAll(helpers.GetJwtContext(ctx), newFilter, newPagination, &input.GetAllRequest)
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
	scheduleList, errCode, err := controller.Service.GetAll(helpers.GetJwtContext(ctx), newFilter, newPagination, &input.GetAllRequest)
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
