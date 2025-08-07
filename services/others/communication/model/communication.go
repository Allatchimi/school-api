package model

import (
	"api/common/types"
	"api/services/others/communication/data"
	modelSchool "api/services/school/common/school/model"
	modelRole "api/services/user/role/model"
)

type Communication struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:CASCADE,onUpdate:CASCADE;"`

	RoleID int64           `gorm:"default:null"`
	Role   *modelRole.Role `gorm:"default:null;foreignKey:RoleID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Subject string `gorm:"default:null"`
	Message string `gorm:"default:null;type:text"`
}

func (item *Communication) ToResponse() *data.CommunicationResponse {
	if item == nil {
		return &data.CommunicationResponse{}
	}
	resp := &data.CommunicationResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.School = item.School.ToPublicResponse()
	resp.Role = item.Role.ToResponse()

	resp.Subject = item.Subject
	resp.Message = item.Message
	return resp
}

func ToResponseList(itemList []Communication) []data.CommunicationResponse {
	resp := make([]data.CommunicationResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
