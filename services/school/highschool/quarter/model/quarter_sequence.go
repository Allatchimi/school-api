package model

import (
	"api/common/types"
	"api/services/school/highschool/quarter/data"
	sequenceModel "api/services/school/highschool/sequence/model"
)

type HighschoolQuarterSequence struct {
	types.BaseGormModel
	QuarterID int64              `gorm:"default:null"`
	Quarter   *HighschoolQuarter `gorm:"default:null;foreignKey:QuarterID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	SequenceID int64                             `gorm:"default:null"`
	Sequence   *sequenceModel.HighschoolSequence `gorm:"default:null;foreignKey:SequenceID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`
}

func (item *HighschoolQuarterSequence) ToQuarterSequenceResponse() *data.QuarterSequenceResponse {
	if item == nil {
		return nil
	}
	resp := &data.QuarterSequenceResponse{}
	resp.Quarter = item.Quarter.ToPublicResponse()
	resp.Sequence = item.Sequence.ToPublicResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *HighschoolQuarterSequence) ToQuarterSequencePublicResponse() *data.QuarterSequencePublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.QuarterSequencePublicResponse{}
	resp.Quarter = item.Quarter.ToPublicResponse()
	resp.Sequence = item.Sequence.ToPublicResponse()
	return resp
}

func ToQuarterSequenceResponseList(itemList []HighschoolQuarterSequence) []data.QuarterSequenceResponse {
	resp := make([]data.QuarterSequenceResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToQuarterSequenceResponse()
	}
	return resp
}
