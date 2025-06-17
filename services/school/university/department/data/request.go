package data

type DepartmentID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Department id" example:"1"`
}

type DepartmentRequest struct {
	SchoolID    int64  `json:"schoolID" required:"true" doc:"School id" example:"1"`
	FacultyID   int64  `json:"facultyID" required:"true" doc:"Faculty id" example:"1"`
	Name        string `json:"name" required:"true" doc:"Department name" example:"Computer Science"`
	Description string `json:"description" required:"false" doc:"Department description" example:""`
}

type GetAllRequest struct {
	SchoolID  int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
	FacultyID int64 `json:"facultyID" query:"facultyID" required:"false" doc:"Faculty id" example:"1"`
}
