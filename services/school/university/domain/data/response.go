package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataDepartment "api/services/school/university/department/data"
)

type DomainResponse struct {
	types.BaseGormModelResponse
	School      *dataSchool.SchoolPublicResponse   `json:"school" required:"false" doc:"School"`
	Department  *dataDepartment.DepartmentResponse `json:"department" required:"false" doc:"Department"`
	Name        string                             `json:"name" required:"false" doc:"Department name"`
	Description string                             `json:"description" required:"false" doc:"Department description"`
}

type DomainResponseList struct {
	types.PaginatedResponse
	Data []DomainResponse `json:"data" required:"false" doc:"List of department"`
}
