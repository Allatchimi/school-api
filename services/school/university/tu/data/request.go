package data

import "time"

type TeachingUnitID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Teaching unit id" example:"1"`
}

type TeachingUnitRequest struct {
	SchoolID   int64 `json:"schoolID" required:"true" doc:"School id" example:"1"`
	DomainID   int64 `json:"domainID" required:"true" doc:"Domain id" example:"1"`
	LevelID    int64 `json:"levelID" required:"true" doc:"Level id" example:"1"`
	SemesterID int64 `json:"semesterID" required:"true" doc:"Semester id" example:"1"`

	Name         string `json:"name" required:"true" doc:"Name" example:"MATH110"`
	Description  string `json:"description" required:"false" doc:"Description" example:""`
	Credit       int    `json:"credit" required:"true" min:"1" doc:"Credit" example:"1"`
	Program      string `json:"program" required:"false" doc:"Program" example:""`
	Requirements string `json:"requirements" required:"false" doc:"Requirements" example:""`

	IsValid     bool       `json:"isValid" required:"false" doc:"Is valid" example:"true"`
	InvalidDate *time.Time `json:"invalidDate" required:"false" doc:"Invalid date" example:""`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
}
