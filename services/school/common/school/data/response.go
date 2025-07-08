package data

import (
	"api/common/types"
	"time"
)

type SchoolResponse struct {
	SchoolPublicResponse

	DeploymentRequest  string `json:"deploymentRequest" required:"false" doc:"Deployment request"`
	DeploymentStatus   string `json:"deploymentStatus" required:"false" doc:"Deployment status"`
	DeploymentFeedback string `json:"deploymentFeedback" required:"false" doc:"Deployment feedback"`
	DeploymentCount    int64  `json:"deploymentVersion" required:"false" doc:"Deployment version"`

	Config *SchoolConfigResponse `json:"config" required:"false" doc:"Configuration"`
}

type SchoolPublicResponse struct {
	types.BaseGormModelResponse
	Name   string `json:"name" required:"false" doc:"School name"`
	Type   string `json:"type" required:"false" doc:"Type"`
	Status string `json:"status" required:"false" doc:"Status"`

	Favicon   string `json:"favicon" required:"false" doc:"Favicon"`
	Logo      string `json:"logo" required:"false" doc:"Logo"`
	LogoWhite string `json:"logoWhite" required:"false" doc:"Logo white"`

	Currency     string `json:"currency" required:"false" doc:"Currency"`
	PaymentCount int64  `json:"paymentCount" required:"false" doc:"Payment count"`

	Info *SchoolInfoResponse `json:"info" required:"false" doc:"Information"`
}

type SchoolInfoResponse struct {
	FullName    string `json:"fullName" required:"false" doc:"Full name"`
	Description string `json:"description" required:"false" doc:"Description"`
	Motto       string `json:"motto" required:"false" doc:"Motto"`

	PhoneNumber1 int64 `json:"phoneNumber1" required:"false" doc:"Phone number 1"`
	PhoneNumber2 int64 `json:"phoneNumber2" required:"false" doc:"Phone number 2"`
	PhoneNumber3 int64 `json:"phoneNumber3" required:"false" doc:"Phone number 3"`

	Email1 string `json:"email1" required:"false" doc:"Email 1"`
	Email2 string `json:"email2" required:"false" doc:"Email 2"`
	Email3 string `json:"email3" required:"false" doc:"Email 3"`

	Founder   string     `json:"founder" required:"false" doc:"Founder name"`
	FoundedAt *time.Time `json:"foundedAt" required:"false" doc:"Founded date time"`

	Address           string  `json:"address" required:"false" doc:"Address"`
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

type SchoolConfigResponse struct {
	DomainName string `json:"domainName" required:"false" doc:"Domain name"`

	SmtpHost         string `json:"smtpHost" required:"false" doc:"SMTP host"`
	SmtpPort         int    `json:"smtpPort" required:"false" doc:"SMTP port"`
	SmtpUsername     string `json:"smtpUsername" required:"false" doc:"SMTP username"`
	SmtpPassword     string `json:"smtpPassword" required:"false" doc:"SMTP password"`
	SmtpNoReplyEmail string `json:"smtpNoReplyEmail" required:"false" doc:"No reply email"`
	SmtpSupportEmail string `json:"smtpSupportEmail" required:"false" doc:"Support email"`

	SmsUserID string `json:"smsUserID" required:"false" doc:"Sms user ID"`

	UserEmailDomain string `json:"userEmailDomain" required:"false" doc:"User email domain"`

	WebsiteTitle       string `json:"websiteTitle" required:"false" doc:"Website title"`
	WebsiteDescription string `json:"websiteDescription" required:"false" doc:"Website description"`

	ColorPrimary        string `json:"colorPrimary" required:"false" doc:"Color primary"`
	ColorPrimaryBg      string `json:"colorPrimaryBg" required:"false" doc:"Color primary bg"`
	ColorPrimaryBgHover string `json:"colorPrimaryBgHover" required:"false" doc:"Color primary bg hover"`
}

type SchoolResponseList struct {
	types.PaginatedResponse
	Data []SchoolResponse `json:"data" required:"false" doc:"List of schools" example:"[]"`
}
