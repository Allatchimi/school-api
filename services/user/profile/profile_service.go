package profile

import (
	"fmt"
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/common/utils"
	"api/common/utils/mail"
	"api/common/utils/security"
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

	result, err = service.UserService.Repository.UpdateUserInfoByID(inputJwtToken.UserID, item)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}

func (service *Service) UpdateProfilePhoneNumber(inputJwtToken *types.JwtToken, phoneNumber uint64) (result *model.User, errCode int, err error) { // Check if user exists
	// Check if this phone number is already taken
	foundUser, err := service.UserService.Repository.GetByPhoneNumber(phoneNumber)
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
	result, err = service.UserService.Repository.UpdatePhoneNumberByID(inputJwtToken.UserID, phoneNumber)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}
	return
}

func (service *Service) UpdateProfilePasswordInit(inputJwtToken *types.JwtToken) (token string, errCode int, err error) {
	// Check if user exists
	var userFound *model.User
	userFound, err = service.UserService.Repository.GetByID(inputJwtToken.UserID)
	if err != nil || userFound.ID <= 0 {
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
	expires := security.NewExpiresDateDefault()
	newJwtToken, newToken, err := security.EncodeJWTToken(
		&types.JwtToken{
			UserID:   userFound.ID,
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
	if utils.IsEmailValid(userFound.Email) {
		go func() {
			err := mail.SendMail(
				fmt.Sprintf("%s - You have requested new password", config.Env.AppName),
				fmt.Sprintf("The code to set new password is %d", randomCode),
				userFound.Email,
			)
			if err != nil {
				return
			}
		}()
	}
	return
}

func (service *Service) UpdateProfilePasswordCheckCode(inputJwtToken *types.JwtToken, inputToken string, inputCode int) (token string, errCode int, err error) {
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
	jwtToken, err := security.DecodeJWTToken(inputToken, config.Keys.JwtPublicKey)
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
	isTokenValid := security.ValidateJWTToken(inputToken, jwtToken, config.GetRedisString)
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

	// Check if user exists
	userFound, err := service.UserService.Repository.GetByID(jwtToken.UserID)
	if err != nil || userFound == nil {
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", "User not found! Please enter valid information.")
		return
	}

	// Invalidate token
	_, _ = config.DeleteRedisString(security.GetJWTCachedKey(jwtToken.UserID, jwtToken.Issuer))

	// Generate new token
	newJwtToken, newToken, err := security.EncodeJWTToken(
		&types.JwtToken{
			UserID:   userFound.ID,
			Platform: "*",
			Device:   "*",
			App:      "*",
		},
		constants.JwtIssuerProfileUpdatePasswordNewPassword,
		security.NewExpiresDateDefault(),
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
	jwtTokenDecoded, err := security.DecodeJWTToken(token, config.Keys.JwtPublicKey)
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
	isTokenValid := security.ValidateJWTToken(token, jwtTokenDecoded, config.GetRedisString)
	if !isTokenValid {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMsg)
		return
	}

	// Check if user exists
	userFound, err := service.UserService.Repository.GetByID(jwtTokenDecoded.UserID)
	if err != nil || userFound == nil {
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", "User not found! Please enter valid information.")
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
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Invalidate token
	_, _ = config.DeleteRedisString(security.GetJWTCachedKey(jwtTokenDecoded.UserID, jwtTokenDecoded.Issuer))

	// Send alert message to email
	if utils.IsEmailValid(userUpdated.Email) {
		go func() {
			err := mail.SendMail(
				fmt.Sprintf("%s - You have changed your password", config.Env.AppName),
				"Your password has been changed successfully.",
				userUpdated.Email,
			)
			if err != nil {
				return
			}
		}()
	}
	return
}

func (service *Service) UpdateProfilePhoneNumberInit(inputJwtToken *types.JwtToken) (token string, errCode int, err error) {
	// Check if user exists
	var userFound *model.User
	userFound, err = service.UserService.Repository.GetByID(inputJwtToken.UserID)
	if err != nil || userFound.ID <= 0 {
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
	expires := security.NewExpiresDateDefault()
	newJwtToken, newToken, err := security.EncodeJWTToken(
		&types.JwtToken{
			UserID:   userFound.ID,
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
	if utils.IsEmailValid(userFound.Email) {
		go func() {
			err := mail.SendMail(
				fmt.Sprintf("%s - You have requested new phone number", config.Env.AppName),
				fmt.Sprintf("The code to set new phone number is %d", randomCode),
				userFound.Email,
			)
			if err != nil {
				return
			}
		}()
	}
	return
}

func (service *Service) UpdateProfilePhoneNumberCheckCode(inputJwtToken *types.JwtToken, inputToken string, inputCode int) (token string, errCode int, err error) {
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
	jwtToken, err := security.DecodeJWTToken(inputToken, config.Keys.JwtPublicKey)
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
	isTokenValid := security.ValidateJWTToken(inputToken, jwtToken, config.GetRedisString)
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

	// Check if user exists
	userFound, err := service.UserService.Repository.GetByID(jwtToken.UserID)
	if err != nil || userFound == nil {
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", "User not found! Please enter valid information.")
		return
	}

	// Invalidate token
	_, _ = config.DeleteRedisString(security.GetJWTCachedKey(jwtToken.UserID, jwtToken.Issuer))

	// Generate new token
	newJwtToken, newToken, err := security.EncodeJWTToken(
		&types.JwtToken{
			UserID:   userFound.ID,
			Platform: "*",
			Device:   "*",
			App:      "*",
		},
		constants.JwtIssuerProfileUpdatePhoneNumberNewPhoneNumber,
		security.NewExpiresDateDefault(),
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
	jwtTokenDecoded, err := security.DecodeJWTToken(token, config.Keys.JwtPublicKey)
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
	isTokenValid := security.ValidateJWTToken(token, jwtTokenDecoded, config.GetRedisString)
	if !isTokenValid {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMsg)
		return
	}

	// Check if user exists
	userFound, err := service.UserService.Repository.GetByID(jwtTokenDecoded.UserID)
	if err != nil || userFound == nil {
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", "User not found! Please enter valid information.")
		return
	}

	// Update user phone number
	userUpdated, err := service.UserService.Repository.UpdatePhoneNumberByID(jwtTokenDecoded.UserID, phoneNumber)
	if err != nil || userUpdated == nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Invalidate token
	_, _ = config.DeleteRedisString(security.GetJWTCachedKey(jwtTokenDecoded.UserID, jwtTokenDecoded.Issuer))

	// Send alert message to email
	if utils.IsEmailValid(userUpdated.Email) {
		go func() {
			err := mail.SendMail(
				fmt.Sprintf("%s - You have changed your phone number", config.Env.AppName),
				"Your phone number has been changed successfully.",
				userUpdated.Email,
			)
			if err != nil {
				return
			}
		}()
	}
	return
}

func (service *Service) UpdateProfileMfaEmailInit(inputJwtToken *types.JwtToken) (token string, errCode int, err error) {
	// Check if user exists
	var userFound *model.User
	userFound, err = service.UserService.Repository.GetByID(inputJwtToken.UserID)
	if err != nil || userFound.ID <= 0 {
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
	expires := security.NewExpiresDateDefault()
	newJwtToken, newToken, err := security.EncodeJWTToken(
		&types.JwtToken{
			UserID:   userFound.ID,
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

	// Send code to email or phone number
	if utils.IsEmailValid(userFound.Email) {
		go func() {
			err := mail.SendMail(
				fmt.Sprintf("%s - You have requested to change Mfa settings for email", config.Env.AppName),
				fmt.Sprintf("The code to change Mfa settings for email is %d", randomCode),
				userFound.Email,
			)
			if err != nil {
				return
			}
		}()
	}
	return
}

func (service *Service) UpdateProfileMfaEmailCheckCode(inputJwtToken *types.JwtToken, inputToken string, inputCode int) (errCode int, err error) {
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
	jwtToken, err := security.DecodeJWTToken(inputToken, config.Keys.JwtPublicKey)
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
	isTokenValid := security.ValidateJWTToken(inputToken, jwtToken, config.GetRedisString)
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

	// Check if user exists
	userFound, err := service.UserService.Repository.GetByID(jwtToken.UserID)
	if err != nil || userFound == nil {
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", "User not found! Please enter valid information.")
		return
	}

	// Invalidate token
	_, _ = config.DeleteRedisString(security.GetJWTCachedKey(jwtToken.UserID, jwtToken.Issuer))

	// Toggle Mfa settings
	mfaUpdated, err := service.UserService.Repository.UpdateUserConfigFieldByID(userFound.UserConfigID, "mfa_email", !userFound.Config.MfaEmail)
	if err != nil || mfaUpdated == nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Send alert message to email
	if utils.IsEmailValid(userFound.Email) {
		go func() {
			err := mail.SendMail(
				fmt.Sprintf("%s - You have changed your Mfa settings for email", config.Env.AppName),
				"Your Mfa settings for email has been changed successfully.",
				userFound.Email,
			)
			if err != nil {
				return
			}
		}()
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
	result = config.Env.WebPushVapidPublicKey
	if len(result) < 1 {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}
