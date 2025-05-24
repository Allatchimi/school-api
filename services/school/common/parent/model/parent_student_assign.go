package model

import (
	"api/common/types"
	"api/services/school/common/parent/data"
)

type ParentStudentAssign struct {
	types.BaseGormModel
	FirstName string `gorm:"default:null"`
	LastName  string `gorm:"default:null"`
	IDCard    string `gorm:"default:null"`
	Document1 string `gorm:"default:null"`
	Document2 string `gorm:"default:null"`

	Status         string `gorm:"default:null"`
	StatusFeedback string `gorm:"default:null"`

	ParentID int64   `gorm:"default:null"`
	Parent   *Parent `gorm:"default:null;foreignKey:ParentID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
}

func (item *ParentStudentAssign) ToResponse() *data.ParentStudentAssignResponse {
	if item == nil {
		return nil
	}
	resp := &data.ParentStudentAssignResponse{}
	resp.FirstName = item.FirstName
	resp.LastName = item.LastName
	resp.IDCard = item.IDCard
	resp.Document1 = item.Document1
	resp.Document2 = item.Document2
	resp.Status = item.Status
	resp.StatusFeedback = item.StatusFeedback

	resp.Parent = item.Parent.ToParentPublicResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *ParentStudentAssign) ToPublicResponse() *data.ParentStudentAssignPublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.ParentStudentAssignPublicResponse{}
	resp.FirstName = item.FirstName
	resp.LastName = item.LastName
	resp.IDCard = item.IDCard
	resp.Document1 = item.Document1
	resp.Document2 = item.Document2
	resp.Status = item.Status
	resp.StatusFeedback = item.StatusFeedback

	resp.Parent = item.Parent.ToParentPublicResponse()
	return resp
}

func ToParentStudentAssignResponseList(itemList []ParentStudentAssign) []data.ParentStudentAssignResponse {
	resp := make([]data.ParentStudentAssignResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
