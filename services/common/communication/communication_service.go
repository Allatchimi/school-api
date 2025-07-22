package communication

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/common/utils"
	"api/services/common/communication/data"
	"api/services/common/communication/model"
	serviceHelper "api/services/helper"
	dataUser "api/services/user/user/data"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

const MODEL_NAME = "communication"
const DEFAULT_ERROR_MESSAGE = "interact with communication model"

func (service *Service) Create(inputJwtToken *types.JwtToken, request *data.CommunicationRequest) (result *model.Communication, errCode int, err error) {
	// Get user
	foundUser, err := serviceHelper.GetUserByID(inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Insert communication
	var foundSchoolID int64 = request.SchoolID
	if foundUser.Role.Feature != constants.FeatureAdmin {
		foundSchoolID = foundUser.SchoolID
	}
	for _, roleID := range request.RoleIDs {
		if foundSchoolID > 0 {
			result, err = service.Repository.Create(&model.Communication{
				SchoolID: foundSchoolID,
				RoleID:   roleID,
				Subject:  request.Subject,
				Message:  request.Message,
			})
		} else {
			result, err = service.Repository.Create(&model.Communication{
				RoleID:  roleID,
				Subject: request.Subject,
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

		// Get all user
		if result == nil {
			continue
		}
		users, _ := serviceHelper.UserService.Repository.GetAll(
			nil,
			nil,
			&dataUser.GetAllRequest{
				SchoolID: foundSchoolID,
				RoleID:   roleID,
			},
		)

		// Send message
		serviceHelper.SendMessage(
			&serviceHelper.MessageRequest{
				PusNotification: true,
				Telegram:        true,
				Whatsapp:        true,
				Mail:            true,
			},
			request.Subject,
			request.Message,
			result.School,
			users,
		)
	}
	return
}

func (service *Service) Get(inputJwtToken *types.JwtToken, id int64) (result *model.Communication, errCode int, err error) {
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

func (service *Service) GetAll(
	inputJwtToken *types.JwtToken,
	filter *types.Filter,
	pagination *types.Pagination,
	request *data.GetAllRequest,
) (result []model.Communication, errCode int, err error) {
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
