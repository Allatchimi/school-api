package student

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"api/common/constants"
	"api/common/types"
	"api/common/utils"
	"api/config"
	"api/services/school/common/student/model"
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

const MODEL_NAME = "student"
const DEFAULT_ERROR_MESSAGE = "interact with student model"

func (service *Service) Create(inputJwtToken *types.JwtToken, schoolID int64, uid string, user *modelUser.User) (result *model.Student, errCode int, err error) {
	// Get student role
	studentRole, errRole := service.RoleService.Repository.GetByName(config.Env.RoleStudent)
	if errRole != nil || studentRole == nil || studentRole.ID < 1 {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get student role")
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
	user.RoleID = studentRole.ID
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

	// Check the uid
	var newUID = uid
	if len(newUID) < 1 {
		prefix := fmt.Sprintf("%02dS", time.Now().Year())
		newUID = utils.GenerateRandomUID(4, prefix, "")
	}
	// Insert student
	result, err = service.Repository.Create(&model.Student{
		SchoolID: schoolID,
		UserID:   createdUser.ID,
		UID:      newUID,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) CreateStudentEnroll(inputJwtToken *types.JwtToken, item *model.StudentEnroll) (result *model.StudentEnroll, errCode int, err error) {
	// Check unique
	foundUnique, err := service.Repository.GetStudentEnrollUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreStudentEnrollSameUniqueObjects(foundUnique, item) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Insert student class/level domain
	result, err = service.Repository.CreateStudentEnroll(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) Update(inputJwtToken *types.JwtToken, id int64, uid string, user *modelUser.User) (result *model.Student, errCode int, err error) {
	// Check if student exists
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

	// Check the uid
	var newUID = uid
	if len(newUID) < 1 {
		prefix := fmt.Sprintf("%02dS", time.Now().Year())
		newUID = utils.GenerateRandomUID(4, prefix, "")
	}
	// Check unique by uid
	foundUnique, err := service.Repository.GetUniqueObjectByUID(&model.Student{
		SchoolID: foundItem.SchoolID,
		UID:      newUID,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreSameUniqueObjectsByUID(
		foundUnique,
		&model.Student{
			SchoolID: foundItem.SchoolID,
			UID:      newUID,
		},
	) && !service.Repository.AreSameUniqueObjectsByUID(
		foundUnique,
		foundItem,
	) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}

	// Get student role
	studentRole, errRole := service.RoleService.Repository.GetByName(config.Env.RoleStudent)
	if errRole != nil || studentRole == nil || studentRole.ID < 1 {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get student role")
		return
	}

	// Update user
	newUser := *user
	newUser.RoleID = studentRole.ID
	_, errCodeUser, errUser := service.UserService.Update(nil, foundItem.UserID, &newUser)
	if errUser != nil {
		errCode = errCodeUser
		err = errUser
		return
	}

	// Update student
	result, err = service.Repository.UpdateByID(id, &model.Student{
		SchoolID: foundItem.SchoolID,
		UserID:   foundItem.UserID,
		UID:      newUID,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) UpdateStudentEnroll(inputJwtToken *types.JwtToken, id int64, item *model.StudentEnroll) (result *model.StudentEnroll, errCode int, err error) {
	// Check if student class/level domain exists
	foundItem, err := service.Repository.GetStudentEnrollByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem == nil || foundItem.ID != id {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Check unique
	foundUnique, err := service.Repository.GetStudentEnrollUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreStudentEnrollSameUniqueObjects(foundUnique, item) && !service.Repository.AreStudentEnrollSameUniqueObjects(foundUnique, foundItem) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Update student
	result, err = service.Repository.UpdateStudentEnrollByID(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
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

func (service *Service) DeleteStudentEnroll(inputJwtToken *types.JwtToken, id int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteStudentEnrollByID(id)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) Get(inputJwtToken *types.JwtToken, id int64) (result *model.Student, errCode int, err error) {
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

func (service *Service) GetStudentEnroll(inputJwtToken *types.JwtToken, id int64) (result *model.StudentEnroll, errCode int, err error) {
	result, err = service.Repository.GetStudentEnrollByID(id)
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

func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, schoolID int64) (result []model.Student, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, schoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}

func (service *Service) GetAllStudentEnroll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, schoolID int64, studentID int64) (result []model.StudentEnroll, errCode int, err error) {
	result, err = service.Repository.GetAllStudentEnroll(filter, pagination, schoolID, studentID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
