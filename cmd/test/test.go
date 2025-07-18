package test

import (
	"context"
	"fmt"

	"api/common/constants"
	"api/common/helpers"
	configDeploy "api/common/helpers/deployment"
	googleMailHelper "api/common/helpers/message/mail/google"
	smtpHelper "api/common/helpers/message/mail/smtp"
	telegramHelper "api/common/helpers/message/telegram"
	whatsappHelper "api/common/helpers/message/whatsapp"
	modelSchool "api/services/school/common/school/model"
	modelUser "api/services/user/user/model"

	"go.uber.org/zap"
)

func Testssssss() {
	// Create Google user
	ctx := context.Background()
	admin := &modelUser.User{
		Email: "admin@emfi.cm",
	}
	user := &modelUser.User{
		Email: "prosper.abouar@gmail.com",
		Info: &modelUser.UserInfo{
			FirstName: "Prosper",
			LastName:  "Abouar",
			Gender:    "male",
		},
	}
	errGoogle := googleMailHelper.CreateGoogleWorkspaceUser(&ctx, "", admin, user, "Cpasbien123!")
	if errGoogle != nil {
		helpers.Logger.Warn(
			"Failed to create Google user!",
			zap.Error(errGoogle))
	}

	// Sent mail
	data := &smtpHelper.EmailDataCheckCode{
		EmailData: smtpHelper.EmailData{
			HomePageLink: fmt.Sprintf("https://%s", "digitcore.cm"),
			Logo:         "https://static-cdn.jtvnw.net/growth-assets/email_twitch_logo_uv",
			Title:        constants.MailVerifyEmailCheckCode.Title,
			Message:      constants.MailVerifyEmailCheckCode.Message,
		},
		Code:            fmt.Sprintf("%d", 234589),
		DurationMinutes: 10,
	}
	body, err := data.LoadTemplate()
	if err != nil {
		helpers.Logger.Error("Failed to load email template", zap.Error(err))
		return
	}
	err = smtpHelper.SendEmailTo(
		"support@digitcore.cm",
		"Digitcore support",
		"prosper.abouar@gmail.com",
		constants.MailVerifyEmailCheckCode.Subject,
		body,
	)
	if err != nil {
		helpers.Logger.Error("Failed to send email", zap.Error(err))
		return
	}
	helpers.Logger.Info("Email sent successfully")

	// Deploy school
	school := &modelSchool.School{
		Type:      "university",
		Favicon:   "https://www.google.com/favicon.ico",
		Logo:      "https://www.gstatic.com/marketing-cms/assets/images/c5/3a/200414104c669203c62270f7884f/google-wordmarks-2x.webp=n-w100-h32-fcrop64=1,00000000ffffffff-rw",
		LogoWhite: "https://www.gstatic.com/marketing-cms/assets/images/c5/3a/200414104c669203c62270f7884f/google-wordmarks-2x.webp=n-w100-h32-fcrop64=1,00000000ffffffff-rw",
		Config: &modelSchool.SchoolConfig{
			DomainName:          "www.digitschool.cm",
			WebsiteTitle:        "Digitschool",
			WebsiteDescription:  "Welcome to Digitschool! The future of education. With Digitschool, you can learn anything you want, whenever you want, from anywhere you want. Just enroll and start now.",
			ColorPrimary:        "#3c6989",
			ColorPrimaryBg:      "#e4e8ea",
			ColorPrimaryBgHover: "#a3b6c6",
		},
	}
	school.ID = 1
	configDeploy.DeploySchool(school)
	// configDeploy.DeleteSchoolDeployment(1)
	// configDeploy.DeleteSchoolDeployment(2)

	// Send Telegram message
	go func() {
		telegramHelper.SendMessage(
			"7676549051:AAF4u-ElGxwzarPY2EAul6YSdCwwKjxLItk",
			"Welcome Prosper! Nice to see you.",
			[]modelUser.User{
				{
					Config: &modelUser.UserConfig{
						TelegramChatID: 123456789,
					},
				},
			},
		)
	}()

	// Send WhatsApp message
	go func() {
		whatsappHelper.SendMessage(
			"ElGxwzarPY2EAul6YSdCwwKjxLItk",
			"7676549051",
			"Welcome Prosper! Nice to see you.",
			[]modelUser.User{
				{
					Config: &modelUser.UserConfig{
						WhatsappPhoneNumber: 237696666666,
					},
				},
			},
		)
	}()
}
