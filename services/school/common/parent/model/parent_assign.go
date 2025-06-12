package model

import (
	"api/common/types"
	"api/services/school/common/parent/data"
)

type ParentAssign struct {
	types.BaseGormModel
	ParentID int64   `gorm:"default:null"`
	Parent   *Parent `gorm:"default:null;foreignKey:ParentID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	ParentAssignStudents []ParentAssignStudent `gorm:"default:null;foreignKey:ParentAssignID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	FirstName string `gorm:"default:null"`
	LastName  string `gorm:"default:null"`
	IDCard    string `gorm:"default:null"`
	Document1 string `gorm:"default:null"`
	Document2 string `gorm:"default:null"`

	Status         string `gorm:"default:null"`
	StatusFeedback string `gorm:"default:null"`
}

func (item *ParentAssign) ToResponse() *data.ParentAssignResponse {
	if item == nil {
		return nil
	}
	resp := &data.ParentAssignResponse{}
	resp.FirstName = item.FirstName
	resp.LastName = item.LastName
	resp.IDCard = item.IDCard
	resp.Document1 = item.Document1
	resp.Document2 = item.Document2
	resp.Status = item.Status
	resp.StatusFeedback = item.StatusFeedback

	resp.Parent = item.Parent.ToParentPublicResponse()
	resp.ParentAssignStudents = ToParentAssignStudentPublicResponseList(item.ParentAssignStudents)

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *ParentAssign) ToPublicResponse() *data.ParentAssignPublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.ParentAssignPublicResponse{}
	resp.FirstName = item.FirstName
	resp.LastName = item.LastName
	resp.IDCard = item.IDCard
	resp.Document1 = item.Document1
	resp.Document2 = item.Document2
	resp.Status = item.Status
	resp.StatusFeedback = item.StatusFeedback

	resp.Parent = item.Parent.ToParentPublicResponse()
	resp.ParentAssignStudents = ToParentAssignStudentPublicResponseList(item.ParentAssignStudents)
	return resp
}

func ToParentAssignResponseList(itemList []ParentAssign) []data.ParentAssignResponse {
	resp := make([]data.ParentAssignResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
