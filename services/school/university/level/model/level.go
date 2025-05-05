package model

import (
	"api/common/types"
	schoolModel "api/services/school/common/school/model"
	"api/services/school/university/level/data"
)

type UniversityLevel struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *schoolModel.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Name        string `gorm:"not null"`
	Description string `gorm:"default:null"`
}

func (item *UniversityLevel) ToResponse() *data.LevelResponse {
	if item == nil {
		return nil
	}
	resp := &data.LevelResponse{}
	resp.School = item.School.ToResponse()
	resp.Name = item.Name
	resp.Description = item.Description

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToResponseList(itemList []UniversityLevel) []data.LevelResponse {
	resp := make([]data.LevelResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
