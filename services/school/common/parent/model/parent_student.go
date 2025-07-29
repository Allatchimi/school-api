package model

import (
	"api/common/types"
	"api/services/school/common/parent/data"
	modelSchool "api/services/school/common/school/model"
	modelStudent "api/services/school/common/student/model"
)

type ParentStudent struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	ParentID int64   `gorm:"default:null"`
	Parent   *Parent `gorm:"default:null;foreignKey:ParentID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	StudentID int64                 `gorm:"default:null"`
	Student   *modelStudent.Student `gorm:"default:null;foreignKey:StudentID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
}

func (item *ParentStudent) ToParentStudentResponse() *data.ParentStudentResponse {
	if item == nil {
		return nil
	}
	resp := &data.ParentStudentResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.School = item.School.ToPublicResponse()
	resp.Parent = item.Parent.ToPublicResponse()
	resp.Student = item.Student.ToPublicResponse()
	return resp
}

func ToParentStudentResponseList(itemList []ParentStudent) []data.ParentStudentResponse {
	resp := make([]data.ParentStudentResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToParentStudentResponse()
	}
	return resp
}
