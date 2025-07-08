package model

import (
	"api/common/types"
	"api/services/school/common/parent/data"
	modelStudent "api/services/school/common/student/model"
)

type ParentStudent struct {
	types.BaseGormModel

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
	resp.Parent = item.Parent.ToParentPublicResponse()
	resp.Student = item.Student.ToStudentPublicResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToParentStudentResponseList(itemList []ParentStudent) []data.ParentStudentResponse {
	resp := make([]data.ParentStudentResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToParentStudentResponse()
	}
	return resp
}
