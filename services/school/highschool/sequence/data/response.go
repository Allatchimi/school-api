package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataQuarter "api/services/school/highschool/quarter/data"
)

type SequenceResponse struct {
	types.BaseGormModelResponse
	School  *dataSchool.SchoolPublicResponse `json:"school" doc:"School"`
	Quarter *dataQuarter.QuarterResponse     `json:"quarter" doc:"Quarter"`

	Name        string `json:"name" required:"false" doc:"Sequence name"`
	Description string `json:"description" required:"false" doc:"Sequence description"`
}

type SequenceResponseList struct {
	types.PaginatedResponse
	Data []SequenceResponse `json:"data" required:"false" doc:"List of section"`
}
