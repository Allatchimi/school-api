package year

import (
	"context"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/common/year/data"
	"api/services/school/common/year/model"
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
		Body data.YearRequest
	},
) (result *model.Year, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		helpers.GetJwtContext(ctx),
		&model.Year{
			SchoolID:  input.Body.SchoolID,
			StartDate: input.Body.StartDate,
			EndDate:   input.Body.EndDate,
		},
	)
	return
}

func (controller *Controller) Update(
	ctx *context.Context,
	input *struct {
		data.YearID
		Body data.YearRequest
	},
) (result *model.Year, errCode int, err error) {
	result, errCode, err = controller.Service.Update(
		helpers.GetJwtContext(ctx), input.ID,
		&model.Year{
			SchoolID:  input.Body.SchoolID,
			StartDate: input.Body.StartDate,
			EndDate:   input.Body.EndDate,
		},
	)
	return
}

func (controller *Controller) Delete(
	ctx *context.Context,
	input *struct {
		data.YearID
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
		data.YearID
	},
) (result *model.Year, errCode int, err error) {
	year, errCode, err := controller.Service.Get(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = year
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
	},
) (result *data.YearResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	yearList, errCode, err := controller.Service.GetAll(helpers.GetJwtContext(ctx), newFilter, newPagination, -1)
	if err != nil {
		return
	}
	result = &data.YearResponseList{
		Data: model.ToResponseList(yearList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}

func (controller *Controller) GetAllBySchool(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllBySchoolRequest
	},
) (result *data.YearResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	yearList, errCode, err := controller.Service.GetAll(helpers.GetJwtContext(ctx), newFilter, newPagination, input.GetAllBySchoolRequest.ID)
	if err != nil {
		return
	}
	result = &data.YearResponseList{
		Data: model.ToResponseList(yearList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
