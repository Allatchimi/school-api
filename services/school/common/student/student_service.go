package student

import (
	"fmt"
	"net/http"
	"time"

	"api/common/constants"
	"api/common/types"
	"api/common/utils"
	"api/common/utils/mail"
	"api/common/utils/password"
	"api/config"
	"api/services/school/common/school"
	"api/services/school/common/student/data"
	"api/services/school/common/student/model"
	"api/services/user/role"
	"api/services/user/user"
	dataUser "api/services/user/user/data"
)

type Service struct {
	Repository    *Repository
	RoleService   *role.Service
	UserService   *user.Service
	SchoolService *school.Service
}

func NewService(
	repository *Repository,
	roleService *role.Service,
	userService *user.Service,
	schoolService *school.Service,
) *Service {
	return &Service{
		Repository:    repository,
		RoleService:   roleService,
		UserService:   userService,
		SchoolService: schoolService,
	}
}

const MODEL_NAME = "student"
const DEFAULT_ERROR_MESSAGE = "interact with student model"
const uidAcceptedLetters = "ABCEFGHIJKLMNOPQRSUVWXYZ"

func (service *Service) Create(
	inputJwtToken *types.JwtToken,
	request *data.StudentRequest,
) (result *model.Student, errCode int, err error) {
	// Get role
	userRole, errRole := service.RoleService.Repository.GetByName(config.Env.RoleStudent)
	if errRole != nil || userRole == nil || userRole.ID < 1 {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Get the school
	foundSchool, errSchool := service.SchoolService.Repository.GetByID(request.SchoolID)
	if errSchool != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundSchool == nil || foundSchool.ID < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("school")
		return
	}

	// Generate the email
	newEmail := request.Email
	if request.AutoGenerateEmail {
		newEmail = mail.GenerateEmailFromFullName(
			request.Info.FirstName,
			request.Info.LastName,
			foundSchool.Config.UserEmailDomain,
		)
	}

	// Format request
	if request == nil || request.Info == nil {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessage()
		return
	}
	var item = &dataUser.UserRequest{
		RoleID:      userRole.ID,
		Email:       newEmail,
		PhoneNumber: request.PhoneNumber,
		IsActivated: true,
		Info: &dataUser.UserInfoRequest{
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
	password := password.GeneratePasswordFromUserInfo(
		item.Info.FirstName,
		item.Info.LastName,
		item.Info.Birthday,
	)

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
		// Get all uids to exclude
		uids, errUids := service.Repository.GetAll(nil, nil, &data.GetAllRequest{SchoolID: request.SchoolID})
		if errUids != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
		formattedUids := make([]string, len(uids))
		for i, uid := range uids {
			formattedUids[i] = uid.UID
		}
		// Generate uid
		var errGenerateUID error
		newUID, errGenerateUID = utils.GenerateRandomUID(
			fmt.Sprintf("%02d", time.Now().Year()),
			uidAcceptedLetters,
			formattedUids,
		)
		if errGenerateUID != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
	}

	// Create
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

func (service *Service) CreateStudentEnroll(
	inputJwtToken *types.JwtToken,
	request *data.StudentEnrollRequest,
) (result *model.StudentEnroll, errCode int, err error) {
	// Format request
	item := &model.StudentEnroll{
		StudentID:     request.StudentID,
		YearID:        request.YearID,
		ClassID:       request.ClassID,
		LevelDomainID: request.LevelDomainID,

		Origin:         "dashboard",
		Status:         request.Status,
		StatusFeedback: request.StatusFeedback,
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

	// Create
	result, err = service.Repository.CreateStudentEnroll(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) CreateStudentEnrollAnonym(
	inputJwtToken *types.JwtToken,
	request *data.StudentEnrollAnonymRequest,
) (result *model.StudentEnroll, errCode int, err error) {
	// Format request
	item := &model.StudentEnroll{
		SchoolID:      request.SchoolID,
		YearID:        request.YearID,
		ClassID:       request.ClassID,
		LevelDomainID: request.LevelDomainID,

		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,

		Origin: "website",
		Status: "initiated",

		Message: request.Message,

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

	// Create
	result, err = service.Repository.CreateStudentEnroll(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) Update(
	inputJwtToken *types.JwtToken,
	id int64,
	request *data.StudentRequest,
) (result *model.Student, errCode int, err error) {
	// Check if exists
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

	// Generate the uid if it is empty
	var newUID = request.UID
	if len(newUID) < 1 {
		// Get all uids to exclude
		uids, errUids := service.Repository.GetAll(nil, nil, &data.GetAllRequest{SchoolID: request.SchoolID})
		if errUids != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
		formattedUids := make([]string, len(uids))
		for i, uid := range uids {
			formattedUids[i] = uid.UID
		}
		// Generate uid
		var errGenerateUID error
		newUID, errGenerateUID = utils.GenerateRandomUID(
			fmt.Sprintf("%02d", time.Now().Year()),
			uidAcceptedLetters,
			formattedUids,
		)
		if errGenerateUID != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
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

	// Get the role
	userRole, errRole := service.RoleService.Repository.GetByName(config.Env.RoleStudent)
	if errRole != nil || userRole == nil || userRole.ID < 1 {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Get the school
	foundSchool, errSchool := service.SchoolService.Repository.GetByID(request.SchoolID)
	if errSchool != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundSchool == nil || foundSchool.ID < 1 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("school")
		return
	}

	// Update
	userRequest := dataUser.UserRequest{
		RoleID:      userRole.ID,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		IsActivated: true,
		Info: &dataUser.UserInfoRequest{
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

	// Update
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

func (service *Service) UpdateStudentEnroll(
	inputJwtToken *types.JwtToken,
	id int64,
	request *data.StudentEnrollRequest,
) (result *model.StudentEnroll, errCode int, err error) {
	// Check if exists
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

		Email:       foundItem.Email,
		PhoneNumber: foundItem.PhoneNumber,

		Message:        foundItem.Message,
		Origin:         foundItem.Origin,
		Status:         request.Status,
		StatusFeedback: request.StatusFeedback,

		Gender:        foundItem.Gender,
		FirstName:     foundItem.FirstName,
		LastName:      foundItem.LastName,
		Birthday:      foundItem.Birthday,
		BirthLocation: foundItem.BirthLocation,

		Document1: foundItem.Document1,
		Document2: foundItem.Document2,
		Document3: foundItem.Document3,
		Document4: foundItem.Document4,
		Document5: foundItem.Document5,
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

	// Update
	result, err = service.Repository.UpdateStudentEnrollByID(id, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) Delete(
	inputJwtToken *types.JwtToken,
	id int64,
) (affectedRows int64, errCode int, err error) {
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

func (service *Service) DeleteStudentEnroll(
	inputJwtToken *types.JwtToken,
	id int64,
) (affectedRows int64, errCode int, err error) {
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

func (service *Service) Get(
	inputJwtToken *types.JwtToken,
	id int64,
) (result *model.Student, errCode int, err error) {
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

func (service *Service) GetStudentEnroll(
	inputJwtToken *types.JwtToken,
	id int64,
) (result *model.StudentEnroll, errCode int, err error) {
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

func (service *Service) GetAll(
	inputJwtToken *types.JwtToken,
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllRequest,
) (result []model.Student, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, request)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}

func (service *Service) GetAllStudentEnroll(
	inputJwtToken *types.JwtToken,
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllStudentEnrollRequest,
) (result []model.StudentEnroll, errCode int, err error) {
	result, err = service.Repository.GetAllStudentEnroll(filter, pagination, request)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
