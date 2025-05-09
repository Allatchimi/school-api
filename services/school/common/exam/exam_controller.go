package exam

import (
	"context"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/common/exam/data"
	"api/services/school/common/exam/model"
)

type Controller struct {
	Service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{Service: service}
}

func (controller *Controller) CreateType(
	ctx *context.Context,
	input *struct {
		Body data.ExamTypeRequest
	},
) (result *model.ExamType, errCode int, err error) {
	result, errCode, err = controller.Service.CreateType(
		helpers.GetJwtContext(ctx),
		&model.ExamType{
			SchoolID:    input.Body.SchoolID,
			Name:        input.Body.Name,
			Description: input.Body.Description,
		},
	)
	return
}

func (controller *Controller) Create(
	ctx *context.Context,
	input *struct {
		Body data.ExamRequest
	},
) (result *model.Exam, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		helpers.GetJwtContext(ctx),
		&model.Exam{
			SchoolID:   input.Body.SchoolID,
			YearID:     input.Body.YearID,
			TypeID:     input.Body.TypeID,
			UnitID:     input.Body.UnitID,
			SubjectID:  input.Body.SubjectID,
			SequenceID: input.Body.SequenceID,

			Percentage:      input.Body.Percentage,
			Description:     input.Body.Description,
			LocationType:    input.Body.LocationType,
			LocationDetails: input.Body.LocationDetails,
			Requirements:    input.Body.Requirements,
			AllowedItems:    input.Body.AllowedItems,
			StartDate:       input.Body.StartDate,
			EndDate:         input.Body.EndDate,
		},
	)
	return
}

func (controller *Controller) UpdateType(
	ctx *context.Context,
	input *struct {
		data.ExamTypeID
		Body data.ExamTypeRequest
	},
) (result *model.ExamType, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateType(
		helpers.GetJwtContext(ctx), input.ID,
		&model.ExamType{
			SchoolID:    input.Body.SchoolID,
			Name:        input.Body.Name,
			Description: input.Body.Description,
		},
	)
	return
}

func (controller *Controller) Update(
	ctx *context.Context,
	input *struct {
		data.ExamID
		Body data.ExamRequest
	},
) (result *model.Exam, errCode int, err error) {
	result, errCode, err = controller.Service.Update(
		helpers.GetJwtContext(ctx), input.ID,
		&model.Exam{
			SchoolID:   input.Body.SchoolID,
			YearID:     input.Body.YearID,
			TypeID:     input.Body.TypeID,
			UnitID:     input.Body.UnitID,
			SubjectID:  input.Body.SubjectID,
			SequenceID: input.Body.SequenceID,

			Percentage:      input.Body.Percentage,
			Description:     input.Body.Description,
			LocationType:    input.Body.LocationType,
			LocationDetails: input.Body.LocationDetails,
			Requirements:    input.Body.Requirements,
			AllowedItems:    input.Body.AllowedItems,
			StartDate:       input.Body.StartDate,
			EndDate:         input.Body.EndDate,
		},
	)
	return
}

func (controller *Controller) DeleteType(
	ctx *context.Context,
	input *struct {
		data.ExamTypeID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteType(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) Delete(
	ctx *context.Context,
	input *struct {
		data.ExamID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.Delete(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) GetType(
	ctx *context.Context,
	input *struct {
		data.ExamTypeID
	},
) (result *model.ExamType, errCode int, err error) {
	exam, errCode, err := controller.Service.GetType(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = exam
	return
}

func (controller *Controller) Get(
	ctx *context.Context,
	input *struct {
		data.ExamID
	},
) (result *model.Exam, errCode int, err error) {
	exam, errCode, err := controller.Service.Get(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = exam
	return
}

func (controller *Controller) GetAllExamType(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllExamTypeRequest
	},
) (result *data.ExamTypeResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	examList, errCode, err := controller.Service.GetAllExamType(helpers.GetJwtContext(ctx), newFilter, newPagination, input.GetAllExamTypeRequest.SchoolID)
	if err != nil {
		return
	}
	result = &data.ExamTypeResponseList{
		Data: model.ToExamTypeResponseList(examList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllRequest
	},
) (result *data.ExamResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	examList, errCode, err := controller.Service.GetAll(helpers.GetJwtContext(ctx), newFilter, newPagination, input.GetAllRequest.SchoolID)
	if err != nil {
		return
	}
	result = &data.ExamResponseList{
		Data: model.ToExamResponseList(examList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
