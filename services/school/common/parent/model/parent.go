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
	resp := &data.ParentResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.User = item.User.ToResponse()
	resp.UID = item.UID
	return resp
}

func ToParentResponseList(itemList []Parent) []data.ParentResponse {
	resp := make([]data.ParentResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToParentResponse()
	}
	return resp
}
