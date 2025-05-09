package data

import (
	"api/common/types"
	schoolData "api/services/school/common/school/data"
	sequenceData "api/services/school/highschool/sequence/data"
)

type QuarterResponse struct {
	types.BaseGormModelResponse
	School      *schoolData.SchoolPublicResponse `json:"school" doc:"School"`
	Name        string                           `json:"name" required:"false" doc:"Quarter name"`
	Description string                           `json:"description" required:"false" doc:"Quarter description"`
}

type QuarterPublicResponse struct {
	School      *schoolData.SchoolPublicResponse `json:"school" doc:"School"`
	Name        string                           `json:"name" required:"false" doc:"Quarter name"`
	Description string                           `json:"description" required:"false" doc:"Quarter description"`
}

type QuarterSequenceResponse struct {
	types.BaseGormModelResponse
	Quarter  *QuarterPublicResponse               `json:"Quarter" doc:"Quarter"`
	Sequence *sequenceData.SequencePublicResponse `json:"Sequence" doc:"Sequence"`
}

type QuarterSequencePublicResponse struct {
	Quarter  *QuarterPublicResponse               `json:"Quarter" doc:"Quarter"`
	Sequence *sequenceData.SequencePublicResponse `json:"Sequence" doc:"Sequence"`
}

type QuarterResponseList struct {
	types.PaginatedResponse
	Data []QuarterResponse `json:"data" required:"false" doc:"List of quarters" example:"[]"`
}

type QuarterSequenceResponseList struct {
	types.PaginatedResponse
	Data []QuarterSequenceResponse `json:"data" required:"false" doc:"List of sequences for matching quarter id" example:"[]"`
}
