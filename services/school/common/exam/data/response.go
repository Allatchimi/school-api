package data

import (
	"api/common/types"
	schoolData "api/services/school/common/school/data"
	yearData "api/services/school/common/year/data"
	classData "api/services/school/highschool/class/data"
	sequenceData "api/services/school/highschool/sequence/data"
	unitData "api/services/school/university/unit/data"
	"time"
)

type ExamResponse struct {
	types.BaseGormModelResponse
	Status          string     `json:"status" required:"false" doc:"Status"`
	Percentage      int        `json:"percentage" required:"false" doc:"Percentage"`
	Description     string     `json:"description" required:"false" doc:"Description"`
	LocationType    string     `json:"locationType" required:"false" doc:"Location type"`
	LocationDetails string     `json:"locationDetails" required:"false" doc:"Location details"`
	Requirements    string     `json:"requirements" required:"false" doc:"Requirements"`
	AllowedItems    string     `json:"allowedItems" required:"false" doc:"Allowed items"`
	StartDate       *time.Time `json:"startDate" required:"false" doc:"Start date"`
	EndDate         *time.Time `json:"endDate" required:"false" doc:"End date"`

	School       *schoolData.SchoolPublicResponse `json:"school" required:"false" doc:"School"`
	Year         *yearData.YearResponse           `json:"Year" required:"false" doc:"Year"`
	Type         *ExamTypeResponse                `json:"type" required:"false" doc:"Type"`
	ClassSubject *classData.ClassSubjectResponse  `json:"classSubject" required:"false" doc:"Class subject"`
	Sequence     *sequenceData.SequenceResponse   `json:"semester" required:"false" doc:"Sequence"`
	Unit         *unitData.UnitResponse           `json:"unit" required:"false" doc:"Unit"`
}

type ExamTypeResponse struct {
	types.BaseGormModelResponse
	Name        string `json:"name" required:"false" doc:"Name"`
	Description string `json:"description" required:"false" doc:"Description"`

	School *schoolData.SchoolPublicResponse `json:"school" required:"false" doc:"School"`
}

type ExamResponseList struct {
	types.PaginatedResponse
	Data []ExamResponse `json:"data" required:"false" doc:"List of exams" example:"[]"`
}

type ExamTypeResponseList struct {
	types.PaginatedResponse
	Data []ExamTypeResponse `json:"data" required:"false" doc:"List of exam types" example:"[]"`
}
