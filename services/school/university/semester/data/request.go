package data

type SemesterID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Semester id" example:"1"`
}

type SemesterRequest struct {
	SchoolID    int64  `json:"schoolID" required:"true" doc:"School id" example:"1"`
	Name        string `json:"name" required:"true" doc:"Semester name" example:"Sciences"`
	Description string `json:"description" required:"false" doc:"Semester description" example:""`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
}
