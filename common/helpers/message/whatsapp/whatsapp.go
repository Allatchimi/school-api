package whatsappHelper

import (
	"api/common/helpers"
	"api/config"
	"api/services/user/user/model"
	"fmt"

	"go.uber.org/zap"
)

func SendMessage(accessToken, phoneID string, message string, users []model.User) error {
	if len(accessToken) < 1 || len(phoneID) < 1 {
		errMsg := "Invalid WhatsApp access token or phone ID"
		return fmt.Errorf("%s", errMsg)
	}

	go safeStartWorker(config.RedisClient, accessToken, phoneID)

	for _, user := range users {
		if user.Config == nil || user.Config.WhatsappPhoneNumber < 1 {
			continue
		}

		helpers.Logger.Info("Sending whatsapp message to user: ", zap.Int64("userID", user.ID), zap.String("email", user.Email))

		job := &WhatsAppJob{
			ReceiverPhoneNumber: fmt.Sprintf("%d", user.Config.WhatsappPhoneNumber),
			Message:             message,
			AccessToken:         accessToken,
			PhoneID:             phoneID,
			Attempts:            0,
			MaxAttempt:          3,
		}

		err := pushJob(config.RedisClient, job)
		if err != nil {
			helpers.Logger.Error("Queue error WhatsApp", zap.Int64("User ID", user.ID), zap.Int64("Phone Number", user.Config.WhatsappPhoneNumber), zap.Error(err))
		}
	}
	return nil
}
