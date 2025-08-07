package model

import (
	"api/common/types"
	modelSchool "api/services/school/common/school/model"
	"api/services/school/university/department/data"
	modelFaculty "api/services/school/university/faculty/model"
)

type UniversityDepartment struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:CASCADE,onUpdate:CASCADE;"`

	FacultyID int64                           `gorm:"default:null"`
	Faculty   *modelFaculty.UniversityFaculty `gorm:"default:null;foreignKey:FacultyID;references:ID;constraint:onDelete:CASCADE,onUpdate:CASCADE;"`

	Name        string `gorm:"default:null"`
	Description string `gorm:"default:null"`
}

func (item *UniversityDepartment) ToResponse() *data.DepartmentResponse {
	if item == nil {
		return &data.DepartmentResponse{}
	}
	resp := &data.DepartmentResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.School = item.School.ToPublicResponse()
	resp.Faculty = item.Faculty.ToResponse()

	resp.Name = item.Name
	resp.Description = item.Description
	return resp
}

func ToResponseList(itemList []UniversityDepartment) []data.DepartmentResponse {
	resp := make([]data.DepartmentResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
