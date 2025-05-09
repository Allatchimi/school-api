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

func (controller *Controller) CreateLevelClass(
	ctx *context.Context,
	input *struct {
		Body data.StudentLevelClassRequest
	},
) (result *model.StudentLevelClass, errCode int, err error) {
	result, errCode, err = controller.Service.CreateLevelClass(
		helpers.GetJwtContext(ctx),
		&model.StudentLevelClass{
			StudentID: input.Body.StudentID,
			YearID:    input.Body.YearID,

			DomainID: input.Body.DomainID,
			LevelID:  input.Body.LevelID,

			ClassID: input.Body.ClassID,
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

func (controller *Controller) UpdateLevelClass(
	ctx *context.Context,
	input *struct {
		data.StudentLevelClassID
		Body data.StudentLevelClassRequest
	},
) (result *model.StudentLevelClass, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateLevelClass(
		helpers.GetJwtContext(ctx), input.ID,
		&model.StudentLevelClass{
			StudentID: input.Body.StudentID,
			YearID:    input.Body.YearID,

			DomainID: input.Body.DomainID,
			LevelID:  input.Body.LevelID,

			ClassID: input.Body.ClassID,
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

func (controller *Controller) DeleteLevelClass(
	ctx *context.Context,
	input *struct {
		data.StudentLevelClassID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteLevelClass(helpers.GetJwtContext(ctx), input.ID)
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

func (controller *Controller) GetLevelClass(
	ctx *context.Context,
	input *struct {
		data.StudentLevelClassID
	},
) (result *model.StudentLevelClass, errCode int, err error) {
	student, errCode, err := controller.Service.GetLevelClass(helpers.GetJwtContext(ctx), input.ID)
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

func (controller *Controller) GetAllLevelClass(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllLevelClassRequest
	},
) (result *data.StudentLevelClassResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	studentList, errCode, err := controller.Service.GetAllLevelClass(helpers.GetJwtContext(ctx), newFilter, newPagination, input.GetAllLevelClassRequest.StudentID)
	if err != nil {
		return
	}
	result = &data.StudentLevelClassResponseList{
		Data: model.ToStudentLevelClassResponseList(studentList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
