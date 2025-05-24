package data

type ApiJoinRoomRequest struct {
	RoomID   string       `json:"room_id" required:"false" doc:"Room ID"`
	UserInfo *ApiUserInfo `json:"user_info" required:"false" doc:"User Information"`
}

type ApiUserInfo struct {
	Name         string           `json:"name" required:"false" doc:"User's name"`
	UserID       string           `json:"user_id" required:"false" doc:"Unique User ID"`
	IsAdmin      bool             `json:"is_admin" required:"false" doc:"Is Admin"`
	IsHidden     bool             `json:"is_hidden" required:"false" doc:"Is Hidden"`
	UserMetadata *ApiUserMetadata `json:"user_metadata" required:"false" doc:"User Metadata"`
}

type ApiUserMetadata struct {
	ProfilePic   string               `json:"profile_pic" required:"false" doc:"Profile Picture URL"`
	LockSettings *ApiUserLockSettings `json:"lock_settings" required:"false" doc:"User Lock Settings"`
}

type ApiUserLockSettings struct {
	LockMicrophone      bool `json:"lock_microphone" required:"false" doc:"Lock Microphone"`
	LockWebcam          bool `json:"lock_webcam" required:"false" doc:"Lock Webcam"`
	LockScreenSharing   bool `json:"lock_screen_sharing" required:"false" doc:"Lock Screen Sharing"`
	LockChat            bool `json:"lock_chat" required:"false" doc:"Lock Chat"`
	LockChatSendMessage bool `json:"lock_chat_send_message" required:"false" doc:"Lock Chat Send Message"`
	LockChatFileShare   bool `json:"lock_chat_file_share" required:"false" doc:"Lock Chat File Share"`
}
