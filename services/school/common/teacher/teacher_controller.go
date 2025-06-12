package teacher

import (
	"context"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/common/teacher/data"
	"api/services/school/common/teacher/model"
	modelUser "api/services/user/user/model"
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
		input.Body.SchoolID,
		input.Body.UID,
		&modelUser.User{
			Email:       input.Body.Email,
			PhoneNumber: input.Body.PhoneNumber,
			Info: &modelUser.UserInfo{
				Gender:        input.Body.Info.Gender,
				Username:      input.Body.Info.Username,
				FirstName:     input.Body.Info.FirstName,
				LastName:      input.Body.Info.LastName,
				Birthday:      input.Body.Info.Birthday,
				BirthLocation: input.Body.Info.BirthLocation,
				Address:       input.Body.Info.Address,
				Language:      input.Body.Info.Language,
				Image:         input.Body.Info.Image,
			},
		},
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
		helpers.GetJwtContext(ctx),
		&model.TeacherClassSubjectUnit{
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
		input.Body.UID,
		&modelUser.User{
			Email:       input.Body.Email,
			PhoneNumber: input.Body.PhoneNumber,
			Info: &modelUser.UserInfo{
				Gender:        input.Body.Info.Gender,
				Username:      input.Body.Info.Username,
				FirstName:     input.Body.Info.FirstName,
				LastName:      input.Body.Info.LastName,
				Birthday:      input.Body.Info.Birthday,
				BirthLocation: input.Body.Info.BirthLocation,
				Address:       input.Body.Info.Address,
				Language:      input.Body.Info.Language,
				Image:         input.Body.Info.Image,
			},
		},
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
		helpers.GetJwtContext(ctx), input.ID,
		&model.TeacherClassSubjectUnit{
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

func (controller *Controller) DeleteTeacherClassSubjectUnit(
	ctx *context.Context,
	input *struct {
		data.UnitSubjectID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteTeacherClassSubjectUnit(helpers.GetJwtContext(ctx), input.ID)
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

func (controller *Controller) GetTeacherClassSubjectUnit(
	ctx *context.Context,
	input *struct {
		data.UnitSubjectID
	},
) (result *model.TeacherClassSubjectUnit, errCode int, err error) {
	teacher, errCode, err := controller.Service.GetTeacherClassSubjectUnit(helpers.GetJwtContext(ctx), input.ID)
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

func (controller *Controller) GetAllTeacherClassSubjectUnit(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllTeacherClassSubjectUnitRequest
	},
) (result *data.TeacherClassSubjectUnitResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	teacherList, errCode, err := controller.Service.GetAllTeacherClassSubjectUnit(helpers.GetJwtContext(ctx), newFilter, newPagination, input.GetAllTeacherClassSubjectUnitRequest.SchoolID, input.GetAllTeacherClassSubjectUnitRequest.TeacherID)
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
