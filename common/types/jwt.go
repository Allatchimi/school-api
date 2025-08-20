package types

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtToken struct {
	jwt.RegisteredClaims
	UserID   int64  `json:"userID"`
	SchoolID int64  `json:"schoolID"`
	Platform string `json:"platform"`
	Device   string `json:"device"`
	App      string `json:"app"`
	Code     string `json:"code"`
}

// GoogleUserProfileResponse Represents claims for a Google user.
//
// Refer to the official Google documentation for more details
// https://developers.google.com/identity/openid-connect/openid-connect
type GoogleUserProfileResponse struct {
	ID            string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	FirstName     string `json:"given_name"`
	LastName      string `json:"family_name"`
	FullName      string `json:"name"`
	Language      string `json:"locale"`
	Picture       string `json:"picture"`
	Expires       int64  `json:"expires"`  // Unix time
	IssuedAt      int64  `json:"issuedAt"` // Unix time
}

type FacebookAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

type FacebookDebugAccessTokenResponse struct {
	Data struct {
		UserID            string   `json:"user_id"`
		AppID             string   `json:"app_id"`
		Type              string   `json:"type"`
		Application       string   `json:"application"`
		DataAccessExpires int64    `json:"data_access_expires_at"`
		Expires           int64    `json:"expires_at"` // Unix time
		IsValid           bool     `json:"is_valid"`
		Scopes            []string `json:"scopes"`
	} `json:"data"`
}

type FacebookUserProfileResponse struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	FullName     string `json:"name"`
	Languages    string `json:"languages"`
	PictureSmall *struct {
		FacebookUserPictureResponse
	} `json:"picture_small"`
	PictureLarge *struct {
		FacebookUserPictureResponse
	} `json:"picture_large"`
	DataAccessExpires int64 `json:"dataAccessExpires"` // Unix time
	Expires           int64 `json:"expires"`           // Unix time
}

type FacebookUserPictureResponse struct {
	Data struct {
		Width        int    `json:"width"`
		Height       int    `json:"height"`
		IsSilhouette bool   `json:"is_silhouette"`
		Url          string `json:"url"`
	} `json:"data"`
}
