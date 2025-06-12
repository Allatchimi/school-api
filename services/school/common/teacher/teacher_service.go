package teacher

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"api/common/constants"
	"api/common/types"
	"api/common/utils"
	"api/config"
	"api/services/school/common/teacher/data"
	"api/services/school/common/teacher/model"
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

const MODEL_NAME = "teacher"
const DEFAULT_ERROR_MESSAGE = "interact with teacher model"
const teacherUIDSeparator = "T"

func (service *Service) Create(inputJwtToken *types.JwtToken, request *data.TeacherRequest) (result *model.Teacher, errCode int, err error) {
	// Get teacher role
	teacherRole, errRole := service.RoleService.Repository.GetByName(config.Env.RoleTeacher)
	if errRole != nil || teacherRole == nil || teacherRole.ID < 1 {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get teacher role")
		return
	}

	// Format request
	if request == nil || request.Info == nil {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessage()
		return
	}
	var item = &userData.UserRequest{
		RoleID:      teacherRole.ID,
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
	// Insert teacher
	result, err = service.Repository.Create(&model.Teacher{
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

func (service *Service) CreateTeacherClassSubjectUnit(inputJwtToken *types.JwtToken, request *data.TeacherClassSubjectUnitRequest) (result *model.TeacherClassSubjectUnit, errCode int, err error) {
	// Retrieve item
	item := &model.TeacherClassSubjectUnit{
		TeacherID:      request.TeacherID,
		YearID:         request.YearID,
		ClassSubjectID: request.ClassSubjectID,
		UnitID:         request.UnitID,
	}

	// Check unique
	foundUnique, err := service.Repository.GetTeacherClassSubjectUnitUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreTeacherClassSubjectUnitSameUniqueObjects(foundUnique, item) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Insert teacher unit/subject
	result, err = service.Repository.CreateTeacherClassSubjectUnit(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) Update(inputJwtToken *types.JwtToken, id int64, request *data.TeacherRequest) (result *model.Teacher, errCode int, err error) {
	// Check if teacher exists
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
		prefix := fmt.Sprintf("%02d%s", time.Now().Year(), teacherUIDSeparator)
		newUID = utils.GenerateRandomUID(4, prefix, "")
	}
	// Check unique by uid
	foundUnique, err := service.Repository.GetUniqueObjectByUID(&model.Teacher{
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
		&model.Teacher{
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

	// Get teacher role
	teacherRole, errRole := service.RoleService.Repository.GetByName(config.Env.RoleTeacher)
	if errRole != nil || teacherRole == nil || teacherRole.ID < 1 {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get teacher role")
		return
	}

	// Update user
	userRequest := userData.UserRequest{
		RoleID:      teacherRole.ID,
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

	// Update teacher
	result, err = service.Repository.UpdateByID(id, &model.Teacher{
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

func (service *Service) UpdateTeacherClassSubjectUnit(inputJwtToken *types.JwtToken, id int64, request *data.TeacherClassSubjectUnitRequest) (result *model.TeacherClassSubjectUnit, errCode int, err error) {
	// Check if teacher unit/subject exists
	foundItem, err := service.Repository.GetTeacherClassSubjectUnitByID(id)
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

	// Retrieve item
	item := &model.TeacherClassSubjectUnit{
		TeacherID:      request.TeacherID,
		YearID:         request.YearID,
		ClassSubjectID: request.ClassSubjectID,
		UnitID:         request.UnitID,
	}

	// Check unique
	foundUnique, err := service.Repository.GetTeacherClassSubjectUnitUniqueObject(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if service.Repository.AreTeacherClassSubjectUnitSameUniqueObjects(foundUnique, item) && !service.Repository.AreTeacherClassSubjectUnitSameUniqueObjects(foundUnique, foundItem) {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Update teacher class subject/unit
	result, err = service.Repository.UpdateTeacherClassSubjectUnitByID(id, item)
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

func (service *Service) DeleteTeacherClassSubjectUnit(inputJwtToken *types.JwtToken, id int64) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.DeleteTeacherClassSubjectUnitByID(id)
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

func (service *Service) Get(inputJwtToken *types.JwtToken, id int64) (result *model.Teacher, errCode int, err error) {
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

func (service *Service) GetTeacherClassSubjectUnit(inputJwtToken *types.JwtToken, id int64) (result *model.TeacherClassSubjectUnit, errCode int, err error) {
	result, err = service.Repository.GetTeacherClassSubjectUnitByID(id)
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

func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, schoolID int64) (result []model.Teacher, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, schoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}

func (service *Service) GetAllTeacherClassSubjectUnit(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, schoolID int64, teacherID int64) (result []model.TeacherClassSubjectUnit, errCode int, err error) {
	result, err = service.Repository.GetAllTeacherClassSubjectUnit(filter, pagination, schoolID, teacherID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
