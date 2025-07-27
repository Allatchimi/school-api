package data

type DomainID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Domain id"`
}

type DomainRequest struct {
	SchoolID     int64 `json:"schoolID" required:"true" doc:"School id"`
	DepartmentID int64 `json:"departmentID" required:"true" doc:"Department id"`

	Name        string `json:"name" required:"true" doc:"Domain name"`
	Description string `json:"description" required:"false" doc:"Domain description"`
}

type GetAllRequest struct {
	SchoolID     int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
	DepartmentID int64 `json:"departmentID" query:"departmentID" required:"false" doc:"Department id"`
}
