package profile

import (
	"fmt"
	"net/http"

	"api/common/constants"
	smtpHelper "api/common/helpers/message/mail/smtp"
	"api/common/types"
	"api/common/utils"
	securityUtil "api/common/utils/security"
	"api/config"
	"api/services/user/profile/data"
	"api/services/user/user"
	"api/services/user/user/model"
)

type Service struct {
	UserService *user.Service
}

func NewService(userService *user.Service) *Service {
	return &Service{UserService: userService}
}

const MODEL_NAME = "user"
const DEFAULT_ERROR_MESSAGE = "interact with user model"

func (service *Service) UpdateProfileInfo(inputJwtToken *types.JwtToken, request *data.UpdateProfileInfoRequest) (result *model.UserInfo, errCode int, err error) {
	// Find user
	userFound, err := service.UserService.Repository.GetByID(inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if userFound == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Format item
	item := &model.UserInfo{
		Username:  request.Username,
		FirstName: request.FirstName,
		LastName:  request.LastName,

		Gender:        request.Gender,
		Birthday:      request.Birthday,
		BirthLocation: request.BirthLocation,
		Address:       request.Address,
		Language:      request.Language,
		Image:         request.Image,
	}

	//Update user info
	result, err = service.UserService.Repository.UpdateUserInfoByID(userFound.UserInfoID, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}

func (service *Service) UpdateProfileConfigMessage(inputJwtToken *types.JwtToken, request *data.UpdateProfileMessageRequest) (result *model.UserConfig, errCode int, err error) {
	// Find user
	userFound, err := service.UserService.Repository.GetByID(inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if userFound == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Format item
	item := &model.UserConfig{
		WhatsappPhoneNumber: request.WhatsappPhoneNumber,
		TelegramChatID:      request.TelegramChatID,
	}

	//Update user config
	result, err = service.UserService.Repository.UpdateUserConfigByID(userFound.UserConfigID, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}

func (service *Service) UpdateProfilePhoneNumber(inputJwtToken *types.JwtToken, phoneNumber uint64) (result *model.User, errCode int, err error) { // Check if user exists
	// Find user
	userFound, err := service.UserService.Repository.GetByID(inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if userFound == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Check if this phone number is already taken
	foundUser, err := service.UserService.Repository.GetByPhoneNumberSchoolID(phoneNumber, inputJwtToken.SchoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if foundUser != nil && foundUser.PhoneNumber == phoneNumber {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(MODEL_NAME)
		return
	}

	// Update
	result, err = service.UserService.Repository.UpdatePhoneNumberByID(userFound.ID, phoneNumber)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}

func (service *Service) UpdateProfilePasswordInit(inputJwtToken *types.JwtToken) (token string, errCode int, err error) {
	// Find user
	userFound, err := service.UserService.Repository.GetByID(inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if userFound == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Generate new random code
	randomCode, err := utils.GenerateRandomCode(6)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	expires := securityUtil.NewExpiresDateDefault()
	newJwtToken, newToken, err := securityUtil.EncodeJWTToken(
		&types.JwtToken{
			UserID:   userFound.ID,
			SchoolID: userFound.SchoolID,
			Platform: "*",
			Device:   "*",
			App:      "*",
			Code:     randomCode,
		},
		constants.JwtIssuerProfileUpdatePasswordCode,
		expires,
		config.Keys.JwtPrivateKey,
		config.SetRedisString,
	)
	if err != nil || newJwtToken == nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	token = newToken

	// Send code to email
	if userFound == nil || userFound.School == nil || userFound.School.Config == nil {
		return
	}
	if utils.IsEmailValid(userFound.Email) {
		go func() {
			fromEmail, fromUsername := userFound.School.SMTPNoReplySender()
			data := &smtpHelper.EmailDataCheckCode{
				EmailData: smtpHelper.EmailData{
					HomePageLink: fmt.Sprintf("https://%s", userFound.School.Config.DomainName),
					Logo:         userFound.School.Logo,
					Title:        constants.MailUpdatePasswordCheckCode.Title,
					Message:      constants.MailUpdatePasswordCheckCode.Message,
				},
				Code:            fmt.Sprintf("%d", randomCode),
				DurationMinutes: 10,
			}
			body, err := data.LoadTemplate()
			if err != nil {
				return
			}
			err = smtpHelper.SendEmailTo(
				fromEmail,
				fromUsername,
				userFound.Email,
				constants.MailUpdatePasswordCheckCode.Subject,
				body,
			)
			if err != nil {
				return
			}
		}()
	}
	return
}

func (service *Service) UpdateProfilePasswordCheckCode(inputJwtToken *types.JwtToken, inputToken string, inputCode int) (token string, errCode int, err error) {
	// Find user
	userFound, err := service.UserService.Repository.GetByID(inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if userFound == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Check input
	if len(inputToken) <= 0 && inputCode < 10000 {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid token and code! Please enter valid information.")
		return
	}
	if len(inputToken) <= 0 {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid token! Please enter valid information.")
		return
	}
	if inputCode < 10000 {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid code! Please enter valid information.")
		return
	}

	// Extract token information and validate the token
	errMsg := "Invalid or expired token! Please enter valid information."
	jwtToken, err := securityUtil.DecodeJWTToken(inputToken, config.Keys.JwtPublicKey)
	if err != nil {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMsg)
		return
	}
	if jwtToken == nil || jwtToken.UserID <= 0 || jwtToken.Issuer != constants.JwtIssuerProfileUpdatePasswordCode {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMsg)
		return
	}
	isTokenValid := securityUtil.ValidateJWTToken(inputToken, jwtToken, config.GetRedisString)
	if !isTokenValid {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMsg)
		return
	}

	// Check if the code is valid
	if jwtToken.Code < 1 || jwtToken.Code != inputCode {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", "Invalid code! Please enter valid information.")
		return
	}

	// Invalidate token
	_, _ = config.DeleteRedisString(securityUtil.GetJWTCachedKey(jwtToken.UserID, jwtToken.Issuer))

	// Generate new token
	newJwtToken, newToken, err := securityUtil.EncodeJWTToken(
		&types.JwtToken{
			UserID:   userFound.ID,
			SchoolID: userFound.SchoolID,
			Platform: "*",
			Device:   "*",
			App:      "*",
		},
		constants.JwtIssuerProfileUpdatePasswordNewPassword,
		securityUtil.NewExpiresDateDefault(),
		config.Keys.JwtPrivateKey,
		config.SetRedisString,
	)
	if err != nil || newJwtToken == nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	token = newToken
	return
}

func (service *Service) UpdateProfilePasswordNewPassword(inputJwtToken *types.JwtToken, token string, currentPassword string, password string) (errCode int, err error) {
	// Find user
	userFound, err := service.UserService.Repository.GetByID(inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if userFound == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Check input
	isPasswordValid, missingPasswordChars := utils.IsPasswordValid(password)
	if len(token) <= 0 && !isPasswordValid {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s %s",
			"Invalid token and password! Please enter valid information.",
			missingPasswordChars,
		)
		return
	}
	if len(token) <= 0 {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid token! Please enter valid information.")
		return
	}
	if !isPasswordValid {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s %s",
			"Invalid password! Please enter valid information.",
			missingPasswordChars,
		)
		return
	}

	// Extract token information and validate the token
	errMsg := "Invalid or expired token! Please enter valid information."
	jwtTokenDecoded, err := securityUtil.DecodeJWTToken(token, config.Keys.JwtPublicKey)
	if err != nil {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMsg)
		return
	}
	if jwtTokenDecoded == nil || jwtTokenDecoded.UserID <= 0 ||
		jwtTokenDecoded.Issuer != constants.JwtIssuerProfileUpdatePasswordNewPassword ||
		jwtTokenDecoded.UserID != inputJwtToken.UserID {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMsg)
		return
	}
	isTokenValid := securityUtil.ValidateJWTToken(token, jwtTokenDecoded, config.GetRedisString)
	if !isTokenValid {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMsg)
		return
	}

	// Check if current password is correct
	if userFound.Password != currentPassword {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Current password is incorrect! Please enter valid information.")
		return
	}

	// Update user password
	userUpdated, err := service.UserService.Repository.UpdatePasswordByID(jwtTokenDecoded.UserID, password)
	if err != nil || userUpdated == nil {
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

	// Invalidate token
	_, _ = config.DeleteRedisString(securityUtil.GetJWTCachedKey(jwtTokenDecoded.UserID, jwtTokenDecoded.Issuer))

	// Send alert message to email
	if userFound == nil || userFound.School == nil || userFound.School.Config == nil {
		return
	}
	if utils.IsEmailValid(userFound.Email) {
		go func() {
			fromEmail, fromUsername := userFound.School.SMTPNoReplySender()
			data := &smtpHelper.EmailData{
				HomePageLink: fmt.Sprintf("https://%s", userFound.School.Config.DomainName),
				Logo:         userFound.School.Logo,
				Title:        constants.MailUpdatePasswordSuccess.Title,
				Message:      constants.MailUpdatePasswordSuccess.Message,
			}
			body, err := data.LoadTemplate()
			if err != nil {
				return
			}
			err = smtpHelper.SendEmailTo(
				fromEmail,
				fromUsername,
				userFound.Email,
				constants.MailUpdatePasswordCheckCode.Subject,
				body,
			)
			if err != nil {
				return
			}
		}()
	}
	return
}

func (service *Service) UpdateProfilePhoneNumberInit(inputJwtToken *types.JwtToken) (token string, errCode int, err error) {
	// Find user
	userFound, err := service.UserService.Repository.GetByID(inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if userFound == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Generate new random code
	randomCode, err := utils.GenerateRandomCode(6)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	expires := securityUtil.NewExpiresDateDefault()
	newJwtToken, newToken, err := securityUtil.EncodeJWTToken(
		&types.JwtToken{
			UserID:   userFound.ID,
			SchoolID: userFound.SchoolID,
			Platform: "*",
			Device:   "*",
			App:      "*",
			Code:     randomCode,
		},
		constants.JwtIssuerProfileUpdatePhoneNumberCode,
		expires,
		config.Keys.JwtPrivateKey,
		config.SetRedisString,
	)
	if err != nil || newJwtToken == nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	token = newToken

	// Send code to email or phone number
	if userFound == nil || userFound.School == nil || userFound.School.Config == nil {
		return
	}
	if utils.IsEmailValid(userFound.Email) {
		go func() {
			fromEmail, fromUsername := userFound.School.SMTPNoReplySender()
			data := &smtpHelper.EmailDataCheckCode{
				EmailData: smtpHelper.EmailData{
					HomePageLink: fmt.Sprintf("https://%s", userFound.School.Config.DomainName),
					Logo:         userFound.School.Logo,
					Title:        constants.MailUpdatePhoneNumberCheckCode.Title,
					Message:      constants.MailUpdatePhoneNumberCheckCode.Message,
				},
				Code:            fmt.Sprintf("%d", randomCode),
				DurationMinutes: 10,
			}
			body, err := data.LoadTemplate()
			if err != nil {
				return
			}
			err = smtpHelper.SendEmailTo(
				fromEmail,
				fromUsername,
				userFound.Email,
				constants.MailUpdatePhoneNumberCheckCode.Subject,
				body,
			)
			if err != nil {
				return
			}
		}()
	}
	return
}

func (service *Service) UpdateProfilePhoneNumberCheckCode(inputJwtToken *types.JwtToken, inputToken string, inputCode int) (token string, errCode int, err error) {
	// Find user
	userFound, err := service.UserService.Repository.GetByID(inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if userFound == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Check input
	if len(inputToken) <= 0 && inputCode < 10000 {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid token and code! Please enter valid information.")
		return
	}
	if len(inputToken) <= 0 {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid token! Please enter valid information.")
		return
	}
	if inputCode < 10000 {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid code! Please enter valid information.")
		return
	}

	// Extract token information and validate the token
	errMsg := "Invalid or expired token! Please enter valid information."
	jwtToken, err := securityUtil.DecodeJWTToken(inputToken, config.Keys.JwtPublicKey)
	if err != nil {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMsg)
		return
	}
	if jwtToken == nil || jwtToken.UserID <= 0 || jwtToken.Issuer != constants.JwtIssuerProfileUpdatePhoneNumberCode {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMsg)
		return
	}
	isTokenValid := securityUtil.ValidateJWTToken(inputToken, jwtToken, config.GetRedisString)
	if !isTokenValid {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMsg)
		return
	}

	// Check if the code is valid
	if jwtToken.Code < 1 || jwtToken.Code != inputCode {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", "Invalid code! Please enter valid information.")
		return
	}

	// Invalidate token
	_, _ = config.DeleteRedisString(securityUtil.GetJWTCachedKey(jwtToken.UserID, jwtToken.Issuer))

	// Generate new token
	newJwtToken, newToken, err := securityUtil.EncodeJWTToken(
		&types.JwtToken{
			UserID:   userFound.ID,
			SchoolID: userFound.SchoolID,
			Platform: "*",
			Device:   "*",
			App:      "*",
		},
		constants.JwtIssuerProfileUpdatePhoneNumberNewPhoneNumber,
		securityUtil.NewExpiresDateDefault(),
		config.Keys.JwtPrivateKey,
		config.SetRedisString,
	)
	if err != nil || newJwtToken == nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	token = newToken
	return
}

func (service *Service) UpdateProfilePhoneNumberNewPhoneNumber(inputJwtToken *types.JwtToken, token string, phoneNumber uint64) (errCode int, err error) {
	// Find user
	userFound, err := service.UserService.Repository.GetByID(inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if userFound == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Check input
	isPhoneNumberValid := utils.IsPhoneNumberValid(phoneNumber)
	if len(token) <= 0 && !isPhoneNumberValid {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s",
			"Invalid token and phone number! Please enter valid information.",
		)
		return
	}
	if len(token) <= 0 {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid token! Please enter valid information.")
		return
	}
	if !isPhoneNumberValid {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s",
			"Invalid phone number! Please enter valid information.",
		)
		return
	}

	// Extract token information and validate the token
	errMsg := "Invalid or expired token! Please enter valid information."
	jwtTokenDecoded, err := securityUtil.DecodeJWTToken(token, config.Keys.JwtPublicKey)
	if err != nil {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMsg)
		return
	}
	if jwtTokenDecoded == nil || jwtTokenDecoded.UserID <= 0 ||
		jwtTokenDecoded.Issuer != constants.JwtIssuerProfileUpdatePhoneNumberNewPhoneNumber ||
		jwtTokenDecoded.UserID != inputJwtToken.UserID {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMsg)
		return
	}
	isTokenValid := securityUtil.ValidateJWTToken(token, jwtTokenDecoded, config.GetRedisString)
	if !isTokenValid {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMsg)
		return
	}

	// Update user phone number
	userUpdated, err := service.UserService.Repository.UpdatePhoneNumberByID(jwtTokenDecoded.UserID, phoneNumber)
	if err != nil || userUpdated == nil {
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

	// Invalidate token
	_, _ = config.DeleteRedisString(securityUtil.GetJWTCachedKey(jwtTokenDecoded.UserID, jwtTokenDecoded.Issuer))
	return
}

func (service *Service) UpdateProfileMfaEmailInit(inputJwtToken *types.JwtToken) (token string, errCode int, err error) {
	// Find user
	userFound, err := service.UserService.Repository.GetByID(inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if userFound == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Generate new random code
	randomCode, err := utils.GenerateRandomCode(6)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	expires := securityUtil.NewExpiresDateDefault()
	newJwtToken, newToken, err := securityUtil.EncodeJWTToken(
		&types.JwtToken{
			UserID:   userFound.ID,
			SchoolID: userFound.SchoolID,
			Platform: "*",
			Device:   "*",
			App:      "*",
			Code:     randomCode,
		},
		constants.JwtIssuerProfileUpdateMfaEmailCode,
		expires,
		config.Keys.JwtPrivateKey,
		config.SetRedisString,
	)
	if err != nil || newJwtToken == nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	token = newToken

	// Send code to email
	if userFound == nil || userFound.School == nil || userFound.School.Config == nil {
		return
	}
	if utils.IsEmailValid(userFound.Email) {
		go func() {
			fromEmail, fromUsername := userFound.School.SMTPNoReplySender()
			data := &smtpHelper.EmailDataCheckCode{
				EmailData: smtpHelper.EmailData{
					HomePageLink: fmt.Sprintf("https://%s", userFound.School.Config.DomainName),
					Logo:         userFound.School.Logo,
					Title:        constants.MailUpdateMfaEmailCheckCode.Title,
					Message:      constants.MailUpdateMfaEmailCheckCode.Message,
				},
				Code:            fmt.Sprintf("%d", randomCode),
				DurationMinutes: 10,
			}
			body, err := data.LoadTemplate()
			if err != nil {
				return
			}
			err = smtpHelper.SendEmailTo(
				fromEmail,
				fromUsername,
				userFound.Email,
				constants.MailUpdateMfaEmailCheckCode.Subject,
				body,
			)
			if err != nil {
				return
			}
		}()
	}
	return
}

func (service *Service) UpdateProfileMfaEmailCheckCode(inputJwtToken *types.JwtToken, inputToken string, inputCode int) (errCode int, err error) {
	// Find user
	userFound, err := service.UserService.Repository.GetByID(inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if userFound == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Check input
	if len(inputToken) <= 0 && inputCode < 10000 {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid token and code! Please enter valid information.")
		return
	}
	if len(inputToken) <= 0 {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid token! Please enter valid information.")
		return
	}
	if inputCode < 10000 {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid code! Please enter valid information.")
		return
	}

	// Extract token information and validate the token
	errMsg := "Invalid or expired token! Please enter valid information."
	jwtToken, err := securityUtil.DecodeJWTToken(inputToken, config.Keys.JwtPublicKey)
	if err != nil {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMsg)
		return
	}
	if jwtToken == nil || jwtToken.UserID <= 0 || jwtToken.Issuer != constants.JwtIssuerProfileUpdateMfaEmailCode {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMsg)
		return
	}
	isTokenValid := securityUtil.ValidateJWTToken(inputToken, jwtToken, config.GetRedisString)
	if !isTokenValid {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMsg)
		return
	}

	// Check if the code is valid
	if jwtToken.Code < 1 || jwtToken.Code != inputCode {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", "Invalid code! Please enter valid information.")
		return
	}

	// Invalidate token
	_, _ = config.DeleteRedisString(securityUtil.GetJWTCachedKey(jwtToken.UserID, jwtToken.Issuer))

	// Toggle Mfa settings
	mfaUpdated, err := service.UserService.Repository.UpdateUserConfigFieldByID(userFound.UserConfigID, "mfa_email", !userFound.Config.MfaEmail)
	if err != nil || mfaUpdated == nil {
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

func (service *Service) UpdateProfileNotification(inputJwtToken *types.JwtToken, enabled bool) (result *model.User, errCode int, err error) {
	// Check if user exists
	userFound, err := service.UserService.Repository.GetByID(inputJwtToken.UserID)
	if err != nil || userFound == nil {
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", "User not found! Please enter valid information.")
		return
	}

	// Update user notification
	tempConfig, err := service.UserService.Repository.UpdateUserConfigAllowNotificationByID(userFound.UserConfigID, enabled)
	if err != nil || tempConfig == nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	result = userFound
	result.Config = tempConfig
	return
}

func (service *Service) UpdateProfileWebPushSubscription(inputJwtToken *types.JwtToken, subscription *data.UpdateProfileWebPushSubscriptionRequest) (errCode int, err error) {
	// Find user
	userFound, err := service.UserService.Repository.GetByID(inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if userFound == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	// Update user web push subscription
	result, err := service.UserService.Repository.UpdateUserConfigWebPushSubscriptionByID(inputJwtToken.UserID, subscription.Endpoint, subscription.Keys.P256dh, subscription.Keys.Auth)
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

func (service *Service) GetProfile(inputJwtToken *types.JwtToken) (result *model.User, errCode int, err error) {
	result, err = service.UserService.Repository.GetByID(inputJwtToken.UserID)
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

func (service *Service) GetWebPushSubscriptionPublicKey(inputJwtToken *types.JwtToken) (result string, errCode int, err error) {
	// Find user
	userFound, err := service.UserService.Repository.GetByID(inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if userFound == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(MODEL_NAME)
		return
	}

	result = config.Env.WebPushVapidPublicKey
	if len(result) < 1 {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}
