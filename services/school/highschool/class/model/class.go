package model

import (
	"api/common/types"
	schoolModel "api/services/school/common/school/model"
	"api/services/school/highschool/class/data"
	specialtyModel "api/services/school/highschool/specialty/model"
)

type HighschoolClass struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *schoolModel.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	SpecialtyID int64                               `gorm:"default:null"`
	Specialty   *specialtyModel.HighschoolSpecialty `gorm:"default:null;foreignKey:SpecialtyID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Name        string `gorm:"not null"`
	Description string `gorm:"default:null"`
	Fees        int64  `gorm:"default:null"`
}

func (item *HighschoolClass) ToResponse() *data.ClassResponse {
	if item == nil {
		return nil
	}
	resp := &data.ClassResponse{}
	resp.Name = item.Name
	resp.Description = item.Description

	resp.School = item.School.ToPublicResponse()
	resp.Specialty = item.Specialty.ToPublicResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *HighschoolClass) ToPublicResponse() *data.ClassPublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.ClassPublicResponse{}
	resp.Name = item.Name
	resp.Description = item.Description

	resp.School = item.School.ToPublicResponse()
	resp.Specialty = item.Specialty.ToPublicResponse()
	return resp
}

func ToResponseList(itemList []HighschoolClass) []data.ClassResponse {
	resp := make([]data.ClassResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
