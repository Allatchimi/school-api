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

type ScheduleResponse struct {
	types.BaseGormModelResponse
	Type           string     `json:"type" required:"false" doc:"Type"`
	DayOfTheWeek   string     `json:"dayOfTheWeek" required:"false" doc:"Day of the week"`
	RepeatCount    int        `json:"repeatCount" required:"false" doc:"Repeat count"`
	RepeatType     string     `json:"repeatType" required:"false" doc:"Repeat type"`
	StartTime      string     `json:"startTime" required:"false" doc:"Start time"`
	EndTime        string     `json:"endTime" required:"false" doc:"End time"`
	StartCountDate *time.Time `json:"startCountDate" required:"false" doc:"Start count date"`

	School       *schoolData.SchoolPublicResponse `json:"school" required:"false" doc:"School"`
	Year         *yearData.YearResponse           `json:"year" required:"false" doc:"Year"`
	ClassSubject *classData.ClassSubjectResponse  `json:"classSubject" required:"false" doc:"Class subject"`
	Sequence     *sequenceData.SequenceResponse   `json:"semester" required:"false" doc:"Sequence"`
	Unit         *unitData.UnitResponse           `json:"unit" required:"false" doc:"Unit"`
	IsValid      bool                             `json:"isValid" required:"false" doc:"Is valid"`
	InvalidDate  *time.Time                       `json:"invalidDate" required:"false" doc:"Invalid date"`
}

type ScheduleWeeklyViewResponse struct {
	StartTime string `json:"startTime" required:"false" doc:"Start time"`
	EndTime   string `json:"endTime" required:"false" doc:"End time"`

	Monday    []ScheduleResponse `json:"monday" required:"false" doc:"Monday"`
	Tuesday   []ScheduleResponse `json:"tuesday" required:"false" doc:"Tuesday"`
	Wednesday []ScheduleResponse `json:"wednesday" required:"false" doc:"Wednesday"`
	Thursday  []ScheduleResponse `json:"thursday" required:"false" doc:"Thursday"`
	Friday    []ScheduleResponse `json:"friday" required:"false" doc:"Friday"`
	Saturday  []ScheduleResponse `json:"saturday" required:"false" doc:"Saturday"`
	Sunday    []ScheduleResponse `json:"sunday" required:"false" doc:"Sunday"`
}

type ScheduleResponseList struct {
	types.PaginatedResponse
	Data []ScheduleResponse `json:"data" required:"false" doc:"List of schedules" example:"[]"`
}

type ScheduleWeeklyViewResponseList struct {
	types.PaginatedResponse
	Data []ScheduleWeeklyViewResponse `json:"data" required:"false" doc:"List of schedules grouped by time and week" example:"[]"`
}
