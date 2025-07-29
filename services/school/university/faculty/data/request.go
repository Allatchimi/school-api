package data

type FacultyID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Faculty id"`
}

type FacultyRequest struct {
	SchoolID int64 `json:"schoolID" required:"true" doc:"School id"`

	Name        string `json:"name" required:"true" doc:"Faculty name"`
	Description string `json:"description" required:"false" doc:"Faculty description"`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
}
