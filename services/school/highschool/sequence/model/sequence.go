package model

import (
	"api/common/types"
	modelSchool "api/services/school/common/school/model"
	modelQuarter "api/services/school/highschool/quarter/model"
	"api/services/school/highschool/sequence/data"
)

type HighschoolSequence struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	QuarterID int64                           `gorm:"default:null"`
	Quarter   *modelQuarter.HighschoolQuarter `gorm:"default:null;foreignKey:QuarterID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Name        string `gorm:"default:null"`
	Description string `gorm:"default:null"`
}

func (item *HighschoolSequence) ToResponse() *data.SequenceResponse {
	if item == nil {
		return &data.SequenceResponse{}
	}
	resp := &data.SequenceResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.School = item.School.ToPublicResponse()
	resp.Quarter = item.Quarter.ToResponse()

	resp.Name = item.Name
	resp.Description = item.Description
	return resp
}

func ToResponseList(itemList []HighschoolSequence) []data.SequenceResponse {
	resp := make([]data.SequenceResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
