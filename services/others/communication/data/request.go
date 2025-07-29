package data

type CommunicationID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Communication id"`
}

type CommunicationRequest struct {
	SchoolID int64   `json:"schoolID" required:"false" doc:"School id"`
	RoleIDs  []int64 `json:"roleIDs" required:"true" doc:"Role id list"`

	Subject string `json:"subject" required:"true" doc:"Subject"`
	Message string `json:"message" required:"true" doc:"Message"`
}

type GetAllRequest struct {
	SchoolID int64 `json:"schoolID" query:"schoolID" required:"false" doc:"School id"`
}
