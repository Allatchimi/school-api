package role

import (
	"context"

	"api/common/helpers"
	"api/common/types"
	"api/services/user/role/data"
	"api/services/user/role/model"
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
		Body data.RoleRequest
	},
) (result *model.Role, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		helpers.GetJwtContext(ctx),
		&model.Role{
			Name:        input.Body.Name,
			Feature:     input.Body.Feature,
			Description: input.Body.Description,
		},
	)
	return
}

func (controller *Controller) Update(
	ctx *context.Context,
	input *struct {
		data.RoleID
		Body data.RoleRequest
	},
) (result *model.Role, errCode int, err error) {
	result, errCode, err = controller.Service.Update(
		helpers.GetJwtContext(ctx), input.ID,
		&model.Role{
			Name:        input.Body.Name,
			Feature:     input.Body.Feature,
			Description: input.Body.Description,
		},
	)
	return
}

func (controller *Controller) Delete(
	ctx *context.Context,
	input *struct {
		data.RoleID
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

func (controller *Controller) GetByID(
	ctx *context.Context,
	input *struct {
		data.RoleID
	},
) (result *model.Role, errCode int, err error) {
	role, errCode, err := controller.Service.GetByID(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = role
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllRequest
	},
) (result *data.RoleResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	roleList, errCode, err := controller.Service.GetAll(helpers.GetJwtContext(ctx), newFilter, newPagination, &input.GetAllRequest)
	if err != nil {
		return
	}
	result = &data.RoleResponseList{
		Data: model.ToResponseList(roleList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
