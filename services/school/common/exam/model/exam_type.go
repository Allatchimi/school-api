package model

import (
	"api/common/types"
	"api/services/school/common/exam/data"
	schoolModel "api/services/school/common/school/model"
)

type ExamType struct {
	types.BaseGormModel

	SchoolID int64               `gorm:"default:null"`
	School   *schoolModel.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Name        string `gorm:"default:null"`
	Description string `gorm:"default:null"`
}

func (item *ExamType) ToResponse() *data.ExamTypeResponse {
	if item == nil {
		return nil
	}
	resp := &data.ExamTypeResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.School = item.School.ToResponse()
	resp.Name = item.Name
	resp.Description = item.Description
	return resp
}

func ToExamTypeResponseList(itemList []ExamType) []data.ExamTypeResponse {
	resp := make([]data.ExamTypeResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
