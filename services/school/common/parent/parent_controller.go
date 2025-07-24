package parent

import (
	"context"

	"api/common/helpers"
	httpHelper "api/common/helpers/http"
	"api/common/types"
	"api/services/school/common/parent/data"
	"api/services/school/common/parent/model"
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
		Body data.ParentRequest
	},
) (result *model.Parent, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		httpHelper.GetContextData(ctx),
		&input.Body,
	)
	return
}

func (controller *Controller) CreateParentStudent(
	ctx *context.Context,
	input *struct {
		Body data.ParentStudentRequest
	},
) (result *model.ParentStudent, errCode int, err error) {
	result, errCode, err = controller.Service.CreateParentStudent(
		httpHelper.GetContextData(ctx),
		&input.Body,
	)
	return
}

func (controller *Controller) Update(
	ctx *context.Context,
	input *struct {
		data.ParentID
		Body data.ParentRequest
	},
) (result *model.Parent, errCode int, err error) {
	result, errCode, err = controller.Service.Update(
		httpHelper.GetContextData(ctx),
		input.ID,
		&input.Body,
	)
	return
}

func (controller *Controller) UpdateParentStudent(
	ctx *context.Context,
	input *struct {
		data.ParentStudentID
		Body data.ParentStudentRequest
	},
) (result *model.ParentStudent, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateParentStudent(
		httpHelper.GetContextData(ctx),
		input.ID,
		&input.Body,
	)
	return
}

func (controller *Controller) Delete(
	ctx *context.Context,
	input *struct {
		data.ParentID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.Delete(httpHelper.GetContextData(ctx), input.ID)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) DeleteParentStudent(
	ctx *context.Context,
	input *struct {
		data.ParentStudentID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteParentStudent(httpHelper.GetContextData(ctx), input.ID)
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
	affectedRows, errCode, err := controller.Service.DeleteMultiple(httpHelper.GetContextData(ctx), input.Body.List)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) Get(
	ctx *context.Context,
	input *struct {
		data.ParentID
	},
) (result *model.Parent, errCode int, err error) {
	parent, errCode, err := controller.Service.Get(httpHelper.GetContextData(ctx), input.ID)
	if err != nil {
		return
	}
	result = parent
	return
}

func (controller *Controller) GetParentStudent(
	ctx *context.Context,
	input *struct {
		data.ParentStudentID
	},
) (result *model.ParentStudent, errCode int, err error) {
	parent, errCode, err := controller.Service.GetParentStudent(httpHelper.GetContextData(ctx), input.ID)
	if err != nil {
		return
	}
	result = parent
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllRequest
	},
) (result *data.ParentResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	parentList, errCode, err := controller.Service.GetAll(httpHelper.GetContextData(ctx), newFilter, newPagination, &input.GetAllRequest)
	if err != nil {
		return
	}
	result = &data.ParentResponseList{
		Data: model.ToParentResponseList(parentList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}

func (controller *Controller) GetAllParentStudent(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllParentStudentRequest
	},
) (result *data.ParentStudentResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	parentList, errCode, err := controller.Service.GetAllParentStudent(httpHelper.GetContextData(ctx), newFilter, newPagination, &input.GetAllParentStudentRequest)
	if err != nil {
		return
	}
	result = &data.ParentStudentResponseList{
		Data: model.ToParentStudentResponseList(parentList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
