package model

import (
	"api/common/types"
	"api/services/school/common/exam/data"
)

type ExamType struct {
	types.BaseGormModel
	SchoolID    int64  `gorm:"default:null"`
	Name        string `gorm:"default:null"`
	Description string `gorm:"default:null"`
}

func (item *ExamType) ToResponse() *data.ExamTypeResponse {
	resp := &data.ExamTypeResponse{}
	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt

	resp.SchoolID = item.SchoolID
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
