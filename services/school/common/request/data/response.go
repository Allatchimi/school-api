package data

import (
	"api/common/types"
	schoolData "api/services/school/common/school/data"
	studentData "api/services/school/common/student/data"
	yearData "api/services/school/common/year/data"
	classData "api/services/school/highschool/class/data"
	sequenceData "api/services/school/highschool/sequence/data"
	unitData "api/services/school/university/unit/data"
)

type RequestResponse struct {
	types.BaseGormModelResponse
	Status         string `json:"status" required:"false" doc:"Status"`
	StatusFeedback string `json:"statusFeedback" required:"false" doc:"Status feedback"`
	Audience       string `json:"audience" required:"false" doc:"Audience"`
	Title          string `json:"title" required:"false" doc:"Title"`
	Message        string `json:"message" required:"false" doc:"Message"`

	Document1 string `json:"document1" required:"false" doc:"Document1"`
	Document2 string `json:"document2" required:"false" doc:"Document2"`
	Document3 string `json:"document3" required:"false" doc:"Document3"`
	Document4 string `json:"document4" required:"false" doc:"Document4"`
	Document5 string `json:"document5" required:"false" doc:"Document5"`

	School       *schoolData.SchoolPublicResponse   `json:"school" required:"false" doc:"School"`
	Year         *yearData.YearResponse             `json:"year" required:"false" doc:"Year"`
	ClassSubject *classData.ClassSubjectResponse    `json:"classSubject" required:"false" doc:"Class subject"`
	Sequence     *sequenceData.SequenceResponse     `json:"semester" required:"false" doc:"Sequence"`
	Unit         *unitData.UnitResponse             `json:"unit" required:"false" doc:"Unit"`
	Student      *studentData.StudentPublicResponse `json:"student" required:"false" doc:"Student"`
}

type RequestResponseList struct {
	types.PaginatedResponse
	Data []RequestResponse `json:"data" required:"false" doc:"List of requests" example:"[]"`
}
