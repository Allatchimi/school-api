package data

type NotificationID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Notification id" example:"1"`
}

type GetAllRequest struct {
}
