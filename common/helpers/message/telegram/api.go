package telegramHelper

import (
	httpHelper "api/common/helpers/http"
	"api/config"
	"fmt"
)

func postHttpMessage(job *TelegramJob) (data any, err error) {
	payload := map[string]any{
		"chat_id": job.UserID,
		"text":    job.Message,
	}
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", job.BotToken)
	headers := []httpHelper.HttpHeader{
		{Label: "Content-Type", Value: "application/json"},
	}
	err = httpHelper.HttpPost(url, headers, payload, data)
	return
}

func postHttpWebhook(schoolID int64, botToken string) (data any, err error) {
	webhookURL := fmt.Sprintf("%s%s/telegram/webhook/%d", config.Env.ApiBaseURL, config.Env.ApiGroup, schoolID)
	payload := map[string]string{
		"url": webhookURL,
	}
	url := fmt.Sprintf("https://api.telegram.org/bot%s/setWebhook", botToken)
	headers := []httpHelper.HttpHeader{
		{Label: "Content-Type", Value: "application/json"},
	}
	err = httpHelper.HttpPost(url, headers, payload, data)
	return
}
