package model

import (
	"api/common/types"
	modelExam "api/services/school/common/exam/model"
	"api/services/school/common/result/data"
	modelSchool "api/services/school/common/school/model"
)

type ResultTable struct {
	types.BaseGormModel
	SchoolID int64               `gorm:"default:null"`
	School   *modelSchool.School `gorm:"default:null;foreignKey:SchoolID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	ExamID int64           `gorm:"default:null"`
	Exam   *modelExam.Exam `gorm:"default:null;foreignKey:ExamID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Status string `gorm:"default:null"`
}

func (item *ResultTable) ToResponse() *data.ResultTableResponse {
	if item == nil {
		return &data.ResultTableResponse{}
	}
	resp := &data.ResultTableResponse{}
	resp.Status = item.Status

	resp.School = item.School.ToResponse()
	resp.Exam = item.Exam.ToResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func ToResultTableResponseList(itemList []ResultTable) []data.ResultTableResponse {
	resp := make([]data.ResultTableResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
