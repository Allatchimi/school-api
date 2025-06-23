package data

import (
	"time"
)

type SchoolID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"School id"`
}

type SchoolRequest struct {
	Name         string `json:"name" required:"true" minLength:"2" maxLength:"50" doc:"School name"`
	Type         string `json:"type" required:"true" minLength:"2" maxLength:"50" enum:"highschool,university" doc:"School type"`
	Logo         string `json:"logo" required:"false" doc:"School logo"`
	Currency     string `json:"currency" required:"true" minLength:"2" maxLength:"20" doc:"Currency"`
	PaymentCount int64  `json:"paymentCount" required:"true" min:"1" max:"10" doc:"Payment count"`

	Info   *SchoolInfoRequest   `json:"info" required:"true" doc:"Information"`
	Config *SchoolConfigRequest `json:"config" required:"true" doc:"Configuration"`
}

type SchoolInfoRequest struct {
	FullName    string `json:"fullName" required:"true" minLength:"2" maxLength:"150" doc:"School name"`
	Description string `json:"description" required:"false" maxLength:"500" doc:"Description"`
	Slogan      string `json:"slogan" required:"false" maxLength:"150" doc:"Slogan"`

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

	Image1 string `json:"image1" required:"false" doc:"Image 1"`
	Image2 string `json:"image2" required:"false" doc:"Image 2"`
	Image3 string `json:"image3" required:"false" doc:"Image 3"`
	Image4 string `json:"image4" required:"false" doc:"Image 4"`
	Image5 string `json:"image5" required:"false" doc:"Image 5"`
}

type SchoolConfigRequest struct {
	Protocol string `json:"protocol" required:"true" enum:"http,https" doc:"Protocol"`

	DomainName string `json:"domainName" required:"true" minLength:"3" maxLength:"150" doc:"Domain name" example:"google.com"`
	DomainCert string `json:"domainCert" required:"true" minLength:"3" maxLength:"150" doc:"Domain cert"`
	DomainKey  string `json:"domainKey" required:"true" minLength:"3" maxLength:"150" doc:"Domain key"`

	SmtpHost         string `json:"smtpHost" required:"true" minLength:"3" maxLength:"150" doc:"SMTP host" example:"smtp.google.com"`
	SmtpPort         int    `json:"smtpPort" required:"true" min:"1" max:"65535" doc:"SMTP port" example:"587"`
	SmtpUsername     string `json:"smtpUsername" required:"true" minLength:"3" maxLength:"150" doc:"SMTP username" example:"user@gmail.com"`
	SmtpPassword     string `json:"smtpPassword" required:"true" minLength:"3" maxLength:"150" doc:"SMTP password" example:"password"`
	SmtpNoReplyEmail string `json:"smtpNoReplyEmail" required:"true" minLength:"3" maxLength:"150" doc:"No reply email" example:"noreply@gmail.com"`
	SmtpSupportEmail string `json:"smtpSupportEmail" required:"true" minLength:"3" maxLength:"150" doc:"Support email" example:"support@gmail.com"`

	UserEmailDomain string `json:"userEmailDomain" required:"true" minLength:"3" maxLength:"150" doc:"User email domain" example:"google.com"`

	WebsiteTitle       string `json:"websiteTitle" required:"true" minLength:"3" maxLength:"150" doc:"Website title" example:"School"`
	WebsiteDescription string `json:"websiteDescription" required:"true" minLength:"3" maxLength:"500" doc:"Website description" example:"School description"`

	ColorPrimary string `json:"colorPrimary" required:"true" minLength:"4" maxLength:"7" doc:"Primary color in HEX" example:"#FFFFFF"`
}

type GetAllRequest struct {
	Type string `json:"type" query:"type" required:"false" enum:"highschool,university" doc:"School type" example:"university"`
}
