package data

import (
	"api/common/types"
	dataSchool "api/services/school/common/school/data"
	dataYear "api/services/school/common/year/data"
	dataClass "api/services/school/highschool/class/data"
	dataSequence "api/services/school/highschool/sequence/data"
	unitData "api/services/school/university/unit/data"
	"time"
)

type ScheduleResponse struct {
	types.BaseGormModelResponse
	IsCommon       bool       `json:"isCommon" required:"false" doc:"Is common"`
	Type           string     `json:"type" required:"false" doc:"Type"`
	DayOfTheWeek   string     `json:"dayOfTheWeek" required:"false" doc:"Day of the week"`
	RepeatCount    int        `json:"repeatCount" required:"false" doc:"Repeat count"`
	RepeatType     string     `json:"repeatType" required:"false" doc:"Repeat type"`
	StartTime      string     `json:"startTime" required:"false" doc:"Start time"`
	EndTime        string     `json:"endTime" required:"false" doc:"End time"`
	StartCountDate *time.Time `json:"startCountDate" required:"false" doc:"Start count date"`
	IsValid        bool       `json:"isValid" required:"false" doc:"Is valid"`
	InvalidDate    *time.Time `json:"invalidDate" required:"false" doc:"Invalid date"`

	School       *dataSchool.SchoolPublicResponse `json:"school" required:"false" doc:"School"`
	Year         *dataYear.YearResponse           `json:"year" required:"false" doc:"Year"`
	ClassSubject *dataClass.ClassSubjectResponse  `json:"classSubject" required:"false" doc:"Class subject"`
	Sequence     *dataSequence.SequenceResponse   `json:"semester" required:"false" doc:"Sequence"`
	Unit         *unitData.UnitResponse           `json:"unit" required:"false" doc:"Unit"`
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
	Data []ScheduleResponse `json:"data" required:"false" doc:"List of schedule" example:"[]"`
}

type ScheduleWeeklyViewResponseList struct {
	types.PaginatedResponse
	Data []ScheduleWeeklyViewResponse `json:"data" required:"false" doc:"List of schedule grouped by time and week" example:"[]"`
}
