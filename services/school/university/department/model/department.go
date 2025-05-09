package model

import (
	"api/common/types"
	schoolModel "api/services/school/common/school/model"
	"api/services/school/university/department/data"
	facultyModel "api/services/school/university/faculty/model"
)

type UniversityDepartment struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *schoolModel.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	FacultyID int64                           `gorm:"default:null"`
	Faculty   *facultyModel.UniversityFaculty `gorm:"default:null;foreignKey:FacultyID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Name        string `gorm:"not null"`
	Description string `gorm:"default:null"`
}

func (item *UniversityDepartment) ToResponse() *data.DepartmentResponse {
	if item == nil {
		return nil
	}
	resp := &data.DepartmentResponse{}
	resp.Name = item.Name
	resp.Description = item.Description

	resp.School = item.School.ToPublicResponse()
	resp.Faculty = item.Faculty.ToPublicResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *UniversityDepartment) ToPublicResponse() *data.DepartmentPublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.DepartmentPublicResponse{}
	resp.Name = item.Name
	resp.Description = item.Description

	resp.School = item.School.ToPublicResponse()
	resp.Faculty = item.Faculty.ToPublicResponse()
	return resp
}

func ToResponseList(itemList []UniversityDepartment) []data.DepartmentResponse {
	resp := make([]data.DepartmentResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
