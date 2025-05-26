package auth

import (
	"context"
	"fmt"
	"time"

	"api/common/constants"
	"api/common/types"
	"api/config"

	googleIDToken "google.golang.org/api/idtoken"
)

// VerifyGoogleIDToken Verifies a Google ID token and returns the associated user information.
//
// Refer to the official Google documentation for more details on token validation
// https://developers.google.com/identity/openid-connect/openid-connect#discovery
func VerifyGoogleIDToken(token string) (*types.GoogleUserProfileResponse, error) {
	// Validate the token
	if len(token) <= 0 {
		return nil, fmt.Errorf("%s", invalidTokenErrMessage)
	}
	tokenValidator, err := googleIDToken.NewValidator(context.Background())
	if err != nil {
		return nil, constants.Http500ErrorMessage("validate Google token")
	}
	payload, err := tokenValidator.Validate(context.Background(), token, config.Env.GooglePlusClientID)
	if err != nil {
		return nil, fmt.Errorf("%s", invalidTokenErrMessage)
	}
	if payload.Expires <= time.Now().Unix() {
		return nil, fmt.Errorf("%s", invalidTokenErrMessage)
	}

	// Retrieve info from claims
	user := &types.GoogleUserProfileResponse{}
	user.ID, _ = payload.Claims["sub"].(string)
	user.Email, _ = payload.Claims["email"].(string)
	user.EmailVerified, _ = payload.Claims["email_verified"].(bool)
	user.LastName, _ = payload.Claims["family_name"].(string)
	user.FirstName, _ = payload.Claims["given_name"].(string)
	user.FullName, _ = payload.Claims["name"].(string)
	user.Language, _ = payload.Claims["locale"].(string)
	user.Picture, _ = payload.Claims["picture"].(string)
	user.Expires = payload.Expires
	user.IssuedAt = payload.IssuedAt
	return user, nil
}
