package model

import (
	"api/common/types"
	examModel "api/services/school/common/exam/model"
	"api/services/school/common/result/data"
	studentModel "api/services/school/common/student/model"
)

type Result struct {
	types.BaseGormModel
	StudentID int64                 `gorm:"default:null"`
	Student   *studentModel.Student `gorm:"default:null;foreignKey:StudentID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	ExamID int64           `gorm:"default:null"`
	Exam   *examModel.Exam `gorm:"default:null;foreignKey:ExamID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Value int `gorm:"default:null"`
}

func (item *Result) ToResponse() *data.ResultResponse {
	if item == nil {
		return nil
	}
	resp := &data.ResultResponse{}
	resp.Value = item.Value

	resp.Student = item.Student.ToStudentPublicResponse()
	resp.Exam = item.Exam.ToPublicResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *Result) ToPublicResponse() *data.ResultPublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.ResultPublicResponse{}
	resp.Value = item.Value

	resp.Student = item.Student.ToStudentPublicResponse()
	resp.Exam = item.Exam.ToPublicResponse()
	return resp
}

func ToResultResponseList(itemList []Result) []data.ResultResponse {
	resp := make([]data.ResultResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
