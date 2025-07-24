package exam

import (
	"context"

	"api/common/helpers"
	httpHelper "api/common/helpers/http"
	"api/common/types"
	"api/services/school/common/exam/data"
	"api/services/school/common/exam/model"
)

type Controller struct {
	Service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{Service: service}
}

func (controller *Controller) CreateType(
	ctx *context.Context,
	input *struct {
		Body data.ExamTypeRequest
	},
) (result *model.ExamType, errCode int, err error) {
	result, errCode, err = controller.Service.CreateType(
		httpHelper.GetContextData(ctx),
		&input.Body,
	)
	return
}

func (controller *Controller) Create(
	ctx *context.Context,
	input *struct {
		Body data.ExamRequest
	},
) (result *model.Exam, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		httpHelper.GetContextData(ctx),
		&input.Body,
	)
	return
}

func (controller *Controller) UpdateType(
	ctx *context.Context,
	input *struct {
		data.ExamTypeID
		Body data.ExamTypeRequest
	},
) (result *model.ExamType, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateType(
		httpHelper.GetContextData(ctx), input.ID,
		&input.Body,
	)
	return
}

func (controller *Controller) Update(
	ctx *context.Context,
	input *struct {
		data.ExamID
		Body data.ExamRequest
	},
) (result *model.Exam, errCode int, err error) {
	result, errCode, err = controller.Service.Update(
		httpHelper.GetContextData(ctx), input.ID,
		&input.Body,
	)
	return
}

func (controller *Controller) DeleteType(
	ctx *context.Context,
	input *struct {
		data.ExamTypeID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteType(httpHelper.GetContextData(ctx), input.ID)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) Delete(
	ctx *context.Context,
	input *struct {
		data.ExamID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.Delete(httpHelper.GetContextData(ctx), input.ID)
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

func (controller *Controller) GetType(
	ctx *context.Context,
	input *struct {
		data.ExamTypeID
	},
) (result *model.ExamType, errCode int, err error) {
	exam, errCode, err := controller.Service.GetType(httpHelper.GetContextData(ctx), input.ID)
	if err != nil {
		return
	}
	result = exam
	return
}

func (controller *Controller) Get(
	ctx *context.Context,
	input *struct {
		data.ExamID
	},
) (result *model.Exam, errCode int, err error) {
	exam, errCode, err := controller.Service.Get(httpHelper.GetContextData(ctx), input.ID)
	if err != nil {
		return
	}
	result = exam
	return
}

func (controller *Controller) GetAllExamType(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllExamTypeRequest
	},
) (result *data.ExamTypeResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	examList, errCode, err := controller.Service.GetAllExamType(httpHelper.GetContextData(ctx), newFilter, newPagination, &input.GetAllExamTypeRequest)
	if err != nil {
		return
	}
	result = &data.ExamTypeResponseList{
		Data: model.ToExamTypeResponseList(examList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllRequest
	},
) (result *data.ExamResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	examList, errCode, err := controller.Service.GetAll(httpHelper.GetContextData(ctx), newFilter, newPagination, &input.GetAllRequest)
	if err != nil {
		return
	}
	result = &data.ExamResponseList{
		Data: model.ToExamResponseList(examList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
