package data

import (
	"api/common/types"
	schoolData "api/services/school/common/school/data"
	domainData "api/services/school/university/domain/data"
)

type SemesterResponse struct {
	types.BaseGormModelResponse
	School      *schoolData.SchoolResponse `json:"school" doc:"School"`
	Domain      *domainData.DomainResponse `json:"domain" doc:"Domain"`
	Name        string                     `json:"name" required:"false" doc:"Semester name"`
	Description string                     `json:"description" required:"false" doc:"Semester description"`
}

type SemesterResponseList struct {
	types.PaginatedResponse
	Data []SemesterResponse `json:"data" required:"false" doc:"List of semesters" example:"[]"`
}
