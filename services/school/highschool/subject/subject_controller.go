package subject

import (
	"context"

	"api/common/helpers"
	httpHelper "api/common/helpers/http"
	"api/common/types"
	"api/services/school/highschool/subject/data"
	"api/services/school/highschool/subject/model"
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
		Body data.SubjectRequest
	},
) (result *model.HighschoolSubject, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		httpHelper.GetContextData(ctx),
		&input.Body,
	)
	return
}

func (controller *Controller) Update(
	ctx *context.Context,
	input *struct {
		data.SubjectID
		Body data.SubjectRequest
	},
) (result *model.HighschoolSubject, errCode int, err error) {
	result, errCode, err = controller.Service.Update(
		httpHelper.GetContextData(ctx),
		input.ID,
		&input.Body,
	)
	return
}

func (controller *Controller) Delete(
	ctx *context.Context,
	input *struct {
		data.SubjectID
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

func (controller *Controller) Get(
	ctx *context.Context,
	input *struct {
		data.SubjectID
	},
) (result *model.HighschoolSubject, errCode int, err error) {
	subject, errCode, err := controller.Service.Get(httpHelper.GetContextData(ctx), input.ID)
	if err != nil {
		return
	}
	result = subject
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllRequest
	},
) (result *data.SubjectResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	subjectList, errCode, err := controller.Service.GetAll(httpHelper.GetContextData(ctx), newFilter, newPagination, &input.GetAllRequest)
	if err != nil {
		return
	}
	result = &data.SubjectResponseList{
		Data: model.ToResponseList(subjectList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
