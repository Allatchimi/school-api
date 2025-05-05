package data

type ClassID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Class id" example:"1"`
}

type ClassRequest struct {
	SchoolID    int64  `json:"schoolID" required:"true" doc:"School id" example:"1"`
	SpecialtyID int64  `json:"specialtyID" required:"true" doc:"Specialty id" example:"1"`
	Name        string `json:"name" required:"true" doc:"Class name" example:"Computer Science"`
	Description string `json:"description" required:"false" doc:"Class description" example:""`
}
type SubjectClassRequest struct {
	SubjectID int64 `json:"subjectID" required:"false" doc:"Subject id" example:"1"`
	ClassID   int64 `json:"classID" required:"false" doc:"Class id" example:"1"`

	Coefficient  int    `json:"Coefficient" required:"false" doc:"Coefficient" example:""`
	Program      string `json:"Program" required:"false" doc:"Program" example:""`
	Requirements string `json:"Requirements" required:"false" doc:"Requirements" example:""`
}
type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
}
type GetAllSubjectClassRequest struct {
	ClassID int64 `json:"id" path:"id" required:"false" doc:"Class id" example:"1"`
}
