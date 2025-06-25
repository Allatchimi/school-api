package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataDomain "api/services/school/university/domain/data"
)

type SemesterResponse struct {
	types.BaseGormModelResponse
	SemesterPublicResponse
}

type SemesterPublicResponse struct {
	School      *dataSchool.SchoolPublicResponse `json:"school" doc:"School"`
	Domain      *dataDomain.DomainPublicResponse `json:"domain" doc:"Domain"`
	Name        string                           `json:"name" required:"false" doc:"Semester name"`
	Description string                           `json:"description" required:"false" doc:"Semester description"`
}

type SemesterResponseList struct {
	types.PaginatedResponse
	Data []SemesterResponse `json:"data" required:"false" doc:"List of semesters" example:"[]"`
}
