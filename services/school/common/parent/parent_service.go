package parent

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"api/common/constants"
	"api/common/types"
	"api/config"
	"api/services/school/common/parent/model"
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

const MODEL_NAME = "parent"
const DEFAULT_ERROR_MESSAGE = "interact with parent model"

func (service *Service) Create(inputJwtToken *types.JwtToken, schoolID int64, user *modelUser.User) (result *model.Parent, errCode int, err error) {
	// Get parent role
	parentRole, errRole := service.RoleService.Repository.GetByName(config.Env.RoleParent)
	if errRole != nil || parentRole == nil || parentRole.ID < 1 {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get parent role")
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
	user.RoleID = parentRole.ID
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

	// Insert parent
	result, err = service.Repository.Create(&model.Parent{
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

func (service *Service) CreateParentStudent(inputJwtToken *types.JwtToken, item *model.ParentStudent) (result *model.ParentStudent, errCode int, err error) {
	// Check if parent level/class already exists
	foundItem, err := service.Repository.GetParentStudentByObject(&model.ParentStudent{
		ParentID:  item.ParentID,
		StudentID: item.StudentID,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}

	// Insert parent level/class
	result, err = service.Repository.CreateParentStudent(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) Update(inputJwtToken *types.JwtToken, id int64, user *modelUser.User) (result *model.Parent, errCode int, err error) {
	// Check if parent exists
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

	// Get parent role
	parentRole, errRole := service.RoleService.Repository.GetByName(config.Env.RoleParent)
	if errRole != nil || parentRole == nil || parentRole.ID < 1 {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get parent role")
		return
	}

	// Update user
	newUser := *user
	newUser.RoleID = parentRole.ID
	_, errCodeUser, errUser := service.UserService.Update(nil, foundItem.UserID, &newUser)
	if errUser != nil {
		errCode = errCodeUser
		err = errUser
		return
	}
	return
}

func (service *Service) UpdateParentStudent(inputJwtToken *types.JwtToken, parentParentStudentID int64, item *model.ParentStudent) (result *model.ParentStudent, errCode int, err error) {
	// Check unique
	foundParentByID, err := service.Repository.GetParentStudentByID(parentParentStudentID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundParentByID == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}
	foundItem, err := service.Repository.GetParentStudentByObject(&model.ParentStudent{
		ParentID:  item.ParentID,
		StudentID: item.StudentID,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem != nil {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}

	// Update parent
	result, err = service.Repository.UpdateParentStudent(parentParentStudentID, item)
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

func (service *Service) DeleteParentStudent(inputJwtToken *types.JwtToken, parentParentStudentID int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteParentStudent(parentParentStudentID)
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

func (service *Service) Get(inputJwtToken *types.JwtToken, id int64) (result *model.Parent, errCode int, err error) {
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

func (service *Service) GetParentStudent(inputJwtToken *types.JwtToken, parentParentStudentID int64) (result *model.ParentStudent, errCode int, err error) {
	result, err = service.Repository.GetParentStudentByID(parentParentStudentID)
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

func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, schoolID int64) (result []model.Parent, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, schoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}

func (service *Service) GetAllParentStudent(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, schoolID int64) (result []model.ParentStudent, errCode int, err error) {
	result, err = service.Repository.GetAllParentStudent(filter, pagination, schoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
