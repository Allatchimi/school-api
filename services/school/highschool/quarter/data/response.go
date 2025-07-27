package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
)

type QuarterResponse struct {
	types.BaseGormModelResponse
	School *dataSchool.SchoolPublicResponse `json:"school" doc:"School"`

	Name        string `json:"name" required:"false" doc:"Quarter name"`
	Description string `json:"description" required:"false" doc:"Quarter description"`
}

type QuarterResponseList struct {
	types.PaginatedResponse
	Data []QuarterResponse `json:"data" required:"false" doc:"List of quarter"`
}
