package model

import (
	"api/common/types"
	"api/services/school/common/parent/data"
	dataSchool "api/services/school/common/school/data"
	modelSchool "api/services/school/common/school/model"
	dataUser "api/services/user/user/data"
	modelUser "api/services/user/user/model"
)

type Parent struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	UserID int64           `gorm:"default:null"`
	User   *modelUser.User `gorm:"default:null;foreignKey:UserID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
}

func (item *Parent) ToResponse() *data.ParentResponse {
	if item == nil {
		return nil
	}
	resp := &data.ParentResponse{
		School: &dataSchool.SchoolPublicResponse{},
		User:   &dataUser.UserResponse{},
	}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	if item.School != nil {
		resp.School = item.School.ToPublicResponse()
	}
	if item.User != nil {
		resp.User = item.User.ToResponse()
	}
	return resp
}

func (item *Parent) ToPublicResponse() *data.ParentPublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.ParentPublicResponse{
		School: &dataSchool.SchoolPublicResponse{},
		User:   &dataUser.UserPublicResponse{},
	}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	if item.School != nil {
		resp.School = item.School.ToPublicResponse()
	}
	if item.User != nil {
		resp.User = item.User.ToPublicResponse()
	}
	return resp
}

func ToParentResponseList(itemList []Parent) []data.ParentResponse {
	resp := make([]data.ParentResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
