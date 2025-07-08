package model

import (
	"api/common/types"
	"api/services/school/common/director/data"
	modelSchool "api/services/school/common/school/model"
	modelUser "api/services/user/user/model"
)

type Director struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	UserID int64           `gorm:"default:null"`
	User   *modelUser.User `gorm:"default:null;foreignKey:UserID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	UID string `gorm:"default:null"`
}

func (item *Director) ToDirectorResponse() *data.DirectorResponse {
	if item == nil {
		return nil
	}
	resp := &data.DirectorResponse{}
	resp.UID = item.UID

	resp.School = item.School.ToPublicResponse()
	resp.User = item.User.ToResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *Director) ToDirectorPublicResponse() *data.DirectorpublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.DirectorpublicResponse{}

	resp.School = item.School.ToPublicResponse()
	resp.User = item.User.ToPublicResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToDirectorResponseList(itemList []Director) []data.DirectorResponse {
	resp := make([]data.DirectorResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToDirectorResponse()
	}
	return resp
}
