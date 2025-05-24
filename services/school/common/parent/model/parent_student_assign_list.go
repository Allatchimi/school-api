package model

import (
	"api/common/types"
	modelParent "api/services/school/common/parent/model"
	"api/services/school/common/school/data"
	modelStudent "api/services/school/common/student/model"
)

type EnrollParentStudent struct {
	types.BaseGormModel
	FirstName string `gorm:"default:null"`
	LastName  string `gorm:"default:null"`
	CNI       string `gorm:"default:null"`
	Document1 string `gorm:"default:null"`
	Document2 string `gorm:"default:null"`

	Status         string `gorm:"default:null"`
	StatusFeedback string `gorm:"default:null"`

	ParentID int64               `gorm:"default:null"`
	Parent   *modelParent.Parent `gorm:"default:null;foreignKey:ParentID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	StudentID int64                 `gorm:"default:null"`
	Student   *modelStudent.Student `gorm:"default:null;foreignKey:StudentID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
}

func (item *EnrollParentStudent) ToResponse() *data.EnrollParentStudentResponse {
	if item == nil {
		return nil
	}
	resp := &data.EnrollParentStudentResponse{}
	resp.Name = item.Name
	resp.Type = item.Type
	resp.Logo = item.Logo
	resp.Currency = item.Currency

	resp.Info = item.Info.ToResponse()
	resp.Config = item.Config.ToResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *EnrollParentStudent) ToPublicResponse() *data.EnrollParentStudentPublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.EnrollParentStudentPublicResponse{}
	resp.Name = item.Name
	resp.Type = item.Type
	resp.Logo = item.Logo
	resp.Currency = item.Currency

	resp.Info = item.Info.ToResponse()
	return resp
}

func ToEnrollParentStudentResponseList(itemList []EnrollParentStudent) []data.EnrollParentStudentResponse {
	resp := make([]data.EnrollParentStudentResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
