package data

type SemesterID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Semester id"`
}

type SemesterRequest struct {
	SchoolID    int64  `json:"schoolID" required:"true" doc:"School id"`
	Name        string `json:"name" required:"true" doc:"Semester name"`
	Description string `json:"description" required:"false" doc:"Semester description"`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
}
