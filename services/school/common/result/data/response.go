package data

import (
	"api/common/types"
	examData "api/services/school/common/exam/data"
	dataSchool "api/services/school/common/school/data"
	studentData "api/services/school/common/student/data"
)

type ResultResponse struct {
	types.BaseGormModelResponse
	Value float64 `json:"value" required:"false" doc:"Value"`

	School  *dataSchool.SchoolPublicResponse   `json:"school" required:"false" doc:"School"`
	Student *studentData.StudentPublicResponse `json:"student" required:"true" doc:"Student"`
	Exam    *examData.ExamResponse             `json:"exam" required:"true" doc:"Exam"`
}

type ResultResponseList struct {
	types.PaginatedResponse
	Data []ResultResponse `json:"data" required:"false" doc:"List of result"`
}
