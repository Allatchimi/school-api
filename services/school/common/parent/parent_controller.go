package parent

import (
	"context"

	"api/common/helpers"
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
		helpers.GetJwtContext(ctx),
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
		helpers.GetJwtContext(ctx),
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
		helpers.GetJwtContext(ctx),
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
		helpers.GetJwtContext(ctx),
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
	affectedRows, errCode, err := controller.Service.Delete(helpers.GetJwtContext(ctx), input.ID)
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
	affectedRows, errCode, err := controller.Service.DeleteParentStudent(helpers.GetJwtContext(ctx), input.ID)
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
	parent, errCode, err := controller.Service.Get(helpers.GetJwtContext(ctx), input.ID)
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
	parent, errCode, err := controller.Service.GetParentStudent(helpers.GetJwtContext(ctx), input.ID)
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
	parentList, errCode, err := controller.Service.GetAll(helpers.GetJwtContext(ctx), newFilter, newPagination, &input.GetAllRequest)
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
	parentList, errCode, err := controller.Service.GetAllParentStudent(helpers.GetJwtContext(ctx), newFilter, newPagination, &input.GetAllParentStudentRequest)
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
