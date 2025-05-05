package data

import (
	"api/common/types"
	schoolData "api/services/school/common/school/data"
	specialtyData "api/services/school/highschool/specialty/data"
)

type ClassResponse struct {
	types.BaseGormModelResponse
	School      *schoolData.SchoolResponse       `json:"school" doc:"School"`
	Specialty   *specialtyData.SpecialtyResponse `json:"specialty" doc:"Specialty"`
	Name        string                           `json:"name" required:"false" doc:"Class name"`
	Description string                           `json:"description" required:"false" doc:"Class description"`
}

type ClassSubjectResponse struct {
	types.BaseGormModelResponse
	SubjectID int64 `json:"subjectID" required:"false" doc:"Subject id"`
	ClassID   int64 `json:"classID" required:"false" doc:"Class id"`

	Coefficient  int    `json:"Coefficient" required:"false" doc:"Coefficient"`
	Program      string `json:"Program" required:"false" doc:"Program"`
	Requirements string `json:"Requirements" required:"false" doc:"Requirements"`
}

type ClassResponseList struct {
	types.PaginatedResponse
	Data []ClassResponse `json:"data" required:"false" doc:"List of classes" example:"[]"`
}

type SubjectClassResponseList struct {
	types.PaginatedResponse
	Data []ClassSubjectResponse `json:"data" required:"false" doc:"List of subject for matching class" example:"[]"`
}
