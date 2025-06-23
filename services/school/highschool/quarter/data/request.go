package data

type QuarterID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Quarter id"`
}

type QuarterRequest struct {
	SchoolID    int64  `json:"schoolID" required:"true" doc:"School id"`
	Name        string `json:"name" required:"true" doc:"Quarter name"`
	Description string `json:"description" required:"false" doc:"Quarter description"`
}

type QuarterSequenceRequest struct {
	QuarterID  int64 `json:"quarterID" required:"true" doc:"Quarter id"`
	SequenceID int64 `json:"sequenceID" required:"true" doc:"Sequence id"`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
}

type GetAllQuarterSequenceRequest struct {
	QuarterID int64 `json:"id" path:"id" required:"true" doc:"Quarter id"`
}
