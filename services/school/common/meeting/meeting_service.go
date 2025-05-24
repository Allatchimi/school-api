package meeting

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/school/common/meeting/model"
	"api/services/user/user"
)

type Service struct {
	Repository     *Repository
	UserRepository *user.Repository
}

const MODEL_NAME = "meeting"
const DEFAULT_ERROR_MESSAGE = "interact with meeting service"

func NewService(repository *Repository, userRepository *user.Repository) *Service {
	return &Service{
		Repository:     repository,
		UserRepository: userRepository,
	}
}

func (service *Service) Create(inputJwtToken *types.JwtToken, item *model.MeetingRoom) (result *model.MeetingRoom, errCode int, err error) {
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
	result, err = service.Repository.GetById(id)
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

func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination) (result []model.MeetingRoom, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}

func (service *Service) Join(inputJwtToken *types.JwtToken, id int64) (result string, errCode int, err error) {
	// Check if the meeting room exists
	meetingRoom, errCode, err := service.Get(inputJwtToken, id)
	if err != nil {
		return
	}

	// Get the teacher

	// Get the user

	// Call external meeting API to get join token
	apiResp, err := service.Repository.ApiJoinRoom(meetingRoom.ApiRoomID, nil)
	if err != nil || apiResp == nil || apiResp.Status == false {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	result = apiResp.Token
	return
}
