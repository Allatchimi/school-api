package telegram

import (
	"context"

	httpHelper "api/common/helpers/http"
	"api/services/others/telegram/data"
)

type Controller struct {
	Service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{Service: service}
}

func (controller *Controller) PostWebhook(
	ctx *context.Context,
	input *struct {
		data.TelegramBotSchoolID
		Body data.TelegramWebhookRequest
	},
) (errCode int, err error) {
	errCode, err = controller.Service.PostWebhook(
		httpHelper.GetContextData(ctx),
		input.TelegramBotSchoolID.SchoolID,
		&input.Body,
	)
	return
}
