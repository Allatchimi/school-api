package data

import (
	"time"
)

type SchoolID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"School id"`
}

type SchoolRequest struct {
	Name   string `json:"name" required:"true" minLength:"2" maxLength:"50" doc:"School name"`
	Type   string `json:"type" required:"true" minLength:"2" maxLength:"50" enum:"highschool,university" doc:"School type"`
	Status string `json:"status" required:"true" minLength:"2" maxLength:"50" enum:"enabled,disabled" doc:"School status"`

	Favicon   string `json:"favicon" required:"false" doc:"School favicon"`
	Logo      string `json:"logo" required:"false" doc:"School logo"`
	LogoWhite string `json:"logoWhite" required:"false" doc:"School logo white"`

	Currency     string `json:"currency" required:"true" minLength:"2" maxLength:"20" doc:"Currency"`
	PaymentCount int64  `json:"paymentCount" required:"true" min:"1" max:"10" doc:"Payment count"`

	Info   *SchoolInfoRequest   `json:"info" required:"true" doc:"Information"`
	Config *SchoolConfigRequest `json:"config" required:"true" doc:"Configuration"`
}

type SchoolInfoRequest struct {
	FullName    string `json:"fullName" required:"true" minLength:"2" maxLength:"150" doc:"School name"`
	Description string `json:"description" required:"false" maxLength:"500" doc:"Description"`
	Motto       string `json:"motto" required:"false" maxLength:"150" doc:"Motto"`

	PhoneNumber1 int64 `json:"phoneNumber1" required:"false" doc:"Phone number 1"`
	PhoneNumber2 int64 `json:"phoneNumber2" required:"false" doc:"Phone number 2"`
	PhoneNumber3 int64 `json:"phoneNumber3" required:"false" doc:"Phone number 3"`

	Email1 string `json:"email1" required:"false" maxLength:"150" doc:"Email 1"`
	Email2 string `json:"email2" required:"false" maxLength:"150" doc:"Email 2"`
	Email3 string `json:"email3" required:"false" maxLength:"150" doc:"Email 3"`

	Founder   string     `json:"founder" required:"false" maxLength:"150" doc:"Founder name"`
	FoundedAt *time.Time `json:"foundedAt" required:"false" doc:"Founded date time"`

	Address           string  `json:"address" required:"false" maxLength:"150" doc:"Address"`
	PoBox             string  `json:"poBox" required:"false" doc:"PO Box"`
	LocationLongitude float64 `json:"locationLongitude" required:"false" doc:"Location longitude"`
	LocationLatitude  float64 `json:"locationLatitude" required:"false" doc:"Location latitude"`

	SocialMediaTelegram string `json:"socialMediaTelegram" required:"false" doc:"Social media telegram"`
	SocialMediaWhasapp  string `json:"socialMediaWhasapp" required:"false" doc:"Social media whasapp"`
	SocialMediaYoutube  string `json:"socialMediaYoutube" required:"false" doc:"Social media youtube"`
	SocialMediaTwitter  string `json:"socialMediaTwitter" required:"false" doc:"Social media twitter"`
	SocialMediaFacebook string `json:"socialMediaFacebook" required:"false" doc:"Social media facebook"`

	Image1 string `json:"image1" required:"false" doc:"Image 1"`
	Image2 string `json:"image2" required:"false" doc:"Image 2"`
	Image3 string `json:"image3" required:"false" doc:"Image 3"`
	Image4 string `json:"image4" required:"false" doc:"Image 4"`
	Image5 string `json:"image5" required:"false" doc:"Image 5"`
}

type SchoolConfigRequest struct {
	DomainName   string `json:"domainName" required:"true" doc:"Domain name" example:".digitschool.cm"`
	SupportEmail string `json:"supportEmail" required:"false" doc:"Support email" example:"support@digitschool.cm"`

	GoogleWorkspaceCredentials     string `json:"googleWorkspaceCredentials" required:"false" doc:"Google workspace credentials" example:"googleWorkspaceCredentials"`
	GoogleWorkspaceUserEmailDomain string `json:"googleWorkspaceUserEmailDomain" required:"false" doc:"Google workspace user email domain" example:"googleWorkspaceUserEmailDomain"`

	SmsUserID        string `json:"smsUserID" required:"false" doc:"Sms user ID" example:"smsUserID"`
	WhatsappToken    string `json:"whatsappToken" required:"false" doc:"Whatsapp token" example:"whatsappToken"`
	WhatsappPhoneID  string `json:"whatsappPhoneID" required:"false" doc:"Whatsapp phone ID" example:"whatsappPhoneID"`
	TelegramBotToken string `json:"telegramBotToken" required:"false" doc:"Telegram bot token" example:"telegramBotToken"`

	WebsiteTitle       string `json:"websiteTitle" required:"true" doc:"Website title" example:"websiteTitle"`
	WebsiteDescription string `json:"websiteDescription" required:"true" doc:"Website description" example:"websiteDescription"`

	ColorPrimary        string `json:"colorPrimary" required:"true" doc:"Primary color in HEX" example:"#FFFFFF"`
	ColorPrimaryBg      string `json:"colorPrimaryBg" required:"true" doc:"Primary color in HEX" example:"#FFFFFF"`
	ColorPrimaryBgHover string `json:"colorPrimaryBgHover" required:"true" doc:"Primary color in HEX" example:"#FFFFFF"`
}

type SchoolDeploymentStatusRequest struct {
	Status   string `json:"status" required:"true" enum:"done,failed" doc:"Deployment status"`
	Feedback string `json:"feedback" required:"false" doc:"Deployment feedback"`
}

type GetAllRequest struct {
	Type string `json:"type" query:"type" required:"false" enum:"highschool,university" doc:"School type"`
}
