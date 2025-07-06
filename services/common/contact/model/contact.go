package model

import (
	"api/common/types"
	"api/services/common/contact/data"
	modelSchool "api/services/school/common/school/model"
)

type Contact struct {
	types.BaseGormModel

	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Subject string `gorm:"default:null"`
	Email   string `gorm:"default:null"`
	Message string `gorm:"default:null;type:text"`
}

func (item *Contact) ToResponse() *data.ContactResponse {
	if item == nil {
		return nil
	}
	resp := &data.ContactResponse{}
	resp.Subject = item.Subject
	resp.Email = item.Email
	resp.Message = item.Message

	resp.School = item.School.ToPublicResponse()

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
