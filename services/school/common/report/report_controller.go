package report

import (
	"context"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/common/report/data"
	"api/services/school/common/report/model"
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
		Body data.ReportRequest
	},
) (result *model.Report, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		helpers.GetJwtContext(ctx),
		&input.Body,
	)
	return
}

func (controller *Controller) CreateGrade(
	ctx *context.Context,
	input *struct {
		Body data.ReportGradeRequest
	},
) (result *model.ReportGrade, errCode int, err error) {
	result, errCode, err = controller.Service.CreateGrade(
		helpers.GetJwtContext(ctx),
		&input.Body,
	)
	return
}

func (controller *Controller) CreateConfig(
	ctx *context.Context,
	input *struct {
		Body data.ReportConfigRequest
	},
) (result *model.ReportConfig, errCode int, err error) {
	result, errCode, err = controller.Service.CreateConfig(
		helpers.GetJwtContext(ctx),
		&input.Body,
	)
	return
}

func (controller *Controller) UpdateGrade(
	ctx *context.Context,
	input *struct {
		data.ReportGradeID
		Body data.ReportGradeRequest
	},
) (result *model.ReportGrade, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateGrade(
		helpers.GetJwtContext(ctx),
		input.ID,
		&input.Body,
	)
	return
}

func (controller *Controller) UpdateConfig(
	ctx *context.Context,
	input *struct {
		data.ReportConfigID
		Body data.ReportConfigRequest
	},
) (result *model.ReportConfig, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateConfig(
		helpers.GetJwtContext(ctx),
		input.ID,
		&input.Body,
	)
	return
}

func (controller *Controller) Delete(
	ctx *context.Context,
	input *struct {
		data.ReportID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.Delete(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) DeleteGrade(
	ctx *context.Context,
	input *struct {
		data.ReportGradeID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteGrade(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) DeleteConfig(
	ctx *context.Context,
	input *struct {
		data.ReportConfigID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteConfig(helpers.GetJwtContext(ctx), input.ID)
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
	affectedRows, errCode, err := controller.Service.DeleteMultiple(helpers.GetJwtContext(ctx), input.Body.List)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) DeleteMultipleGrade(
	ctx *context.Context,
	input *struct {
		Body types.DeleteMultipleRequest
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteMultipleGrade(helpers.GetJwtContext(ctx), input.Body.List)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) DeleteMultipleConfig(
	ctx *context.Context,
	input *struct {
		Body types.DeleteMultipleRequest
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteMultipleConfig(helpers.GetJwtContext(ctx), input.Body.List)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) Get(
	ctx *context.Context,
	input *struct {
		data.ReportID
	},
) (result *model.Report, errCode int, err error) {
	result, errCode, err = controller.Service.Get(helpers.GetJwtContext(ctx), input.ID)
	return
}

func (controller *Controller) GetGrade(
	ctx *context.Context,
	input *struct {
		data.ReportGradeID
	},
) (result *model.ReportGrade, errCode int, err error) {
	result, errCode, err = controller.Service.GetGrade(helpers.GetJwtContext(ctx), input.ID)
	return
}

func (controller *Controller) GetConfig(
	ctx *context.Context,
	input *struct {
		data.ReportConfigID
	},
) (result *model.ReportConfig, errCode int, err error) {
	result, errCode, err = controller.Service.GetConfig(helpers.GetJwtContext(ctx), input.ID)
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllRequest
	},
) (result *data.ReportResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	resultList, errCode, err := controller.Service.GetAll(helpers.GetJwtContext(ctx), newFilter, newPagination, &input.GetAllRequest)
	if err != nil {
		return
	}
	result = &data.ReportResponseList{
		Data: model.ToReportResponseList(resultList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}

func (controller *Controller) GetAllGrade(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllReportGradeRequest
	},
) (result *data.ReportGradeResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	resultList, errCode, err := controller.Service.GetAllGrade(helpers.GetJwtContext(ctx), newFilter, newPagination, &input.GetAllReportGradeRequest)
	if err != nil {
		return
	}
	result = &data.ReportGradeResponseList{
		Data: model.ToReportGradeResponseList(resultList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}

func (controller *Controller) GetAllConfig(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllReportConfigRequest
	},
) (result *data.ReportConfigResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	resultList, errCode, err := controller.Service.GetAllConfig(helpers.GetJwtContext(ctx), newFilter, newPagination, &input.GetAllReportConfigRequest)
	if err != nil {
		return
	}
	result = &data.ReportConfigResponseList{
		Data: model.ToReportConfigResponseList(resultList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
