package data

type ApiCreateRoomResponse struct {
	Status   bool                 `json:"status" required:"false" doc:"Status"`
	Message  string               `json:"msg" required:"false" doc:"Message"`
	Token    string               `json:"token" required:"false" doc:"Token"`
	RoomInfo *ApiRoomInfoResponse `json:"room_info" required:"false" doc:"Room info"`
}

type ApiJoinRoomResponse struct {
	Status  bool   `json:"status" required:"false" doc:"Status"`
	Message string `json:"msg" required:"false" doc:"Message"`
	Token   string `json:"token" required:"false" doc:"Token"`
}

type ApiRoomInfoResponse struct {
	RoomID             string  `json:"room_id" required:"false" doc:"ID"`
	RoomTitle          string  `json:"room_title" required:"false" doc:"Title"`
	SID                string  `json:"sid" required:"false" doc:"Room SID"`
	JoinedParticipants int64   `json:"joined_participants" required:"false" doc:"Joined participants"`
	IsRunning          bool    `json:"is_running" required:"false" doc:"Is running"`
	IsRecording        bool    `json:"is_recording" required:"false" doc:"Is recording"`
	IsActiveRtmp       bool    `json:"is_active_rtmp" required:"false" doc:"Is active rtmp"`
	CreationTime       int64   `json:"creation_time" required:"false" doc:"Creation time"`
	Metadata           *string `json:"metadata" required:"false" doc:"Metadata"`
	WebhookUrl         string  `json:"webhook_url" required:"false" doc:"webhook url"`
}
