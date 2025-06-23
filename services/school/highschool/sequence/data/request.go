package data

type SequenceID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Sequence id"`
}

type SequenceRequest struct {
	SchoolID    int64  `json:"schoolID" required:"true" doc:"School id"`
	Name        string `json:"name" required:"true" doc:"Sequence name"`
	Description string `json:"description" required:"false" doc:"Sequence description"`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
}
