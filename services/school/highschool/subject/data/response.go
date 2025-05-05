package data

import (
	"api/common/types"
	"api/services/school/common/school/data"
)

type SubjectResponse struct {
	types.BaseGormModelResponse
	School      *data.SchoolResponse `json:"school" doc:"School"`
	Name        string               `json:"name" required:"false" doc:"Subject name"`
	Description string               `json:"description" required:"false" doc:"Subject description"`
}

type SubjectResponseList struct {
	types.PaginatedResponse
	Data []SubjectResponse `json:"data" required:"false" doc:"List of sections" example:"[]"`
}
