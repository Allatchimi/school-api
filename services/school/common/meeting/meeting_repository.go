package meeting

import (
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"api/common/helpers"
	httpHelper "api/common/helpers/http"
	"api/common/types"
	"api/common/utils"
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
	apiResp := &data.ApiCreateRoomResponse{}
	url := fmt.Sprintf("%s/room/create", config.Env.MeetingApiUrl)
	headers := []httpHelper.HttpHeader{
		{Label: "Content-Type", Value: "application/json"},
	}
	err := httpHelper.HttpPost(
		url,
		headers,
		room,
		apiResp,
	)
	return apiResp, err
}

func (repository *Repository) DeleteByID(id int64) (int64, error) {
	result := repository.Db.Where("id = ?", id).Delete(&model.MeetingRoom{})
	return result.RowsAffected, result.Error
}

func (repository *Repository) DeleteMultipleByID(list []int64, schoolID int64) (result int64, err error) {
	where := fmt.Sprintf("id IN (%s)", utils.ListIntToString(list))
	var query *gorm.DB = repository.Db.Where(where)
	if schoolID > 0 {
		query = query.Where("school_id = ?", schoolID)
	}
	query = query.Delete(&model.MeetingRoom{})

	result = query.RowsAffected
	err = query.Error
	return
}

func (repository *Repository) GetByID(id int64) (*model.MeetingRoom, error) {
	result := &model.MeetingRoom{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).Limit(1).Find(result).Error
}

func (repository *Repository) GetByIDSchoolID(id int64, schoolID int64) (*model.MeetingRoom, error) {
	result := &model.MeetingRoom{}
	return result, repository.Db.Preload(clause.Associations).
		Where("id = ?", id).
		Where("school_id = ?", schoolID).
		Limit(1).Find(result).Error
}

func (repository *Repository) GetUniqueObject(item *model.MeetingRoom) (*model.MeetingRoom, error) {
	result := &model.MeetingRoom{}
	return result, repository.Db.Preload(clause.Associations).Where(&model.MeetingRoom{
		SchoolID:       item.SchoolID,
		UnitID:         item.UnitID,
		ClassSubjectID: item.ClassSubjectID,
	}).Limit(1).Find(result).Error
}

func (repository *Repository) AreSameUniqueObjects(item1 *model.MeetingRoom, item2 *model.MeetingRoom) bool {
	if item1 != nil && item2 != nil &&
		(item1.SchoolID == item2.SchoolID &&
			item1.ClassSubjectID == item2.ClassSubjectID &&
			item1.UnitID == item2.UnitID) {
		return true
	}
	return false
}

func (repository *Repository) GetAll(
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllRequest,
) (result []model.MeetingRoom, err error) {
	result = make([]model.MeetingRoom, 0)

	// Build secure WHERE conditions
	where := ""
	args := []any{}
	if request != nil {
		if request.SchoolID > 0 {
			where = helpers.AppendWhereClause(where, "meetings.school_id = ?")
			args = append(args, request.SchoolID)
		}
		if request.ClassID > 0 {
			where = helpers.AppendWhereClause(where, "highschool_class_subjects.class_id = ?")
			args = append(args, request.ClassID)
		}
		if request.LevelDomainID > 0 {
			where = helpers.AppendWhereClause(where, "university_units.level_domain_id = ?")
			args = append(args, request.LevelDomainID)
		}
	}

	// Handle search filter securely
	if filter != nil && len(filter.Search) > 0 {
		search := filter.Search
		like := "%" + search + "%"

		// Securely append search conditions
		searchClause := `(
			CAST(meetings.id AS TEXT) = ? OR
			meetings.api_room_id ILIKE ? OR
			schools.name ILIKE ? OR
			schools.type ILIKE ? OR
			highschool_classes.name ILIKE ? OR
			highschool_classes.description ILIKE ? OR
			highschool_subjects.name ILIKE ? OR
			highschool_subjects.description ILIKE ? OR
			university_units.name ILIKE ? OR
			university_units.description ILIKE ? OR
			university_levels.name ILIKE ? OR
			university_levels.description ILIKE ? OR
			university_domains.name ILIKE ? OR
			university_domains.description ILIKE ?
		)`

		where = helpers.AppendWhereClause(where, searchClause)
		args = append(args, search, like, like, like, like, like, like, like, like, like, like, like, like, like)
	}

	// Perform query with preloads and custom pagination scope
	err = repository.Db.
		Preload(clause.Associations).
		Preload("ClassSubject.Class").
		Preload("ClassSubject.Subject").
		Preload("Unit.Level").
		Preload("Unit.Domain").
		Preload("Unit.Domain.Department").
		Preload("Unit.Semester").
		Scopes(
			helpers.PaginationScopeV2(
				repository.Db,
				`SELECT meetings.*
				FROM meeting_rooms meetings
				LEFT JOIN schools ON meetings.school_id = schools.id
				LEFT JOIN highschool_class_subjects ON meetings.class_subject_id = highschool_class_subjects.id
				LEFT JOIN university_units ON meetings.unit_id = university_units.id
				LEFT JOIN highschool_classes ON highschool_class_subjects.class_id = highschool_classes.id
				LEFT JOIN highschool_subjects ON highschool_class_subjects.subject_id = highschool_subjects.id
				LEFT JOIN university_level_domains ON university_units.level_domain_id = university_level_domains.id
				LEFT JOIN university_levels ON university_level_domains.level_id = university_levels.id
				LEFT JOIN university_domains ON university_level_domains.domain_id = university_domains.id `,
				where,
				pagination,
				filter,
				args...,
			),
		).Find(&result).Error

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
	url := fmt.Sprintf("%s/room/getJoinToken", config.Env.MeetingApiUrl)
	headers := []httpHelper.HttpHeader{
		{Label: "Content-Type", Value: "application/json"},
	}
	err := httpHelper.HttpPost(
		url,
		headers,
		join,
		apiResp,
	)
	return apiResp, err
}
