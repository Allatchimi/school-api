package model

import (
	"api/common/types"
	"api/services/school/common/school/model"
	"api/services/school/highschool/quarter/data"
)

type HighschoolQuarter struct {
	types.BaseGormModel
	SchoolID int64         `gorm:"default:null"`
	School   *model.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Name        string `gorm:"not null"`
	Description string `gorm:"default:null"`
}

func (item *HighschoolQuarter) ToResponse() *data.QuarterResponse {
	if item == nil {
		return nil
	}
	resp := &data.QuarterResponse{}
	resp.Name = item.Name
	resp.Description = item.Description

	resp.School = item.School.ToResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToResponseList(itemList []HighschoolQuarter) []data.QuarterResponse {
	resp := make([]data.QuarterResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
