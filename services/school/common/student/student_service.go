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
	"api/services/school/common/student/data"
	"api/services/school/common/student/model"
	"api/services/user/role"
	"api/services/user/user"
	userData "api/services/user/user/data"
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
const studentUIDSeparator = "S"

func (service *Service) Create(inputJwtToken *types.JwtToken, request *data.StudentRequest) (result *model.Student, errCode int, err error) {
	// Get student role
	studentRole, errRole := service.RoleService.Repository.GetByName(config.Env.RoleStudent)
	if errRole != nil || studentRole == nil || studentRole.ID < 1 {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get student role")
		return
	}

	// Format request
	if request == nil || request.Info == nil {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessage()
		return
	}
	var item = &userData.UserRequest{
		RoleID:      studentRole.ID,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		IsActivated: true,
		Info: &userData.UserInfoRequest{
			Gender:        request.Info.Gender,
			Username:      request.Info.Username,
			FirstName:     request.Info.FirstName,
			LastName:      request.Info.LastName,
			Birthday:      request.Info.Birthday,
			BirthLocation: request.Info.BirthLocation,
			Address:       request.Info.Address,
			Language:      request.Info.Language,
			Image:         request.Info.Image,
		},
	}

	// Generate password
	var firstName, lastName string
	var birthYear = time.Now().Year()
	firstNameParts := strings.Split(request.Info.FirstName, " ")
	if len(firstNameParts) > 0 {
		firstName = firstNameParts[0]
	}
	lastNameParts := strings.Split(request.Info.LastName, " ")
	if len(lastNameParts) > 0 {
		lastName = lastNameParts[0]
	}
	if request.Info.Birthday != nil {
		birthYear = request.Info.Birthday.Year()
	}
	var password string = ""
	if len(firstName) > 0 && len(lastName) > 0 && birthYear > 0 {
		password = fmt.Sprintf("%s%s%d", firstName, lastName, birthYear)
	}

	// Create user
	createdUser, errCodeCreate, errCreate := service.UserService.Create(nil, item, &password)
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

	// Generate the uid if it is empty
	var newUID = request.UID
	if len(newUID) < 1 {
		prefix := fmt.Sprintf("%02dT", time.Now().Year())
		newUID = utils.GenerateRandomUID(4, prefix, "")
	}
	// Insert student
	result, err = service.Repository.Create(&model.Student{
		SchoolID: request.SchoolID,
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

func (service *Service) CreateStudentEnroll(inputJwtToken *types.JwtToken, request *data.StudentEnrollRequest) (result *model.StudentEnroll, errCode int, err error) {
	// Format request
	item := &model.StudentEnroll{
		StudentID:     request.StudentID,
		YearID:        request.YearID,
		ClassID:       request.ClassID,
		LevelDomainID: request.LevelDomainID,

		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,

		Message:       request.Message,
		Gender:        request.Gender,
		FirstName:     request.FirstName,
		LastName:      request.LastName,
		Birthday:      request.Birthday,
		BirthLocation: request.BirthLocation,

		Document1: request.Document1,
		Document2: request.Document2,
		Document3: request.Document3,
		Document4: request.Document4,
		Document5: request.Document5,
	}

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

func (service *Service) Update(inputJwtToken *types.JwtToken, id int64, request *data.StudentRequest) (result *model.Student, errCode int, err error) {
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
	var newUID = request.UID
	if len(newUID) < 1 {
		prefix := fmt.Sprintf("%02d%s", time.Now().Year(), studentUIDSeparator)
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
	userRequest := userData.UserRequest{
		RoleID:      studentRole.ID,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		IsActivated: true,
		Info: &userData.UserInfoRequest{
			Gender:        request.Info.Gender,
			Username:      request.Info.Username,
			FirstName:     request.Info.FirstName,
			LastName:      request.Info.LastName,
			Birthday:      request.Info.Birthday,
			BirthLocation: request.Info.BirthLocation,
			Address:       request.Info.Address,
			Language:      request.Info.Language,
			Image:         request.Info.Image,
		},
	}
	_, errCodeUser, errUser := service.UserService.Update(nil, foundItem.UserID, &userRequest)
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

func (service *Service) UpdateStudentEnroll(inputJwtToken *types.JwtToken, id int64, request *data.StudentEnrollRequest) (result *model.StudentEnroll, errCode int, err error) {
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

	// Format request
	item := &model.StudentEnroll{
		StudentID:     request.StudentID,
		YearID:        request.YearID,
		ClassID:       request.ClassID,
		LevelDomainID: request.LevelDomainID,

		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,

		Message:       request.Message,
		Gender:        request.Gender,
		FirstName:     request.FirstName,
		LastName:      request.LastName,
		Birthday:      request.Birthday,
		BirthLocation: request.BirthLocation,

		Document1: request.Document1,
		Document2: request.Document2,
		Document3: request.Document3,
		Document4: request.Document4,
		Document5: request.Document5,
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
