package model

import (
	"api/common/types"
	"api/services/school/common/school/model"
	"api/services/school/highschool/subject/data"
)

type HighschoolSubject struct {
	types.BaseGormModel
	SchoolID int64         `gorm:"default:null"`
	School   *model.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Name        string `gorm:"default:null"`
	Description string `gorm:"default:null"`
}

func (item *HighschoolSubject) ToResponse() *data.SubjectResponse {
	if item == nil {
		return nil
	}
	resp := &data.SubjectResponse{}
	resp.Name = item.Name
	resp.Description = item.Description

	resp.School = item.School.ToPublicResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *HighschoolSubject) ToPublicResponse() *data.SubjectPublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.SubjectPublicResponse{}
	resp.Name = item.Name
	resp.Description = item.Description

	resp.School = item.School.ToPublicResponse()
	return resp
}

func ToResponseList(itemList []HighschoolSubject) []data.SubjectResponse {
	resp := make([]data.SubjectResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
