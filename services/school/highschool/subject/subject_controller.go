package subject

import (
	"context"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/highschool/subject/data"
	"api/services/school/highschool/subject/model"
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
		Body data.CreateSubjectRequest
	},
) (result *model.HighschoolSubject, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		helpers.GetJwtContext(ctx),
		&model.HighschoolSubject{
			SchoolID:     input.Body.SchoolID,
			ClassID:      input.Body.ClassID,
			Name:         input.Body.Name,
			Description:  input.Body.Description,
			Coefficient:  input.Body.Coefficient,
			Program:      input.Body.Program,
			Requirements: input.Body.Requirements,
		},
	)
	return
}

func (controller *Controller) Update(
	ctx *context.Context,
	input *struct {
		data.SubjectID
		Body data.UpdateSubjectRequest
	},
) (result *model.HighschoolSubject, errCode int, err error) {
	result, errCode, err = controller.Service.Update(
		helpers.GetJwtContext(ctx), input.ID,
		&model.HighschoolSubject{
			Name:         input.Body.Name,
			Description:  input.Body.Description,
			Coefficient:  input.Body.Coefficient,
			Program:      input.Body.Program,
			Requirements: input.Body.Requirements,
		},
	)
	return
}

func (controller *Controller) Delete(
	ctx *context.Context,
	input *struct {
		data.SubjectID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.Delete(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) Get(
	ctx *context.Context,
	input *struct {
		data.SubjectID
	},
) (result *model.HighschoolSubject, errCode int, err error) {
	subject, errCode, err := controller.Service.Get(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = subject
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllRequest
	},
) (result *data.SubjectResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	subjectList, errCode, err := controller.Service.GetAll(helpers.GetJwtContext(ctx), newFilter, newPagination, input.GetAllRequest.SchoolID)
	if err != nil {
		return
	}
	result = &data.SubjectResponseList{
		Data: model.ToSubjectResponseList(subjectList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
