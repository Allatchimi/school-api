package data

import (
	"api/common/types"
	dataUser "api/services/user/user/data"
)

type NotificationResponse struct {
	types.BaseGormModelResponse
	User    *dataUser.UserPublicResponse `json:"user" required:"false" doc:"User"`
	Title   string                       `json:"title" required:"true" doc:"Title"`
	Message string                       `json:"message" required:"true" doc:"Message"`
}

type NotificationResponseList struct {
	types.PaginatedResponse
	Data []NotificationResponse `json:"data" required:"false" doc:"List of notifications" example:"[]"`
}
