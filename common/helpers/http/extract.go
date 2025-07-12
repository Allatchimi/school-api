package httpHelper

import (
	"context"

	"api/common/constants"
	"api/common/types"
)

// GetJwtContext Returns JWT token from standard context (idem Gin)
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

// ExtractBearerContext Extracts Bearer token from standard context
func ExtractBearerContext(ctx *context.Context) string {
	if token, ok := (*ctx).Value(constants.TokenKey).(string); ok {
		return token
	}
	return ""
}
