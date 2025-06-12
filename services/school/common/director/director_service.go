package director

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"api/common/constants"
	"api/common/types"
	"api/config"
	"api/services/school/common/director/data"
	"api/services/school/common/director/model"
	"api/services/user/role"
	"api/services/user/user"
	modelUser "api/services/user/user/model"
)

type Service struct {
	Repository  *Repository
	RoleService *role.Service
	UserService *user.Service
}

func NewService(repository *Repository, roleService *role.Service, userService *user.Service) *Service {
	return &Service{
		Repository:  repository,
		RoleService: roleService,
		UserService: userService,
	}
}

const MODEL_NAME = "director"
const DEFAULT_ERROR_MESSAGE = "interact with director model"

func (service *Service) Create(inputJwtToken *types.JwtToken, schoolID int64, user *modelUser.User) (result *model.Director, errCode int, err error) {
	// Get director role
	directorRole, errRole := service.RoleService.Repository.GetByName(config.Env.RoleDirector)
	if errRole != nil || directorRole == nil || directorRole.ID < 1 {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get director role")
		return
	}

	// Create user
	var password string = ""
	if user != nil && user.Info != nil {
		var firstName, lastName string
		var birthYear = time.Now().Year()
		firstNameParts := strings.Split(user.Info.FirstName, " ")
		if len(firstNameParts) > 0 {
			firstName = firstNameParts[0]
		}
		lastNameParts := strings.Split(user.Info.LastName, " ")
		if len(lastNameParts) > 0 {
			lastName = lastNameParts[0]
		}
		if user.Info.Birthday != nil {
			birthYear = user.Info.Birthday.Year()
		}
		if len(firstName) > 0 && len(lastName) > 0 && birthYear > 0 {
			password = fmt.Sprintf("%s%s%d", firstName, lastName, birthYear)
		}
	}
	user.RoleID = directorRole.ID
	createdUser, errCodeCreate, errCreate := service.UserService.Create(nil, user, &password)
	if errCreate != nil {
		errCode = errCodeCreate
		err = errCreate
		return
	}
	if createdUser == nil || createdUser.ID < 1 {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Insert director
	result, err = service.Repository.Create(&model.Director{
		SchoolID: schoolID,
		UserID:   createdUser.ID,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) Update(inputJwtToken *types.JwtToken, id int64, user *modelUser.User) (result *model.Director, errCode int, err error) {
	// Check if director exists
	foundItem, err := service.Repository.GetByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil || foundItem.ID < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Get director role
	directorRole, errRole := service.RoleService.Repository.GetByName(config.Env.RoleDirector)
	if errRole != nil || directorRole == nil || directorRole.ID < 1 {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get director role")
		return
	}

	// Update user
	newUser := *user
	newUser.RoleID = directorRole.ID
	_, errCodeUser, errUser := service.UserService.Update(nil, foundItem.UserID, &newUser)
	if errUser != nil {
		errCode = errCodeUser
		err = errUser
		return
	}
	return
}

func (service *Service) Delete(inputJwtToken *types.JwtToken, id int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteByID(id)
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
	affectedRows, err = service.Repository.DeleteMultipleByID(list)
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

func (service *Service) Get(inputJwtToken *types.JwtToken, id int64) (result *model.Director, errCode int, err error) {
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

func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, request *data.GetAllRequest) (result []model.Director, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, request)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
