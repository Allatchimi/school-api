package model

import (
	"api/common/types"
	"api/services/common/notification/data"
	userModel "api/services/user/user/model"
)

type Notification struct {
	types.BaseGormModel

	UserID int64           `gorm:"default:null"`
	User   *userModel.User `gorm:"default:null;foreignKey:UserID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Title   string `gorm:"default:null"`
	Message string `gorm:"default:null"`
}

func (item *Notification) ToResponse() *data.NotificationResponse {
	if item == nil {
		return nil
	}
	resp := &data.NotificationResponse{}
	resp.Title = item.Title
	resp.Message = item.Message

	resp.User = item.User.ToPublicResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToResponseList(itemList []Notification) []data.NotificationResponse {
	resp := make([]data.NotificationResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
