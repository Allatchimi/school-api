package data

type SectionID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Section id"`
}

type SectionRequest struct {
	SchoolID    int64  `json:"schoolID" required:"true" doc:"School id"`
	Name        string `json:"name" required:"true" doc:"Section name"`
	Description string `json:"description" required:"false" doc:"Section description"`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
}
