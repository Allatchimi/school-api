package data

import (
	"api/common/types"
	"api/services/school/common/school/data"
)

type FacultyResponse struct {
	types.BaseGormModelResponse
	FacultyPublicResponse
}

type FacultyPublicResponse struct {
	School      *data.SchoolPublicResponse `json:"school" doc:"School"`
	Name        string                     `json:"name" required:"false" doc:"Faculty name"`
	Description string                     `json:"description" required:"false" doc:"Faculty description"`
}

type FacultyResponseList struct {
	types.PaginatedResponse
	Data []FacultyResponse `json:"data" required:"false" doc:"List of faculties" example:"[]"`
}
