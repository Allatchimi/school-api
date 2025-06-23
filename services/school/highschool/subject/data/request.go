package data

type SubjectID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Subject id"`
}

type SubjectRequest struct {
	SchoolID int64 `json:"schoolID" required:"true" doc:"School id"`

	Name        string `json:"name" required:"true" doc:"Name"`
	Description string `json:"description" required:"false" doc:"Description"`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
}
