package auth

import (
	"fmt"
	"net/http"
	"slices"
	"time"

	"api/common/constants"
	"api/common/helpers"
	authHelper "api/common/helpers/auth"
	smtpHelper "api/common/helpers/message/mail/smtp"
	"api/common/types"
	"api/common/utils"
	securityUtil "api/common/utils/security"
	"api/config"
	"api/services/user/auth/data"
	"api/services/user/role"
	modelRole "api/services/user/role/model"
	"api/services/user/user"
	"api/services/user/user/model"

	"go.uber.org/zap"
)

type Service struct {
	UserService *user.Service
	RoleService *role.Service
}

func NewAuthService(userService *user.Service, roleService *role.Service) *Service {
	return &Service{UserService: userService, RoleService: roleService}
}

const MODEL_NAME = "user"
const DEFAULT_ERROR_MESSAGE = "interact with auth service"

func (service *Service) Login(
	schoolID int64,
	request *data.LoginRequest,
	device *data.LoginDevice,
) (accessToken string, accessExpires *time.Time, activateAccountToken string, errCode int, err error) {
	// Find user
	var userFound *model.User
	var errMsg string
	if utils.IsEmailValid(request.Email) {
		userFound, err = service.UserService.Repository.GetByEmailSchoolID(request.Email, schoolID)
		errMsg = "Invalid email or password! Please enter valid information."
	}
	if err != nil || userFound == nil || userFound.Email != request.Email {
		errCode = http.StatusNotFound
		err = fmt.Errorf("%s", errMsg)
		return
	}
	isPasswordMatches, err := securityUtil.CompareArgon2id(request.Password, userFound.Password)
	if err != nil || !isPasswordMatches {
		errCode = http.StatusNotFound
		err = fmt.Errorf("%s", errMsg)
		return
	}
	// Check if the status is enabled
	if userFound.Status != constants.USER_STATUS_ENABLED {
		errCode = http.StatusUnavailableForLegalReasons
		err = fmt.Errorf("%s", "User account is disabled! Please contact support.")
		return
	}

	// Check if account is activated
	if userFound.IsActivated {
		// Generate new token
		var accessJwtToken *types.JwtToken
		accessJwtToken, accessToken, err = securityUtil.EncodeJWTToken(
			&types.JwtToken{
				UserID:   userFound.ID,
				SchoolID: userFound.SchoolID,

				Platform: device.Platform,
				Device:   device.DeviceName,
				App:      device.App,
			},
			constants.JwtIssuerSession,
			securityUtil.NewExpiresDateLogin(request.StayConnected),
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
	// issuer JWT_ISSUER_AUTH_ACTIVATE and send code to email
	randomCode := 0
	randomCode, err = utils.GenerateRandomCode(6)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Generate new token
	var activateAccountJwtToken *types.JwtToken
	activateAccountJwtToken, activateAccountToken, err = securityUtil.EncodeJWTToken(
		&types.JwtToken{
			UserID:   userFound.ID,
			SchoolID: userFound.SchoolID,

			Platform: "*",
			Device:   "*",
			App:      "*",
			Code:     randomCode,
		},
		constants.JwtIssuerAuthActivate,
		securityUtil.NewExpiresDateDefault(),
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

	// Send code to email
	go func() {
		if userFound == nil || userFound.ID < 1 {
			return
		}
		fromEmail, fromUsername := userFound.School.SMTPNoReplySender()
		data := &smtpHelper.EmailDataCheckCode{
			EmailData: smtpHelper.EmailData{
				HomePageLink: userFound.School.WebsiteUrl(),
				Logo:         userFound.School.LogoUrl(),
				Title:        constants.MailVerifyEmailCheckCode.Title,
				Message:      constants.MailVerifyEmailCheckCode.Message,
			},
			Code:            fmt.Sprintf("%d", randomCode),
			DurationMinutes: 10,
		}
		mailBody, errTemplate := data.LoadTemplate()
		if errTemplate != nil {
			helpers.Logger.Error("Failed to load email template!", zap.Error(errTemplate))
		}
		if len(mailBody) < 1 {
			helpers.Logger.Warn("Empty message body!")
		}
		errMail := smtpHelper.SendEmailTo(
			fromEmail,
			fromUsername,
			userFound.Email,
			constants.MailVerifyEmailCheckCode.Subject,
			mailBody,
		)
		if errMail != nil {
			helpers.Logger.Error("Failed to send mail!", zap.Error(errMail))
		}
	}()
	return
}

func (service *Service) LoginWithProvider(
	schoolID int64,
	request *data.LoginWithProviderRequest,
	device *data.LoginDevice,
) (accessToken string, accessExpires *time.Time, errCode int, err error) {
	// Validate provider token and update user
	isProviderValid := utils.IsAuthProviderValid(request.Provider)
	if !isProviderValid {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid or empty provider! Please enter valid information.")
		return
	}
	var newUser = &model.User{
		SchoolID: schoolID,
		Provider: request.Provider,
		Info:     &model.UserInfo{},
		Config:   &model.UserConfig{},
	}
	var expires int64 = 0
	switch request.Provider {
	case constants.AuthProviderGoogle:
		googleUser, errGoogleUser := authHelper.VerifyGoogleIDToken(request.Token)
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
	case constants.AuthProviderFacebook:
		facebookUser, errFacebookUser := authHelper.VerifyFacebookToken(request.Token)
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
	default:
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", "Invalid provider or token! Please enter valid information.")
		return
	}

	// Save user if it's not in database
	userFound, err := service.UserService.Repository.GetByProviderSchoolID(request.Provider, newUser.ProviderUserID, schoolID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if userFound == nil || userFound.ID < 1 {
		// Add info
		var userInfo *model.UserInfo
		userInfo, err = service.UserService.Repository.CreateUserInfo(newUser.Info)
		if err != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}
		// Add mfa
		var userConfig *model.UserConfig
		userConfig, err = service.UserService.Repository.CreateUserConfig(newUser.Config)
		if err != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}

		// Get the default role
		var defaultRole *modelRole.Role
		defaultRole, err = service.RoleService.Repository.GetByName(config.Env.FixtureRoleDefault)
		if err != nil {
			errCode = http.StatusInternalServerError
			err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
			return
		}

		// Create user
		tmpActivatedAt := time.Now()
		userFound, err = service.UserService.Repository.Create(
			&model.User{
				SchoolID:       schoolID,
				Email:          newUser.Email,
				Status:         constants.USER_STATUS_ENABLED,
				Provider:       request.Provider,
				ProviderUserID: newUser.ProviderUserID,
				LoginMethod:    constants.AuthLoginMethodProvider,
				RoleID:         defaultRole.ID,
				IsActivated:    true,
				ActivatedAt:    &tmpActivatedAt,
				InfoID:         userInfo.ID,
				ConfigID:       userConfig.ID,
			},
		)
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
	}

	// Check if the status is enabled
	if userFound.Status != constants.USER_STATUS_ENABLED {
		errCode = http.StatusUnavailableForLegalReasons
		err = fmt.Errorf("%s", "User account is disabled! Please contact support.")
		return
	}

	// Generate new token
	expiresTime := time.Unix(expires, 0)
	jwtToken, accessToken, err := securityUtil.EncodeJWTToken(
		&types.JwtToken{
			UserID:   userFound.ID,
			SchoolID: userFound.SchoolID,

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

func (service *Service) Register(
	schoolID int64,
	request *data.RegisterRequest,
) (activateAccountToken string, errCode int, err error) {
	// Check inputs
	isEmailValid := utils.IsEmailValid(request.Email)
	isPasswordValid, missingPasswordChars := utils.IsPasswordValid(request.Password)
	if !isEmailValid && !isPasswordValid {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s %s",
			"Invalid email and password! Password missing",
			missingPasswordChars,
		)
		return
	}
	if !isEmailValid {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid email! Please enter valid information.")
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

	// Check if user exists
	var userFound *model.User
	var errMsg string
	if utils.IsEmailValid(request.Email) {
		userFound, err = service.UserService.Repository.GetByEmailSchoolID(request.Email, schoolID)
		errMsg = "user email"
	}
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	if userFound != nil && userFound.Email == request.Email {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage(errMsg)
		return
	}

	// Get the default role
	var defaultRole *modelRole.Role
	defaultRole, err = service.RoleService.Repository.GetByName(config.Env.FixtureRoleDefault)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	// Create new user
	userFound = &model.User{}
	userFound.Email = request.Email
	userFound.Password = request.Password
	userFound.LoginMethod = constants.AuthLoginMethodDefault
	userFound.SchoolID = schoolID
	userFound.RoleID = defaultRole.ID
	userFound.Status = constants.USER_STATUS_ENABLED
	createdUser, err := service.UserService.Repository.Create(userFound)
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

	// Since the new user account is not activated, we generate code with
	// issuer JWT_ISSUER_AUTH_ACTIVATE and send code to email
	randomCode := 0
	randomCode, err = utils.GenerateRandomCode(6)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Generate new token
	var activateAccountJwtToken *types.JwtToken
	activateAccountJwtToken, activateAccountToken, err = securityUtil.EncodeJWTToken(
		&types.JwtToken{
			UserID:   createdUser.ID,
			SchoolID: userFound.SchoolID,

			Platform: "*",
			Device:   "*",
			App:      "*",
			Code:     randomCode,
		},
		constants.JwtIssuerAuthActivate,
		securityUtil.NewExpiresDateDefault(),
		config.Keys.JwtPrivateKey,
		config.SetRedisString,
	)
	if err != nil || activateAccountJwtToken == nil || len(activateAccountToken) <= 0 {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
	}

	// Send code to email
	go func() {
		if createdUser == nil || createdUser.ID < 1 {
			return
		}
		fromEmail, fromUsername := createdUser.School.SMTPNoReplySender()
		data := &smtpHelper.EmailDataCheckCode{
			EmailData: smtpHelper.EmailData{
				HomePageLink: createdUser.School.WebsiteUrl(),
				Logo:         createdUser.School.LogoUrl(),
				Title:        constants.MailVerifyEmailCheckCode.Title,
				Message:      constants.MailVerifyEmailCheckCode.Message,
			},
			Code:            fmt.Sprintf("%d", randomCode),
			DurationMinutes: 10,
		}
		mailBody, errTemplate := data.LoadTemplate()
		if errTemplate != nil {
			helpers.Logger.Error("Failed to load email template!", zap.Error(errTemplate))
		}
		if len(mailBody) < 1 {
			helpers.Logger.Warn("Empty message body!")
		}
		errMail := smtpHelper.SendEmailTo(
			fromEmail,
			fromUsername,
			createdUser.Email,
			constants.MailVerifyEmailCheckCode.Subject,
			mailBody,
		)
		if errMail != nil {
			helpers.Logger.Error("Failed to send mail!", zap.Error(errMail))
		}
	}()
	return
}

func (service *Service) ActivateAccount(
	request *data.ActivateAccountRequest,
) (activatedAt *time.Time, errCode int, err error) {
	// Extract token information and validate the token
	errMsg := "Invalid or expired token! Please enter valid information."
	jwtToken, err := securityUtil.DecodeJWTToken(request.Token, config.Keys.JwtPublicKey)
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
	isTokenValid := securityUtil.ValidateJWTToken(request.Token, jwtToken, config.GetRedisString)
	if !isTokenValid {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMsg)
		return
	}

	// Check if code is valid
	if jwtToken.Code < 1 || jwtToken.Code != request.Code {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", "Invalid code! Please enter valid information.")
		return
	}

	// Check if account is activated
	userFound, err := service.UserService.Repository.GetByID(jwtToken.UserID)
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
	newUserInfo, err := service.UserService.Repository.CreateUserInfo(&model.UserInfo{})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	newUserConfig, err := service.UserService.Repository.CreateUserConfig(&model.UserConfig{})
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}

	// Update account
	tmpActivatedAt := time.Now()
	userFound.InfoID = newUserInfo.ID
	userFound.ConfigID = newUserConfig.ID
	userFound.ActivatedAt = &tmpActivatedAt
	userFound.IsActivated = true
	updatedUser, err := service.UserService.Repository.UpdateActivationByID(userFound.ID, userFound)
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
	activatedAt = updatedUser.ActivatedAt

	// Invalidate the token
	_, _ = config.DeleteRedisString(securityUtil.GetJWTCachedKey(jwtToken.UserID, jwtToken.Issuer))

	// Send welcome message
	go func() {
		if updatedUser == nil || updatedUser.ID < 1 {
			return
		}
		fromEmail, fromUsername := updatedUser.School.SMTPNoReplySender()
		data := &smtpHelper.EmailData{
			HomePageLink: updatedUser.School.WebsiteUrl(),
			Logo:         updatedUser.School.LogoUrl(),
			Title:        constants.MailWelcomeVerifiedEmail.Title,
			Message:      constants.MailWelcomeVerifiedEmail.Message,
		}
		mailBody, errTemplate := data.LoadTemplate()
		if errTemplate != nil {
			helpers.Logger.Error("Failed to load email template!", zap.Error(errTemplate))
		}
		if len(mailBody) < 1 {
			helpers.Logger.Warn("Empty message body!")
		}
		errMail := smtpHelper.SendEmailTo(
			fromEmail,
			fromUsername,
			updatedUser.Email,
			constants.MailWelcomeVerifiedEmail.Subject,
			mailBody,
		)
		if errMail != nil {
			helpers.Logger.Error("Failed to send mail!", zap.Error(errMail))
		}
	}()
	return
}

func (service *Service) ForgotPasswordInit(
	ctxData *types.ContextData,
	request *data.ForgotPasswordInitRequest,
) (token string, errCode int, err error) {
	// Check request
	var errMsg string
	var isInputValid bool
	if utils.IsEmailValid(request.Email) {
		errMsg = "email"
		isInputValid = utils.IsEmailValid(request.Email)
	}
	if !isInputValid {
		errCode = http.StatusBadRequest
		errMsg = fmt.Sprintf("Invalid %s! Please enter valid information", errMsg)
		err = fmt.Errorf("%s", errMsg)
		return
	}

	// Check if user exists
	var userFound *model.User
	if utils.IsEmailValid(request.Email) {
		errMsg = "User with this email"
		userFound, err = service.UserService.Repository.GetByEmailSchoolID(request.Email, ctxData.Jwt.SchoolID)
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

	// Send code to email
	go func() {
		if userFound == nil || userFound.ID < 1 {
			return
		}
		fromEmail, fromUsername := userFound.School.SMTPNoReplySender()
		data := &smtpHelper.EmailDataCheckCode{
			EmailData: smtpHelper.EmailData{
				HomePageLink: userFound.School.WebsiteUrl(),
				Logo:         userFound.School.LogoUrl(),
				Title:        constants.MailForgotPasswordCheckCode.Title,
				Message:      constants.MailForgotPasswordCheckCode.Message,
			},
			Code:            fmt.Sprintf("%d", randomCode),
			DurationMinutes: 10,
		}
		mailBody, errTemplate := data.LoadTemplate()
		if errTemplate != nil {
			helpers.Logger.Error("Failed to load email template!", zap.Error(errTemplate))
		}
		if len(mailBody) < 1 {
			helpers.Logger.Warn("Empty message body!")
		}
		errMail := smtpHelper.SendEmailTo(
			fromEmail,
			fromUsername,
			userFound.Email,
			constants.MailForgotPasswordCheckCode.Subject,
			mailBody,
		)
		if errMail != nil {
			helpers.Logger.Error("Failed to send mail!", zap.Error(errMail))
		}
	}()
	return
}

func (service *Service) ForgotPasswordCode(
	request *data.ForgotPasswordCodeRequest,
) (token string, errCode int, err error) {
	// Check request
	if len(request.Token) < 1 && request.Code < 1 {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid token and code! Please enter valid information.")
		return
	}
	if len(request.Token) < 1 {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid token! Please enter valid information.")
		return
	}
	if request.Code < 10000 {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid code! Please enter valid information.")
		return
	}

	// Extract token information and validate the token
	errMsg := "Invalid or expired token! Please enter valid information."
	jwtToken, err := securityUtil.DecodeJWTToken(request.Token, config.Keys.JwtPublicKey)
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
	isTokenValid := securityUtil.ValidateJWTToken(request.Token, jwtToken, config.GetRedisString)
	if !isTokenValid {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMsg)
		return
	}

	// Check if the code is valid
	if jwtToken.Code < 1 || jwtToken.Code != request.Code {
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

	// Invalidate the token
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
		constants.JwtIssuerAuthForgotPasswordNewPassword,
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

func (service *Service) ForgotPasswordNewPassword(
	request *data.ForgotPasswordNewPasswordRequest,
) (errCode int, err error) {
	// Check request
	if len(request.Token) < 1 {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", "Invalid token! Please enter valid information.")
		return
	}

	// Extract token information and validate the token
	errMsg := "Invalid or expired token! Please enter valid information."
	jwtToken, err := securityUtil.DecodeJWTToken(request.Token, config.Keys.JwtPublicKey)
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
	isTokenValid := securityUtil.ValidateJWTToken(request.Token, jwtToken, config.GetRedisString)
	if !isTokenValid {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", errMsg)
		return
	}

	// Check if user exists
	userFound, err := service.UserService.Repository.GetByID(jwtToken.UserID)
	if err != nil || userFound == nil {
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", "User not found! Please enter valid information.")
		return
	}

	// Update user password
	hasedPassword, err := securityUtil.EncodeArgon2id(request.NewPassword)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	userUpdated, err := service.UserService.Repository.UpdatePasswordByID(jwtToken.UserID, hasedPassword)
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

	// Invalidate the token
	_, _ = config.DeleteRedisString(securityUtil.GetJWTCachedKey(jwtToken.UserID, jwtToken.Issuer))
	return
}

func (service *Service) Logout(
	ctxData *types.ContextData,
	bearerToken string,
) (errCode int, err error) {
	// Invalidate the token
	sessions, err := config.GetRedisStringList(securityUtil.GetJWTCachedKey(ctxData.Jwt.UserID, ctxData.Jwt.Issuer))
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
	err = config.RemoveFromRedisStringList(fmt.Sprintf("%d", ctxData.Jwt.UserID), int64(tokenIndex))
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage(DEFAULT_ERROR_MESSAGE)
		return
	}
	return
}
