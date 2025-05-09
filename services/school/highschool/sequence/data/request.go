package data

type SequenceID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Sequence id" example:"1"`
}

type SequenceRequest struct {
	SchoolID    int64  `json:"schoolID" required:"true" doc:"School id" example:"1"`
	Name        string `json:"name" required:"true" doc:"Sequence name" example:"Sciences"`
	Description string `json:"description" required:"false" doc:"Sequence description" example:""`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
}
