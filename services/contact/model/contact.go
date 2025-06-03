package model

import (
	"api/common/types"
	"api/services/contact/data"
)

type Contact struct {
	types.BaseGormModel
	Subject string `gorm:"default:null"`
	Email   string `gorm:"default:null"`
	Message string `gorm:"default:null"`
}

func (item *Contact) ToResponse() *data.ContactResponse {
	if item == nil {
		return nil
	}
	resp := &data.ContactResponse{}
	resp.Subject = item.Subject
	resp.Email = item.Email
	resp.Message = item.Message

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToResponseList(itemList []Contact) []data.ContactResponse {
	resp := make([]data.ContactResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
