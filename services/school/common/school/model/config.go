package model

import (
	"api/common/types"
	"api/services/school/common/school/data"
)

type SchoolConfig struct {
	types.BaseGormModel
	DomainName   string `gorm:"default:null"`
	SupportEmail string `gorm:"default:null"`

	GoogleWorkspaceCredentials     string `gorm:"default:null"`
	GoogleWorkspaceUserEmailDomain string `gorm:"default:null"`

	SmsUserID        string `gorm:"default:null"`
	WhatsappToken    string `gorm:"default:null"`
	WhatsappPhoneID  string `gorm:"default:null"`
	TelegramBotToken string `gorm:"default:null"`

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
	resp.DomainName = item.DomainName
	resp.SupportEmail = item.SupportEmail
	resp.GoogleWorkspaceCredentials = item.GoogleWorkspaceCredentials
	resp.GoogleWorkspaceUserEmailDomain = item.GoogleWorkspaceUserEmailDomain
	resp.SmsUserID = item.SmsUserID
	resp.WhatsappToken = item.WhatsappToken
	resp.WhatsappPhoneID = item.WhatsappPhoneID
	resp.TelegramBotToken = item.TelegramBotToken
	resp.WebsiteTitle = item.WebsiteTitle
	resp.WebsiteDescription = item.WebsiteDescription
	resp.ColorPrimary = item.ColorPrimary
	resp.ColorPrimaryBg = item.ColorPrimaryBg
	resp.ColorPrimaryBgHover = item.ColorPrimaryBgHover
	return resp
}

func FromConfigRequest(item *data.SchoolConfigRequest) *SchoolConfig {
	resp := &SchoolConfig{
		DomainName:                     item.DomainName,
		SupportEmail:                   item.SupportEmail,
		GoogleWorkspaceCredentials:     item.GoogleWorkspaceCredentials,
		GoogleWorkspaceUserEmailDomain: item.GoogleWorkspaceUserEmailDomain,
		SmsUserID:                      item.SmsUserID,
		WhatsappToken:                  item.WhatsappToken,
		WhatsappPhoneID:                item.WhatsappPhoneID,
		TelegramBotToken:               item.TelegramBotToken,
		WebsiteTitle:                   item.WebsiteTitle,
		WebsiteDescription:             item.WebsiteDescription,
		ColorPrimary:                   item.ColorPrimary,
		ColorPrimaryBg:                 item.ColorPrimaryBg,
		ColorPrimaryBgHover:            item.ColorPrimaryBgHover,
	}
	return resp
}
