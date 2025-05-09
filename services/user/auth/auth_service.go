package auth

import (
	"fmt"
	"net/http"
	"slices"
	"time"

	"api/common/constants"
	"api/common/types"
	"api/common/utils"
	"api/common/utils/auth"
	"api/common/utils/mail"
	"api/common/utils/security"
	"api/common/utils/sms"
	"api/config"
	"api/services/user/auth/data"
	"api/services/user/role"
	modelRole "api/services/user/role/model"
	"api/services/user/user"
	"api/services/user/user/model"
)

type Service struct {
	Repository     *user.Repository
	RoleRepository *role.Repository
}

func NewAuthService(repository *user.Repository, roleRepository *role.Repository) *Service {
	return &Service{Repository: repository, RoleRepository: roleRepository}
}

const MODEL_NAME = "user"
const DEFAULT_ERROR_MESSAGE = "interact with auth service"

func (service *Service) Login(input *data.LoginRequest, device *data.LoginDevice) (accessToken string, accessExpires *time.Time, activateAccountToken string, errCode int, err error) {
	// Check if user exists
	var userFound *model.User
	var errMsg string
	if utils.IsEmailValid(input.Email) {
		userFound, err = service.Repository.GetByEmail(input.Email)
		errMsg = "Invalid email or password! Please enter valid information."
	} else {
		errMsg = "Invalid phone number or password! Please enter valid information."
		userFound, err = service.Repository.GetByPhoneNumber(input.PhoneNumber)
	}
	if err != nil || userFound == nil || userFound.Email != input.Email {
		errCode = http.StatusNotFound
		err = fmt.Errorf("%s", errMsg)
		return
	}
	isPasswordMatches, err := security.CompareArgon2id(input.Password, userFound.Password)
	if err != nil || !isPasswordMatches {
		errCode = http.StatusNotFound
		err = fmt.Errorf("%s", errMsg)
		return
	}

	// Check if account is activated
	if userFound.IsActivated {
		// Generate new token
		var accessJwtToken *types.JwtToken
		accessJwtToken, accessToken, err = security.EncodeJWTToken(
			&types.JwtToken{
				UserID:   userFound.ID,
				RoleID:   userFound.RoleID,
				Platform: device.Platform,
				Device:   device.DeviceName,
				App:      device.App,
			},
			constants.JwtIssuerSession,
			security.NewExpiresDateLogin(input.StayConnected),
			config.Keys.JwtPrivateKey,
			config.AppendToRedisStringList,
		)
		if err != nil || accessJwtToken == nil || len(accessToken) <= 0 {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}

		accessExpires = &accessJwtToken.ExpiresAt.Time
		return
	}

	// For non activated user account, generate new random code and token with
	// issuer JWT_ISSUER_AUTH_ACTIVATE and send code to email or phone number
	randomCode := 0
	randomCode, err = utils.GenerateRandomCode(6)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Generate new token
	var activateAccountJwtToken *types.JwtToken
	activateAccountJwtToken, activateAccountToken, err = security.EncodeJWTToken(
		&types.JwtToken{
			UserID:   userFound.ID,
			RoleID:   userFound.RoleID,
			Platform: "*",
			Device:   "*",
			App:      "*",
			Code:     randomCode,
		},
		constants.JwtIssuerAuthActivate,
		security.NewExpiresDateDefault(),
		config.Keys.JwtPrivateKey,
		config.SetRedisString,
	)
	if err != nil || activateAccountJwtToken == nil || len(activateAccountToken) <= 0 {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	errCode = http.StatusForbidden
	err = fmt.Errorf("%s", "Account found but not activated! Please activate your account to start using your services.")

	// Send code to email or phone number
	if utils.IsEmailValid(input.Email) {
		go func() {
			err := mail.SendMail(
				fmt.Sprintf("%s - Activate your account", config.Env.AppName),
				fmt.Sprintf("The code to activate your account is %d", randomCode),
				input.Email,
			)
			if err != nil {
				return
			}
		}()
	} else {
		go func() {
			err := sms.SendSMS(
				fmt.Sprintf("The code to activate your account is %d.", randomCode),
				fmt.Sprintf("+%d", input.PhoneNumber),
			)
			if err != nil {
				return
			}
		}()
	}
	return
}

func (service *Service) LoginWithProvider(input *data.LoginWithProviderRequest, device *data.LoginDevice) (accessToken string, accessExpires *time.Time, errCode int, err error) {
	// Validate provider token and update user
	var newUser = &model.User{
		Provider: input.Provider,
		Info:     &model.UserInfo{},
		Mfa:      &model.UserMfa{},
	}
	var expires int64 = 0
	if input.Provider == constants.AuthProviderGoogle {
		googleUser, errGoogleUser := auth.VerifyGoogleIDToken(input.Token)
		if errGoogleUser != nil || googleUser == nil || len(googleUser.ID) <= 0 {
			errCode = http.StatusUnprocessableEntity
			err = fmt.Errorf("%s", "Invalid provider or token! Please enter valid information.")
			return
		}
		if googleUser.Expires <= time.Now().Unix() {
			errCode = http.StatusUnprocessableEntity
			err = fmt.Errorf("%s", "Token already expired! Please enter valid information.")
			return
		}
		expires = googleUser.Expires
		newUser.FromGoogleUser(googleUser)
	} else if input.Provider == constants.AuthProviderFacebook {
		facebookUser, errFacebookUser := auth.VerifyFacebookToken(input.Token)
		if errFacebookUser != nil || facebookUser == nil || len(facebookUser.ID) <= 0 {
			errCode = http.StatusUnprocessableEntity
			err = fmt.Errorf("%s", "Invalid provider or token! Please enter valid information.")
			return
		}
		if facebookUser.Expires <= time.Now().Unix() {
			errCode = http.StatusUnprocessableEntity
			err = fmt.Errorf("%s", "Token already expired! Please enter valid information.")
			return
		}
		expires = facebookUser.Expires
		newUser.FromFacebookUser(facebookUser)
	} else {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", "Invalid provider or token! Please enter valid information.")
		return
	}

	// Save user if it's not in database
	userFound, err := service.Repository.GetByProvider(input.Provider, newUser.ProviderUserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if userFound == nil || userFound.ID < 1 {
		// Add info
		var userInfo *model.UserInfo
		userInfo, err = service.Repository.CreateUserInfo(newUser.Info)
		if err != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
		// Add mfa
		var userMfa *model.UserMfa
		userMfa, err = service.Repository.CreateUserMfa(newUser.Mfa)
		if err != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}

		// Get the default role
		var defaultRole *modelRole.Role
		defaultRole, err = service.RoleRepository.GetByName(config.Env.RoleDefault)
		if err != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}

		// Create user
		tmpActivatedAt := time.Now()
		userFound, err = service.Repository.Create(
			&model.User{
				Email:          newUser.Email,
				Provider:       input.Provider,
				ProviderUserID: newUser.ProviderUserID,
				LoginMethod:    constants.AuthLoginMethodProvider,
				RoleID:         defaultRole.ID,
				IsActivated:    true,
				ActivatedAt:    &tmpActivatedAt,
				UserInfoID:     userInfo.ID,
				UserMfaID:      userMfa.ID,
			},
		)
		if err != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
	}

	// Generate new token
	expiresTime := time.Unix(expires, 0)
	jwtToken, accessToken, err := security.EncodeJWTToken(
		&types.JwtToken{
			UserID:   userFound.ID,
			RoleID:   userFound.RoleID,
			Platform: device.Platform,
			Device:   device.DeviceName,
			App:      device.App,
		},
		constants.JwtIssuerSession,
		&expiresTime,
		config.Keys.JwtPrivateKey,
		config.AppendToRedisStringList,
	)
	if err != nil || jwtToken == nil || len(accessToken) <= 0 {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	accessExpires = &jwtToken.ExpiresAt.Time
	return
}

func (service *Service) Register(input *data.RegisterRequest) (activateAccountToken string, errCode int, err error) {
	// Check if user exists
	var userFound *model.User
	var errMsg string
	if utils.IsEmailValid(input.Email) {
		userFound, err = service.Repository.GetByEmail(input.Email)
		errMsg = "user email"
	} else {
		errMsg = "user phone number"
		userFound, err = service.Repository.GetByPhoneNumber(input.PhoneNumber)
	}
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if userFound != nil && userFound.Email == input.Email {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(errMsg)
		return
	}

	// Get the default role
	var defaultRole *modelRole.Role
	defaultRole, err = service.RoleRepository.GetByName(config.Env.RoleDefault)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	// Create new user
	userFound.Email = input.Email
	userFound.PhoneNumber = input.PhoneNumber
	userFound.Password = input.Password
	userFound.LoginMethod = constants.AuthLoginMethodDefault
	userFound.RoleID = defaultRole.ID
	createdUser, err := service.Repository.Create(userFound)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Since the new user account is not activated, we generate code with
	// issuer JWT_ISSUER_AUTH_ACTIVATE and send code to email or phone number
	randomCode := 0
	randomCode, err = utils.GenerateRandomCode(6)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Generate new token
	var activateAccountJwtToken *types.JwtToken
	activateAccountJwtToken, activateAccountToken, err = security.EncodeJWTToken(
		&types.JwtToken{
			UserID:   createdUser.ID,
			RoleID:   userFound.RoleID,
			Platform: "*",
			Device:   "*",
			App:      "*",
			Code:     randomCode,
		},
		constants.JwtIssuerAuthActivate,
		security.NewExpiresDateDefault(),
		config.Keys.JwtPrivateKey,
		config.SetRedisString,
	)
	if err != nil || activateAccountJwtToken == nil || len(activateAccountToken) <= 0 {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}

	// Send code to email or phone number
	if utils.IsEmailValid(input.Email) {
		go func() {
			err := mail.SendMail(
				fmt.Sprintf("%s - Activate your account", config.Env.AppName),
				fmt.Sprintf("The code to activate your account is %d", randomCode),
				input.Email,
			)
			if err != nil {
				return
			}
		}()
	} else {
		go func() {
			err := sms.SendSMS(
				fmt.Sprintf("The code to activate your account is %d.", randomCode),
				fmt.Sprintf("+%d", input.PhoneNumber),
			)
			if err != nil {
				return
			}
		}()
	}
	return
}

func (service *Service) ActivateAccount(input *data.ActivateAccountRequest) (activatedAt *time.Time, errCode int, err error) {
	// Extract token information and validate the token
	errMsg := "Invalid or expired token! Please enter valid information."
	jwtToken, err := security.DecodeJWTToken(input.Token, config.Keys.JwtPublicKey)
	if err != nil {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMsg)
		return
	}
	if jwtToken == nil || jwtToken.UserID <= 0 || jwtToken.Issuer != constants.JwtIssuerAuthActivate {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMsg)
		return
	}
	isTokenValid := security.ValidateJWTToken(input.Token, jwtToken, config.GetRedisString)
	if !isTokenValid {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMsg)
		return
	}

	// Check if code is valid
	if jwtToken.Code <= 0 || jwtToken.Code != input.Code {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", "Invalid code! Please enter valid information.")
		return
	}

	// Check if account is activated
	userFound, err := service.Repository.GetByID(jwtToken.UserID)
	if err != nil || userFound == nil {
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", "User not found! Please enter valid information.")
		return
	}
	if userFound.IsActivated {
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", "User account is already activated! Please sign in and start using our services.")
		return
	}

	// Create user info and MFA
	newUserInfo, err := service.Repository.CreateUserInfo(&model.UserInfo{})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	newUserMfa, err := service.Repository.CreateUserMfa(&model.UserMfa{})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Update account
	tmpActivatedAt := time.Now()
	userFound.UserInfoID = newUserInfo.ID
	userFound.UserMfaID = newUserMfa.ID
	userFound.ActivatedAt = &tmpActivatedAt
	userFound.IsActivated = true
	updatedUser, err := service.Repository.UpdateUserActivation(userFound.ID, userFound)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	activatedAt = updatedUser.ActivatedAt

	// Invalidate the token
	_, _ = config.DeleteRedisString(security.GetJWTCachedKey(jwtToken.UserID, jwtToken.Issuer))

	// Send welcome message
	if utils.IsEmailValid(updatedUser.Email) {
		go func() {
			err := mail.SendMail(
				fmt.Sprintf("%s - Welcome", config.Env.AppName),
				"Welcome",
				updatedUser.Email,
			)
			if err != nil {
				return
			}
		}()
	}
	return
}

func (service *Service) ForgotPasswordInit(input *data.ForgotPasswordInitRequest) (token string, errCode int, err error) {
	// Check input
	var errMsg string
	var isInputValid bool
	if utils.IsEmailValid(input.Email) {
		errMsg = "email"
		isInputValid = utils.IsEmailValid(input.Email)
	} else {
		errMsg = "phone number"
		isInputValid = utils.IsPhoneNumberValid(input.PhoneNumber)
	}
	if !isInputValid {
		errCode = http.StatusBadRequest
		errMsg = fmt.Sprintf("Invalid %s! Please enter valid information", errMsg)
		err = fmt.Errorf("%s", errMsg)
		return
	}

	// Check if user exists
	var userFound *model.User
	if utils.IsEmailValid(input.Email) {
		errMsg = "User with this email"
		userFound, err = service.Repository.GetByEmail(input.Email)
	} else {
		errMsg = "User with this phone number"
		userFound, err = service.Repository.GetByPhoneNumber(input.PhoneNumber)
	}
	if err != nil || userFound.ID <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage(errMsg)
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
			RoleID:   userFound.RoleID,
			Platform: "*",
			Device:   "*",
			App:      "*",
			Code:     randomCode,
		},
		constants.JwtIssuerAuthForgotPasswordCode,
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
	if utils.IsEmailValid(input.Email) {
		go func() {
			err := mail.SendMail(
				fmt.Sprintf("%s - Forgot your password", config.Env.AppName),
				fmt.Sprintf("The code to reset your password is %d", randomCode),
				input.Email,
			)
			if err != nil {
				return
			}
		}()
	} else {
		go func() {
			err := sms.SendSMS(
				fmt.Sprintf("The code to reset your password is %d.", randomCode),
				fmt.Sprintf("+%d", input.PhoneNumber),
			)
			if err != nil {
				return
			}
		}()
	}
	return
}

func (service *Service) ForgotPasswordCode(input *data.ForgotPasswordCodeRequest) (token string, errCode int, err error) {
	// Check input
	if len(input.Token) <= 0 && input.Code < 10000 {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid token and code! Please enter valid information.")
		return
	}
	if len(input.Token) <= 0 {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid token! Please enter valid information.")
		return
	}
	if input.Code < 10000 {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid code! Please enter valid information.")
		return
	}

	// Extract token information and validate the token
	errMsg := "Invalid or expired token! Please enter valid information."
	jwtToken, err := security.DecodeJWTToken(input.Token, config.Keys.JwtPublicKey)
	if err != nil {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMsg)
		return
	}
	if jwtToken == nil || jwtToken.UserID <= 0 || jwtToken.Issuer != constants.JwtIssuerAuthForgotPasswordCode {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMsg)
		return
	}
	isTokenValid := security.ValidateJWTToken(input.Token, jwtToken, config.GetRedisString)
	if !isTokenValid {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMsg)
		return
	}

	// Check if the code is valid
	if jwtToken.Code <= 0 || jwtToken.Code != input.Code {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", "Invalid code! Please enter valid information.")
		return
	}

	// Check if user exists
	userFound, err := service.Repository.GetByID(jwtToken.UserID)
	if err != nil || userFound == nil {
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", "User not found! Please enter valid information.")
		return
	}

	// Invalidate the token
	_, _ = config.DeleteRedisString(security.GetJWTCachedKey(jwtToken.UserID, jwtToken.Issuer))

	// Generate new token
	newJwtToken, newToken, err := security.EncodeJWTToken(
		&types.JwtToken{
			UserID:   userFound.ID,
			RoleID:   userFound.RoleID,
			Platform: "*",
			Device:   "*",
			App:      "*",
		},
		constants.JwtIssuerAuthForgotPasswordNewPassword,
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

func (service *Service) ForgotPasswordNewPassword(input *data.ForgotPasswordNewPasswordRequest) (errCode int, err error) {
	// Check input
	isPasswordValid, missingPasswordChars := utils.IsPasswordValid(input.NewPassword)
	if len(input.Token) <= 0 && !isPasswordValid {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s %s",
			"Invalid token and password! Password missing",
			missingPasswordChars,
		)
		return
	}
	if len(input.Token) <= 0 {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid token! Please enter valid information.")
		return
	}
	if !isPasswordValid {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s %s",
			"Invalid password! Password missing",
			missingPasswordChars,
		)
		return
	}

	// Extract token information and validate the token
	errMsg := "Invalid or expired token! Please enter valid information."
	jwtToken, err := security.DecodeJWTToken(input.Token, config.Keys.JwtPublicKey)
	if err != nil {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMsg)
		return
	}
	if jwtToken == nil || jwtToken.UserID <= 0 || jwtToken.Issuer != constants.JwtIssuerAuthForgotPasswordNewPassword {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMsg)
		return
	}
	isTokenValid := security.ValidateJWTToken(input.Token, jwtToken, config.GetRedisString)
	if !isTokenValid {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMsg)
		return
	}

	// Check if user exists
	userFound, err := service.Repository.GetByID(jwtToken.UserID)
	if err != nil || userFound == nil {
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", "User not found! Please enter valid information.")
		return
	}

	// Update user password
	userUpdated, err := service.Repository.UpdateUserPassword(jwtToken.UserID, input.NewPassword)
	if err != nil || userUpdated == nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Invalidate the token
	_, _ = config.DeleteRedisString(security.GetJWTCachedKey(jwtToken.UserID, jwtToken.Issuer))
	return
}

func (service *Service) Logout(jwtToken *types.JwtToken, bearerToken string) (errCode int, err error) {
	// Invalidate the token
	sessions, err := config.GetRedisStringList(security.GetJWTCachedKey(jwtToken.UserID, jwtToken.Issuer))
	if err != nil {
		errCode = http.StatusUnauthorized
		err = constants.Http401InvalidTokenErrorMessage()
		return
	}
	tokenIndex := slices.Index(sessions, bearerToken)
	if tokenIndex < 0 {
		errCode = http.StatusUnauthorized
		err = constants.Http401InvalidTokenErrorMessage()
		return
	}
	err = config.RemoveFromRedisStringList(fmt.Sprintf("%d", jwtToken.UserID), int64(tokenIndex))
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}
