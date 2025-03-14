package profile

import (
	"api/common/utils/mail"
	"api/common/utils/security"
	"api/common/utils/sms"
	"api/config"
	"fmt"
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/common/utils"
	"api/services/user/user"
	"api/services/user/user/model"
)

type Service struct {
	Repository *user.Repository
}

func NewService(repository *user.Repository) *Service {
	return &Service{Repository: repository}
}

// UpdateProfileEmail Updates email
func (service *Service) UpdateProfileEmail(inputJwtToken *types.JwtToken, email string) (result *model.User, errCode int, err error) { // Check if user exists
	// Check if this email is already taken
	foundUser, err := service.Repository.GetByEmail(email)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get user by email from database")
		return
	}
	if foundUser != nil && foundUser.Email == email {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("user email")
		return
	}

	// Update
	result, err = service.Repository.UpdateEmail(inputJwtToken.UserID, email)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update user from database")
	}
	return
}

// UpdateProfilePhoneNumber Updates phone number
func (service *Service) UpdateProfilePhoneNumber(inputJwtToken *types.JwtToken, phoneNumber uint64) (result *model.User, errCode int, err error) { // Check if user exists
	// Check if this phone number is already taken
	foundUser, err := service.Repository.GetByPhoneNumber(phoneNumber)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get user by phone number from database")
		return
	}
	if foundUser != nil && foundUser.PhoneNumber == phoneNumber {
		errCode = http.StatusFound
		err = constants.Http302ErrorMessage("user phone number")
		return
	}

	// Update
	result, err = service.Repository.UpdatePhoneNumber(inputJwtToken.UserID, phoneNumber)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update user from database")
	}
	return
}

// UpdateProfilePasswordInit Updates password step 1 - initialize the process
func (service *Service) UpdateProfilePasswordInit(inputJwtToken *types.JwtToken) (token string, errCode int, err error) {
	// Check if user exists
	var userFound *model.User
	userFound, err = service.Repository.GetByID(inputJwtToken.UserID)
	if err != nil || userFound.ID <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("User not found")
		return
	}

	// Generate new random code
	randomCode, err := utils.GenerateRandomCode(6)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("generate random code")
		return
	}
	expires := security.NewExpiresDateDefault()
	var tempRoleID int64 = userFound.RoleID
	newJwtToken, newToken, err := security.EncodeJWTToken(
		&types.JwtToken{
			UserID:   userFound.ID,
			RoleID:   tempRoleID,
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
		err = constants.Http500ErrorMessage("generate new JWT token")
		return
	}
	token = newToken

	// Send code to email or phone number
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
	} else {
		go func() {
			err := sms.SendSMS(
				fmt.Sprintf("The code to reset your password is %d.", randomCode),
				fmt.Sprintf("+%d", userFound.PhoneNumber),
			)
			if err != nil {
				return
			}
		}()
	}
	return
}

// UpdateProfilePasswordCheckCode Updates password step 2 - check the code
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
	if jwtToken.Code <= 0 || jwtToken.Code != inputCode {
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

	// Invalidate token
	_, _ = config.DeleteRedisString(security.GetJWTCachedKey(jwtToken.UserID, jwtToken.Issuer))

	// Generate new token
	var tempRoleID int64 = userFound.RoleID
	newJwtToken, newToken, err := security.EncodeJWTToken(
		&types.JwtToken{
			UserID:   userFound.ID,
			RoleID:   tempRoleID,
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
		err = constants.Http500ErrorMessage("generate new JWT token")
	}
	token = newToken
	return
}

// UpdateProfilePasswordNewPassword Updates password step 3 - set new password
func (service *Service) UpdateProfilePasswordNewPassword(inputJwtToken *types.JwtToken, token string, password string) (errCode int, err error) {
	// Check input
	isPasswordValid, missingPasswordChars := utils.IsPasswordValid(password)
	if len(token) <= 0 && !isPasswordValid {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s %s",
			"Invalid token and password! Password missing",
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
			"Invalid password! Password missing",
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
		jwtTokenDecoded.UserID != inputJwtToken.UserID || jwtTokenDecoded.RoleID != inputJwtToken.RoleID {
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
	userFound, err := service.Repository.GetByID(jwtTokenDecoded.UserID)
	if err != nil || userFound == nil {
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", "User not found! Please enter valid information.")
		return
	}

	// Update user password
	userUpdated, err := service.Repository.UpdateUserPassword(jwtTokenDecoded.UserID, password)
	if err != nil || userUpdated == nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update password")
		return
	}

	// Invalidate token
	_, _ = config.DeleteRedisString(security.GetJWTCachedKey(jwtTokenDecoded.UserID, jwtTokenDecoded.Issuer))
	return
}

// UpdateProfileInfo Update profile info
func (service *Service) UpdateProfileInfo(inputJwtToken *types.JwtToken, userInfo *model.UserInfo) (result *model.UserInfo, errCode int, err error) {
	result, err = service.Repository.UpdateProfileInfo(inputJwtToken.UserID, userInfo)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update profile info from database")
	}
	return
}

// UpdateProfileMfa Update profile MFA
func (service *Service) UpdateProfileMfa(inputJwtToken *types.JwtToken, method string, value bool) (result *model.UserMfa, errCode int, err error) {
	if !utils.IsMfaMethodValid(method) {
		errCode = http.StatusUnprocessableEntity
		err = fmt.Errorf("%s", "Invalid MFA method! Please enter valid method.")
		return
	}
	result, err = service.Repository.UpdateProfileMfa(inputJwtToken.UserID, method, value)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("update profile MFA from database")
	}
	return
}

// DeleteProfile Delete user account
func (service *Service) DeleteProfile(inputJwtToken *types.JwtToken) (affectedRows int64, errCode int, err error) {
	affectedRows, err = service.Repository.Delete(inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("delete account from database")
		return
	}
	if affectedRows <= 0 {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("User")
		return
	}
	return
}

// GetProfile Return profile information
func (service *Service) GetProfile(inputJwtToken *types.JwtToken) (result *model.User, errCode int, err error) {
	result, err = service.Repository.GetByID(inputJwtToken.UserID)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.Http500ErrorMessage("get profile from database")
		return
	}
	if result == nil {
		errCode = http.StatusNotFound
		err = constants.Http404ErrorMessage("User")
		return
	}
	return
}
