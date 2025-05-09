package data

type DocumentID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Document id" example:"1"`
}

type DocumentRequest struct {
	SchoolID       int64 `json:"schoolID" required:"true" doc:"School id"`
	YearID         int64 `json:"yearID" required:"true" doc:"Year id"`
	ClassSubjectID int64 `json:"classSubjectID" required:"true" doc:"Subject class id"`
	UnitID         int64 `json:"unitID" required:"true" doc:"Unit id"`

	URL         string `json:"url" required:"false" doc:"URL"`
	Name        string `json:"name" required:"true" doc:"Name"`
	Description string `json:"description" required:"false" doc:"Description"`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
}
