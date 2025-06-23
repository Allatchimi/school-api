package data

type SpecialtyID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Specialty id"`
}

type SpecialtyRequest struct {
	SchoolID    int64  `json:"schoolID" required:"true" doc:"School id"`
	SectionID   int64  `json:"sectionID" required:"true" doc:"Section id"`
	Name        string `json:"name" required:"true" doc:"Specialty name"`
	Description string `json:"description" required:"false" doc:"Specialty description"`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
}
