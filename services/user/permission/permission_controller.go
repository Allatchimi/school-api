package permission

import (
	"context"

	"api/common/helpers"
	httpHelper "api/common/helpers/http"
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
		data.PermissionRoleID
		Body data.UpdatePermissionRequest
	},
) (result *data.PermissionResponse, errCode int, err error) {
	tmpResult, errCode, err := controller.Service.Update(
		httpHelper.GetJwtContext(ctx),
		input.PermissionRoleID.RoleID,
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
	affectedRows, errCode, err := controller.Service.Delete(httpHelper.GetJwtContext(ctx), input.ID)
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
	affectedRows, errCode, err := controller.Service.DeleteMultiple(httpHelper.GetJwtContext(ctx), input.Body.List)
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
	permissionList, errCode, err := controller.Service.GetAll(httpHelper.GetJwtContext(ctx), newFilter, newPagination, &input.GetAllRequest)
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
