package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	"time"
)

type YearResponse struct {
	types.BaseGormModelResponse
	YearPublicResponse
}

type YearPublicResponse struct {
	Name      string                           `json:"name" doc:"Name"`
	StartDate *time.Time                       `json:"startDate" doc:"Academic year start date"`
	EndDate   *time.Time                       `json:"endDate" doc:"Academic year end date"`
	School    *dataSchool.SchoolPublicResponse `json:"school" doc:"School"`
}

type YearResponseList struct {
	types.PaginatedResponse
	Data []YearResponse `json:"data" required:"false" doc:"List of academic years" example:"[]"`
}
