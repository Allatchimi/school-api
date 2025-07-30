package model

import (
	"api/common/types"
	"api/services/others/notification/data"
	modelUser "api/services/user/user/model"
	"time"
)

type Notification struct {
	types.BaseGormModel

	UserID int64           `gorm:"default:null"`
	User   *modelUser.User `gorm:"default:null;foreignKey:UserID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Title   string     `gorm:"default:null"`
	Message string     `gorm:"default:null"`
	Seen    bool       `gorm:"default:null"`
	SeenAt  *time.Time `gorm:"default:null"`
}

func (item *Notification) ToResponse() *data.NotificationResponse {
	if item == nil {
		return &data.NotificationResponse{}
	}
	resp := &data.NotificationResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.User = item.User.ToPublicResponse()

	resp.Title = item.Title
	resp.Message = item.Message
	resp.Seen = item.Seen
	resp.SeenAt = item.SeenAt
	return resp
}

func ToResponseList(itemList []Notification) []data.NotificationResponse {
	resp := make([]data.NotificationResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
