package class

import (
	"context"

	"api/common/helpers"
	httpHelper "api/common/helpers/http"
	"api/common/types"
	"api/services/school/highschool/class/data"
	"api/services/school/highschool/class/model"
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
		Body data.ClassRequest
	},
) (result *model.HighschoolClass, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		httpHelper.GetJwtContext(ctx),
		&input.Body,
	)
	return
}

func (controller *Controller) CreateClassSubject(
	ctx *context.Context,
	input *struct {
		Body data.ClassSubjectRequest
	},
) (result *model.HighschoolClassSubject, errCode int, err error) {
	result, errCode, err = controller.Service.CreateClassSubject(
		httpHelper.GetJwtContext(ctx),
		&input.Body,
	)
	return
}

func (controller *Controller) Update(
	ctx *context.Context,
	input *struct {
		data.ClassID
		Body data.ClassRequest
	},
) (result *model.HighschoolClass, errCode int, err error) {
	result, errCode, err = controller.Service.Update(
		httpHelper.GetJwtContext(ctx),
		input.ID,
		&input.Body,
	)
	return
}

func (controller *Controller) UpdateClassSubject(
	ctx *context.Context,
	input *struct {
		data.ClassSubjectID
		Body data.ClassSubjectRequest
	},
) (result *model.HighschoolClassSubject, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateClassSubject(
		httpHelper.GetJwtContext(ctx),
		input.ID,
		&input.Body,
	)
	return
}

func (controller *Controller) Delete(
	ctx *context.Context,
	input *struct {
		data.ClassID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.Delete(httpHelper.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) DeleteClassSubject(
	ctx *context.Context,
	input *struct {
		data.ClassSubjectID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteClassSubject(httpHelper.GetJwtContext(ctx), input.ID)
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

func (controller *Controller) DeleteMultipleClassSubject(
	ctx *context.Context,
	input *struct {
		Body types.DeleteMultipleRequest
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteMultipleClassSubject(httpHelper.GetJwtContext(ctx), input.Body.List)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) Get(
	ctx *context.Context,
	input *struct {
		data.ClassID
	},
) (result *model.HighschoolClass, errCode int, err error) {
	class, errCode, err := controller.Service.Get(httpHelper.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = class
	return
}

func (controller *Controller) GetClassSubject(
	ctx *context.Context,
	input *struct {
		data.ClassSubjectID
	},
) (result *model.HighschoolClassSubject, errCode int, err error) {
	classSubject, errCode, err := controller.Service.GetClassSubject(httpHelper.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = classSubject
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllRequest
	},
) (result *data.ClassResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	classList, errCode, err := controller.Service.GetAll(httpHelper.GetJwtContext(ctx), newFilter, newPagination, &input.GetAllRequest)
	if err != nil {
		return
	}
	result = &data.ClassResponseList{
		Data: model.ToResponseList(classList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}

func (controller *Controller) GetAllClassSubject(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllClassSubjectRequest
	},
) (result *data.ClassSubjectResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	classList, errCode, err := controller.Service.GetAllClassSubject(httpHelper.GetJwtContext(ctx), newFilter, newPagination, &input.GetAllClassSubjectRequest)
	if err != nil {
		return
	}
	result = &data.ClassSubjectResponseList{
		Data: model.ToClassSubjectResponseList(classList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
