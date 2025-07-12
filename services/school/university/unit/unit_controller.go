package unit

import (
	"context"

	"api/common/helpers"
	httpHelper "api/common/helpers/http"
	"api/common/types"
	"api/services/school/university/unit/data"
	"api/services/school/university/unit/model"
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
		Body data.UnitRequest
	},
) (result *model.UniversityUnit, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		httpHelper.GetJwtContext(ctx),
		&input.Body,
	)
	return
}

func (controller *Controller) Update(
	ctx *context.Context,
	input *struct {
		data.UnitID
		Body data.UnitRequest
	},
) (result *model.UniversityUnit, errCode int, err error) {
	result, errCode, err = controller.Service.Update(
		httpHelper.GetJwtContext(ctx),
		input.ID,
		&input.Body,
	)
	return
}

func (controller *Controller) Delete(
	ctx *context.Context,
	input *struct {
		data.UnitID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.Delete(httpHelper.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) Get(
	ctx *context.Context,
	input *struct {
		data.UnitID
	},
) (result *model.UniversityUnit, errCode int, err error) {
	unit, errCode, err := controller.Service.Get(httpHelper.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = unit
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllRequest
	},
) (result *data.UnitResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	unitList, errCode, err := controller.Service.GetAll(httpHelper.GetJwtContext(ctx), newFilter, newPagination, &input.GetAllRequest)
	if err != nil {
		return
	}
	result = &data.UnitResponseList{
		Data: model.ToUnitResponseList(unitList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
