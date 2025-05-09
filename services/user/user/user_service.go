package user

import (
	"net/http"
	"time"

	"api/common/constants"
	"api/common/types"
	"api/common/utils"
	"api/services/user/user/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

const MODEL_NAME = "user"
const DEFAULT_ERROR_MESSAGE = "interact with user model"

func (service *Service) Create(inputJwtToken *types.JwtToken, item *model.User) (result *model.User, errCode int, err error) {
	// Check if user exists
	var foundItem *model.User
	var isEmailValid = utils.IsEmailValid(item.Email)
	if isEmailValid {
		foundItem, err = service.Repository.GetByEmail(item.Email)
	} else {
		foundItem, err = service.Repository.GetByPhoneNumber(item.PhoneNumber)
	}
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem != nil {
		if (isEmailValid && foundItem.Email == item.Email) || (!isEmailValid && foundItem.PhoneNumber == item.PhoneNumber) {
			errCode = http.StatusFound
			err = constants.Http302ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
	}

	// Create new user
	randomPassword := utils.GenerateRandomPassword(8)
	var activatedAt *time.Time = nil
	if item.IsActivated {
		tmpTime := time.Now()
		activatedAt = &tmpTime
	}
	newUser := &model.User{
		Email:       item.Email,
		PhoneNumber: item.PhoneNumber,
		RoleID:      item.RoleID,
		IsActivated: item.IsActivated,
		ActivatedAt: activatedAt,
		LoginMethod: constants.AuthLoginMethodDefault,
		Password:    randomPassword,
	}
	result, err = service.Repository.Create(newUser)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) AssignRole(inputJwtToken *types.JwtToken, userID int64, roleID int64) (result *model.User, errCode int, err error) {
	result, err = service.Repository.AssignRole(userID, roleID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) Update(inputJwtToken *types.JwtToken, id int64, item *model.User) (result *model.User, errCode int, err error) {
	// Check if user exists
	foundItem, err := service.Repository.GetByEmail(item.Email)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem != nil {
		if foundItem.Email != item.Email {
			errCode = http.StatusFound
			err = constants.Http302ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
	}

	foundItem, err = service.Repository.GetByPhoneNumber(item.PhoneNumber)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem != nil {
		if foundItem.PhoneNumber != item.PhoneNumber {
			errCode = http.StatusFound
			err = constants.Http302ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
	}

	// Update
	if !item.IsActivated {
		item.ActivatedAt = nil
	} else if item.IsActivated && !foundItem.IsActivated {
		tmpTime := time.Now()
		item.ActivatedAt = &tmpTime
	}
	result, err = service.Repository.Update(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
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

func (service *Service) DeleteRole(inputJwtToken *types.JwtToken, userID int64, roleID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteRole(userID, roleID)
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

func (service *Service) Get(inputJwtToken *types.JwtToken, id int64) (result *model.User, errCode int, err error) {
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

func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, roleName string) (result []model.User, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, roleName)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
