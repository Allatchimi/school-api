package communication

import (
	"net/http"
	"time"

	"api/common/constants"
	"api/common/helpers"
	smtpHelper "api/common/helpers/message/mail/smtp"
	telegramHelper "api/common/helpers/message/telegram"
	whatsappHelper "api/common/helpers/message/whatsapp"
	"api/common/types"
	"api/common/utils"
	webpushConfig "api/config/webpush"
	"api/services/common/communication/data"
	"api/services/common/communication/model"
	serviceHelper "api/services/helper"
	dataUser "api/services/user/user/data"

	"go.uber.org/zap"
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
		if result == nil || result.School == nil || result.School.Config == nil {
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

		// Send push notification
		go func() {
			createdAt := new(time.Time)
			*createdAt = time.Now()
			webpushConfig.SendPushNotificationToUserBulk(
				users,
				&webpushConfig.WebPushPayload{
					Title:     request.Subject,
					Body:      request.Message,
					Icon:      result.School.LogoUrl(),
					Url:       result.School.WebsiteUrl(),
					CreatedAt: createdAt,
				},
				nil,
				serviceHelper.UserService.Repository,
			)
		}()
		// Send telegram
		go func() {
			telegramHelper.SendMessage(
				result.School.Config.TelegramBotToken,
				request.Message,
				users,
			)
		}()
		// Send whatsapp
		go func() {
			whatsappHelper.SendMessage(
				result.School.Config.WhatsappToken,
				result.School.Config.WhatsappPhoneID,
				request.Message,
				users,
			)
		}()
		// Send mail
		go func() {
			mailData := &smtpHelper.EmailData{
				HomePageLink: result.School.WebsiteUrl(),
				Logo:         result.School.LogoUrl(),
				Title:        request.Subject,
				Message:      request.Message,
			}
			mailBody, errTemplate := mailData.LoadTemplate()
			if errTemplate != nil {
				helpers.Logger.Error("Failed to load email template!", zap.Error(errTemplate))
			}
			if len(mailBody) < 1 {
				helpers.Logger.Warn("Empty message body!")
			}
			mailUsers := make([]string, len(users))
			for i := range users {
				mailUsers[i] = users[i].Email
			}
			errMail := smtpHelper.SendEmailBCC(mailUsers, request.Subject, mailBody)
			if errMail != nil {
				helpers.Logger.Error("Failed to send mail!", zap.Error(errMail))
			}
		}()
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
