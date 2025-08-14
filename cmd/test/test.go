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

func TestAll() {
	go func() {
		TestSpecificSendMail("prosper.abouar@gmail.com")
		TestSpecificSendMail("prosper.abouar@yahoo.fr")
		TestSpecificSendWhatsappMessage()
	}()
}

func TestSpecificSendMail(to string) {
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
		to,
		constants.MailVerifyEmailCheckCode.Subject,
		body,
	)
	if err != nil {
		helpers.Logger.Error("Failed to send email", zap.Error(err))
		return
	}
	helpers.Logger.Info("Email sent successfully")
}

func TestSpecificCreateGoogleUser() {
	// Create Google user
	ctx := context.Background()
	user := &modelUser.User{
		Email: "prosper1.abouar1@digitschool.com",
		School: &modelSchool.School{
			Config: &modelSchool.SchoolConfig{
				GoogleWorkspaceCredentials:     "googleWorkspaceCredentials",
				GoogleWorkspaceUserEmailDomain: "googleWorkspaceUserEmailDomain",
			},
		},
		Info: &modelUser.UserInfo{
			FirstName: "Prosper",
			LastName:  "Abouar",
			Gender:    "male",
		},
	}
	createUser, errGoogle := googleMailHelper.CreateGoogleWorkspaceUser(
		ctx,
		"creds.json",
		"admin@digitschool.cm",
		user,
		"Cpasbien123!",
	)
	if errGoogle != nil || createUser == nil {
		helpers.Logger.Warn(
			"Failed to create Google user!",
			zap.Error(errGoogle))
	}
}

func TestSpecificSendTelegramMessage() {
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
}

func TestSpecificSendWhatsappMessage() {
	// Send WhatsApp message
	go func() {
		whatsappHelper.SendMessage(
			"EAAVClhRHZAuYBPCQXj5IPhtHFqRmERdeZBLm72L9gyUyDeUki5aYajYiVHaycxoxpTZAHFlusjHZB87UvYOObTdCbjYYhq7io6cmZA7ZBahl9kpdCdq9AYcicVFqJtKLQC3ldVpQRTshxNOLXCyJ86JCjbhgZBdXa8SUArWAZBkFH1ggBqjh96j5oe5my2H0CMBSUJwoNTNSExHspsIuShHmFGwZChf07mB5ZAMbzzrhGbtZAr4WygZD",
			"774627652396866",
			"request_status",
			[]string{"Problem with score of INF101 ", "rejected", "DIGIT-School"},
			[]modelUser.User{
				{
					Info: &modelUser.UserInfo{
						FirstName: "Prosper",
						LastName:  "Abouar",
						Language:  "en",
					},
					Config: &modelUser.UserConfig{
						WhatsappPhoneNumber: 237693264668,
					},
				},
			},
		)
	}()
}

func TestSpecificDeploySchool() {
	// Deploy school
	school := &modelSchool.School{
		Type:      "university",
		Favicon:   "https://www.google.com/favicon.ico",
		Logo:      "https://www.gstatic.com/marketing-cms/assets/images/c5/3a/200414104c669203c62270f7884f/google-wordmarks-2x.webp=n-w100-h32-fcrop64=1,00000000ffffffff-rw",
		LogoWhite: "https://www.gstatic.com/marketing-cms/assets/images/c5/3a/200414104c669203c62270f7884f/google-wordmarks-2x.webp=n-w100-h32-fcrop64=1,00000000ffffffff-rw",
		Config: &modelSchool.SchoolConfig{
			WebsiteDomainName:   "www.digitschool.cm",
			WebsiteTitle:        "Digitschool",
			WebsiteDescription:  "Welcome to Digitschool! The future of education. With Digitschool, you can learn anything you want, whenever you want, from anywhere you want. Just enroll and start now.",
			ColorPrimary:        "#b6b43b",
			ColorPrimaryBg:      "#ebebe1",
			ColorPrimaryBgHover: "#c5c6a3",
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
}
