package meeting

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/common/meeting/data"
	"api/services/school/common/meeting/model"
	"api/services/school/common/teacher"
	"api/services/user/user"
)

type Service struct {
	Repository     *Repository
	UserService    *user.Service
	TeacherService *teacher.Service
}

const MODEL_NAME = "meeting"
const DEFAULT_ERROR_MESSAGE = "interact with meeting service"

func NewService(repository *Repository, userService *user.Service, teacherService *teacher.Service) *Service {
	return &Service{
		Repository:     repository,
		UserService:    userService,
		TeacherService: teacherService,
	}
}

func (service *Service) Create(inputJwtToken *types.JwtToken, request *data.MeetingRoomRequest) (result *model.MeetingRoom, errCode int, err error) {
	// Format request
	item := &model.MeetingRoom{
		SchoolID:       request.SchoolID,
		ClassSubjectID: request.ClassSubjectID,
		UnitID:         request.UnitID,
	}

	// Check unique
	foundUnique, err := service.Repository.GetUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreSameUniqueObjects(foundUnique, item) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}

	// Call external meeting API to create a new room
	apiResp, err := service.Repository.ApiCreateRoom()
	if err != nil || apiResp == nil || apiResp.RoomInfo == nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	// Update the room id
	item.ApiRoomID = apiResp.RoomInfo.RoomID

	// Insert
	result, err = service.Repository.Create(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) Delete(inputJwtToken *types.JwtToken, id int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	return
}

func (service *Service) DeleteMultiple(inputJwtToken *types.JwtToken, list []int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultiple(list)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	return
}

func (service *Service) Get(inputJwtToken *types.JwtToken, id int64) (result *model.MeetingRoom, errCode int, err error) {
	result, err = service.Repository.GetByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	return
}

func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, request *data.GetAllRequest) (result []model.MeetingRoom, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, request)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}

func (service *Service) Join(inputJwtToken *types.JwtToken, id int64) (result string, errCode int, err error) {
	// Check if the meeting room exists
	meetingRoom, errCode, err := service.Get(inputJwtToken, id)
	if err != nil || meetingRoom == nil || meetingRoom.ID <= 0 || meetingRoom.ID != id {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Get the user
	user, err := service.UserService.Repository.GetByID(inputJwtToken.UserID)
	if err != nil || user == nil || user.ID <= 0 {
		errCode = http.StatusForbidden
		err = constants.Http403InvalidPermissionErrorMessage()
		return
	}

	// Get the teacher and check if it's the teacher for this room(room is associated to unit/class subject)
	var isAdmin bool = false
	teacher, _ := service.TeacherService.Repository.GetByUserID(inputJwtToken.UserID)
	if teacher != nil && teacher.ID > 0 {
		if meetingRoom.School.Type == constants.SCHOOL_TYPE_HIGHSCHOOL {
			teacherClassSubject, _ := service.TeacherService.Repository.GetTeacherClassSubjectUnitByUserIDClassSubjectID(teacher.ID, meetingRoom.ClassSubjectID)
			if teacherClassSubject != nil && teacherClassSubject.ID > 0 {
				isAdmin = true
			}
		} else {
			teacherUnit, _ := service.TeacherService.Repository.GetTeacherClassSubjectUnitByUserIDUnitID(teacher.ID, meetingRoom.UnitID)
			if teacherUnit != nil && teacherUnit.ID > 0 {
				isAdmin = true
			}
		}
	}

	// Call external meeting API to get join token
	apiResp, err := service.Repository.ApiJoinRoom(meetingRoom.ApiRoomID, user, isAdmin)
	if err != nil || apiResp == nil || !apiResp.Status {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	result = apiResp.Token
	return
}
