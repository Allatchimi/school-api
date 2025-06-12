package model

import (
	"api/common/types"
	"api/services/school/common/school/data"
)

type SchoolConfig struct {
	types.BaseGormModel
	Protocol string `gorm:"default:null"`

	DomainName string `gorm:"default:null"`
	DomainCert string `gorm:"default:null"`
	DomainKey  string `gorm:"default:null"`

	SmtpHost     string `gorm:"default:null"`
	SmtpPort     int    `gorm:"default:null"`
	SmtpUsername string `gorm:"default:null"`
	SmtpPassword string `gorm:"default:null"`
	SmtpSender   string `gorm:"default:null"`

	EmailDomain  string `gorm:"default:null"`
	NoReplyEmail string `gorm:"default:null"`
	SupportEmail string `gorm:"default:null"`

	WebsiteTitle       string `gorm:"default:null"`
	WebsiteDescription string `gorm:"default:null"`

	ColorPrimary string `gorm:"default:null"`
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
	resp.SmtpSender = item.SmtpSender
	resp.EmailDomain = item.EmailDomain
	resp.NoReplyEmail = item.NoReplyEmail
	resp.SupportEmail = item.SupportEmail
	resp.WebsiteTitle = item.WebsiteTitle
	resp.WebsiteDescription = item.WebsiteDescription
	resp.ColorPrimary = item.ColorPrimary
	return resp
}

func FromConfigRequest(item *data.SchoolConfigRequest) *SchoolConfig {
	resp := &SchoolConfig{
		Protocol:           item.Protocol,
		DomainName:         item.DomainName,
		DomainCert:         item.DomainCert,
		DomainKey:          item.DomainKey,
		SmtpHost:           item.SmtpHost,
		SmtpPort:           item.SmtpPort,
		SmtpUsername:       item.SmtpUsername,
		SmtpPassword:       item.SmtpPassword,
		SmtpSender:         item.SmtpSender,
		EmailDomain:        item.EmailDomain,
		NoReplyEmail:       item.NoReplyEmail,
		SupportEmail:       item.SupportEmail,
		WebsiteTitle:       item.WebsiteTitle,
		WebsiteDescription: item.WebsiteDescription,
		ColorPrimary:       item.ColorPrimary,
	}
	return resp
}
