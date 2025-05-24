package meeting

import (
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/common/utils/meeting"
	"api/config"
	"api/services/school/common/meeting/data"
	"api/services/school/common/meeting/model"
	modelUser "api/services/user/user/model"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}

func (repository *Repository) Create(item *model.MeetingRoom) (*model.MeetingRoom, error) {
	result := *item
	return &result, repository.Db.Preload(clause.Associations).Create(&result).Error
}

func (repository *Repository) ApiCreateRoom() (*data.ApiCreateRoomResponse, error) {
	roomID := fmt.Sprintf("%d_%s", time.Now().Unix(), utils.GenerateRandomAlphaNumeric(10))
	room := &data.ApiCreateRoomRequest{
		RoomID: roomID,
		Metadata: &data.ApiRoomMetadata{
			RoomTitle:      "Classroom",
			WelcomeMessage: "",
			RoomFeatures: &data.ApiRoomFeatures{
				AllowWebcams:            true,
				MuteOnStart:             true,
				AllowScreenShare:        true,
				AllowRtmp:               true,
				AdminOnlyWebcams:        false,
				AllowViewOtherWebcams:   true,
				AllowViewOtherUsersList: true,
				AllowPolls:              true,
				EnableAnalytics:         false,
				AllowVirtualBg:          true,
				AllowRaiseHand:          true,
				AutoGenUserID:           false,
				RoomDuration:            0,
				RecordingFeatures: &data.ApiRecordingFeatures{
					IsAllow:                  true,
					IsAllowCloud:             true,
					IsAllowLocal:             true,
					EnableAutoCloudRecording: false,
				},
				ChatFeatures: &data.ApiChatFeatures{
					AllowChat:       true,
					AllowFileUpload: true,
				},
				SharedNotePadFeatures: &data.ApiSharedNotePadFeatures{
					AllowedSharedNotePad: true,
				},
				WhiteboardFeatures: &data.ApiWhiteboardFeatures{
					AllowedWhiteboard: true,
				},
				ExternalMediaPlayerFeatures: &data.ApiExternalMediaPlayerFeatures{
					AllowedExternalMediaPlayer: true,
				},
				WaitingRoomFeatures: &data.ApiWaitingRoomFeatures{
					IsActive: false,
				},
				BreakoutRoomFeatures: &data.ApiBreakoutRoomFeatures{
					IsAllow:            true,
					AllowedNumberRooms: 2,
				},
				DisplayExternalLinkFeatures: &data.ApiDisplayExternalLinkFeatures{
					IsAllow: true,
				},
				IngressFeatures: &data.ApiIngressFeatures{
					IsAllow: true,
				},
				SpeechToTextTranslationFeatures: &data.ApiSpeechToTextTranslationFeatures{
					IsAllow:            true,
					IsAllowTranslation: true,
				},
				EndToEndEncryptionFeatures: &data.ApiEndToEndEncryptionFeatures{
					IsEnabled: false,
				},
			},
			DefaultLockSettings: &data.ApiDefaultLockSettings{
				LockMicrophone:      false,
				LockWebcam:          false,
				LockScreenSharing:   false,
				LockWhiteboard:      false,
				LockSharedNotepad:   false,
				LockChat:            false,
				LockChatSendMessage: false,
				LockChatFileShare:   false,
				LockPrivateChat:     false,
			},
		},
	}
	var apiResp = &data.ApiCreateRoomResponse{}
	err := utils.HttpPost(
		fmt.Sprintf("%s/room/create", config.Env.MeetingApiUrl),
		"application/json",
		meeting.GetPlugnMeetPostHeaders(meeting.GetPlugnMeetSignature(fmt.Sprintf("{\"room_id\":\"%s\"}", roomID))),
		room,
		apiResp,
	)
	return apiResp, err
}

func (repository *Repository) Delete(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.MeetingRoom{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultiple(list []int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	tmpResult := repository.Db.Where(where).Delete(&model.MeetingRoom{})

	result = tmpResult.RowsAffected
	err = tmpResult.Error
	return
}

func (repository *Repository) GetById(id int64) (*model.MeetingRoom, error) {
	result := &model.MeetingRoom{}
	return result, repository.Db.Preload(clause.Associations).Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetUniqueObject(item *model.MeetingRoom) (*model.MeetingRoom, error) {
	result := &model.MeetingRoom{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.MeetingRoom{
		SchoolID:      item.SchoolID,
		LevelDomainID: item.LevelDomainID,
		ClassID:       item.ClassID,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjects(item1 *model.MeetingRoom, item2 *model.MeetingRoom) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.LevelDomainID == item2.LevelDomainID &&
			item1.ClassID == item2.ClassID) {
		return true
	}
	return false
}

func (repository *Repository) GetAll(filter *types.Filter, pagination *types.Pagination) (result []model.MeetingRoom, err error) {
	result = make([]model.MeetingRoom, 0)
	var where string = ""
	if filter != nil && len(filter.Search) >= 1 {
		where = fmt.Sprintf(
			"WHERE CAST(id AS TEXT) = '%s' OR name ILIKE '%s' OR start_date ILIKE '%s' OR end_date ILIKE '%s'",
			filter.Search,
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
		)
	}
	tmpErr := repository.Db.Preload(clause.Associations).Scopes(
		helpers.PaginationScope(
			repository.Db,
			"SELECT * FROM meetings",
			where,
			pagination,
			filter,
		),
	).Find(&result).Error

	err = tmpErr
	return
}

func (repository *Repository) ApiJoinRoom(roomID string, user *modelUser.User, isAdmin bool) (*data.ApiJoinRoomResponse, error) {
	apiResp := &data.ApiJoinRoomResponse{}
	var join = &data.ApiJoinRoomRequest{
		RoomID: roomID,
		UserInfo: &data.ApiUserInfo{
			Name:     user.Info.FirstName,
			UserID:   fmt.Sprintf("user_%d", user.ID),
			IsAdmin:  isAdmin,
			IsHidden: false,
			UserMetadata: &data.ApiUserMetadata{
				ProfilePic: "",
				LockSettings: &data.ApiUserLockSettings{
					LockMicrophone:      !isAdmin,
					LockWebcam:          !isAdmin,
					LockScreenSharing:   !isAdmin,
					LockChat:            false,
					LockChatSendMessage: false,
					LockChatFileShare:   false,
				},
			},
		},
	}
	err := utils.HttpPost(
		fmt.Sprintf("%s/room/getJoinToken", config.Env.MeetingApiUrl),
		"application/json",
		meeting.GetPlugnMeetPostHeaders(meeting.GetPlugnMeetSignature(fmt.Sprintf("{\"room_id\":\"%s\"}", roomID))),
		join,
		apiResp,
	)
	return apiResp, err
}
