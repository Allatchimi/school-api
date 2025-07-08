package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
)

type SequenceResponse struct {
	types.BaseGormModelResponse
	School      *dataSchool.SchoolPublicResponse `json:"school" doc:"School"`
	Name        string                           `json:"name" required:"false" doc:"Sequence name"`
	Description string                           `json:"description" required:"false" doc:"Sequence description"`
}

type SequenceResponseList struct {
	types.PaginatedResponse
	Data []SequenceResponse `json:"data" required:"false" doc:"List of sections" example:"[]"`
}
