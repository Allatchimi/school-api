package data

type DocumentID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Document id" example:"1"`
}

type DocumentRequest struct {
	SchoolID       int64  `json:"schoolID" required:"true" doc:"School id"`
	YearID         int64  `json:"yearID" required:"true" doc:"Year id"`
	SubjectID      int64  `json:"subjectID" required:"true" doc:"Subject id"`
	TeachingUnitID int64  `json:"teachingUnitID" required:"true" doc:"Teaching unit id"`
	Type           string `json:"type" required:"true" doc:"Type"`
	URL            string `json:"url" required:"false" doc:"URL"`
	Name           string `json:"name" required:"true" doc:"Name"`
	Description    string `json:"description" required:"false" doc:"Description"`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
}
