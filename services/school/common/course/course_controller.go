package course

import (
	"context"

	"api/common/helpers"
	httpHelper "api/common/helpers/http"
	"api/common/types"
	"api/services/school/common/course/data"
	"api/services/school/common/course/model"
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
		Body data.CourseRequest
	},
) (result *model.Course, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		httpHelper.GetContextData(ctx),
		&input.Body,
	)
	return
}

func (controller *Controller) CreateCourseComment(
	ctx *context.Context,
	input *struct {
		data.CourseID
		Body data.CourseCommentRequest
	},
) (result *model.CourseComment, errCode int, err error) {
	result, errCode, err = controller.Service.CreateCourseComment(
		httpHelper.GetContextData(ctx),
		input.CourseID.ID,
		&input.Body,
	)
	return
}

func (controller *Controller) Update(
	ctx *context.Context,
	input *struct {
		data.CourseID
		Body data.CourseRequest
	},
) (result *model.Course, errCode int, err error) {
	result, errCode, err = controller.Service.Update(
		httpHelper.GetContextData(ctx),
		input.ID,
		&input.Body,
	)
	return
}

func (controller *Controller) UpdateCourseComment(
	ctx *context.Context,
	input *struct {
		data.CourseID
		Body data.CourseCommentRequest
	},
) (result *model.CourseComment, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateCourseComment(
		httpHelper.GetContextData(ctx),
		input.ID,
		&input.Body,
	)
	return
}

func (controller *Controller) Delete(
	ctx *context.Context,
	input *struct {
		data.CourseID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.Delete(httpHelper.GetContextData(ctx), input.ID)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) DeleteCourseComment(
	ctx *context.Context,
	input *struct {
		data.CourseID
		data.CourseCommentID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteCourseComment(httpHelper.GetContextData(ctx), input.CourseID.ID, input.CourseCommentID.ID)
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
		data.CourseID
	},
) (result *model.Course, errCode int, err error) {
	course, errCode, err := controller.Service.Get(httpHelper.GetContextData(ctx), input.ID)
	if err != nil {
		return
	}
	result = course
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllRequest
	},
) (result *data.CourseResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	courseList, errCode, err := controller.Service.GetAll(
		httpHelper.GetContextData(ctx), newFilter, newPagination,
		&input.GetAllRequest,
	)
	if err != nil {
		return
	}
	result = &data.CourseResponseList{
		Data: model.ToResponseList(courseList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}

func (controller *Controller) GetAllCourseComment(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.CourseID
		data.GetAllCourseCommentRequest
	},
) (result *data.CourseCommentResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	courseList, errCode, err := controller.Service.GetAllCourseComment(
		httpHelper.GetContextData(ctx), newFilter, newPagination,
		input.CourseID.ID,
		&input.GetAllCourseCommentRequest,
	)
	if err != nil {
		return
	}
	result = &data.CourseCommentResponseList{
		Data: model.ToCourseCommentResponseList(courseList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
