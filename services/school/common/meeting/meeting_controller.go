package meeting

import (
	"context"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/common/meeting/data"
	"api/services/school/common/meeting/model"
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
		Body data.MeetingRoomRequest
	},
) (result *model.MeetingRoom, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		helpers.GetJwtContext(ctx),
		&model.MeetingRoom{
			SchoolID:       input.Body.SchoolID,
			ClassSubjectID: input.Body.ClassSubjectID,
			UnitID:         input.Body.UnitID,
		},
	)
	return
}

func (controller *Controller) Delete(
	ctx *context.Context,
	input *struct {
		data.MeetingRoomID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.Delete(helpers.GetJwtContext(ctx), input.ID)
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
	affectedRows, errCode, err := controller.Service.DeleteMultiple(helpers.GetJwtContext(ctx), input.Body.List)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) Get(
	ctx *context.Context,
	input *struct {
		data.MeetingRoomID
	},
) (result *model.MeetingRoom, errCode int, err error) {
	meeting, errCode, err := controller.Service.Get(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = meeting
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllRequest
	},
) (result *data.MeetingRoomResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	meetingList, errCode, err := controller.Service.GetAll(helpers.GetJwtContext(ctx), newFilter, newPagination, &input.GetAllRequest)
	if err != nil {
		return
	}
	result = &data.MeetingRoomResponseList{
		Data: model.ToResponseList(meetingList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}

func (controller *Controller) Join(
	ctx *context.Context,
	input *struct {
		data.MeetingRoomID
	},
) (result string, errCode int, err error) {
	meeting, errCode, err := controller.Service.Join(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = meeting
	return
}
