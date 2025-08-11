package director

import (
	"fmt"
	"net/http"
	"time"

	"api/common/constants"
	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/config"
	serviceHelperMessage "api/services/helper/message"
	"api/services/school/common/director/data"
	"api/services/school/common/director/model"
	"api/services/school/common/school"
	"api/services/user/role"
	"api/services/user/user"
	dataUser "api/services/user/user/data"
	modelUser "api/services/user/user/model"
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

const MODEL_NAME = "director"
const DEFAULT_ERROR_MESSAGE = "interact with director model"
const uidAcceptedLetters = "D"

func (service *Service) Create(
	ctxData *types.ContextData,
	request *data.DirectorRequest,
) (result *model.Director, errCode int, err error) {
	// Get role
	userRole, errRole := service.RoleService.Repository.GetByName(config.Env.FixtureRoleDirector)
	if errRole != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if userRole == nil || userRole.ID < 1 {
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

	// Check request
	if request == nil || request.Info == nil {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessage()
		return
	}

	// Check email
	newEmail := request.Email
	if request.AutoGenerateEmail {
		if foundSchool.Config == nil {
			errCode = http.StatusNotFound
			err = constants.Http404ErrorMessage("school")
			return
		}
		newEmail = helpers.GenerateEmailFromFullName(
			request.Info.FirstName,
			request.Info.LastName,
			foundSchool.Config.UserEmailDomainName,
		)
	}

	// Format request
	var item = &dataUser.UserRequest{
		SchoolID:    request.SchoolID,
		RoleID:      userRole.ID,
		Email:       newEmail,
		PhoneNumber: request.PhoneNumber,
		IsActivated: true,
		Status:      request.Status,
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
	password := helpers.GeneratePasswordFromUser(
		item.Info.FirstName,
		item.Info.LastName,
		item.Info.Birthday,
	)

	// Create user
	createdUser, errCodeCreate, errCreate := service.UserService.Create(ctxData, item, &password)
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

	// Insert
	result, err = service.Repository.Create(&model.Director{
		SchoolID: request.SchoolID,
		UserID:   createdUser.ID,
		UID:      newUID,
	})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if result == nil || result.ID < 1 {
		return
	}

	// Send message
	var msgTitle, msgBody string
	msgTitle = fmt.Sprintf("Welcome to %s - Director Assignment", result.School.Name)
	msgBody = fmt.Sprintf(`Dear %s %s,
	<br><br>
	Congratulations on your appointment as Director of %s!
	We are delighted to welcome you to our educational community. 
	Your leadership and expertise will be invaluable as we continue to provide excellent education to our students.

	<br><br>
	Your account has been set up with the following credentials:
	<br>
	• Email: %s
	<br>
	• Temporary Password: %s

	<br><br>
	For security reasons, please log in to the administrative portal at your earliest convenience and update your password. 
	You will have full administrative access to manage school operations, staff, students, and resources.
	If you need any assistance getting started or have questions about the platform, please don't hesitate to contact our support team.

	<br><br>
	We look forward to working with you and wish you great success in your new role.
	
	<br><br>
	Best regards,
	The Administration Team`,
		result.User.Info.FirstName, result.User.Info.LastName, result.School.Name, result.User.Email, password)

	go serviceHelperMessage.SendMessage(
		&serviceHelperMessage.MessageRequest{
			Mail: true,
		},
		msgTitle,
		msgBody,
		result.School,
		"",
		[]modelUser.User{*result.User},
	)
	return
}

func (service *Service) Update(
	ctxData *types.ContextData,
	id int64,
	request *data.DirectorRequest,
) (result *model.Director, errCode int, err error) {
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

	// Get role
	userRole, errRole := service.RoleService.Repository.GetByName(config.Env.FixtureRoleDirector)
	if errRole != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if userRole == nil || userRole.ID < 1 {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Update user
	user := dataUser.UserRequest{
		SchoolID:    request.SchoolID,
		RoleID:      userRole.ID,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		IsActivated: true,
		Status:      request.Status,
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
	_, errCodeUser, errUser := service.UserService.Update(ctxData, foundItem.UserID, &user)
	if errUser != nil {
		errCode = errCodeUser
		err = errUser
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
	foundUnique, err := service.Repository.GetUniqueObjectByUID(&model.Director{
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
		&model.Director{
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

	// Update
	result, err = service.Repository.UpdateByID(id, &model.Director{
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

func (service *Service) Delete(
	ctxData *types.ContextData,
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

func (service *Service) DeleteMultiple(ctxData *types.ContextData, list []int64) (affectedRows int64, errCode int, err error) {
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

func (service *Service) Get(ctxData *types.ContextData, id int64) (result *model.Director, errCode int, err error) {
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

func (service *Service) GetAll(
	ctxData *types.ContextData,
	filter *types.Filter,
	pagination *types.Pagination,
	schoolID int64,
) (result []model.Director, errCode int, err error) {
	result, err = service.Repository.GetAll(filter, pagination, &data.GetAllRequest{SchoolID: schoolID})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
