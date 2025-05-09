package unit

import (
	"context"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/university/unit/data"
	"api/services/school/university/unit/model"
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
		Body data.UnitRequest
	},
) (result *model.UniversityUnit, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		helpers.GetJwtContext(ctx),
		&model.UniversityUnit{
			SchoolID:      input.Body.SchoolID,
			LevelDomainID: input.Body.LevelDomainID,
			SemesterID:    input.Body.SemesterID,

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
		data.UnitID
		Body data.UnitRequest
	},
) (result *model.UniversityUnit, errCode int, err error) {
	result, errCode, err = controller.Service.Update(
		helpers.GetJwtContext(ctx), input.ID,
		&model.UniversityUnit{
			SchoolID:      input.Body.SchoolID,
			LevelDomainID: input.Body.LevelDomainID,
			SemesterID:    input.Body.SemesterID,

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
		data.UnitID
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
		data.UnitID
	},
) (result *model.UniversityUnit, errCode int, err error) {
	unit, errCode, err := controller.Service.Get(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = unit
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllRequest
	},
) (result *data.UnitResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	unitList, errCode, err := controller.Service.GetAll(helpers.GetJwtContext(ctx), newFilter, newPagination, input.GetAllRequest.SchoolID)
	if err != nil {
		return
	}
	result = &data.UnitResponseList{
		Data: model.ToUnitResponseList(unitList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
