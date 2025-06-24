package model

import (
	"api/common/types"
	examModel "api/services/school/common/exam/model"
	"api/services/school/common/report/data"
	studentModel "api/services/school/common/student/model"
)

type Report struct {
	types.BaseGormModel
	StudentID int64                 `gorm:"default:null"`
	Student   *studentModel.Student `gorm:"default:null;foreignKey:StudentID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	ExamID int64           `gorm:"default:null"`
	Exam   *examModel.Exam `gorm:"default:null;foreignKey:ExamID;references:ID;constraint:onDelete:SET NULL,onUpdate:CASCADE;"`

	Value float64 `gorm:"default:null"`
}

func (item *Report) ToResponse() *data.ReportResponse {
	if item == nil {
		return nil
	}
	resp := &data.ReportResponse{}
	resp.Value = item.Value

	resp.Student = item.Student.ToStudentPublicResponse()
	resp.Exam = item.Exam.ToPublicResponse()

	resp.ID = item.ID
	resp.CreatedAt = item.CreatedAt
	resp.UpdatedAt = item.UpdatedAt
	return resp
}

func (item *Report) ToPublicResponse() *data.ReportPublicResponse {
	if item == nil {
		return nil
	}
	resp := &data.ReportPublicResponse{}
	resp.Value = item.Value

	resp.Student = item.Student.ToStudentPublicResponse()
	resp.Exam = item.Exam.ToPublicResponse()
	return resp
}

func ToReportResponseList(itemList []Report) []data.ReportResponse {
	resp := make([]data.ReportResponse, len(itemList))
	for index, item := range itemList {
		resp[index] = *item.ToResponse()
	}
	return resp
}
