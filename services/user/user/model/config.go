package model

import (
	"api/common/types"
	"api/services/user/user/data"
)

type UserConfig struct {
	types.BaseGormModel
	WhatsappPhoneNumber int64 `gorm:"default:null"`
	TelegramChatID      int64 `gorm:"default:null"`

	AllowNotifications           bool   `gorm:"default:true"`
	MfaEmail                     bool   `gorm:"default:false"`
	MfaAuthenticator             bool   `gorm:"default:false"`
	WebPushSubscriptionEndpoint  string `gorm:"default:null"`
	WebPushSubscriptionKeyP256dh string `gorm:"default:null"`
	WebPushSubscriptionKeyAuth   string `gorm:"default:null"`
}

func (item *UserConfig) ToResponse() *data.UserConfigResponse {
	if item == nil {
		return nil
	}
	resp := &data.UserConfigResponse{}
	resp.WhatsappPhoneNumber = item.WhatsappPhoneNumber
	resp.TelegramChatID = item.TelegramChatID
	resp.AllowNotifications = item.AllowNotifications
	resp.MfaEmail = item.MfaEmail
	resp.MfaAuthenticator = item.MfaAuthenticator
	resp.WebPushSubscriptionEndpoint = item.WebPushSubscriptionEndpoint
	resp.WebPushSubscriptionKeyP256dh = item.WebPushSubscriptionKeyP256dh
	resp.WebPushSubscriptionKeyAuth = item.WebPushSubscriptionKeyAuth
	return resp
}
