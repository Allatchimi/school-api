package health

import (
	"context"

	"api/common/helpers"
	"api/services/health/data"
)

type Controller struct {
	Service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{Service: service}
}

func (controller *Controller) GetPostgresHealth(
	ctx *context.Context,
	input *struct {
		data.HealthRequest
	},
) (result bool, errCode int, err error) {
	health, errCode, err := controller.Service.GetPostgresHealth(helpers.GetJwtContext(ctx))
	if err != nil {
		return
	}
	result = health
	return
}

func (controller *Controller) GetRedisHealth(
	ctx *context.Context,
	input *struct {
		data.HealthRequest
	},
) (result bool, errCode int, err error) {
	health, errCode, err := controller.Service.GetRedisHealth(helpers.GetJwtContext(ctx))
	if err != nil {
		return
	}
	result = health
	return
}
