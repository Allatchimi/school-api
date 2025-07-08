package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataSequence "api/services/school/highschool/sequence/data"
)

type QuarterResponse struct {
	types.BaseGormModelResponse
	School      *dataSchool.SchoolPublicResponse `json:"school" doc:"School"`
	Name        string                           `json:"name" required:"false" doc:"Quarter name"`
	Description string                           `json:"description" required:"false" doc:"Quarter description"`
}

type QuarterSequenceResponse struct {
	types.BaseGormModelResponse
	Quarter  *QuarterResponse               `json:"Quarter" doc:"Quarter"`
	Sequence *dataSequence.SequenceResponse `json:"Sequence" doc:"Sequence"`
}

type QuarterResponseList struct {
	types.PaginatedResponse
	Data []QuarterResponse `json:"data" required:"false" doc:"List of quarters" example:"[]"`
}

type QuarterSequenceResponseList struct {
	types.PaginatedResponse
	Data []QuarterSequenceResponse `json:"data" required:"false" doc:"List of sequences for matching quarter id" example:"[]"`
}
