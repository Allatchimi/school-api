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

func (controller *Controller) CreateTeacherUnitSubject(
	ctx *context.Context,
	input *struct {
		Body data.TeacherUnitSubjectRequest
	},
) (result *model.TeacherUnitSubject, errCode int, err error) {
	result, errCode, err = controller.Service.CreateTeacherUnitSubject(
		helpers.GetJwtContext(ctx),
		&model.TeacherUnitSubject{
			TeacherID: input.Body.TeacherID,
			YearID:    input.Body.YearID,

			UnitID:         input.Body.UnitID,
			ClassSubjectID: input.Body.ClassSubjectID,
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

func (controller *Controller) UpdateTeacherUnitSubject(
	ctx *context.Context,
	input *struct {
		data.UnitSubjectID
		Body data.TeacherUnitSubjectRequest
	},
) (result *model.TeacherUnitSubject, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateTeacherUnitSubject(
		helpers.GetJwtContext(ctx), input.ID,
		&model.TeacherUnitSubject{
			TeacherID: input.Body.TeacherID,
			YearID:    input.Body.YearID,

			UnitID:         input.Body.UnitID,
			ClassSubjectID: input.Body.ClassSubjectID,
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

func (controller *Controller) DeleteTeacherUnitSubject(
	ctx *context.Context,
	input *struct {
		data.UnitSubjectID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteTeacherUnitSubject(helpers.GetJwtContext(ctx), input.ID)
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

func (controller *Controller) GetTeacherUnitSubject(
	ctx *context.Context,
	input *struct {
		data.UnitSubjectID
	},
) (result *model.TeacherUnitSubject, errCode int, err error) {
	teacher, errCode, err := controller.Service.GetTeacherUnitSubject(helpers.GetJwtContext(ctx), input.ID)
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

func (controller *Controller) GetAllTeacherUnitSubject(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllTeacherUnitSubjectRequest
	},
) (result *data.TeacherUnitSubjectResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	teacherList, errCode, err := controller.Service.GetAllTeacherUnitSubject(helpers.GetJwtContext(ctx), newFilter, newPagination, input.GetAllTeacherUnitSubjectRequest.SchoolID, input.GetAllTeacherUnitSubjectRequest.TeacherID)
	if err != nil {
		return
	}
	result = &data.TeacherUnitSubjectResponseList{
		Data: model.ToTeacherUnitSubjectResponseList(teacherList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
