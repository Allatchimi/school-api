package model

import (
	"api/common/types"
	"api/services/school/common/document/data"
)

type Document struct {
	types.BaseGormModel
	SchoolID       int64  `gorm:"not null"`
	YearID         int64  `gorm:"not null"`
	ClassSubjectID int64  `gorm:"default:null"`
	UnitID         int64  `gorm:"default:null"`
	URL            string `gorm:"default:null"`
	Name           string `gorm:"not null"`
	Description    string `gorm:"default:null"`
}

func (item *Document) ToResponse() *data.DocumentResponse {
	if item == nil {
		return nil
	}
	resp := &data.DocumentResponse{}
	resp.SchoolID = item.SchoolID
	resp.YearID = item.YearID
	resp.ClassSubjectID = item.ClassSubjectID
	resp.UnitID = item.UnitID
	resp.URL = item.URL
	resp.Description = item.Description

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToResponseList(itemList []Document) []data.DocumentResponse {
	resp := make([]data.DocumentResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
