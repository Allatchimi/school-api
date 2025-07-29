package health

import (
	"context"

	httpHelper "api/common/helpers/http"
	"api/services/others/health/data"
)

type Controller struct {
	Service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{Service: service}
}

func (controller *Controller) HealthLive(
	ctx *context.Context,
	input *struct {
		data.HealthRequest
	},
) (result bool, errCode int, err error) {
	health, errCode, err := controller.Service.HealthLive(httpHelper.GetContextData(ctx))
	if err != nil {
		return
	}
	result = health
	return
}

func (controller *Controller) HealthDepencencies(
	ctx *context.Context,
	input *struct {
		data.HealthRequest
	},
) (result bool, errCode int, err error) {
	health, errCode, err := controller.Service.HealthDepencencies(httpHelper.GetContextData(ctx))
	if err != nil {
		return
	}
	result = health
	return
}
