package data

import (
	"api/common/types"
	subjectData "api/services/school/highschool/subject/data"
	tuData "api/services/school/university/tu/data"
)

type ExamResponse struct {
	types.BaseGormModelResponse
	SchoolID     int64                        `json:"schoolID" required:"true" doc:"School id"`
	TeachingUnit *tuData.TeachingUnitResponse `json:"teachingUnit" required:"false" doc:"Teaching unit"`
	Subject      *subjectData.SubjectResponse `json:"subject" required:"false" doc:"Subject"`
	Type         *ExamTypeResponse            `json:"type" required:"false" doc:"Type"`
	Percentage   int                          `json:"percentage" required:"false" doc:"Percentage"`
	Description  string                       `json:"description" required:"false" doc:"Description"`
}

type ExamTypeResponse struct {
	types.BaseGormModelResponse
	SchoolID    int64  `json:"schoolID" required:"false" doc:"School id" example:"1"`
	Name        string `json:"name" required:"false" doc:"Name"`
	Description string `json:"description" required:"false" doc:"Description"`
}

type ExamResponseList struct {
	types.PaginatedResponse
	Data []ExamResponse `json:"data" required:"false" doc:"List of exams" example:"[]"`
}

type ExamTypeResponseList struct {
	types.PaginatedResponse
	Data []ExamTypeResponse `json:"data" required:"false" doc:"List of exam types" example:"[]"`
}
