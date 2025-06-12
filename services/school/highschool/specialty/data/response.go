package data

import (
	"api/common/types"
	schoolData "api/services/school/common/school/data"
	sectionData "api/services/school/highschool/section/data"
)

type SpecialtyResponse struct {
	types.BaseGormModelResponse
	SpecialtyPublicResponse
}

type SpecialtyPublicResponse struct {
	School      *schoolData.SchoolPublicResponse   `json:"school" doc:"School"`
	Section     *sectionData.SectionPublicResponse `json:"section" doc:"Section"`
	Name        string                             `json:"name" required:"false" doc:"Specialty name"`
	Description string                             `json:"description" required:"false" doc:"Specialty description"`
}

type SpecialtyResponseList struct {
	types.PaginatedResponse
	Data []SpecialtyResponse `json:"data" required:"false" doc:"List of specialties" example:"[]"`
}
