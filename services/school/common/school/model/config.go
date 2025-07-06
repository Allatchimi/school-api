package model

import (
	"api/common/types"
	"api/services/school/common/school/data"
)

type SchoolConfig struct {
	types.BaseGormModel
	Protocol string `gorm:"default:null"`

	DomainName string `gorm:"default:null"`
	DomainCert string `gorm:"default:null;type:text"`
	DomainKey  string `gorm:"default:null;type:text"`

	SmtpHost         string `gorm:"default:null"`
	SmtpPort         int    `gorm:"default:null"`
	SmtpUsername     string `gorm:"default:null"`
	SmtpPassword     string `gorm:"default:null"`
	SmtpNoReplyEmail string `gorm:"default:null"`
	SmtpSupportEmail string `gorm:"default:null"`

	SmsUserID string `gorm:"default:null"`

	UserEmailDomain string `gorm:"default:null"`

	WebsiteTitle       string `gorm:"default:null"`
	WebsiteDescription string `gorm:"default:null"`

	ColorPrimary        string `gorm:"default:null"`
	ColorPrimaryBg      string `gorm:"default:null"`
	ColorPrimaryBgHover string `gorm:"default:null"`
}

func (item *SchoolConfig) ToResponse() *data.SchoolConfigResponse {
	if item == nil {
		return nil
	}
	resp := &data.SchoolConfigResponse{}
	resp.Protocol = item.Protocol
	resp.DomainName = item.DomainName
	resp.DomainCert = item.DomainCert
	resp.DomainKey = item.DomainKey
	resp.SmtpHost = item.SmtpHost
	resp.SmtpPort = item.SmtpPort
	resp.SmtpUsername = item.SmtpUsername
	resp.SmtpPassword = item.SmtpPassword
	resp.SmtpNoReplyEmail = item.SmtpNoReplyEmail
	resp.SmtpSupportEmail = item.SmtpSupportEmail
	resp.SmsUserID = item.SmsUserID
	resp.UserEmailDomain = item.UserEmailDomain
	resp.WebsiteTitle = item.WebsiteTitle
	resp.WebsiteDescription = item.WebsiteDescription
	resp.ColorPrimary = item.ColorPrimary
	resp.ColorPrimaryBg = item.ColorPrimaryBg
	resp.ColorPrimaryBgHover = item.ColorPrimaryBgHover
	return resp
}

func FromConfigRequest(item *data.SchoolConfigRequest) *SchoolConfig {
	resp := &SchoolConfig{
		Protocol:            item.Protocol,
		DomainName:          item.DomainName,
		DomainCert:          item.DomainCert,
		DomainKey:           item.DomainKey,
		SmtpHost:            item.SmtpHost,
		SmtpPort:            item.SmtpPort,
		SmtpUsername:        item.SmtpUsername,
		SmtpPassword:        item.SmtpPassword,
		SmtpNoReplyEmail:    item.SmtpNoReplyEmail,
		SmtpSupportEmail:    item.SmtpSupportEmail,
		SmsUserID:           item.SmsUserID,
		UserEmailDomain:     item.UserEmailDomain,
		WebsiteTitle:        item.WebsiteTitle,
		WebsiteDescription:  item.WebsiteDescription,
		ColorPrimary:        item.ColorPrimary,
		ColorPrimaryBg:      item.ColorPrimaryBg,
		ColorPrimaryBgHover: item.ColorPrimaryBgHover,
	}
	return resp
}
