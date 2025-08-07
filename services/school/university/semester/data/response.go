package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataDomain "api/services/school/university/domain/data"
)

type SemesterResponse struct {
	types.BaseGormModelResponse
	School *dataSchool.SchoolPublicResponse `json:"school,omitempty" required:"false" doc:"School"`
	Domain *dataDomain.DomainResponse       `json:"domain,omitempty" required:"false" doc:"Domain"`

	Name        string `json:"name" required:"false" doc:"Semester name"`
	Description string `json:"description" required:"false" doc:"Semester description"`
}

type SemesterResponseList struct {
	types.PaginatedResponse
	Data []SemesterResponse `json:"data" required:"false" doc:"List of semester"`
}
