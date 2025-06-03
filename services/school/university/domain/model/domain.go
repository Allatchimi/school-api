package model

import (
	"api/common/types"
	schoolModel "api/services/school/common/school/model"
	departmentModel "api/services/school/university/department/model"
	"api/services/school/university/domain/data"
)

type UniversityDomain struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *schoolModel.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	DepartmentID int64                                 `gorm:"default:null"`
	Department   *departmentModel.UniversityDepartment `gorm:"default:null;foreignKey:DepartmentID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Name        string `gorm:"default:null"`
	Description string `gorm:"default:null"`
}

func (item *UniversityDomain) ToResponse() *data.DomainResponse {
	if item == nil {
		return nil
	}
	resp := &data.DomainResponse{}
	resp.Name = item.Name
	resp.Description = item.Description

	resp.School = item.School.ToPublicResponse()
	resp.Department = item.Department.ToPublicResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *UniversityDomain) ToPublicResponse() *data.DomainPublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.DomainPublicResponse{}
	resp.Name = item.Name
	resp.Description = item.Description

	resp.School = item.School.ToPublicResponse()
	resp.Department = item.Department.ToPublicResponse()
	return resp
}

func ToResponseList(itemList []UniversityDomain) []data.DomainResponse {
	resp := make([]data.DomainResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
