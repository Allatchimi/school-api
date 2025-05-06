package tu

import (
	"context"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/university/tu/data"
	"api/services/school/university/tu/model"
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
		Body data.TeachingUnitRequest
	},
) (result *model.UniversityTeachingUnit, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		helpers.GetJwtContext(ctx),
		&model.UniversityTeachingUnit{
			SchoolID:   input.Body.SchoolID,
			DomainID:   input.Body.DomainID,
			LevelID:    input.Body.LevelID,
			SemesterID: input.Body.SemesterID,

			Name:         input.Body.Name,
			Description:  input.Body.Description,
			Credit:       input.Body.Credit,
			Program:      input.Body.Program,
			Requirements: input.Body.Requirements,

			IsValid: input.Body.IsValid,
		},
	)
	return
}

func (controller *Controller) Update(
	ctx *context.Context,
	input *struct {
		data.TeachingUnitID
		Body data.TeachingUnitRequest
	},
) (result *model.UniversityTeachingUnit, errCode int, err error) {
	result, errCode, err = controller.Service.Update(
		helpers.GetJwtContext(ctx), input.ID,
		&model.UniversityTeachingUnit{
			SchoolID:   input.Body.SchoolID,
			DomainID:   input.Body.DomainID,
			LevelID:    input.Body.LevelID,
			SemesterID: input.Body.SemesterID,

			Name:         input.Body.Name,
			Description:  input.Body.Description,
			Credit:       input.Body.Credit,
			Program:      input.Body.Program,
			Requirements: input.Body.Requirements,

			IsValid: input.Body.IsValid,
		},
	)
	return
}

func (controller *Controller) Delete(
	ctx *context.Context,
	input *struct {
		data.TeachingUnitID
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
		data.TeachingUnitID
	},
) (result *model.UniversityTeachingUnit, errCode int, err error) {
	teachingUnit, errCode, err := controller.Service.Get(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = teachingUnit
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllRequest
	},
) (result *data.TeachingUnitResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	teachingUnitList, errCode, err := controller.Service.GetAll(helpers.GetJwtContext(ctx), newFilter, newPagination, input.GetAllRequest.SchoolID)
	if err != nil {
		return
	}
	result = &data.TeachingUnitResponseList{
		Data: model.ToTeachingUnitResponseList(teachingUnitList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
