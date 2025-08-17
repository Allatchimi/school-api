package data

import (
	"api/common/types"
	"time"
)

type SchoolResponse struct {
	SchoolPublicResponse
	Config *SchoolConfigResponse `json:"config,omitempty" required:"false" doc:"Configuration"`

	DeploymentRequest  string `json:"deploymentRequest" required:"false" doc:"Deployment request"`
	DeploymentStatus   string `json:"deploymentStatus" required:"false" doc:"Deployment status"`
	DeploymentFeedback string `json:"deploymentFeedback" required:"false" doc:"Deployment feedback"`
	DeploymentExtra    string `json:"deploymentExtra" required:"false" doc:"Deployment exta"`
	DeploymentCount    int64  `json:"deploymentVersion" required:"false" doc:"Deployment version"`
}

type SchoolPublicResponse struct {
	types.BaseGormModelResponse
	Info *SchoolInfoResponse `json:"info,omitempty" required:"false" doc:"Information"`

	Name         string `json:"name" required:"false" doc:"School name"`
	Type         string `json:"type" required:"false" doc:"Type"`
	Status       string `json:"status" required:"false" doc:"Status"`
	Favicon      string `json:"favicon" required:"false" doc:"Favicon"`
	Logo         string `json:"logo" required:"false" doc:"Logo"`
	LogoWhite    string `json:"logoWhite" required:"false" doc:"Logo white"`
	Currency     string `json:"currency" required:"false" doc:"Currency"`
	PaymentCount int64  `json:"paymentCount" required:"false" doc:"Payment count"`
}

type SchoolInfoResponse struct {
	FullName            string     `json:"fullName" required:"false" doc:"Full name"`
	Description         string     `json:"description" required:"false" doc:"Description"`
	Motto               string     `json:"motto" required:"false" doc:"Motto"`
	PhoneNumber1        int64      `json:"phoneNumber1" required:"false" doc:"Phone number 1"`
	PhoneNumber2        int64      `json:"phoneNumber2" required:"false" doc:"Phone number 2"`
	PhoneNumber3        int64      `json:"phoneNumber3" required:"false" doc:"Phone number 3"`
	Email1              string     `json:"email1" required:"false" doc:"Email 1"`
	Email2              string     `json:"email2" required:"false" doc:"Email 2"`
	Email3              string     `json:"email3" required:"false" doc:"Email 3"`
	Founder             string     `json:"founder" required:"false" doc:"Founder name"`
	FoundedAt           *time.Time `json:"foundedAt" required:"false" doc:"Founded date time"`
	Address             string     `json:"address" required:"false" doc:"Address"`
	PoBox               string     `json:"poBox" required:"false" doc:"PO Box"`
	LocationLongitude   string     `json:"locationLongitude" required:"false" doc:"Location longitude"`
	LocationLatitude    string     `json:"locationLatitude" required:"false" doc:"Location latitude"`
	SocialMediaTelegram string     `json:"socialMediaTelegram" required:"false" doc:"Social media telegram"`
	SocialMediaWhasapp  string     `json:"socialMediaWhasapp" required:"false" doc:"Social media whasapp"`
	SocialMediaYoutube  string     `json:"socialMediaYoutube" required:"false" doc:"Social media youtube"`
	SocialMediaTwitter  string     `json:"socialMediaTwitter" required:"false" doc:"Social media twitter"`
	SocialMediaFacebook string     `json:"socialMediaFacebook" required:"false" doc:"Social media facebook"`
	Image1              string     `json:"image1" required:"false" doc:"Image 1"`
	Image2              string     `json:"image2" required:"false" doc:"Image 2"`
	Image3              string     `json:"image3" required:"false" doc:"Image 3"`
	Image4              string     `json:"image4" required:"false" doc:"Image 4"`
	Image5              string     `json:"image5" required:"false" doc:"Image 5"`
}

type SchoolConfigResponse struct {
	WebsiteDomainName              string `json:"websiteDomainName" required:"false" doc:"Website domain name"`
	UserEmailDomainName            string `json:"userEmailDomainName" required:"false" doc:"User email domain name"`
	SupportEmail                   string `json:"supportEmail" required:"false" doc:"Support email"`
	GoogleWorkspaceCredentials     string `json:"googleWorkspaceCredentials" required:"false" doc:"Google workspace credentials"`
	GoogleWorkspaceUserEmailDomain string `json:"googleWorkspaceUserEmailDomain" required:"false" doc:"Google workspace user email domain"`
	SmsUserID                      string `json:"smsUserID" required:"false" doc:"Sms user ID"`
	WhatsappToken                  string `json:"whatsappToken" required:"false" doc:"Whatsapp token"`
	WhatsappPhoneID                string `json:"whatsappPhoneID" required:"false" doc:"Whatsapp phone ID"`
	TelegramBotToken               string `json:"telegramBotToken" required:"false" doc:"Telegram bot token"`
	WebsiteTitle                   string `json:"websiteTitle" required:"false" doc:"Website title"`
	WebsiteDescription             string `json:"websiteDescription" required:"false" doc:"Website description"`
	ColorPrimary                   string `json:"colorPrimary" required:"false" doc:"Primary color in HEX"`
	ColorPrimaryBg                 string `json:"colorPrimaryBg" required:"false" doc:"Primary color in HEX"`
	ColorPrimaryBgHover            string `json:"colorPrimaryBgHover" required:"false" doc:"Primary color in HEX"`
}

type SchoolResponseList struct {
	types.PaginatedResponse
	Data []SchoolResponse `json:"data" required:"false" doc:"List of school"`
}
