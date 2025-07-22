package serviceHelper

import (
	"api/common/helpers"
	smtpHelper "api/common/helpers/message/mail/smtp"
	telegramHelper "api/common/helpers/message/telegram"
	whatsappHelper "api/common/helpers/message/whatsapp"
	"api/common/utils"
	webpushConfig "api/config/webpush"
	"api/services/common/notification"
	modelSchool "api/services/school/common/school/model"
	"api/services/user/user"
	modelUser "api/services/user/user/model"
	"time"

	"go.uber.org/zap"
)

type MessageRequest struct {
	PusNotification bool
	Telegram        bool
	Whatsapp        bool
	Mail            bool
}

var UserService *user.Service
var NotificationService *notification.Service

func InjectServices(
	userSvc *user.Service,
	notificationSvc *notification.Service,
) {
	UserService = userSvc
	NotificationService = notificationSvc
}

func GetUserByID(userID int64) (result *modelUser.User, err error) {
	result, err = UserService.Repository.GetByID(userID)
	return
}

func SendMessage(
	request *MessageRequest,
	messageTitle string,
	messageBody string,
	school *modelSchool.School,
	users []modelUser.User,
) {
	if request == nil || len(users) < 1 {
		return
	}

	helpers.Logger.Info("Preparing to send message to users: ", zap.Int("count", len(users)), zap.Any("data", users))

	// Send push notification
	go func() {
		if len(messageTitle) > 0 {
			createdAt := new(time.Time)
			*createdAt = time.Now()
			webpushConfig.SendPushNotificationToUserBulk(
				users,
				&webpushConfig.WebPushPayload{
					Title:     messageTitle,
					Body:      messageBody,
					Icon:      school.LogoUrl(),
					Url:       school.WebsiteUrl(),
					CreatedAt: createdAt,
				},
				nil,
				UserService.Repository,
				NotificationService.Repository,
			)
		}
	}()
	// Send telegram
	go func() {
		if school != nil && school.Config != nil && len(school.Config.TelegramBotToken) > 0 {
			if len(messageBody) > 1 {
				telegramHelper.SendMessage(
					school.Config.TelegramBotToken,
					messageBody,
					users,
				)
			}
		}
	}()
	// Send whatsapp
	go func() {
		if school != nil && school.Config != nil && len(school.Config.WhatsappToken) > 0 && len(school.Config.WhatsappPhoneID) > 0 {
			if len(messageBody) > 1 {
				whatsappHelper.SendMessage(
					school.Config.WhatsappToken,
					school.Config.WhatsappPhoneID,
					messageBody,
					users,
				)
			}
		}
	}()
	// Send mail
	go func() {
		mailData := &smtpHelper.EmailData{
			HomePageLink: school.WebsiteUrl(),
			Logo:         school.LogoUrl(),
			Title:        messageTitle,
			Message:      messageBody,
		}
		mailBody, errTemplate := mailData.LoadTemplate()
		if errTemplate != nil {
			helpers.Logger.Error("Failed to load email template!", zap.Error(errTemplate))
		}
		if len(mailBody) < 1 {
			helpers.Logger.Warn("Empty message body!")
		}
		mailUsers := make([]string, 0, len(users))
		for _, user := range users {
			if utils.IsEmailValid(user.Email) {
				mailUsers = append(mailUsers, user.Email)
			}
		}
		errMail := smtpHelper.SendEmailBCC(mailUsers, messageTitle, mailBody)
		if errMail != nil {
			helpers.Logger.Error("Failed to send mail!", zap.Error(errMail))
		}
	}()
}
