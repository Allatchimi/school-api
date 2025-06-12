package permission

import (
	"context"

	"api/common/helpers"
	"api/common/types"
	"api/services/user/permission/data"
	"api/services/user/permission/model"
)

type Controller struct {
	Service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{Service: service}
}

func (controller *Controller) Update(
	ctx *context.Context,
	input *struct {
		data.PermissionPathRequest
		Body data.UpdatePermissionRequest
	},
) (result *data.PermissionResponse, errCode int, err error) {
	tmpResult, errCode, err := controller.Service.Update(
		helpers.GetJwtContext(ctx),
		input.PermissionPathRequest.RoleID,
		&input.Body,
	)
	result = tmpResult.ToResponse()
	return
}

func (controller *Controller) Delete(
	ctx *context.Context,
	input *struct {
		data.PermissionID
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

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllRequest
	},
) (result *data.PermissionListResponse, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	permissionList, errCode, err := controller.Service.GetAll(helpers.GetJwtContext(ctx), newFilter, newPagination, &input.GetAllRequest)
	if err != nil {
		return
	}
	result = &data.PermissionListResponse{
		Data: model.ToResponseList(permissionList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
