package data

import (
	"api/common/types"
	schoolData "api/services/school/common/school/data"
	facultyData "api/services/school/university/faculty/data"
)

type DepartmentResponse struct {
	types.BaseGormModelResponse
	School      *schoolData.SchoolPublicResponse   `json:"school" doc:"School"`
	Faculty     *facultyData.FacultyPublicResponse `json:"faculty" doc:"Faculty"`
	Name        string                             `json:"name" required:"false" doc:"Department name"`
	Description string                             `json:"description" required:"false" doc:"Department description"`
}

type DepartmentPublicResponse struct {
	School      *schoolData.SchoolPublicResponse   `json:"school" doc:"School"`
	Faculty     *facultyData.FacultyPublicResponse `json:"faculty" doc:"Faculty"`
	Name        string                             `json:"name" required:"false" doc:"Department name"`
	Description string                             `json:"description" required:"false" doc:"Department description"`
}

type DepartmentResponseList struct {
	types.PaginatedResponse
	Data []DepartmentResponse `json:"data" required:"false" doc:"List of departments" example:"[]"`
}
