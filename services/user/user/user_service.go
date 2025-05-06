package user

import (
	"fmt"
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

// Create Creates a new user and returns the newly created
func (service *Service) Create(inputJwtToken *types.JwtToken, item *model.User) (result *model.User, errCode int, err error) {
	// Check if user exists
	var foundItem *model.User
	var errMsg string = ""
	var isEmailValid = utils.IsEmailValid(item.Email)
	if isEmailValid {
		errMsg = "email"
		foundItem, err = service.Repository.GetByEmail(item.Email)
	} else {
		errMsg = "phone number"
		foundItem, err = service.Repository.GetByPhoneNumber(item.PhoneNumber)
	}
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(
			fmt.Sprintf("get user by %s from database", errMsg),
		)
		return
	}
	if foundItem != nil {
		if (isEmailValid && foundItem.Email == item.Email) || (!isEmailValid && foundItem.PhoneNumber == item.PhoneNumber) {
			errCode = http.StatusFound
			err = constants.Http302ErrorMessage(
				fmt.Sprintf("user %s", errMsg),
			)
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
		err = constants.Http500ErrorMessage("create user from database")
		return
	}
	return
}

// CreateUserRole assigns role to user and returns the updated user version
func (service *Service) AssignRole(inputJwtToken *types.JwtToken, userID int64, roleID int64) (result *model.User, errCode int, err error) {
	result, err = service.Repository.AssignRole(userID, roleID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("assign role to user from database")
		return
	}
	return
}

// UpdateUser Updates user and returns the updated version
func (service *Service) Update(inputJwtToken *types.JwtToken, id int64, item *model.User) (result *model.User, errCode int, err error) {
	// Check if user exists
	var errMsg string = ""
	errMsg = "email"
	foundItem, err := service.Repository.GetByEmail(item.Email)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(
			fmt.Sprintf("get user by %s from database", errMsg),
		)
		return
	}
	if foundItem != nil {
		if foundItem.Email != item.Email {
			errCode = http.StatusFound
			err = constants.Http302ErrorMessage(
				fmt.Sprintf("user %s", errMsg),
			)
			return
		}
	}

	errMsg = "phone number"
	foundItem, err = service.Repository.GetByPhoneNumber(item.PhoneNumber)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(
			fmt.Sprintf("get user by %s from database", errMsg),
		)
		return
	}
	if foundItem != nil {
		if foundItem.PhoneNumber != item.PhoneNumber {
			errCode = http.StatusFound
			err = constants.Http302ErrorMessage(
				fmt.Sprintf("user %s", errMsg),
			)
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
		err = constants.Http500ErrorMessage("update user from database")
	}
	return
}

// Delete Deletes user with matching id and returns affected rows
func (service *Service) Delete(inputJwtToken *types.JwtToken, id int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete user from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("user")
		return
	}
	return
}

// DeleteUserRole Deletes user role and returns affected rows
func (service *Service) DeleteRole(inputJwtToken *types.JwtToken, userID int64, roleID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteRole(userID, roleID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete user role from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("User role")
		return
	}
	return
}

// Delete Deletes selection with matching ids and returns the affected rows
func (service *Service) DeleteMultiple(inputJwtToken *types.JwtToken, list []int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteMultiple(list)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete multiple user from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("user selection")
		return
	}
	return
}

// Get Returns user with matching id
func (service *Service) Get(inputJwtToken *types.JwtToken, id int64) (result *model.User, errCode int, err error) {
	result, err = service.Repository.GetByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get user from database")
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("user")
		return
	}
	return
}

// GetAll Returns user list with matching search, filter and pagination
func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, roleName string) (result []model.User, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, roleName)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get user list from database")
	}
	return
}
