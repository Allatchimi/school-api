package initialize

import (
	"context"

	httpHelper "api/common/helpers/http"
	"api/services/others/initialize/data"
)

type Controller struct {
	Service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{Service: service}
}

func (controller *Controller) GetInitial(
	ctx *context.Context,
) (result *data.InitializeResponse, errCode int, err error) {
	school, errCode, err := controller.Service.GetInitial(httpHelper.GetContextData(ctx))
	if err != nil {
		return
	}
	result = school
	return
}
