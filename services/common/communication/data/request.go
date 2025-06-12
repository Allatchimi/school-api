package data

type CommunicationID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Communication id" example:"1"`
}

type CommunicationRequest struct {
	SchoolID int64 `json:"schoolID" required:"false" doc:"School id" example:"1"`

	Subject  string `json:"subject" required:"true" doc:"Subject" example:""`
	Message  string `json:"message" required:"true" doc:"Message" example:""`
	Audience string `json:"audience" required:"true" doc:"Audience" example:""`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
}
