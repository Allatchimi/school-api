package model

import (
	"api/common/types"
	modelSchool "api/services/school/common/school/model"
	"api/services/school/university/student/data"
	modelUser "api/services/user/user/model"
)

type Student struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	UserID int64           `gorm:"default:null"`
	User   *modelUser.User `gorm:"default:null;foreignKey:UserID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	SerialNumber string `gorm:"default:null"`
}

func (item *Student) ToStudentResponse() *data.StudentResponse {
	resp := &data.StudentResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.School = item.School.ToResponse()
	resp.User = item.User.ToResponse()
	return resp
}

func ToStudentResponseList(itemList []Student) []data.StudentResponse {
	resp := make([]data.StudentResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToStudentResponse()
	}
	return resp
}
