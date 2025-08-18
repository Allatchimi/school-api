package telegram

import (
	"api/common/constants"
	"api/common/helpers"
	httpHelper "api/common/helpers/http"
	"api/common/types"
	"api/services/others/telegram/data"
	"api/services/school/common/school"
	"fmt"
	"net/http"

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

	msg, ok := (*request)["message"].(map[string]any)
	if !ok {
		return
	}
	text, _ := msg["text"].(string)

	// Handle command
	switch text {
	case "/id":
		errCode, err = handleCommandID(service, msg, SchoolID)
	case "/start":
		errCode, err = handleCommandStart(service, msg, SchoolID)
	}

	return
}

func handleCommandID(service *Service, msg map[string]any, SchoolID int64) (errCode int, err error) {
	// Get data
	chat, ok := msg["chat"].(map[string]any)
	if !ok {
		return
	}
	chatIDFloat, ok := chat["id"].(float64)
	if !ok {
		return
	}
	chatID := int64(chatIDFloat)
	firstName, _ := chat["first_name"].(string)
	username, _ := chat["username"].(string)

	// Get school
	schoolFound, err := service.SchoolService.Repository.GetByID(SchoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if schoolFound == nil || schoolFound.ID < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Send message
	go func() {
		firstName := firstName
		username := username
		message := fmt.Sprintf(
			"Hello %s (@%s) 👋\nYour Telegram Chat ID is: %d\n\nPlease enter this ID in your school profile.",
			firstName, username, chatID,
		)

		err = sendTelegramMessage(schoolFound.Config.TelegramBotToken, chatID, message)
		if err != nil {
			helpers.Logger.Error("Error sending telegram message: %v", zap.Error(err))
			return
		}
	}()
	return
}

func handleCommandStart(service *Service, msg map[string]any, SchoolID int64) (errCode int, err error) {
	// Get data
	chat, ok := msg["chat"].(map[string]any)
	if !ok {
		return
	}
	chatIDFloat, ok := chat["id"].(float64)
	if !ok {
		return
	}
	chatID := int64(chatIDFloat)
	firstName, _ := chat["first_name"].(string)
	username, _ := chat["username"].(string)

	// Get school
	schoolFound, err := service.SchoolService.Repository.GetByID(SchoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if schoolFound == nil || schoolFound.ID < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Send message
	go func() {
		firstName := firstName
		username := username
		message := fmt.Sprintf(
			`👋 Welcome %s (@%s)!\n

			I'm here to assist you with notifications and updates.  
			If you haven't set your Chat ID yet, please use the "/id" command so we can link your account to this chat.  

			Once your Chat ID is defined, you'll start receiving updates directly here.
			`,
			firstName, username, chatID,
		)

		err = sendTelegramMessage(schoolFound.Config.TelegramBotToken, chatID, message)
		if err != nil {
			helpers.Logger.Error("Error sending telegram message: %v", zap.Error(err))
			return
		}
	}()
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
