package model

import (
	"api/common/types"
	"api/services/common/communication/data"
	modelSchool "api/services/school/common/school/model"
)

type Communication struct {
	types.BaseGormModel

	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Audience string `gorm:"default:null"`

	Subject string `gorm:"default:null"`
	Message string `gorm:"default:null;type:text"`
}

func (item *Communication) ToResponse() *data.CommunicationResponse {
	if item == nil {
		return nil
	}
	resp := &data.CommunicationResponse{}
	resp.Subject = item.Subject
	resp.Message = item.Message
	resp.Audience = item.Audience

	resp.School = item.School.ToPublicResponse()

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
