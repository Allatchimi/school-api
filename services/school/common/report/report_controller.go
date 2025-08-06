package report

import (
	"context"

	"api/common/helpers"
	httpHelper "api/common/helpers/http"
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

func (controller *Controller) CreateEntry(
	ctx *context.Context,
	input *struct {
		Body data.ReportEntryRequest
	},
) (result *model.ReportEntry, errCode int, err error) {
	result, errCode, err = controller.Service.CreateEntry(
		httpHelper.GetContextData(ctx),
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
		httpHelper.GetContextData(ctx),
		&input.Body,
	)
	return
}

func (controller *Controller) CreateCorrespondence(
	ctx *context.Context,
	input *struct {
		Body data.ReportCorrespondenceRequest
	},
) (result *model.ReportCorrespondence, errCode int, err error) {
	result, errCode, err = controller.Service.CreateCorrespondence(
		httpHelper.GetContextData(ctx),
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
		httpHelper.GetContextData(ctx),
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
		httpHelper.GetContextData(ctx),
		input.ID,
		&input.Body,
	)
	return
}

func (controller *Controller) UpdateCorrespondence(
	ctx *context.Context,
	input *struct {
		data.ReportCorrespondenceID
		Body data.ReportCorrespondenceRequest
	},
) (result *model.ReportCorrespondence, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateCorrespondence(
		httpHelper.GetContextData(ctx),
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
		httpHelper.GetContextData(ctx),
		input.ID,
		&input.Body,
	)
	return
}

func (controller *Controller) UpdateTableStatus(
	ctx *context.Context,
	input *struct {
		data.ReportTableID
		Body data.ReportTableStatusRequest
	},
) (result *model.ReportTable, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateTableStatus(
		httpHelper.GetContextData(ctx),
		input.ID,
		&input.Body,
	)
	return
}

func (controller *Controller) DeleteGrade(
	ctx *context.Context,
	input *struct {
		data.ReportGradeID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteGrade(httpHelper.GetContextData(ctx), input.ID)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) DeleteCorrespondence(
	ctx *context.Context,
	input *struct {
		data.ReportCorrespondenceID
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteCorrespondence(httpHelper.GetContextData(ctx), input.ID)
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
	affectedRows, errCode, err := controller.Service.DeleteConfig(httpHelper.GetContextData(ctx), input.ID)
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
	affectedRows, errCode, err := controller.Service.DeleteMultipleGrade(httpHelper.GetContextData(ctx), input.Body.List)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) DeleteMultipleCorrespondence(
	ctx *context.Context,
	input *struct {
		Body types.DeleteMultipleRequest
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.DeleteMultipleCorrespondence(httpHelper.GetContextData(ctx), input.Body.List)
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
	affectedRows, errCode, err := controller.Service.DeleteMultipleConfig(httpHelper.GetContextData(ctx), input.Body.List)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) GetEntry(
	ctx *context.Context,
	input *struct {
		data.ReportEntryID
	},
) (result *model.ReportEntry, errCode int, err error) {
	result, errCode, err = controller.Service.GetEntry(httpHelper.GetContextData(ctx), input.ID)
	return
}

func (controller *Controller) GetGrade(
	ctx *context.Context,
	input *struct {
		data.ReportGradeID
	},
) (result *model.ReportGrade, errCode int, err error) {
	result, errCode, err = controller.Service.GetGrade(httpHelper.GetContextData(ctx), input.ID)
	return
}

func (controller *Controller) GetCorrespondence(
	ctx *context.Context,
	input *struct {
		data.ReportCorrespondenceID
	},
) (result *model.ReportCorrespondence, errCode int, err error) {
	result, errCode, err = controller.Service.GetCorrespondence(httpHelper.GetContextData(ctx), input.ID)
	return
}

func (controller *Controller) GetConfig(
	ctx *context.Context,
	input *struct {
		data.ReportConfigID
	},
) (result *model.ReportConfig, errCode int, err error) {
	result, errCode, err = controller.Service.GetConfig(httpHelper.GetContextData(ctx), input.ID)
	return
}

func (controller *Controller) GetAllEntry(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllReportEntryRequest
	},
) (result *data.ReportEntryResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	resultList, errCode, err := controller.Service.GetAllEntry(httpHelper.GetContextData(ctx), newFilter, newPagination, &input.GetAllReportEntryRequest)
	if err != nil {
		return
	}
	result = &data.ReportEntryResponseList{
		Data: model.ToReportEntryResponseList(resultList),
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
	resultList, errCode, err := controller.Service.GetAllGrade(httpHelper.GetContextData(ctx), newFilter, newPagination, &input.GetAllReportGradeRequest)
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

func (controller *Controller) GetAllCorrespondence(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllReportCorrespondenceRequest
	},
) (result *data.ReportCorrespondenceResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	resultList, errCode, err := controller.Service.GetAllCorrespondence(httpHelper.GetContextData(ctx), newFilter, newPagination, &input.GetAllReportCorrespondenceRequest)
	if err != nil {
		return
	}
	result = &data.ReportCorrespondenceResponseList{
		Data: model.ToReportCorrespondenceResponseList(resultList),
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
	resultList, errCode, err := controller.Service.GetAllConfig(httpHelper.GetContextData(ctx), newFilter, newPagination, &input.GetAllReportConfigRequest)
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

func (controller *Controller) GetAllAverage(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllReportAverageRequest
	},
) (result *data.ReportAverageResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	resultList, errCode, err := controller.Service.GetAllAverage(httpHelper.GetContextData(ctx), newFilter, newPagination, &input.GetAllReportAverageRequest)
	if err != nil {
		return
	}
	result = &data.ReportAverageResponseList{
		Data: model.ToReportAverageResponseList(resultList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}

func (controller *Controller) GetAllTable(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllReportTableRequest
	},
) (result *data.ReportTableResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	resultList, errCode, err := controller.Service.GetAllTable(httpHelper.GetContextData(ctx), newFilter, newPagination, &input.GetAllReportTableRequest)
	if err != nil {
		return
	}
	result = &data.ReportTableResponseList{
		Data: model.ToReportTableResponseList(resultList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
