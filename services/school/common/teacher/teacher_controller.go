package teacher

import (
	"context"

	"api/common/helpers"
	httpHelper "api/common/helpers/http"
	"api/common/types"
	"api/services/school/common/teacher/data"
	"api/services/school/common/teacher/model"
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
		Body data.TeacherRequest
	},
) (result *model.Teacher, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		httpHelper.GetContextData(ctx),
		&input.Body,
	)
	return
}

func (controller *Controller) CreateTeacherClassSubjectUnit(
	ctx *context.Context,
	input *struct {
		Body data.TeacherClassSubjectUnitRequest
	},
) (result *model.TeacherClassSubjectUnit, errCode int, err error) {
	result, errCode, err = controller.Service.CreateTeacherClassSubjectUnit(
		httpHelper.GetContextData(ctx),
		&input.Body,
	)
	return
}

func (controller *Controller) Update(
	ctx *context.Context,
	input *struct {
		data.TeacherID
		Body data.TeacherRequest
	},
) (result *model.Teacher, errCode int, err error) {
	result, errCode, err = controller.Service.Update(
		httpHelper.GetContextData(ctx),
		input.ID,
		&input.Body,
	)
	return
}

func (controller *Controller) UpdateTeacherClassSubjectUnit(
	ctx *context.Context,
	input *struct {
		data.UnitSubjectID
		Body data.TeacherClassSubjectUnitRequest
	},
) (result *model.TeacherClassSubjectUnit, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateTeacherClassSubjectUnit(
		httpHelper.GetContextData(ctx), input.ID,
		&input.Body,
	)
	return
}

func (controller *Controller) Delete(
	ctx *context.Context,
	input *struct {
		data.TeacherID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.Delete(httpHelper.GetContextData(ctx), input.ID)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) DeleteTeacherClassSubjectUnit(
	ctx *context.Context,
	input *struct {
		data.UnitSubjectID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteTeacherClassSubjectUnit(httpHelper.GetContextData(ctx), input.ID)
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

func (controller *Controller) DeleteMultipleTeacherClassSubjectUnit(
	ctx *context.Context,
	input *struct {
		Body types.DeleteMultipleRequest
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteMultipleTeacherClassSubjectUnit(httpHelper.GetContextData(ctx), input.Body.List)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) Get(
	ctx *context.Context,
	input *struct {
		data.TeacherID
	},
) (result *model.Teacher, errCode int, err error) {
	teacher, errCode, err := controller.Service.Get(httpHelper.GetContextData(ctx), input.ID)
	if err != nil {
		return
	}
	result = teacher
	return
}

func (controller *Controller) GetTeacherClassSubjectUnit(
	ctx *context.Context,
	input *struct {
		data.UnitSubjectID
	},
) (result *model.TeacherClassSubjectUnit, errCode int, err error) {
	teacher, errCode, err := controller.Service.GetTeacherClassSubjectUnit(httpHelper.GetContextData(ctx), input.ID)
	if err != nil {
		return
	}
	result = teacher
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllRequest
	},
) (result *data.TeacherResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	teacherList, errCode, err := controller.Service.GetAll(
		httpHelper.GetContextData(ctx),
		newFilter,
		newPagination,
		&input.GetAllRequest,
	)
	if err != nil {
		return
	}
	result = &data.TeacherResponseList{
		Data: model.ToTeacherResponseList(teacherList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}

func (controller *Controller) GetAllTeacherClassSubjectUnit(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllTeacherClassSubjectUnitRequest
	},
) (result *data.TeacherClassSubjectUnitResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	teacherList, errCode, err := controller.Service.GetAllTeacherClassSubjectUnit(
		httpHelper.GetContextData(ctx),
		newFilter,
		newPagination,
		&input.GetAllTeacherClassSubjectUnitRequest,
	)
	if err != nil {
		return
	}
	result = &data.TeacherClassSubjectUnitResponseList{
		Data: model.ToTeacherClassSubjectUnitResponseList(teacherList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
