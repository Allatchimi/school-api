package securityUtil

import (
	"api/common/constants"
	"api/common/types"
	"api/config"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// NewExpiresDateDefault Returns the default expiration time for JWT. It's 10 min.
func NewExpiresDateDefault() *time.Time {
	tempDate := time.Now().Add(time.Minute * 10) // 10 minutes
	return &tempDate
}

// NewExpiresDateLogin Returns the login expiration time for JWT.
// With default ENV var, the JWT login expiration time is
// 30 days if stayConnected is true, otherwise it's 1 hour.
func NewExpiresDateLogin(stayConnected bool) (date *time.Time) {
	if stayConnected {
		tempDate := time.Now().Add(time.Hour * 24 * 30) // 30 days
		return &tempDate
	}
	tempDate := time.Now().Add(time.Minute * 60) // 1 hours
	return &tempDate
}

// GetJWTCachedKey Returns the key of Redis entry(combining userID and Issuer)
func GetJWTCachedKey(userID int64, issuer string) (key string) {
	return fmt.Sprintf("%d_%s", userID, issuer)
}

// EncodeJWTToken Encodes the JWT token and store it in the cache. This returns a signed string token.
func EncodeJWTToken(jwtToken *types.JwtToken, issuer string, expires *time.Time, privateKey *string, cacheFunc func(string, string) error) (*types.JwtToken, string, error) {
	// Encode token with claims
	jwtToken.Issuer = issuer
	jwtToken.ExpiresAt = jwt.NewNumericDate(*expires)
	jwtToken.IssuedAt = jwt.NewNumericDate(time.Now())
	jwtTokenClaimed := jwt.NewWithClaims(jwt.SigningMethodES512, *jwtToken)

	signedKey, errParse := jwt.ParseECPrivateKeyFromPEM([]byte(*privateKey))
	if errParse != nil {
		return nil, "", errParse
	}

	token, errSigning := jwtTokenClaimed.SignedString(signedKey)
	if errSigning != nil {
		return nil, "", errSigning
	}

	// Cache new token
	errCache := cacheFunc(GetJWTCachedKey(jwtToken.UserID, jwtToken.Issuer), token)
	if errCache != nil {
		return nil, "", errCache
	}
	return jwtToken, token, nil
}

// DecodeJWTToken Decodes the JWT token using the signed string token and public key.
// It also validates to token if this one has not expired
// This returns a JWT token object.
func DecodeJWTToken(token string, publicKey *string) (*types.JwtToken, error) {
	// Parse the token and get claims
	jwtToken, err := jwt.ParseWithClaims(token, &types.JwtToken{}, func(token *jwt.Token) (signedKey interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("%s", fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
		}
		signedKey, err = jwt.ParseECPublicKeyFromPEM([]byte(*publicKey))
		return
	})
	if err != nil {
		return nil, err
	} else if claims, ok := jwtToken.Claims.(*types.JwtToken); ok && jwtToken.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("%s", "Unknown error!")
}

// ValidateJWTToken Validates the token by checking if it is cached.
func ValidateJWTToken(token string, jwtToken *types.JwtToken, loadCachedFunc func(string) (string, error)) bool {
	tokenCached, errCached := loadCachedFunc(GetJWTCachedKey(jwtToken.UserID, jwtToken.Issuer))
	if errCached != nil || len(tokenCached) <= 0 {
		return false
	}
	if token != tokenCached {
		return false
	}
	return true
}

// ValidateAuthToken Validates the token by checking if it is cached.
func ValidateAuthToken(token string) (jwtToken *types.JwtToken, errCode int, err error) {
	if len(token) < 1 {
		errMsg := "Missing or bad authorization header! Please enter valid information."
		return nil, http.StatusUnauthorized, fmt.Errorf("%s", errMsg)
	}
	jwtDecoded, errDecoded := DecodeJWTToken(
		token,
		config.Keys.JwtPublicKey,
	)
	if errDecoded != nil || jwtDecoded == nil {
		return nil, http.StatusUnauthorized, errDecoded
	}

	// Validate the token by checking if it's cached
	isTokenCached := false
	if jwtDecoded.Issuer == constants.JwtIssuerSession {
		isTokenCached = ValidateJWTToken(
			token,
			jwtDecoded,
			config.CheckValueInRedisList(token),
		)
	} else if slices.Contains(constants.JwtIssuerAuthList, jwtDecoded.Issuer) {
		isTokenCached = ValidateJWTToken(
			token,
			jwtDecoded,
			config.GetRedisString,
		)
	}
	if isTokenCached &&
		jwtDecoded.Issuer == constants.JwtIssuerSession &&
		jwtDecoded.UserID > 0 {
		return jwtDecoded, http.StatusOK, nil
	}
	errMsg := "Invalid token! Please enter valid information."
	return nil, http.StatusUnauthorized, fmt.Errorf("%s", errMsg)
}
