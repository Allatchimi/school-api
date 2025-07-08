package data

import (
	"api/common/types"
	"api/services/school/common/school/data"
)

type SectionResponse struct {
	types.BaseGormModelResponse
	School      *data.SchoolResponse `json:"school" doc:"School"`
	Name        string               `json:"name" required:"false" doc:"Section name"`
	Description string               `json:"description" required:"false" doc:"Section description"`
}

type SectionResponseList struct {
	types.PaginatedResponse
	Data []SectionResponse `json:"data" required:"false" doc:"List of sections" example:"[]"`
}
