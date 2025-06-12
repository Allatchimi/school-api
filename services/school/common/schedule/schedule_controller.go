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
) (schedule *model.Schedule, errCode int, err error) {
	schedule, errCode, err = controller.Service.Create(
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
) (schedule *model.Schedule, errCode int, err error) {
	schedule, errCode, err = controller.Service.Update(
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
) (schedule int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.Delete(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	schedule = affectedRows
	return
}

func (controller *Controller) Get(
	ctx *context.Context,
	input *struct {
		data.ScheduleID
	},
) (schedule *model.Schedule, errCode int, err error) {
	schedule, errCode, err = controller.Service.Get(helpers.GetJwtContext(ctx), input.ID)
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllRequest
	},
) (schedule *data.ScheduleResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	scheduleList, errCode, err := controller.Service.GetAll(helpers.GetJwtContext(ctx), newFilter, newPagination, &input.GetAllRequest)
	if err != nil {
		return
	}
	schedule = &data.ScheduleResponseList{
		Data: model.ToScheduleResponseList(scheduleList),
	}
	schedule.Filter = newFilter
	schedule.Pagination = newPagination
	return
}
