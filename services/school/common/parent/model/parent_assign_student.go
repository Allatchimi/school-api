package model

import (
	"api/common/types"
	"api/services/school/common/parent/data"
	modelStudent "api/services/school/common/student/model"
)

type ParentAssignStudent struct {
	types.BaseGormModel
	ParentAssignID int64         `gorm:"default:null"`
	ParentAssign   *ParentAssign `gorm:"default:null;foreignKey:ParentAssignID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	StudentID int64                 `gorm:"default:null"`
	Student   *modelStudent.Student `gorm:"default:null;foreignKey:StudentID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
}

func (item *ParentAssignStudent) ToResponse() *data.ParentAssignStudentResponse {
	if item == nil {
		return nil
	}
	resp := &data.ParentAssignStudentResponse{}
	resp.ParentAssign = item.ParentAssign.ToPublicResponse()
	resp.Student = item.Student.ToStudentPublicResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *ParentAssignStudent) ToPublicResponse() *data.ParentAssignStudentPublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.ParentAssignStudentPublicResponse{}
	resp.ParentAssign = item.ParentAssign.ToPublicResponse()
	resp.Student = item.Student.ToStudentPublicResponse()
	return resp
}

func ToParentAssignStudentResponseList(itemList []ParentAssignStudent) []data.ParentAssignStudentResponse {
	resp := make([]data.ParentAssignStudentResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}

func ToParentAssignStudentPublicResponseList(itemList []ParentAssignStudent) []data.ParentAssignStudentPublicResponse {
	resp := make([]data.ParentAssignStudentPublicResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToPublicResponse()
	}
	return resp
}
