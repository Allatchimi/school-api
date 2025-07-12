package test

import (
	"context"

	"api/common/helpers"
	googleMailHelper "api/common/helpers/message/mail/google"
	smtpMailHelper "api/common/helpers/message/mail/smtp"
	telegramHelper "api/common/helpers/message/telegram"
	whatsappHelper "api/common/helpers/message/whatsapp"
	"api/services/user/user/model"

	"go.uber.org/zap"
)

func Testssssss() {
	// Send Telegram message
	telegramHelper.SendMessage(
		"7676549051:AAF4u-ElGxwzarPY2EAul6YSdCwwKjxLItk",
		"Welcome Prosper! Nice to see you.",
		[]model.User{
			{
				Config: &model.UserConfig{
					TelegramChatID: 123456789,
				},
			},
		},
	)

	// Send WhatsApp message
	whatsappHelper.SendMessage(
		"ElGxwzarPY2EAul6YSdCwwKjxLItk",
		"7676549051",
		"Welcome Prosper! Nice to see you.",
		[]model.User{
			{
				Config: &model.UserConfig{
					WhatsappPhoneNumber: 237696666666,
				},
			},
		},
	)

	// Create Google user
	ctx := context.Background()
	admin := &model.User{
		Email: "admin@emfi.cm",
	}
	user := &model.User{
		Email: "prosper.abouar@gmail.com",
		Info: &model.UserInfo{
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
	data := &smtpMailHelper.EmailDataCheckCode{
		EmailData: smtpMailHelper.EmailData{
			HomePageLink: "https://digitcore.cm",
			Logo:         "https://static-cdn.jtvnw.net/growth-assets/email_twitch_logo_uv",
			Title:        "Welcome to EMFI!",
			Message:      "Thank you for signing up. Use the code below to verify your account.",
		},
		Code:            "728491",
		DurationMinutes: 10,
	}
	htmlBody, errMail := data.LoadTemplate()
	if errMail != nil {
		helpers.Logger.Warn(
			"Failed to load template!",
			zap.Error(errMail))
	}
	errMail = smtpMailHelper.SendEmailTo("support@emfi.cm", "EMFI support", "prosper.abouar@gmail.com", "Account verification", htmlBody)
	if errMail != nil {
		helpers.Logger.Warn(
			"Failed to send email!",
			zap.Error(errMail))
	}
	helpers.Logger.Info("Email sent!")

	// school := &model.School{
	// 	Type:      "university",
	// 	Favicon:   "https://www.google.com/favicon.ico",
	// 	Logo:      "https://www.gstatic.com/marketing-cms/assets/images/c5/3a/200414104c669203c62270f7884f/google-wordmarks-2x.webp=n-w100-h32-fcrop64=1,00000000ffffffff-rw",
	// 	LogoWhite: "https://www.gstatic.com/marketing-cms/assets/images/c5/3a/200414104c669203c62270f7884f/google-wordmarks-2x.webp=n-w100-h32-fcrop64=1,00000000ffffffff-rw",
	// 	Config: &model.SchoolConfig{
	// 		DomainName:          "www.uy1.cm",
	// 		WebsiteTitle:        "UY1",
	// 		WebsiteDescription:  "School management app",
	// 		ColorPrimary:        "#111111",
	// 		ColorPrimaryBg:      "#F1F1F1",
	// 		ColorPrimaryBgHover: "#D1D1D1",
	// 	},
	// }
	// school.ID = 2
	// configDeploy.DeploySchool(school)
	// configDeploy.DeleteSchoolDeployment(1)
	// configDeploy.DeleteSchoolDeployment(2)
}
