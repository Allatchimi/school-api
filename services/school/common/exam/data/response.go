package data

import (
	"api/common/types"
	schoolData "api/services/school/common/school/data"
	yearData "api/services/school/common/year/data"
	subjectData "api/services/school/highschool/subject/data"
	tuData "api/services/school/university/tu/data"
)

type ExamResponse struct {
	types.BaseGormModelResponse
	Percentage  int    `json:"percentage" required:"false" doc:"Percentage"`
	Description string `json:"description" required:"false" doc:"Description"`

	School *schoolData.SchoolResponse `json:"school" required:"true" doc:"School"`
	Year   *yearData.YearResponse     `json:"Year" required:"true" doc:"Year"`
	Type   *ExamTypeResponse          `json:"type" required:"false" doc:"Type"`

	TeachingUnit *tuData.TeachingUnitResponse `json:"teachingUnit" required:"false" doc:"Teaching unit"`
	Subject      *subjectData.SubjectResponse `json:"subject" required:"false" doc:"Subject"`
}

type ExamTypeResponse struct {
	types.BaseGormModelResponse
	Name        string `json:"name" required:"false" doc:"Name"`
	Description string `json:"description" required:"false" doc:"Description"`

	School *schoolData.SchoolResponse `json:"school" required:"true" doc:"School"`
}

type ExamResponseList struct {
	types.PaginatedResponse
	Data []ExamResponse `json:"data" required:"false" doc:"List of exams" example:"[]"`
}

type ExamTypeResponseList struct {
	types.PaginatedResponse
	Data []ExamTypeResponse `json:"data" required:"false" doc:"List of exam types" example:"[]"`
}
