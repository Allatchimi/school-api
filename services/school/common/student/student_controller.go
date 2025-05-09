package student

import (
	"context"

	"api/common/helpers"
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
		helpers.GetJwtContext(ctx),
		&model.Student{
			SchoolID: input.Body.SchoolID,
			UserID:   input.Body.UserID,
			UID:      input.Body.UID,
		},
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
		helpers.GetJwtContext(ctx),
		&model.StudentEnroll{
			StudentID:     input.Body.StudentID,
			YearID:        input.Body.YearID,
			LevelDomainID: input.Body.LevelDomainID,
			ClassID:       input.Body.ClassID,
		},
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
		helpers.GetJwtContext(ctx), input.ID,
		&model.Student{
			SchoolID: input.Body.SchoolID,
			UserID:   input.Body.UserID,
			UID:      input.Body.UID,
		},
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
		helpers.GetJwtContext(ctx), input.ID,
		&model.StudentEnroll{
			StudentID:     input.Body.StudentID,
			YearID:        input.Body.YearID,
			LevelDomainID: input.Body.LevelDomainID,
			ClassID:       input.Body.ClassID,
		},
	)
	return
}

func (controller *Controller) Delete(
	ctx *context.Context,
	input *struct {
		data.StudentID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.Delete(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) DeleteStudentEnroll(
	ctx *context.Context,
	input *struct {
		data.EnrollID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteStudentEnroll(helpers.GetJwtContext(ctx), input.ID)
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
	student, errCode, err := controller.Service.Get(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = student
	return
}

func (controller *Controller) GetStudentEnroll(
	ctx *context.Context,
	input *struct {
		data.EnrollID
	},
) (result *model.StudentEnroll, errCode int, err error) {
	student, errCode, err := controller.Service.GetStudentEnroll(helpers.GetJwtContext(ctx), input.ID)
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
	studentList, errCode, err := controller.Service.GetAll(helpers.GetJwtContext(ctx), newFilter, newPagination, input.GetAllRequest.SchoolID)
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
	studentList, errCode, err := controller.Service.GetAllStudentEnroll(helpers.GetJwtContext(ctx), newFilter, newPagination, input.GetAllStudentEnrollRequest.SchoolID, input.GetAllStudentEnrollRequest.StudentID)
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
