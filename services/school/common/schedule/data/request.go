package data

import (
	"api/common/types"
	"time"
)

type ScheduleID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Schedule id"`
}

type ScheduleRequest struct {
	SchoolID       int64 `json:"schoolID" required:"true" doc:"School id"`
	YearID         int64 `json:"yearID" required:"true" doc:"Year id"`
	ClassSubjectID int64 `json:"classSubjectID" required:"false" doc:"Class Subject id"`
	UnitID         int64 `json:"unitID" required:"false" doc:"Unit id"`

	IsCommon       bool       `json:"isCommon" required:"true" doc:"Is common"`
	Type           string     `json:"type" required:"true" enum:"course,tp,td,pause,event" doc:"Type"`
	DayOfTheWeek   string     `json:"dayOfTheWeek" required:"true" doc:"Day of the week"`
	RepeatCount    int        `json:"repeatCount" required:"true" doc:"Repeat count"`
	RepeatType     string     `json:"repeatType" required:"true"  enum:"daily,weekly,monthly,yearly,onetime" doc:"Repeat type"`
	StartTime      string     `json:"startTime" required:"true" doc:"Start time"`
	EndTime        string     `json:"endTime" required:"true" doc:"End time"`
	StartCountDate *time.Time `json:"startCountDate" required:"true" doc:"Start count date"`
	IsValid        bool       `json:"isValid" required:"true" doc:"Is valid"`
}

type GetAllRequest struct {
	types.FilterSchoolYearClassSubjectUnitRequest
	types.FilterTeacherStudentRequest
	Type string `json:"type" query:"type" required:"false" doc:"Type"`
}
