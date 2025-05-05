package teacher

import (
	"context"

	"api/common/helpers"
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
		helpers.GetJwtContext(ctx),
		&model.Teacher{
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
		Body data.TeacherLevelClassRequest
	},
) (result *model.TeacherLevelClass, errCode int, err error) {
	result, errCode, err = controller.Service.CreateLevelClass(
		helpers.GetJwtContext(ctx),
		&model.TeacherLevelClass{
			TeacherID: input.Body.TeacherID,
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
		data.TeacherID
		Body data.TeacherRequest
	},
) (result *model.Teacher, errCode int, err error) {
	result, errCode, err = controller.Service.Update(
		helpers.GetJwtContext(ctx), input.ID,
		&model.Teacher{
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
		data.TeacherLevelClassID
		Body data.TeacherLevelClassRequest
	},
) (result *model.TeacherLevelClass, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateLevelClass(
		helpers.GetJwtContext(ctx), input.ID,
		&model.TeacherLevelClass{
			TeacherID: input.Body.TeacherID,
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
		data.TeacherID
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
		data.TeacherLevelClassID
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
		data.TeacherID
	},
) (result *model.Teacher, errCode int, err error) {
	teacher, errCode, err := controller.Service.Get(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = teacher
	return
}

func (controller *Controller) GetLevelClass(
	ctx *context.Context,
	input *struct {
		data.TeacherLevelClassID
	},
) (result *model.TeacherLevelClass, errCode int, err error) {
	teacher, errCode, err := controller.Service.GetLevelClass(helpers.GetJwtContext(ctx), input.ID)
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
	teacherList, errCode, err := controller.Service.GetAll(helpers.GetJwtContext(ctx), newFilter, newPagination, input.GetAllRequest.SchoolID)
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

func (controller *Controller) GetAllLevelClass(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllLevelClassRequest
	},
) (result *data.TeacherLevelClassResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	teacherList, errCode, err := controller.Service.GetAllLevelClass(helpers.GetJwtContext(ctx), newFilter, newPagination, input.GetAllLevelClassRequest.TeacherID)
	if err != nil {
		return
	}
	result = &data.TeacherLevelClassResponseList{
		Data: model.ToTeacherLevelClassResponseList(teacherList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
