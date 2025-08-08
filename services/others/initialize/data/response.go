package data

import (
	dataSchool "api/services/school/common/school/data"
	dataYear "api/services/school/common/year/data"
)

type InitializeResponse struct {
	School *dataSchool.SchoolPublicResponse `json:"school,omitempty" required:"false" doc:"School"`
	Years  []dataYear.YearResponse          `json:"years" required:"false" doc:"Years"`
}
