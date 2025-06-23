package data

type ContactID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Contact id"`
}

type ContactRequest struct {
	SchoolID int64 `json:"schoolID" required:"false" doc:"School id"`

	Subject string `json:"subject" required:"true" doc:"Subject"`
	Email   string `json:"email" required:"true" doc:"Email"`
	Message string `json:"message" required:"true" doc:"Message"`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
}
