package data

import (
	"api/common/types"
	dataUser "api/services/user/user/data"
	"time"
)

type TelegramResponse struct {
	types.BaseGormModelResponse
	User    *dataUser.UserPublicResponse `json:"user" required:"false" doc:"User"`
	Title   string                       `json:"title" required:"false" doc:"Title"`
	Message string                       `json:"message" required:"false" doc:"Message"`
	Seen    bool                         `json:"isReaded" required:"false" doc:"Seen"`
	SeenAt  *time.Time                   `json:"seenAt" required:"false" doc:"Seen at"`
}

type TelegramNotSeenResponse struct {
	Count int64 `json:"count" required:"false" doc:"Count"`
}

type TelegramResponseList struct {
	types.PaginatedResponse
	Data []TelegramResponse `json:"data" required:"false" doc:"List of notifications" example:"[]"`
}
