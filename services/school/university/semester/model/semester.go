package model

import (
	"api/common/types"
	schoolModel "api/services/school/common/school/model"
	"api/services/school/university/semester/data"
)

type UniversitySemester struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *schoolModel.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Name        string `gorm:"default:null"`
	Description string `gorm:"default:null"`
}

func (item *UniversitySemester) ToResponse() *data.SemesterResponse {
	if item == nil {
		return nil
	}
	resp := &data.SemesterResponse{}
	resp.Name = item.Name
	resp.Description = item.Description

	resp.School = item.School.ToPublicResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *UniversitySemester) ToPublicResponse() *data.SemesterPublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.SemesterPublicResponse{}
	resp.Name = item.Name
	resp.Description = item.Description

	resp.School = item.School.ToPublicResponse()
	return resp
}

func ToResponseList(itemList []UniversitySemester) []data.SemesterResponse {
	resp := make([]data.SemesterResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
