package data

type DepartmentID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Department id"`
}

type DepartmentRequest struct {
	SchoolID  int64 `json:"schoolID" required:"true" doc:"School id"`
	FacultyID int64 `json:"facultyID" required:"true" doc:"Faculty id"`

	Name        string `json:"name" required:"true" doc:"Department name"`
	Description string `json:"description" required:"false" doc:"Department description"`
}

type GetAllRequest struct {
	SchoolID  int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
	FacultyID int64 `json:"facultyID" query:"facultyID" required:"false" doc:"Faculty id"`
}
