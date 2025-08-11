package serviceHelperMessage

import (
	"api/common/helpers"
	smtpHelper "api/common/helpers/message/mail/smtp"
	telegramHelper "api/common/helpers/message/telegram"
	whatsappHelper "api/common/helpers/message/whatsapp"
	"api/common/utils"
	webpushConfig "api/config/webpush"
	"api/services/others/notification"
	modelSchool "api/services/school/common/school/model"
	"api/services/user/user"
	modelUser "api/services/user/user/model"
	"time"

	"go.uber.org/zap"
)

type MessageRequest struct {
	Audience        string
	PusNotification bool
	Telegram        bool
	Whatsapp        bool
	Mail            bool
}

var UserService *user.Service
var NotificationService *notification.Service

func InjectServices(
	userService *user.Service,
	notificationService *notification.Service,
) {
	UserService = userService
	NotificationService = notificationService
}

func SendMessage(
	request *MessageRequest,
	messageTitle string,
	messageBody string,
	school *modelSchool.School,
	urlPath string,
	users []modelUser.User,
) {
	if request == nil || len(users) < 1 {
		helpers.Logger.Info("No users to send message! Skipped!")
		return
	}

	helpers.Logger.Info("Preparing to send message to users: ", zap.Int("count", len(users)))

	// Send push notification
	if request.PusNotification && len(messageTitle) > 0 {
		createdAt := new(time.Time)
		*createdAt = time.Now()
		go webpushConfig.SendPushNotificationToUserBulk(
			users,
			&webpushConfig.WebPushPayload{
				Title:     messageTitle,
				Body:      messageBody,
				Icon:      school.LogoUrl(),
				Url:       school.WebsiteUrl() + urlPath,
				CreatedAt: createdAt,
			},
			nil,
			UserService.Repository,
			NotificationService.Repository,
		)
	}

	// Send telegram
	if request.Telegram && school != nil && school.Config != nil &&
		len(school.Config.TelegramBotToken) > 0 && len(messageBody) > 0 {
		go telegramHelper.SendMessage(
			school.Config.TelegramBotToken,
			messageBody,
			users,
		)
	}

	// Send whatsapp
	if request.Whatsapp && school.Config != nil && len(school.Config.WhatsappToken) > 0 &&
		len(school.Config.WhatsappPhoneID) > 0 && len(messageBody) > 0 {
		go whatsappHelper.SendMessage(
			school.Config.WhatsappToken,
			school.Config.WhatsappPhoneID,
			messageBody,
			users,
		)
	}

	// Send mail
	if request.PusNotification {
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
}
