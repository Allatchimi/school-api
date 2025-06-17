package data

type DomainID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Domain id" example:"1"`
}

type DomainRequest struct {
	SchoolID     int64  `json:"schoolID" required:"true" doc:"School id" example:"1"`
	DepartmentID int64  `json:"departmentID" required:"true" doc:"Department id" example:"1"`
	Name         string `json:"name" required:"true" doc:"Domain name" example:"Computer Science"`
	Description  string `json:"description" required:"false" doc:"Domain description" example:""`
}

type GetAllRequest struct {
	SchoolID     int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
	DepartmentID int64 `json:"departmentID" query:"departmentID" required:"false" doc:"Department id" example:"1"`
}
