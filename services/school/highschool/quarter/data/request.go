package data

type QuarterID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Quarter id" example:"1"`
}

type QuarterRequest struct {
	SchoolID    int64  `json:"schoolID" required:"true" doc:"School id" example:"1"`
	Name        string `json:"name" required:"true" doc:"Quarter name" example:"Sciences"`
	Description string `json:"description" required:"false" doc:"Quarter description" example:""`
}

type QuarterSequenceRequest struct {
	QuarterID  int64 `json:"quarterID" required:"true" doc:"Quarter id" example:"1"`
	SequenceID int64 `json:"sequenceID" required:"true" doc:"Sequence id" example:"1"`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
}

type GetAllQuarterSequenceRequest struct {
	QuarterID int64 `json:"id" path:"id" required:"true" doc:"Quarter id" example:"1"`
}
