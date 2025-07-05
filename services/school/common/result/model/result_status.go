package model

import (
	"api/common/types"
	modelExam "api/services/school/common/exam/model"
	"api/services/school/common/result/data"
)

type ResultStatus struct {
	types.BaseGormModel
	ExamID int64           `gorm:"default:null"`
	Exam   *modelExam.Exam `gorm:"default:null;foreignKey:ExamID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Status string `gorm:"default:null"`
}

func (item *ResultStatus) ToResponse() *data.ResultStatusResponse {
	if item == nil {
		return nil
	}
	resp := &data.ResultStatusResponse{}
	resp.Status = item.Status

	resp.Exam = item.Exam.ToPublicResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *ResultStatus) ToPublicResponse() *data.ResultStatusPublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.ResultStatusPublicResponse{}
	resp.Status = item.Status

	resp.Exam = item.Exam.ToPublicResponse()
	return resp
}

func ToResultStatusResponseList(itemList []ResultStatus) []data.ResultStatusResponse {
	resp := make([]data.ResultStatusResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
