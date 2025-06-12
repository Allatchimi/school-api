package contact

import (
	"context"

	"api/common/helpers"
	"api/common/types"
	"api/services/common/contact/data"
	"api/services/common/contact/model"
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
		Body data.ContactRequest
	},
) (result *model.Contact, errCode int, err error) {
	result, errCode, err = controller.Service.Create(
		helpers.GetJwtContext(ctx),
		&model.Contact{
			SchoolID: input.Body.SchoolID,
			Subject:  input.Body.Subject,
			Email:    input.Body.Email,
			Message:  input.Body.Message,
		},
	)
	return
}

func (controller *Controller) Get(
	ctx *context.Context,
	input *struct {
		data.ContactID
	},
) (result *model.Contact, errCode int, err error) {
	contact, errCode, err := controller.Service.Get(helpers.GetJwtContext(ctx), input.ID)
	if err != nil {
		return
	}
	result = contact
	return
}

func (controller *Controller) GetAll(
	ctx *context.Context,
	input *struct {
		types.Filter
		types.PaginationRequest
		data.GetAllRequest
	},
) (result *data.ContactResponseList, errCode int, err error) {
	newPagination, newFilter := helpers.GetPaginationFiltersFromQuery(&input.Filter, &input.PaginationRequest)
	contactList, errCode, err := controller.Service.GetAll(helpers.GetJwtContext(ctx), newFilter, newPagination, &input.GetAllRequest)
	if err != nil {
		return
	}
	result = &data.ContactResponseList{
		Data: model.ToResponseList(contactList),
	}
	result.Filter = newFilter
	result.Pagination = newPagination
	return
}
