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
	Audience string

	PusNotification bool
	Telegram        bool
	Whatsapp        bool
	Mail            bool

	WhatsappTemplate   string
	WhatsappBodyParams []string
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

	helpers.Logger.Info(
		"Preparing to send message to users!",
		zap.Int("count", len(users)),
		zap.String("audience", request.Audience),
		zap.Bool("pushNotification", request.PusNotification),
		zap.Bool("telegram", request.Telegram),
		zap.Bool("whatsapp", request.Whatsapp),
		zap.Bool("mail", request.Mail),
	)

	// Send push notification
	go func() {
		if !request.PusNotification {
			return
		}
		createdAt := new(time.Time)
		*createdAt = time.Now()
		helpers.Logger.Info("Sending push notification message!")
		webpushConfig.SendPushNotificationToUserBulk(
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
	}()

	// Send telegram
	go func() {
		if !(request.Telegram && school != nil && school.Config != nil &&
			len(school.Config.TelegramBotToken) > 0) {
			return
		}
		helpers.Logger.Info("Sending telegram message!")
		telegramHelper.SendMessage(
			school.Config.TelegramBotToken,
			messageBody,
			users,
		)
	}()

	// Send whatsapp
	go func() {
		if !(request.Whatsapp &&
			school != nil &&
			school.Config != nil &&
			len(school.Config.WhatsappToken) > 0 &&
			len(school.Config.WhatsappPhoneID) > 0 &&
			users != nil && len(users) > 0) {
			return
		}
		helpers.Logger.Info("Sending whatsapp template message!",
			zap.String("token", school.Config.WhatsappToken),
			zap.String("phoneID", school.Config.WhatsappPhoneID),
			zap.String("template", request.WhatsappTemplate),
			zap.Int("userCount", len(users)),
		)
		whatsappHelper.SendMessage(
			school.Config.WhatsappToken,
			school.Config.WhatsappPhoneID,
			request.WhatsappTemplate,
			request.WhatsappBodyParams,
			users,
		)
	}()

	// Send mail
	go func() {
		if !request.Mail {
			return
		}
		helpers.Logger.Info("Sending mail!")
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
