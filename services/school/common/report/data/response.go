package data

import (
	"api/common/types"
	examData "api/services/school/common/exam/data"
	studentData "api/services/school/common/student/data"
)

type ReportResponse struct {
	types.BaseGormModelResponse
	ReportPublicResponse
}

type ReportPublicResponse struct {
	Value float64 `json:"value" required:"false" doc:"Value"`

	Student *studentData.StudentPublicResponse `json:"Student" required:"true" doc:"Student"`
	Exam    *examData.ExamPublicResponse       `json:"Exam" required:"true" doc:"Exam"`
}

type ReportResponseList struct {
	types.PaginatedResponse
	Data []ReportResponse `json:"data" required:"false" doc:"List of reports" example:"[]"`
}
