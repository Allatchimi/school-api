package data

type SubjectID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Subject id" example:"1"`
}

type SubjectRequest struct {
	SchoolID    int64  `json:"schoolID" required:"true" doc:"School id" example:"1"`
	Name        string `json:"name" required:"true" doc:"Subject name" example:"Sciences"`
	Description string `json:"description" required:"false" doc:"Subject description" example:""`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
}
