package data

import "api/common/types"

type ScheduleID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Schedule id" example:"1"`
}

type ScheduleRequest struct {
	SchoolID       int64 `json:"schoolID" required:"false" doc:"School id" example:"1"`
	YearID         int64 `json:"yearID" required:"false" doc:"Year id" example:"1"`
	ClassSubjectID int64 `json:"classSubjectID" required:"false" doc:"Class Subject id" example:"1"`
	UnitID         int64 `json:"unitID" required:"false" doc:"Unit id" example:"1"`

	Type         string `json:"type" required:"false" doc:"Type" example:"Normal"`
	DayOfTheWeek string `json:"dayOfTheWeek" required:"false" doc:"Day of the week" example:"MONDAY"`
	RepeatCount  string `json:"repeatCount" required:"false" doc:"Repeat count" example:"1"`
	RepeatType   string `json:"repeatType" required:"false" doc:"Repeat type" example:"WEEKLY"`
	StartTime    string `json:"startTime" required:"false" doc:"Start time" example:"09:00:00"`
	EndTime      string `json:"endTime" required:"false" doc:"End time" example:"10:00:00"`
}

type GetAllRequest struct {
	types.FilterSchoolYearClassSubjectUnitRequest
}
