package telegram

import (
	"api/common/helpers"
	httpHelper "api/common/helpers/http"
	"api/common/types"
	"api/services/others/telegram/data"
	"api/services/school/common/school"
	"fmt"

	"go.uber.org/zap"
)

type Service struct {
	SchoolService *school.Service
}

func NewService(schoolService *school.Service) *Service {
	return &Service{
		SchoolService: schoolService,
	}
}

const MODEL_NAME = "telegram"
const DEFAULT_ERROR_MESSAGE = "interact with telegram"

func (service *Service) PostWebhook(
	ctxData *types.ContextData,
	SchoolID int64,
	request *data.TelegramWebhookRequest,
) (errCode int, err error) {
	helpers.Logger.Info("Received telegram message", zap.Any("request", request))

	return
}

func sendTelegramMessage(botToken string, chatID int64, message string) (err error) {
	if len(botToken) < 1 {
		errMsg := "Invalid bot token!"
		err = fmt.Errorf("%s", errMsg)
		return
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	var resp any
	payload := map[string]any{
		"chat_id": chatID,
		"text":    message,
	}
	headers := []httpHelper.HttpHeader{
		{Label: "Content-Type", Value: "application/json"},
	}
	err = httpHelper.HttpPost(url, headers, payload, &resp)
	if err != nil {
		return
	}

	if resp.(map[string]any)["ok"] != true {
		errMsg := "Failed to send message!"
		err = fmt.Errorf("%s", errMsg)
		return
	}
	return
}
