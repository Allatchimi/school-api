package data

import (
	"api/common/types"
	dataUser "api/services/user/user/data"
	"time"
)

type NotificationResponse struct {
	types.BaseGormModelResponse
	User *dataUser.UserPublicResponse `json:"user" required:"false" doc:"User"`

	Title   string     `json:"title" required:"false" doc:"Title"`
	Message string     `json:"message" required:"false" doc:"Message"`
	Seen    bool       `json:"seen" required:"false" doc:"Seen"`
	SeenAt  *time.Time `json:"seenAt" required:"false" doc:"Seen at"`
}

type NotificationNotSeenResponse struct {
	Count int64 `json:"count" required:"false" doc:"Count"`
}

type NotificationResponseList struct {
	types.PaginatedResponse
	Data []NotificationResponse `json:"data" required:"false" doc:"List of notification"`
}
