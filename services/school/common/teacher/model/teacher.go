package model

import (
	"api/common/types"
	modelSchool "api/services/school/common/school/model"
	"api/services/school/common/teacher/data"
	modelUser "api/services/user/user/model"
)

type Teacher struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	UserID int64           `gorm:"default:null"`
	User   *modelUser.User `gorm:"default:null;foreignKey:UserID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	UID string `gorm:"default:null"`
}

func (item *Teacher) ToResponse() *data.TeacherResponse {
	if item == nil {
		return nil
	}
	resp := &data.TeacherResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.School = item.School.ToPublicResponse()
	resp.User = item.User.ToResponse()

	resp.UID = item.UID
	return resp
}

func (item *Teacher) ToPublicResponse() *data.TeacherPublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.TeacherPublicResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.School = item.School.ToPublicResponse()
	resp.User = item.User.ToPublicResponse()

	resp.UID = item.UID
	return resp
}

func ToTeacherResponseList(itemList []Teacher) []data.TeacherResponse {
	resp := make([]data.TeacherResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
