package data

import "time"

type YearID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Academic year id" example:"1"`
}

type YearRequest struct {
	SchoolID  int64      `json:"schoolID" required:"true" doc:"School id" example:"1"`
	StartDate *time.Time `json:"startDate" required:"true" doc:"Academic year start date" example:""`
	EndDate   *time.Time `json:"endDate" required:"true" doc:"Academic year end date" example:""`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
}
