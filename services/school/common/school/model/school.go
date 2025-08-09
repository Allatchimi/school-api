package model

import (
	"api/common/types"
	"api/common/utils"
	"api/config"
	"api/services/school/common/school/data"
	"fmt"
)

type School struct {
	types.BaseGormModel
	ConfigID int64         `gorm:"default:null"`
	Config   *SchoolConfig `gorm:"default:null;foreignKey:ConfigID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	InfoID int64       `gorm:"default:null"`
	Info   *SchoolInfo `gorm:"default:null;foreignKey:InfoID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Name               string `gorm:"unique;default:null"`
	Type               string `gorm:"default:null"`
	Status             string `gorm:"default:null"`
	DeploymentRequest  string `gorm:"default:null"`
	DeploymentStatus   string `gorm:"default:null"`
	DeploymentFeedback string `gorm:"default:null;type:text"`
	DeploymentCount    int64  `gorm:"default:null"`
	Favicon            string `gorm:"default:null"`
	Logo               string `gorm:"default:null"`
	LogoWhite          string `gorm:"default:null"`
	Currency           string `gorm:"default:null"`
	PaymentCount       int64  `gorm:"default:1"`
}

func (item *School) ToResponse() *data.SchoolResponse {
	if item == nil {
		return &data.SchoolResponse{}
	}
	resp := &data.SchoolResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.Info = item.Info.ToResponse()
	resp.Config = item.Config.ToResponse()

	resp.Name = item.Name
	resp.Type = item.Type
	resp.Status = item.Status
	resp.DeploymentRequest = item.DeploymentRequest
	resp.DeploymentStatus = item.DeploymentStatus
	resp.DeploymentFeedback = item.DeploymentFeedback
	resp.DeploymentCount = item.DeploymentCount
	resp.Favicon = item.Favicon
	resp.Logo = item.Logo
	resp.LogoWhite = item.LogoWhite
	resp.Currency = item.Currency
	resp.PaymentCount = item.PaymentCount
	return resp
}

func (item *School) ToPublicResponse() *data.SchoolPublicResponse {
	if item == nil {
		return &data.SchoolPublicResponse{}
	}
	resp := &data.SchoolPublicResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.Info = item.Info.ToResponse()

	resp.Name = item.Name
	resp.Type = item.Type
	resp.Status = item.Status
	resp.Favicon = item.Favicon
	resp.Logo = item.Logo
	resp.LogoWhite = item.LogoWhite
	resp.Currency = item.Currency
	resp.PaymentCount = item.PaymentCount
	return resp
}

func ToSchoolResponseList(itemList []School) []data.SchoolResponse {
	resp := make([]data.SchoolResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}

func (item *School) WebsiteUrl() (url string) {
	if item == nil || item.Config == nil {
		url = config.Env.WebsiteBaseURL
		return
	}
	url = fmt.Sprintf("https://%s", item.Config.WebsiteDomainName)
	return
}

func (item *School) LogoUrl() (url string) {
	if item == nil {
		url = fmt.Sprintf("%s/assets/images/logos/logo.png", config.Env.WebsiteBaseURL)
		return
	}
	url = item.Logo
	return
}

func (item *School) SMTPNoReplySender() (senderEmail string, senderName string) {
	if item == nil {
		senderEmail = config.Env.SmtpUserNoReply + "@" + config.Env.SmtpDomainName
		senderName = config.Env.AppName
		return
	}
	senderEmail = config.Env.SmtpUserNoReply + "@" + item.Config.UserEmailDomainName
	senderName = item.Name
	return
}

func (item *School) SMTPSupportSender() (senderEmail string, senderName string) {
	if item == nil {
		senderEmail = config.Env.SmtpUserSupport + "@" + config.Env.SmtpDomainName
		senderName = "Support " + config.Env.AppName
		return
	}

	if utils.IsEmailValid(item.Config.SupportEmail) {
		senderEmail = item.Config.SupportEmail
		senderName = "Support " + item.Name
		return
	}
	senderEmail = config.Env.SmtpUserSupport + "@" + item.Config.UserEmailDomainName
	senderName = "Support " + item.Name
	return
}

func (item *School) IsSameDeploymentAsRequest(itemRequest *data.SchoolRequest) bool {
	if item == nil || itemRequest == nil {
		return false
	}
	return (item.Favicon == itemRequest.Favicon) &&
		(item.Logo == itemRequest.Logo) &&
		(item.LogoWhite == itemRequest.LogoWhite)
}
