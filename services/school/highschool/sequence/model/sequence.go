package model

import (
	"api/common/types"
	"api/services/school/common/school/model"
	"api/services/school/highschool/sequence/data"
)

type HighschoolSequence struct {
	types.BaseGormModel
	SchoolID int64         `gorm:"default:null"`
	School   *model.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Name        string `gorm:"not null"`
	Description string `gorm:"default:null"`
}

func (item *HighschoolSequence) ToResponse() *data.SequenceResponse {
	if item == nil {
		return nil
	}
	resp := &data.SequenceResponse{}
	resp.Name = item.Name
	resp.Description = item.Description

	resp.School = item.School.ToPublicResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *HighschoolSequence) ToPublicResponse() *data.SequencePublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.SequencePublicResponse{}
	resp.Name = item.Name
	resp.Description = item.Description

	resp.School = item.School.ToPublicResponse()
	return resp
}

func ToResponseList(itemList []HighschoolSequence) []data.SequenceResponse {
	resp := make([]data.SequenceResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
