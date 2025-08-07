package model

import (
	"api/common/types"
	modelSchool "api/services/school/common/school/model"
	"api/services/school/university/semester/data"
)

type UniversitySemester struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:CASCADE,onUpdate:CASCADE;"`

	Name        string `gorm:"default:null"`
	Description string `gorm:"default:null"`
}

func (item *UniversitySemester) ToResponse() *data.SemesterResponse {
	if item == nil {
		return &data.SemesterResponse{}
	}
	resp := &data.SemesterResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.School = item.School.ToPublicResponse()

	resp.Name = item.Name
	resp.Description = item.Description
	return resp
}

func ToResponseList(itemList []UniversitySemester) []data.SemesterResponse {
	resp := make([]data.SemesterResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
