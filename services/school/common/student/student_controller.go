package student

import (
	"context"

	"api/common/helpers"
	httpHelper "api/common/helpers/http"
	"api/common/types"
	"api/services/school/common/student/data"
	"api/services/school/common/student/model"
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
		Body data.StudentRequest
	},
) (result *model.Student, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		httpHelper.GetContextData(ctx),
		&input.Body,
	)
	return
}

func (controller *Controller) CreateStudentEnroll(
	ctx *context.Context,
	input *struct {
		Body data.StudentEnrollRequest
	},
) (result *model.StudentEnroll, errCode int, err error) {
	result, errCode, err = controller.Service.CreateStudentEnroll(
		httpHelper.GetContextData(ctx),
		&input.Body,
	)
	return
}

func (controller *Controller) CreateStudentPreEnroll(
	ctx *context.Context,
	input *struct {
		Body data.StudentPreEnrollRequest
	},
) (result *model.StudentPreEnroll, errCode int, err error) {
	result, errCode, err = controller.Service.CreateStudentPreEnroll(
		httpHelper.GetContextData(ctx),
		&input.Body,
	)
	return
}

func (controller *Controller) Update(
	ctx *context.Context,
	input *struct {
		data.StudentID
		Body data.StudentRequest
	},
) (result *model.Student, errCode int, err error) {
	result, errCode, err = controller.Service.Update(
		httpHelper.GetContextData(ctx),
		input.ID,
		&input.Body,
	)
	return
}

func (controller *Controller) UpdateStudentEnroll(
	ctx *context.Context,
	input *struct {
		data.StudentEnrollID
		Body data.StudentEnrollRequest
	},
) (result *model.StudentEnroll, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateStudentEnroll(
		httpHelper.GetContextData(ctx),
		input.ID,
		&input.Body,
	)
	return
}

func (controller *Controller) UpdateStudentPreEnroll(
	ctx *context.Context,
	input *struct {
		data.StudentPreEnrollID
		Body data.StudentPreEnrollRequest
	},
) (result *model.StudentPreEnroll, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateStudentPreEnroll(
		httpHelper.GetContextData(ctx),
		input.ID,
		&input.Body,
	)
	return
}

func (controller *Controller) UpdateStudentPreEnrollStatus(
	ctx *context.Context,
	input *struct {
		data.StudentPreEnrollID
		Body data.StudentPreEnrollStatusRequest
	},
) (result *model.StudentPreEnroll, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateStudentPreEnrollStatus(
		httpHelper.GetContextData(ctx),
		input.ID,
		&input.Body,
	)
	return
}

func (controller *Controller) Delete(
	ctx *context.Context,
	input *struct {
		data.StudentID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.Delete(httpHelper.GetContextData(ctx), input.ID)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) DeleteStudentEnroll(
	ctx *context.Context,
	input *struct {
		data.StudentEnrollID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteStudentEnroll(httpHelper.GetContextData(ctx), input.ID)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) DeleteStudentPreEnroll(
	ctx *context.Context,
	input *struct {
		data.StudentPreEnrollID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteStudentPreEnroll(httpHelper.GetContextData(ctx), input.ID)
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

func (controller *Controller) DeleteMultipleStudentEnroll(
	ctx *context.Context,
	input *struct {
		Body types.DeleteMultipleRequest
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteMultipleStudentEnroll(httpHelper.GetContextData(ctx), input.Body.List)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) DeleteMultipleStudentPreEnroll(
	ctx *context.Context,
	input *struct {
		Body types.DeleteMultipleRequest
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteMultipleStudentPreEnroll(httpHelper.GetContextData(ctx), input.Body.List)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) Get(
	ctx *context.Context,
	input *struct {
		data.StudentID
	},
) (result *model.Student, errCode int, err error) {
	student, errCode, err := controller.Service.Get(httpHelper.GetContextData(ctx), input.ID)
	if err != nil {
		return
	}
	result = student
	return
}

func (controller *Controller) GetStudentEnroll(
	ctx *context.Context,
	input *struct {
		data.StudentEnrollID
	},
) (result *model.StudentEnroll, errCode int, err error) {
	student, errCode, err := controller.Service.GetStudentEnroll(httpHelper.GetContextData(ctx), input.ID)
	if err != nil {
		return
	}
	result = student
	return
}

func (controller *Controller) GetStudentPreEnroll(
	ctx *context.Context,
	input *struct {
		data.StudentPreEnrollID
	},
) (result *model.StudentPreEnroll, errCode int, err error) {
	student, errCode, err := controller.Service.GetStudentPreEnroll(httpHelper.GetContextData(ctx), input.ID)
	if err != nil {
		return
	}
	result = student
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllRequest
	},
) (result *data.StudentResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	studentList, errCode, err := controller.Service.GetAll(httpHelper.GetContextData(ctx), newFilter, newPagination, &input.GetAllRequest)
	if err != nil {
		return
	}
	result = &data.StudentResponseList{
		Data: model.ToStudentResponseList(studentList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}

func (controller *Controller) GetAllStudentEnroll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllStudentEnrollRequest
	},
) (result *data.StudentEnrollResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	studentList, errCode, err := controller.Service.GetAllStudentEnroll(httpHelper.GetContextData(ctx), newFilter, newPagination, &input.GetAllStudentEnrollRequest)
	if err != nil {
		return
	}
	result = &data.StudentEnrollResponseList{
		Data: model.ToStudentEnrollResponseList(studentList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}

func (controller *Controller) GetAllStudentPreEnroll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllStudentPreEnrollRequest
	},
) (result *data.StudentPreEnrollResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	studentList, errCode, err := controller.Service.GetAllStudentPreEnroll(httpHelper.GetContextData(ctx), newFilter, newPagination, &input.GetAllStudentPreEnrollRequest)
	if err != nil {
		return
	}
	result = &data.StudentPreEnrollResponseList{
		Data: model.ToStudentPreEnrollResponseList(studentList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}

func (controller *Controller) GetAllPublic(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllRequest
	},
) (result *data.StudentResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	studentList, errCode, err := controller.Service.GetAll(httpHelper.GetContextData(ctx), newFilter, newPagination, &input.GetAllRequest)
	if err != nil {
		return
	}
	result = &data.StudentResponseList{
		Data: model.ToStudentResponseList(studentList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
