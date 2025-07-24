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
	resp.ParentAssign = item.ParentAssign.ToResponse()
	resp.Student = item.Student.ToResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToParentAssignStudentResponseList(itemList []ParentAssignStudent) []data.ParentAssignStudentResponse {
	resp := make([]data.ParentAssignStudentResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
