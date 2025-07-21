package contact

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/common/utils"
	"api/services/common/contact/data"
	"api/services/common/contact/model"
	serviceHelper "api/services/helper"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

const MODEL_NAME = "contact"
const DEFAULT_ERROR_MESSAGE = "interact with contact model"

func (service *Service) Create(inputJwtToken *types.JwtToken, request *data.ContactRequest) (result *model.Contact, errCode int, err error) {
	// Check inputs
	if inputJwtToken.SchoolID != request.SchoolID {
		errCode = http.StatusBadRequest
		err = constants.Http422LockedErrorMessage()
		return
	}

	// Insert contact
	if inputJwtToken.SchoolID > 0 {
		result, err = service.Repository.Create(&model.Contact{
			SchoolID: request.SchoolID,
			Subject:  request.Subject,
			Email:    request.Email,
			Message:  request.Message,
		})
	} else {
		result, err = service.Repository.Create(&model.Contact{
			Subject: request.Subject,
			Email:   request.Email,
			Message: request.Message,
		})
	}
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

func (service *Service) Get(inputJwtToken *types.JwtToken, id int64) (result *model.Contact, errCode int, err error) {
	// Get user
	foundUser, err := serviceHelper.GetUserByID(inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Proceed by feature
	if foundUser.Role.Feature != constants.FeatureAdmin {
		result, err = service.Repository.GetByIDSchoolID(id, foundUser.SchoolID)
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

func (service *Service) GetAll(inputJwtToken *types.JwtToken, filter *types.Filter, pagination *types.Pagination, request *data.GetAllRequest) (result []model.Contact, errCode int, err error) {
	// Get user
	foundUser, err := serviceHelper.GetUserByID(inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Proceed by feature
	if foundUser.Role.Feature != constants.FeatureAdmin {
		newRequest := *request
		newRequest.SchoolID = foundUser.SchoolID
		result, err = service.Repository.GetAll(filter, pagination, &newRequest)
	} else {
		result, err = service.Repository.GetAll(filter, pagination, request)
	}
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}
