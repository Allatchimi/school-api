package data

import (
	"api/common/types"
	"api/services/school/common/school/data"
)

type SubjectResponse struct {
	types.BaseGormModelResponse
	School *data.SchoolPublicResponse `json:"school" doc:"School"`

	Name        string `json:"name" required:"false" doc:"Name"`
	Description string `json:"description" required:"false" doc:"Description"`
}

type SubjectPublicResponse struct {
	School *data.SchoolPublicResponse `json:"school" doc:"School"`

	Name        string `json:"name" required:"false" doc:"Name"`
	Description string `json:"description" required:"false" doc:"Description"`
}

type SubjectResponseList struct {
	types.PaginatedResponse
	Data []SubjectResponse `json:"data" required:"false" doc:"List of sections" example:"[]"`
}
