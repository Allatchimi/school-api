package model

import (
	"api/common/types"
	modelSchool "api/services/school/common/school/model"
	"api/services/school/highschool/class/data"
	modelSpecialty "api/services/school/highschool/specialty/model"
)

type HighschoolClass struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	SpecialtyID int64                               `gorm:"default:null"`
	Specialty   *modelSpecialty.HighschoolSpecialty `gorm:"default:null;foreignKey:SpecialtyID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Fees        int64  `gorm:"default:null"`
	Name        string `gorm:"default:null"`
	Description string `gorm:"default:null"`
}

func (item *HighschoolClass) ToResponse() *data.ClassResponse {
	if item == nil {
		return nil
	}
	resp := &data.ClassResponse{}
	resp.Fees = item.Fees
	resp.Name = item.Name
	resp.Description = item.Description

	resp.School = item.School.ToPublicResponse()
	resp.Specialty = item.Specialty.ToResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToResponseList(itemList []HighschoolClass) []data.ClassResponse {
	resp := make([]data.ClassResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
