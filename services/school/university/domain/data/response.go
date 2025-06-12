package data

import (
	"api/common/types"
	schoolData "api/services/school/common/school/data"
	departmentData "api/services/school/university/department/data"
)

type DomainResponse struct {
	types.BaseGormModelResponse
	DomainPublicResponse
}

type DomainPublicResponse struct {
	School      *schoolData.SchoolPublicResponse         `json:"school" doc:"School"`
	Department  *departmentData.DepartmentPublicResponse `json:"department" doc:"Department"`
	Name        string                                   `json:"name" required:"false" doc:"Department name"`
	Description string                                   `json:"description" required:"false" doc:"Department description"`
}

type DomainResponseList struct {
	types.PaginatedResponse
	Data []DomainResponse `json:"data" required:"false" doc:"List of departments" example:"[]"`
}
