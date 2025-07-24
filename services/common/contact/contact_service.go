package contact

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/common/utils"
	"api/services/common/contact/data"
	"api/services/common/contact/model"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

const MODEL_NAME = "contact"
const DEFAULT_ERROR_MESSAGE = "interact with contact model"

func (service *Service) Create(
	ctxData *types.ContextData,
	request *data.ContactRequest,
) (result *model.Contact, errCode int, err error) {
	// Check inputs
	if ctxData.Jwt.SchoolID != request.SchoolID {
		errCode = http.StatusBadRequest
		err = constants.Http422InvalidInputsErrorMessage()
		return
	}

	// Insert contact
	if ctxData.Jwt.SchoolID > 0 {
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

func (service *Service) Get(ctxData *types.ContextData, id int64) (result *model.Contact, errCode int, err error) {
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
) (result []model.Contact, errCode int, err error) {
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
