package data

import (
	"api/common/types"
	schoolData "api/services/school/common/school/data"
	yearData "api/services/school/common/year/data"
	classData "api/services/school/highschool/class/data"
	sequenceData "api/services/school/highschool/sequence/data"
	unitData "api/services/school/university/unit/data"
)

type ScheduleResponse struct {
	types.BaseGormModelResponse
	SchedulePublicResponse
}

type SchedulePublicResponse struct {
	Type         string `json:"type" required:"false" doc:"Type"`
	DayOfTheWeek string `json:"dayOfTheWeek" required:"false" doc:"Day of the week"`
	RepeatCount  string `json:"repeatCount" required:"false" doc:"Repeat count"`
	RepeatType   string `json:"repeatType" required:"false" doc:"Repeat type"`
	StartTime    string `json:"startTime" required:"false" doc:"Start time"`
	EndTime      string `json:"endTime" required:"false" doc:"End time"`

	School       *schoolData.SchoolPublicResponse      `json:"school" required:"false" doc:"School"`
	Year         *yearData.YearPublicResponse          `json:"Year" required:"false" doc:"Year"`
	ClassSubject *classData.ClassSubjectPublicResponse `json:"classSubject" required:"false" doc:"Class subject"`
	Sequence     *sequenceData.SequencePublicResponse  `json:"semester" required:"false" doc:"Sequence"`
	Unit         *unitData.UnitPublicResponse          `json:"unit" required:"false" doc:"Unit"`
}

type ScheduleResponseList struct {
	types.PaginatedResponse
	Data []ScheduleResponse `json:"data" required:"false" doc:"List of schedules" example:"[]"`
}
