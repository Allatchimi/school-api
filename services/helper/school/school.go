package serviceHelperSchool

import (
	"api/services/user/user"
	dataUser "api/services/user/user/data"
	"api/services/user/user/model"
)

var UserService *user.Service

func InjectServices(
	userService *user.Service,
) {
	UserService = userService
}

func GetAllUserByFeature(feature string) (users []model.User, err error) {
	users, err = UserService.Repository.GetAll(
		nil, nil,
		&dataUser.GetAllRequest{
			Feature:                feature,
			RemoveAdminRestriction: true,
		},
	)
	return
}
