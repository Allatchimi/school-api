package telegramHelper

import (
	"api/common/helpers"
	"api/config"
	"api/services/user/user/model"
	"fmt"

	"go.uber.org/zap"
)

func SendMessage(botToken string, message string, users []model.User) (err error) {
	if len(botToken) < 1 {
		errMsg := "Invalid bot token!"
		err = fmt.Errorf("%s", errMsg)
		return
	}
	go safeStartWorker(config.RedisClient, botToken)

	for _, user := range users {
		job := &TelegramJob{
			BotToken:   botToken,
			ChatID:     user.Config.TelegramChatID,
			Message:    message,
			Attempts:   0,
			MaxAttempt: 3,
		}
		errJob := pushJob(config.RedisClient, job)
		if errJob != nil {
			helpers.Logger.Error("Queue error: user %d - %v", zap.Int64("User ID", user.ID), zap.Error(errJob))
		}

	}
	return
}

func SetTelegramWebhook(schoolID int64, botToken string) (data any, err error) {
	if len(botToken) < 1 {
		errMsg := "Invalid bot token!"
		err = fmt.Errorf("%s", errMsg)
		return
	}
	data, err = postHttpWebhook(schoolID, botToken)
	return
}
