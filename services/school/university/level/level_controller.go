package level

import (
	"context"

	"api/common/helpers"
	httpHelper "api/common/helpers/http"
	"api/common/types"
	"api/services/school/university/level/data"
	"api/services/school/university/level/model"
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
		Body data.LevelRequest
	},
) (result *model.UniversityLevel, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		httpHelper.GetJwtContext(ctx),
		&input.Body,
	)
	return
}

func (controller *Controller) CreateLevelDomain(
	ctx *context.Context,
	input *struct {
		Body data.LevelDomainRequest
	},
) (result *model.UniversityLevelDomain, errCode int, err error) {
	result, errCode, err = controller.Service.CreateLevelDomain(
		httpHelper.GetJwtContext(ctx),
		&input.Body,
	)
	return
}

func (controller *Controller) Update(
	ctx *context.Context,
	input *struct {
		data.LevelID
		Body data.LevelRequest
	},
) (result *model.UniversityLevel, errCode int, err error) {
	result, errCode, err = controller.Service.Update(
		httpHelper.GetJwtContext(ctx),
		input.ID,
		&input.Body,
	)
	return
}

func (controller *Controller) UpdateLevelDomain(
	ctx *context.Context,
	input *struct {
		data.LevelDomainID
		Body data.LevelDomainRequest
	},
) (result *model.UniversityLevelDomain, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateLevelDomain(
		httpHelper.GetJwtContext(ctx),
		input.ID,
		&input.Body,
	)
	return
}

func (controller *Controller) Delete(
	ctx *context.Context,
	input *struct {
		data.LevelID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.Delete(httpHelper.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) DeleteLevelDomain(
	ctx *context.Context,
	input *struct {
		data.LevelDomainID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteLevelDomain(httpHelper.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) DeleteMultiple(
	ctx *context.Context,
	input *struct {
		Body types.DeleteMultipleRequest
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteMultiple(httpHelper.GetJwtContext(ctx), input.Body.List)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) DeleteMultipleLevelDomain(
	ctx *context.Context,
	input *struct {
		Body types.DeleteMultipleRequest
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteMultipleLevelDomain(httpHelper.GetJwtContext(ctx), input.Body.List)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) Get(
	ctx *context.Context,
	input *struct {
		data.LevelID
	},
) (result *model.UniversityLevel, errCode int, err error) {
	level, errCode, err := controller.Service.Get(httpHelper.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = level
	return
}

func (controller *Controller) GetLevelDomain(
	ctx *context.Context,
	input *struct {
		data.LevelDomainID
	},
) (result *model.UniversityLevelDomain, errCode int, err error) {
	levelDomain, errCode, err := controller.Service.GetLevelDomain(httpHelper.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = levelDomain
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllRequest
	},
) (result *data.LevelResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	levelList, errCode, err := controller.Service.GetAll(httpHelper.GetJwtContext(ctx), newFilter, newPagination, &input.GetAllRequest)
	if err != nil {
		return
	}
	result = &data.LevelResponseList{
		Data: model.ToResponseList(levelList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}

func (controller *Controller) GetAllLevelDomain(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllLevelDomainRequest
	},
) (result *data.LevelDomainResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	levelList, errCode, err := controller.Service.GetAllLevelDomain(httpHelper.GetJwtContext(ctx), newFilter, newPagination, &input.GetAllLevelDomainRequest)
	if err != nil {
		return
	}
	result = &data.LevelDomainResponseList{
		Data: model.ToLevelDomainResponseList(levelList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
