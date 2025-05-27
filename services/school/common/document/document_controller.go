package document

import (
	"context"

	"api/common/helpers"
	"api/common/types"
	"api/services/school/common/document/data"
	"api/services/school/common/document/model"
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
		Body data.DocumentRequest
	},
) (result *model.Document, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		helpers.GetJwtContext(ctx),
		&model.Document{
			SchoolID:       input.Body.SchoolID,
			YearID:         input.Body.YearID,
			ClassSubjectID: input.Body.ClassSubjectID,
			UnitID:         input.Body.UnitID,
		},
	)
	return
}

func (controller *Controller) Delete(
	ctx *context.Context,
	input *struct {
		data.DocumentID
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
		data.DocumentID
	},
) (result *model.Document, errCode int, err error) {
	document, errCode, err := controller.Service.Get(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = document
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		types.FilterlSchoolYearUnitClassSubjectRequest
	},
) (result *data.DocumentResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	documentList, errCode, err := controller.Service.GetAll(
		helpers.GetJwtContext(ctx), newFilter, newPagination,
		input.SchoolID,
		input.YearID,
		input.UnitID,
		input.ClassSubjectID,
	)
	if err != nil {
		return
	}
	result = &data.DocumentResponseList{
		Data: model.ToResponseList(documentList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
