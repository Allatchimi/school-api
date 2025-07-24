package user

import (
	"net/http"
	"time"

	"api/common/constants"
	"api/common/types"
	"api/common/utils"
	"api/services/user/user/data"
	"api/services/user/user/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

const MODEL_NAME = "user"
const DEFAULT_ERROR_MESSAGE = "interact with user model"

func (service *Service) Create(
	ctxData *types.ContextData,
	request *data.UserRequest,
	password *string,
) (result *model.User, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.Jwt.SchoolID > 0 {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Format item
	item := &model.User{
		SchoolID:    newRequest.SchoolID,
		RoleID:      newRequest.RoleID,
		Email:       newRequest.Email,
		PhoneNumber: newRequest.PhoneNumber,
		Status:      newRequest.Status,
		IsActivated: newRequest.IsActivated,
		Info: &model.UserInfo{
			Username:  newRequest.Info.Username,
			FirstName: newRequest.Info.FirstName,
			LastName:  newRequest.Info.LastName,

			Gender:        newRequest.Info.Gender,
			Birthday:      newRequest.Info.Birthday,
			BirthLocation: newRequest.Info.BirthLocation,
			Address:       newRequest.Info.Address,
			Language:      newRequest.Info.Language,
			Image:         newRequest.Info.Image,
		},
	}

	// Check if user exists
	var foundItem *model.User
	var isEmailValid = utils.IsEmailValid(item.Email)
	var isPhoneNumberValid = utils.IsPhoneNumberValid(item.PhoneNumber)
	if isEmailValid {
		foundItem, err = service.Repository.GetByEmailSchoolID(item.Email, item.SchoolID)
	} else if isPhoneNumberValid {
		foundItem, err = service.Repository.GetByPhoneNumberSchoolID(item.PhoneNumber, item.SchoolID)
	} else {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessage()
		return
	}
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundItem != nil && foundItem.ID > 0 {
		if (isEmailValid && foundItem.Email == item.Email) || (!isEmailValid && foundItem.PhoneNumber == item.PhoneNumber) {
			errCode = http.StatusFound
			err = constants.Http302ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
	}

	// Create user info
	tempInfo, err := service.Repository.CreateUserInfo(&model.UserInfo{
		Username:  item.Info.Username,
		FirstName: item.Info.FirstName,
		LastName:  item.Info.LastName,

		Gender:        item.Info.Gender,
		Birthday:      item.Info.Birthday,
		BirthLocation: item.Info.BirthLocation,
		Address:       item.Info.Address,
		Language:      item.Info.Language,
		Image:         item.Info.Image,
	})
	if err != nil || tempInfo == nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Create user config
	tempConfig, err := service.Repository.CreateUserConfig(&model.UserConfig{
		AllowNotifications: true,
	})
	if err != nil || tempConfig == nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Create user
	var randomPassword string = ""
	if password == nil || len(*password) < 1 {
		randomPassword = utils.GenerateRandomPassword(8)
	} else {
		randomPassword = *password
	}
	var activatedAt *time.Time = nil
	if item.IsActivated {
		tmpTime := time.Now()
		activatedAt = &tmpTime
	}
	result, err = service.Repository.Create(&model.User{
		RoleID:       item.RoleID,
		SchoolID:     item.SchoolID,
		Email:        item.Email,
		PhoneNumber:  item.PhoneNumber,
		Status:       item.Status,
		IsActivated:  item.IsActivated,
		ActivatedAt:  activatedAt,
		LoginMethod:  constants.AuthLoginMethodDefault,
		Password:     randomPassword,
		UserInfoID:   tempInfo.ID,
		UserConfigID: tempConfig.ID,
	})
	if err != nil {
		pgState, errPgState := utils.ExtractSQLState(err.Error())
		if errPgState == nil {
			if pgState == constants.PG_ERROR_CONSTRAINT_COLUMN {
				errCode = http.StatusConflict
				err = constants.Http409ConflictErrorMessage()
				return
			}
		}
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}

func (service *Service) Update(
	ctxData *types.ContextData,
	id int64,
	request *data.UserRequest,
) (result *model.User, errCode int, err error) {
	// Check if the item exists
	var foundItem *model.User
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

	// Format item
	item := &model.User{
		SchoolID:    newRequest.SchoolID,
		RoleID:      newRequest.RoleID,
		Email:       newRequest.Email,
		PhoneNumber: newRequest.PhoneNumber,
		Status:      newRequest.Status,
		IsActivated: newRequest.IsActivated,
		Info: &model.UserInfo{
			Username:  newRequest.Info.Username,
			FirstName: newRequest.Info.FirstName,
			LastName:  newRequest.Info.LastName,

			Gender:        newRequest.Info.Gender,
			Birthday:      newRequest.Info.Birthday,
			BirthLocation: newRequest.Info.BirthLocation,
			Address:       newRequest.Info.Address,
			Language:      newRequest.Info.Language,
			Image:         newRequest.Info.Image,
		},
	}

	// Check if user exists
	var isEmailValid = utils.IsEmailValid(item.Email)
	var isPhoneNumberValid = utils.IsPhoneNumberValid(item.PhoneNumber)
	if isEmailValid {
		foundItem, err = service.Repository.GetByEmailSchoolID(item.Email, item.SchoolID)
	} else if isPhoneNumberValid {
		foundItem, err = service.Repository.GetByPhoneNumberSchoolID(item.PhoneNumber, item.SchoolID)
	} else {
		errCode = http.StatusBadRequest
		err = constants.Http400BadRequestErrorMessage()
		return
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

	// Update user info
	_, err = service.Repository.UpdateUserInfoByID(foundItem.UserInfoID, item.Info)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	// Update user
	if !item.IsActivated {
		item.ActivatedAt = nil
	} else if item.IsActivated && !foundItem.IsActivated {
		tmpTime := time.Now()
		item.ActivatedAt = &tmpTime
	}
	result, err = service.Repository.UpdateByID(id, item)
	if err != nil {
		pgState, errPgState := utils.ExtractSQLState(err.Error())
		if errPgState == nil {
			if pgState == constants.PG_ERROR_CONSTRAINT_COLUMN {
				errCode = http.StatusConflict
				err = constants.Http409ConflictErrorMessage()
				return
			}
		}
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
	var foundItem *model.User
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
) (result *model.User, errCode int, err error) {
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

func (service *Service) GetAll(
	ctxData *types.ContextData,
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllRequest,
) (result []model.User, errCode int, err error) {
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
