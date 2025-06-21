package data

import (
	"api/common/types"
	"time"
)

type ScheduleID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Schedule id" example:"1"`
}

type ScheduleRequest struct {
	SchoolID       int64 `json:"schoolID" required:"true" doc:"School id" example:"1"`
	YearID         int64 `json:"yearID" required:"true" doc:"Year id" example:"1"`
	ClassSubjectID int64 `json:"classSubjectID" required:"false" doc:"Class Subject id" example:"1"`
	UnitID         int64 `json:"unitID" required:"false" doc:"Unit id" example:"1"`

	IsGeneric      bool       `json:"isGeneric" required:"true" doc:"Is generic" example:"false"`
	Type           string     `json:"type" required:"true" doc:"Type" example:"Normal"`
	DayOfTheWeek   string     `json:"dayOfTheWeek" required:"true" doc:"Day of the week" example:"MONDAY"`
	RepeatCount    int        `json:"repeatCount" required:"true" doc:"Repeat count" example:"1"`
	RepeatType     string     `json:"repeatType" required:"true" doc:"Repeat type" example:"WEEKLY"`
	StartTime      string     `json:"startTime" required:"true" doc:"Start time" example:"09:00:00"`
	EndTime        string     `json:"endTime" required:"true" doc:"End time" example:"10:00:00"`
	StartCountDate *time.Time `json:"startCountDate" required:"true" doc:"Start count date" example:"2025-06-17"`
	IsValid        bool       `json:"isValid" required:"true" doc:"Is valid" example:"true"`
}

type GetAllRequest struct {
	types.FilterSchoolYearClassSubjectUnitRequest
	Type string `json:"type" query:"type" required:"false" enum:"all,default,generic" doc:"Type" example:"default"`
}
