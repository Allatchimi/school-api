package model

import (
	"api/common/types"
	"api/services/communication/data"
)

type Communication struct {
	types.BaseGormModel
	Subject       string `gorm:"default:null"`
	Message       string `gorm:"default:null"`
	AudienceType  string `gorm:"default:null"`
	AudienceValue string `gorm:"default:null"`
}

func (item *Communication) ToResponse() *data.CommunicationResponse {
	if item == nil {
		return nil
	}
	resp := &data.CommunicationResponse{}
	resp.Subject = item.Subject
	resp.Message = item.Message
	resp.AudienceType = item.AudienceType
	resp.AudienceValue = item.AudienceValue

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToResponseList(itemList []Communication) []data.CommunicationResponse {
	resp := make([]data.CommunicationResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
