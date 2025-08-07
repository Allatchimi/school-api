package data

import (
	"api/common/types"
	examData "api/services/school/common/exam/data"
	dataSchool "api/services/school/common/school/data"
	studentData "api/services/school/common/student/data"
)

type ResultResponse struct {
	types.BaseGormModelResponse
	School  *dataSchool.SchoolPublicResponse   `json:"school,omitempty" required:"false" doc:"School"`
	Exam    *examData.ExamResponse             `json:"exam,omitempty" required:"true" doc:"Exam"`
	Student *studentData.StudentPublicResponse `json:"student,omitempty" required:"true" doc:"Student"`

	Score float64 `json:"score" required:"false" doc:"Score"`
}

type ResultTableResponse struct {
	types.BaseGormModelResponse
	School *dataSchool.SchoolResponse `json:"school,omitempty" required:"false" doc:"School"`
	Exam   *examData.ExamResponse     `json:"exam,omitempty" required:"true" doc:"Exam"`

	Status string `json:"status" required:"false" doc:"Status"`
}

type ResultResponseList struct {
	types.PaginatedResponse
	Data []ResultResponse `json:"data" required:"false" doc:"List of result"`
}

type ResultTableResponseList struct {
	types.PaginatedResponse
	Data []ResultTableResponse `json:"data" required:"false" doc:"List of result table"`
}
