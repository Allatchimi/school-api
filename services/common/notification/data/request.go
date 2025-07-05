package data

type NotificationID struct {
	ID int64 `json:"id" path:"id" required:"true" doc:"Notification id"`
}

type NotificationSeenRequest struct {
	Seen bool `json:"seen" required:"true" doc:"Seen"`
}

type NotificationSeenAllRequest struct {
	Seen bool `json:"seen" required:"true" doc:"Seen"`
}

type NotificationNotSeenCountRequest struct {
}

type GetAllRequest struct {
}
