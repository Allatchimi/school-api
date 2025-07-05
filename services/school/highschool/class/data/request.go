package data

import "time"

type ClassID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Class id"`
}

type ClassSubjectID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Class subject id"`
}

type ClassRequest struct {
	SchoolID    int64  `json:"schoolID" required:"true" doc:"School id"`
	SpecialtyID int64  `json:"specialtyID" required:"true" doc:"Specialty id"`
	Name        string `json:"name" required:"true" doc:"Class name"`
	Description string `json:"description" required:"false" doc:"Class description"`
}
type ClassSubjectRequest struct {
	SchoolID  int64 `json:"schoolID" required:"true" doc:"School id"`
	SubjectID int64 `json:"subjectID" required:"false" doc:"Subject id"`
	ClassID   int64 `json:"classID" required:"false" doc:"Class id"`

	Coefficient  int    `json:"coefficient" required:"false" doc:"Coefficient"`
	Program      string `json:"program" required:"false" doc:"Program"`
	Requirements string `json:"requirements" required:"false" doc:"Requirements"`

	IsValid     bool       `json:"isValid" required:"false" doc:"Is valid"`
	InvalidDate *time.Time `json:"invalidDate" required:"false" doc:"Invalid date"`
}
type GetAllRequest struct {
	SchoolID    int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
	SpecialtyID int64 `json:"specialtyID" query:"specialtyID" required:"false" doc:"Specialty id"`
}

type GetAllClassSubjectRequest struct {
	SchoolID  int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
	ClassID   int64 `json:"classID" query:"classID" required:"false" doc:"Class id"`
	SubjectID int64 `json:"subjectID" query:"subjectID" required:"false" doc:"Subject id"`
}
