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
}

func (item *Parent) ToParentResponse() *data.ParentResponse {
	if item == nil {
		return nil
	}
	resp := &data.ParentResponse{}

	resp.User = item.User.ToPublicResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *Parent) ToParentPublicResponse() *data.ParentPublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.ParentPublicResponse{}
	resp.User = item.User.ToPublicResponse()
	return resp
}

func ToParentResponseList(itemList []Parent) []data.ParentResponse {
	resp := make([]data.ParentResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToParentResponse()
	}
	return resp
}
