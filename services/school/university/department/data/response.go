package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataFaculty "api/services/school/university/faculty/data"
)

type DepartmentResponse struct {
	types.BaseGormModelResponse
	School      *dataSchool.SchoolPublicResponse `json:"school" required:"false" doc:"School"`
	Faculty     *dataFaculty.FacultyResponse     `json:"faculty" required:"false" doc:"Faculty"`
	Name        string                           `json:"name" required:"false" doc:"Department name"`
	Description string                           `json:"description" required:"false" doc:"Department description"`
}

type DepartmentResponseList struct {
	types.PaginatedResponse
	Data []DepartmentResponse `json:"data" required:"false" doc:"List of departments" example:"[]"`
}
