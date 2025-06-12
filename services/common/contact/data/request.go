package data

type ContactID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Contact id" example:"1"`
}

type ContactRequest struct {
	SchoolID int64 `json:"schoolID" required:"false" doc:"School id" example:"1"`

	Subject string `json:"subject" required:"true" doc:"Subject" example:""`
	Email   string `json:"email" required:"true" doc:"Email" example:""`
	Message string `json:"message" required:"true" doc:"Message" example:""`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id" example:"1"`
}
