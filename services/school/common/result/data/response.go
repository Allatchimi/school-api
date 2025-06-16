package data

import (
	"api/common/types"
	examData "api/services/school/common/exam/data"
	studentData "api/services/school/common/student/data"
)

type ResultResponse struct {
	types.BaseGormModelResponse
	ResultPublicResponse
}

type ResultPublicResponse struct {
	Value float64 `json:"value" required:"false" doc:"Value"`

	Student *studentData.StudentPublicResponse `json:"Student" required:"true" doc:"Student"`
	Exam    *examData.ExamPublicResponse       `json:"Exam" required:"true" doc:"Exam"`
}

type ResultResponseList struct {
	types.PaginatedResponse
	Data []ResultResponse `json:"data" required:"false" doc:"List of results" example:"[]"`
}
