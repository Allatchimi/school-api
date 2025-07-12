package httpHelper

import (
	"context"
	"fmt"

	"api/common/constants"
	"api/common/types"
	securityUtil "api/common/utils/security"

	"github.com/gin-gonic/gin"
)

// GetJwtContext Returns JWT token from context
func GetJwtContext(ctx *context.Context) *types.JwtToken {
	result := &types.JwtToken{}
	if id, okID := (*ctx).Value(constants.UserIDKey).(int64); okID {
		result.UserID = id
	}
	if iss, okIss := (*ctx).Value(constants.IssuerKey).(string); okIss {
		result.Issuer = iss
	}
	if platform, okPlatform := (*ctx).Value(constants.PlatformKey).(string); okPlatform {
		result.Platform = platform
	}
	if device, okDevice := (*ctx).Value(constants.DeviceKey).(string); okDevice {
		result.Device = device
	}
	if app, okApp := (*ctx).Value(constants.AppKey).(string); okApp {
		result.App = app
	}
	if code, okCode := (*ctx).Value(constants.CodeKey).(int); okCode {
		result.Code = code
	}
	return result
}

// GetJwtContextFromQuery Extracts JWT token from query
func GetJwtContextFromQuery(c *gin.Context) (*types.JwtToken, error) {
	token := c.Query("token")
	if len(token) < 1 {
		errMsg := "No token found! Please enter valid information."
		return nil, fmt.Errorf("%s", errMsg)
	}
	jwtToken, _, err := securityUtil.ValidateAuthToken(token)
	if err != nil || jwtToken == nil || jwtToken.UserID < 1 {
		errMsg := "Invalid token! Please enter valid information."
		return nil, fmt.Errorf("%s", errMsg)
	}
	return jwtToken, nil
}

// ExtractBearerContext Extracts Bearer token from context
func ExtractBearerContext(ctx *context.Context) string {
	if token, ok := (*ctx).Value(constants.TokenKey).(string); ok {
		return token
	}
	return ""
}
