package data

import "time"

type ClassID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Class id" example:"1"`
}

type ClassSubjectID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Class subject id" example:"1"`
}

type ClassRequest struct {
	SchoolID    int64  `json:"schoolID" required:"true" doc:"School id" example:"1"`
	SpecialtyID int64  `json:"specialtyID" required:"true" doc:"Specialty id" example:"1"`
	Name        string `json:"name" required:"true" doc:"Class name" example:"Computer Science"`
	Description string `json:"description" required:"false" doc:"Class description" example:""`
}
type ClassSubjectRequest struct {
	SubjectID int64 `json:"subjectID" required:"false" doc:"Subject id" example:"1"`
	ClassID   int64 `json:"classID" required:"false" doc:"Class id" example:"1"`

	Coefficient  int    `json:"coefficient" required:"false" doc:"Coefficient" example:""`
	Program      string `json:"program" required:"false" doc:"Program" example:""`
	Requirements string `json:"requirements" required:"false" doc:"Requirements" example:""`

	IsValid     bool       `json:"isValid" required:"false" doc:"Is valid" example:"true"`
	InvalidDate *time.Time `json:"invalidDate" required:"false" doc:"Invalid date" example:""`
}
type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
}

type GetAllClassSubjectRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
	ClassID  int64 `json:"classID" query:"classID" required:"false" doc:"Class id" example:"1"`
}
