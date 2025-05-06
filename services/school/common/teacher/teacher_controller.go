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

func (controller *Controller) CreateTeacherTUSubject(
	ctx *context.Context,
	input *struct {
		Body data.TeacherTUSubjectRequest
	},
) (result *model.TeacherTeachingUnitSubject, errCode int, err error) {
	result, errCode, err = controller.Service.CreateTeacherTUSubject(
		helpers.GetJwtContext(ctx),
		&model.TeacherTeachingUnitSubject{
			TeacherID: input.Body.TeacherID,
			YearID:    input.Body.YearID,

			TeachingUnitID: input.Body.TeachingUnitID,
			SubjectID:      input.Body.SubjectID,
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

func (controller *Controller) UpdateTeacherTUSubject(
	ctx *context.Context,
	input *struct {
		data.TUSubjectID
		Body data.TeacherTUSubjectRequest
	},
) (result *model.TeacherTeachingUnitSubject, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateTeacherTUSubject(
		helpers.GetJwtContext(ctx), input.ID,
		&model.TeacherTeachingUnitSubject{
			TeacherID: input.Body.TeacherID,
			YearID:    input.Body.YearID,

			TeachingUnitID: input.Body.TeachingUnitID,
			SubjectID:      input.Body.SubjectID,
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

func (controller *Controller) DeleteTeacherTUSubject(
	ctx *context.Context,
	input *struct {
		data.TUSubjectID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteTeacherTUSubject(helpers.GetJwtContext(ctx), input.ID)
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

func (controller *Controller) GetTeacherTUSubject(
	ctx *context.Context,
	input *struct {
		data.TUSubjectID
	},
) (result *model.TeacherTeachingUnitSubject, errCode int, err error) {
	teacher, errCode, err := controller.Service.GetTeacherTUSubject(helpers.GetJwtContext(ctx), input.ID)
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

func (controller *Controller) GetAllTeacherTUSubject(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllTeacherTUSubjectRequest
	},
) (result *data.TeacherTUSubjectResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	teacherList, errCode, err := controller.Service.GetAllTeacherTUSubject(helpers.GetJwtContext(ctx), newFilter, newPagination, input.GetAllTeacherTUSubjectRequest.TeacherID)
	if err != nil {
		return
	}
	result = &data.TeacherTUSubjectResponseList{
		Data: model.ToTeacherTUSubjectResponseList(teacherList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
