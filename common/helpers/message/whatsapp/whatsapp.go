package whatsappHelper

import (
	"api/common/helpers"
	"api/config"
	"api/services/user/user/model"
	"fmt"

	"go.uber.org/zap"
)

func SendMessage(accessToken, phoneID string, template string, bodyParams []string, users []model.User) error {
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

		var language string
		if user.Info != nil {
			language = user.Info.Language
		}
		job := &WhatsAppJob{
			AccessToken:         accessToken,
			PhoneID:             phoneID,
			ReceiverPhoneNumber: fmt.Sprintf("%d", user.Config.WhatsappPhoneNumber),

			Template:   template,
			TTLSeconds: 3600,
			BodyParams: bodyParams,
			Language:   language,

			Attempts:   0,
			MaxAttempt: 3,
		}

		err := pushJob(config.RedisClient, job)
		if err != nil {
			helpers.Logger.Error("Queue error WhatsApp", zap.Int64("User ID", user.ID), zap.Int64("Phone Number", user.Config.WhatsappPhoneNumber), zap.Error(err))
		}
	}
	return nil
}
