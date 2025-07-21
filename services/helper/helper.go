package serviceHelper

import (
	"api/services/user/user"
	"api/services/user/user/model"
)

var UserService *user.Service

func InjectServices(
	userService *user.Service,
) {
	UserService = userService
}

func GetUserByID(userID int64) (result *model.User, err error) {
	result, err = UserService.Repository.GetByID(userID)
	return
}
