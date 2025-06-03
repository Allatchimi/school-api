package model

import (
	"api/common/types"
	"api/services/school/common/director/data"
	dataSchool "api/services/school/common/school/data"
	schoolModel "api/services/school/common/school/model"
	dataUser "api/services/user/user/data"
	userModel "api/services/user/user/model"
)

type Director struct {
	types.BaseGormModel

	UserID int64           `gorm:"default:null"`
	User   *userModel.User `gorm:"default:null;foreignKey:UserID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	SchoolID int64               `gorm:"default:null"`
	School   *schoolModel.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
}

func (item *Director) ToResponse() *data.DirectorResponse {
	if item == nil {
		return nil
	}
	resp := &data.DirectorResponse{
		DirectorPublicResponse: data.DirectorPublicResponse{
			User:   &dataUser.UserPublicResponse{},
			School: &dataSchool.SchoolPublicResponse{},
		},
	}
	if item.User != nil {
		resp.User = item.User.ToPublicResponse()
	}
	if item.School != nil {
		resp.School = item.School.ToPublicResponse()
	}

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *Director) ToPublicResponse() *data.DirectorPublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.DirectorPublicResponse{
		User:   &dataUser.UserPublicResponse{},
		School: &dataSchool.SchoolPublicResponse{},
	}
	if item.User != nil {
		resp.User = item.User.ToPublicResponse()
	}
	if item.School != nil {
		resp.School = item.School.ToPublicResponse()
	}
	return resp
}

func ToResponseList(itemList []Director) []data.DirectorResponse {
	resp := make([]data.DirectorResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
