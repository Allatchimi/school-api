package model

import (
	"api/common/types"
	modelSchool "api/services/school/common/school/model"
	modelSection "api/services/school/highschool/section/model"
	"api/services/school/highschool/specialty/data"
)

type HighschoolSpecialty struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	SectionID int64                           `gorm:"default:null"`
	Section   *modelSection.HighschoolSection `gorm:"default:null;foreignKey:SectionID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Name        string `gorm:"default:null"`
	Description string `gorm:"default:null"`
}

func (item *HighschoolSpecialty) ToResponse() *data.SpecialtyResponse {
	if item == nil {
		return nil
	}
	resp := &data.SpecialtyResponse{}
	resp.Name = item.Name
	resp.Description = item.Description

	resp.School = item.School.ToPublicResponse()
	resp.Section = item.Section.ToPublicResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *HighschoolSpecialty) ToPublicResponse() *data.SpecialtyPublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.SpecialtyPublicResponse{}
	resp.Name = item.Name
	resp.Description = item.Description

	resp.School = item.School.ToPublicResponse()
	resp.Section = item.Section.ToPublicResponse()
	return resp
}

func ToResponseList(itemList []HighschoolSpecialty) []data.SpecialtyResponse {
	resp := make([]data.SpecialtyResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
