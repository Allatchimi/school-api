package communication

import (
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/common/utils"
	serviceHelperMessage "api/services/helper/message"
	"api/services/others/communication/data"
	"api/services/others/communication/model"
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

func (service *Service) Create(
	ctxData *types.ContextData,
	request *data.CommunicationRequest,
) (result *model.Communication, errCode int, err error) {
	// Check school
	newRequest := *request
	if ctxData.User.Feature != constants.FeatureAdmin {
		newRequest.SchoolID = ctxData.Jwt.SchoolID
	}

	// Insert
	for _, roleID := range request.RoleIDs {
		if newRequest.SchoolID > 0 {
			result, err = service.Repository.Create(&model.Communication{
				SchoolID: newRequest.SchoolID,
				RoleID:   roleID,

				Subject: request.Subject,
				Message: request.Message,
			})
		} else {
			result, err = service.Repository.Create(&model.Communication{
				RoleID: roleID,

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
		go func() {
			if result == nil {
				return
			}
			users, _ := serviceHelperMessage.UserService.Repository.GetAll(
				nil,
				nil,
				&dataUser.GetAllRequest{
					SchoolID: newRequest.SchoolID,
					RoleID:   roleID,
				},
			)

			// Send message
			serviceHelperMessage.SendMessage(
				&serviceHelperMessage.MessageRequest{
					PusNotification: true,
					Mail:            true,
				},
				request.Subject,
				request.Message,
				result.School,
				"",
				users,
			)
		}()
	}
	return
}

func (service *Service) Get(ctxData *types.ContextData, id int64) (result *model.Communication, errCode int, err error) {
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
) (result []model.Communication, errCode int, err error) {
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
