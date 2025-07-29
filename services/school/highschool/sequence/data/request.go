package data

type SequenceID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Sequence id"`
}

type SequenceRequest struct {
	SchoolID  int64 `json:"schoolID" required:"true" doc:"School id"`
	QuarterID int64 `json:"quarterID" required:"true" doc:"Quarter id"`

	Name        string `json:"name" required:"true" doc:"Sequence name"`
	Description string `json:"description" required:"false" doc:"Sequence description"`
}

type GetAllRequest struct {
	SchoolID  int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
	QuarterID int64 `json:"quarterID" query:"quarterID" required:"false" doc:"Quarter id"`
}
