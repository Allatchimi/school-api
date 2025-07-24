package teacher

import (
	"fmt"
	"net/http"
	"time"

	"api/common/constants"
	"api/common/helpers"
	"api/common/types"
	"api/common/utils"
	"api/config"
	"api/services/school/common/school"
	"api/services/school/common/teacher/data"
	"api/services/school/common/teacher/model"
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

const MODEL_NAME = "teacher"
const DEFAULT_ERROR_MESSAGE = "interact with teacher model"
const uidAcceptedLetters = "T"

func (service *Service) Create(
	ctxData *types.ContextData,
	request *data.TeacherRequest,
) (result *model.Teacher, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Get role
	userRole, errRole := service.RoleService.Repository.GetByName(config.Env.FixtureRoleTeacher)
	if errRole != nil || userRole == nil || userRole.ID < 1 {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Get the school
	foundSchool, errSchool := service.SchoolService.Repository.GetByID(newRequest.SchoolID)
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
	if request == nil || newRequest.Info == nil {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessage()
		return
	}

	// Generate the email
	newEmail := newRequest.Email
	if newRequest.AutoGenerateEmail {
		if foundSchool.Config == nil {
			errCode = http.StatusNotFound
			err = constants.Http404ErrorMessage("school")
			return
		}
		newEmail = helpers.GenerateEmailFromFullName(
			newRequest.Info.FirstName,
			newRequest.Info.LastName,
			foundSchool.Config.UserEmailDomainName,
		)
	}

	// Format request
	var item = &dataUser.UserRequest{
		SchoolID:    newRequest.SchoolID,
		RoleID:      userRole.ID,
		Email:       newEmail,
		PhoneNumber: newRequest.PhoneNumber,
		IsActivated: true,
		Status:      newRequest.Status,
		Info: &dataUser.UserInfoRequest{
			Gender:        newRequest.Info.Gender,
			Username:      newRequest.Info.Username,
			FirstName:     newRequest.Info.FirstName,
			LastName:      newRequest.Info.LastName,
			Birthday:      newRequest.Info.Birthday,
			BirthLocation: newRequest.Info.BirthLocation,
			Address:       newRequest.Info.Address,
			Language:      newRequest.Info.Language,
			Image:         newRequest.Info.Image,
		},
	}

	// Generate password
	password := helpers.GeneratePasswordFromUser(
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
	var newUID = newRequest.UID
	if len(newUID) < 1 {
		// Get all uids to exclude
		uids, errUids := service.Repository.GetAll(nil, nil, &data.GetAllRequest{SchoolID: newRequest.SchoolID})
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
	result, err = service.Repository.Create(&model.Teacher{
		SchoolID: newRequest.SchoolID,
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

func (service *Service) CreateTeacherClassSubjectUnit(
	ctxData *types.ContextData,
	request *data.TeacherClassSubjectUnitRequest,
) (result *model.TeacherClassSubjectUnit, errCode int, err error) {
	// Format request
	newRequest := *request
	item := &model.TeacherClassSubjectUnit{
		TeacherID:      newRequest.TeacherID,
		YearID:         newRequest.YearID,
		ClassSubjectID: newRequest.ClassSubjectID,
		UnitID:         newRequest.UnitID,
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

	// Insert
	result, err = service.Repository.CreateTeacherClassSubjectUnit(item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) Update(
	ctxData *types.ContextData,
	id int64,
	request *data.TeacherRequest,
) (result *model.Teacher, errCode int, err error) {
	// Check if the item exists
	var foundItem *model.Teacher
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetByIDSchoolID(id, ctxData.Jwt.SchoolID)
	} else {
		foundItem, err = service.Repository.GetByID(id)
	}
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

	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Generate the uid if it is empty
	var newUID = newRequest.UID
	if len(newUID) < 1 {
		// Get all uids to exclude
		uids, errUids := service.Repository.GetAll(nil, nil, &data.GetAllRequest{SchoolID: newRequest.SchoolID})
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

	// Get role
	teacherRole, errRole := service.RoleService.Repository.GetByName(config.Env.FixtureRoleTeacher)
	if errRole != nil || teacherRole == nil || teacherRole.ID < 1 {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Update user
	userRequest := dataUser.UserRequest{
		SchoolID:    newRequest.SchoolID,
		RoleID:      teacherRole.ID,
		Email:       newRequest.Email,
		PhoneNumber: newRequest.PhoneNumber,
		IsActivated: true,
		Status:      newRequest.Status,
		Info: &dataUser.UserInfoRequest{
			Gender:        newRequest.Info.Gender,
			Username:      newRequest.Info.Username,
			FirstName:     newRequest.Info.FirstName,
			LastName:      newRequest.Info.LastName,
			Birthday:      newRequest.Info.Birthday,
			BirthLocation: newRequest.Info.BirthLocation,
			Address:       newRequest.Info.Address,
			Language:      newRequest.Info.Language,
			Image:         newRequest.Info.Image,
		},
	}
	_, errCodeUser, errUser := service.UserService.Update(nil, foundItem.UserID, &userRequest)
	if errUser != nil {
		errCode = errCodeUser
		err = errUser
		return
	}

	// Update
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

func (service *Service) UpdateTeacherClassSubjectUnit(
	ctxData *types.ContextData,
	id int64,
	request *data.TeacherClassSubjectUnitRequest,
) (result *model.TeacherClassSubjectUnit, errCode int, err error) {
	// Check if the item exists
	var foundItem *model.TeacherClassSubjectUnit
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetByIDSchoolID(id, ctxData.Jwt.SchoolID)
	} else {
		foundItem, err = service.Repository.GetByID(id)
	}
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

	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Format request
	item := &model.TeacherClassSubjectUnit{
		TeacherID:      newRequest.TeacherID,
		YearID:         newRequest.YearID,
		ClassSubjectID: newRequest.ClassSubjectID,
		UnitID:         newRequest.UnitID,
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

	// Update
	result, err = service.Repository.UpdateTeacherClassSubjectUnitByID(id, item)
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
	// Check if the item exists
	var foundItem *model.Teacher
	if ctxData.User.Feature != constants.FeatureAdmin {
		foundItem, err = service.Repository.GetByIDSchoolID(id, ctxData.Jwt.SchoolID)
	} else {
		foundItem, err = service.Repository.GetByID(id)
	}
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

	// Delete
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

func (service *Service) DeleteTeacherClassSubjectUnit(
	ctxData *types.ContextData,
	id int64,
) (affectedRows int64, errCode int, err error) {
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

func (service *Service) DeleteMultiple(
	ctxData *types.ContextData,
	list []int64,
) (affectedRows int64, errCode int, err error) {
	// Check school
	var foundSchool int64
	if ctxData.Jwt.SchoolID > 0 {
		foundSchool = ctxData.Jwt.SchoolID
	}

	affectedRows, err = service.Repository.DeleteMultipleByID(list, foundSchool)
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
	ctxData *types.ContextData,
	id int64,
) (result *model.Teacher, errCode int, err error) {
	if ctxData.User.Feature != constants.FeatureAdmin {
		result, err = service.Repository.GetByIDSchoolID(id, ctxData.Jwt.SchoolID)
	} else {
		result, err = service.Repository.GetByID(id)
	}
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

func (service *Service) GetTeacherClassSubjectUnit(
	ctxData *types.ContextData,
	id int64,
) (result *model.TeacherClassSubjectUnit, errCode int, err error) {
	if ctxData.User.Feature != constants.FeatureAdmin {
		result, err = service.Repository.GetTeacherClassSubjectUnitByIDShoolID(id, ctxData.Jwt.SchoolID)
	} else {
		result, err = service.Repository.GetTeacherClassSubjectUnitByID(id)
	}
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
	request *data.GetAllRequest,
) (result []model.Teacher, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Get
	result, err = service.Repository.GetAll(filter, pagination, &newRequest)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}

func (service *Service) GetAllTeacherClassSubjectUnit(
	ctxData *types.ContextData,
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllTeacherClassSubjectUnitRequest,
) (result []model.TeacherClassSubjectUnit, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Get
	result, err = service.Repository.GetAllTeacherClassSubjectUnit(filter, pagination, &newRequest)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
