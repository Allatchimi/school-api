package data

import "time"

type YearID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Academic year id"`
}

type YearRequest struct {
	SchoolID int64 `json:"schoolID" required:"true" doc:"School id"`

	StartDate *time.Time `json:"startDate" required:"true" doc:"Academic year start date"`
	EndDate   *time.Time `json:"endDate" required:"true" doc:"Academic year end date"`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
}
