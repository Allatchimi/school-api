package model

import (
	"api/common/types"
	"api/services/school/common/parent/data"
	modelUser "api/services/user/user/model"
)

type Parent struct {
	types.BaseGormModel
	UserID int64           `gorm:"default:null"`
	User   *modelUser.User `gorm:"default:null;foreignKey:UserID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	UID string `gorm:"default:null"`
}

func (item *Parent) ToParentResponse() *data.ParentResponse {
	if item == nil {
		return nil
	}
	resp := &data.ParentResponse{}
	resp.UID = item.UID

	resp.User = item.User.ToResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToParentResponseList(itemList []Parent) []data.ParentResponse {
	resp := make([]data.ParentResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToParentResponse()
	}
	return resp
}
