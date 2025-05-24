package data

type ApiCreateRoomRequest struct {
	RoomID   string           `json:"room_id" required:"false" doc:"ID"`
	Metadata *ApiRoomMetadata `json:"metadata" required:"false" doc:"Metadata"`
}

type ApiRoomMetadata struct {
	RoomTitle           string                  `json:"room_title" required:"false" doc:"Room title"`
	WelcomeMessage      string                  `json:"welcome_message" required:"false" doc:"Welcome message"`
	WebhookUrl          string                  `json:"webhook_url" required:"false" doc:"Webhook url"`
	LogoutUrl           string                  `json:"logout_url" required:"false" doc:"Logout url"`
	RoomFeatures        *ApiRoomFeatures        `json:"room_features" required:"false" doc:"Room features"`
	DefaultLockSettings *ApiDefaultLockSettings `json:"default_lock_settings" required:"false" doc:"Default lock settings"`
}

type ApiRoomFeatures struct {
	AllowWebcams                    bool                                `json:"allow_webcams" required:"false" doc:"Allow webcams"`
	MuteOnStart                     bool                                `json:"mute_on_start" required:"false" doc:"Mute on start"`
	AllowScreenShare                bool                                `json:"allow_screen_share" required:"false" doc:"Allow screen share"`
	AllowRtmp                       bool                                `json:"allow_rtmp" required:"false" doc:"Allow RTMP"`
	AdminOnlyWebcams                bool                                `json:"admin_only_webcams" required:"false" doc:"Admin only webcams"`
	AllowViewOtherWebcams           bool                                `json:"allow_view_other_webcams" required:"false" doc:"Allow view other webcams"`
	AllowViewOtherUsersList         bool                                `json:"allow_view_other_users_list" required:"false" doc:"Allow view other users list"`
	AllowPolls                      bool                                `json:"allow_polls" required:"false" doc:"Allow polls"`
	EnableAnalytics                 bool                                `json:"enable_analytics" required:"false" doc:"Enable analytics"`
	AllowVirtualBg                  bool                                `json:"allow_virtual_bg" required:"false" doc:"Allow virtual background"`
	AllowRaiseHand                  bool                                `json:"allow_raise_hand" required:"false" doc:"Allow raise hand"`
	AutoGenUserID                   bool                                `json:"auto_gen_user_id" required:"false" doc:"Auto generate user ID"`
	RoomDuration                    int64                               `json:"room_duration" required:"false" doc:"Room duration"`
	RecordingFeatures               *ApiRecordingFeatures               `json:"recording_features" required:"false" doc:"Recording features"`
	ChatFeatures                    *ApiChatFeatures                    `json:"chat_features" required:"false" doc:"Chat features"`
	SharedNotePadFeatures           *ApiSharedNotePadFeatures           `json:"shared_note_pad_features" required:"false" doc:"Shared notepad features"`
	WhiteboardFeatures              *ApiWhiteboardFeatures              `json:"whiteboard_features" required:"false" doc:"Whiteboard features"`
	ExternalMediaPlayerFeatures     *ApiExternalMediaPlayerFeatures     `json:"external_media_player_features" required:"false" doc:"External media player features"`
	WaitingRoomFeatures             *ApiWaitingRoomFeatures             `json:"waiting_room_features" required:"false" doc:"Waiting room features"`
	BreakoutRoomFeatures            *ApiBreakoutRoomFeatures            `json:"breakout_room_features" required:"false" doc:"Breakout room features"`
	DisplayExternalLinkFeatures     *ApiDisplayExternalLinkFeatures     `json:"display_external_link_features" required:"false" doc:"Display external link features"`
	IngressFeatures                 *ApiIngressFeatures                 `json:"ingress_features" required:"false" doc:"Ingress features"`
	SpeechToTextTranslationFeatures *ApiSpeechToTextTranslationFeatures `json:"speech_to_text_translation_features" required:"false" doc:"Speech to text translation features"`
	EndToEndEncryptionFeatures      *ApiEndToEndEncryptionFeatures      `json:"end_to_end_encryption_features" required:"false" doc:"End to end encryption features"`
}

type ApiRecordingFeatures struct {
	IsAllow                  bool `json:"is_allow" required:"false" doc:"Is allowed"`
	IsAllowCloud             bool `json:"is_allow_cloud" required:"false" doc:"Is allowed in cloud"`
	IsAllowLocal             bool `json:"is_allow_local" required:"false" doc:"Is allowed locally"`
	EnableAutoCloudRecording bool `json:"enable_auto_cloud_recording" required:"false" doc:"Enable auto cloud recording"`
}

type ApiChatFeatures struct {
	AllowChat       bool `json:"allow_chat" required:"false" doc:"Allow chat"`
	AllowFileUpload bool `json:"allow_file_upload" required:"false" doc:"Allow file upload"`
}

type ApiSharedNotePadFeatures struct {
	AllowedSharedNotePad bool `json:"allowed_shared_note_pad" required:"false" doc:"Allowed shared notepad"`
}

type ApiWhiteboardFeatures struct {
	AllowedWhiteboard bool `json:"allowed_whiteboard" required:"false" doc:"Allowed whiteboard"`
}

type ApiExternalMediaPlayerFeatures struct {
	AllowedExternalMediaPlayer bool `json:"allowed_external_media_player" required:"false" doc:"Allowed external media player"`
}

type ApiWaitingRoomFeatures struct {
	IsActive bool `json:"is_active" required:"false" doc:"Is active"`
}

type ApiBreakoutRoomFeatures struct {
	IsAllow            bool  `json:"is_allow" required:"false" doc:"Is allowed"`
	AllowedNumberRooms int64 `json:"allowed_number_rooms" required:"false" doc:"Allowed number of rooms"`
}

type ApiDisplayExternalLinkFeatures struct {
	IsAllow bool `json:"is_allow" required:"false" doc:"Is allowed"`
}

type ApiIngressFeatures struct {
	IsAllow bool `json:"is_allow" required:"false" doc:"Is allowed"`
}

type ApiSpeechToTextTranslationFeatures struct {
	IsAllow            bool `json:"is_allow" required:"false" doc:"Is allowed"`
	IsAllowTranslation bool `json:"is_allow_translation" required:"false" doc:"Is allowed translation"`
}

type ApiEndToEndEncryptionFeatures struct {
	IsEnabled bool `json:"is_enabled" required:"false" doc:"Is enabled"`
}

type ApiDefaultLockSettings struct {
	LockMicrophone      bool `json:"lock_microphone" required:"false" doc:"Lock microphone"`
	LockWebcam          bool `json:"lock_webcam" required:"false" doc:"Lock webcam"`
	LockScreenSharing   bool `json:"lock_screen_sharing" required:"false" doc:"Lock screen sharing"`
	LockWhiteboard      bool `json:"lock_whiteboard" required:"false" doc:"Lock whiteboard"`
	LockSharedNotepad   bool `json:"lock_shared_notepad" required:"false" doc:"Lock shared notepad"`
	LockChat            bool `json:"lock_chat" required:"false" doc:"Lock chat"`
	LockChatSendMessage bool `json:"lock_chat_send_message" required:"false" doc:"Lock chat send message"`
	LockChatFileShare   bool `json:"lock_chat_file_share" required:"false" doc:"Lock chat file share"`
	LockPrivateChat     bool `json:"lock_private_chat" required:"false" doc:"Lock private chat"`
}
