package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataDepartment "api/services/school/university/department/data"
)

type DomainResponse struct {
	types.BaseGormModelResponse
	DomainPublicResponse
}

type DomainPublicResponse struct {
	School      *dataSchool.SchoolPublicResponse         `json:"school" doc:"School"`
	Department  *dataDepartment.DepartmentPublicResponse `json:"department" doc:"Department"`
	Name        string                                   `json:"name" required:"false" doc:"Department name"`
	Description string                                   `json:"description" required:"false" doc:"Department description"`
}

type DomainResponseList struct {
	types.PaginatedResponse
	Data []DomainResponse `json:"data" required:"false" doc:"List of departments" example:"[]"`
}
