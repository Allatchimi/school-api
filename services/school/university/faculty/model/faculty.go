package model

import (
	"api/common/types"
	modelSchool "api/services/school/common/school/model"
	"api/services/school/university/faculty/data"
)

type UniversityFaculty struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Name        string `gorm:"default:null"`
	Description string `gorm:"default:null"`
}

func (item *UniversityFaculty) ToResponse() *data.FacultyResponse {
	if item == nil {
		return nil
	}
	resp := &data.FacultyResponse{}
	resp.Name = item.Name
	resp.Description = item.Description

	resp.School = item.School.ToPublicResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToResponseList(itemList []UniversityFaculty) []data.FacultyResponse {
	resp := make([]data.FacultyResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
