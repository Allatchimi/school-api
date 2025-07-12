package telegram

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"api/common/constants"
	"api/common/types"
	"api/services/common/telegram/data"
)

func RegisterEndpoints(
	humaApi *huma.API,
	controller *Controller,
) {
	var endpointConfig = types.ApiEndpointConfig{
		Group: "/telegram",
		Tag:   []string{"Telegram"},
	}

	// Bot webhook
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "post-telegram-webhook",
			Summary:       "Post telegram webhook",
			Description:   "Post existing telegram webhook from telegram bot",
			Method:        http.MethodPost,
			Path:          fmt.Sprintf("%s/webhook/{schoolID}", endpointConfig.Group),
			Tags:          endpointConfig.Tag,
			MaxBodyBytes:  constants.DefaultBodySize,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				data.TelegramBotSchoolID
				Body data.TelegramWebhookRequest
			},
		) (*struct{ Body types.DefaultResponse }, error) {
			errCode, err := controller.PostWebhook(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DefaultResponse }{Body: types.DefaultResponse{
				Message: "Successful received!",
			}}, nil
		},
	)
}
