package telegram

import (
	"api/common/constants"
	httpHelper "api/common/helpers/http"
	"api/common/types"
	"api/services/common/telegram/data"
	"api/services/school/common/school"
	"fmt"
	"net/http"
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
	// Safety check
	if request.Update.Message.Text != "/id" || request.Update.Message.Chat.ID == 0 {
		return
	}

	// Get school
	schoolFound, err := service.SchoolService.Repository.GetByID(SchoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if schoolFound == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Send message
	chatID := request.Update.Message.Chat.ID
	firstName := request.Update.Message.From.FirstName
	username := request.Update.Message.From.Username
	message := fmt.Sprintf(
		"Hello %s (@%s) 👋\nYour Telegram Chat ID is: %d\n\nPlease enter this ID in your school profile.",
		firstName, username, chatID,
	)

	err = sendTelegramMessage(schoolFound.Config.TelegramBotToken, chatID, message)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
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
