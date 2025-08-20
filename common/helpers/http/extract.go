package httpHelper

import (
	"context"

	"api/common/constants"
	"api/common/types"
)

// GetContextData Returns JWT token from standard context (idem Gin)
func GetContextData(ctx *context.Context) (result *types.ContextData) {
	// JWT
	jwt := &types.JwtToken{}
	if value, ok := (*ctx).Value(constants.UserIDKey).(int64); ok {
		jwt.UserID = value
	}
	if value, ok := (*ctx).Value(constants.SchoolIDKey).(int64); ok {
		jwt.SchoolID = value
	}
	if value, ok := (*ctx).Value(constants.PlatformKey).(string); ok {
		jwt.Platform = value
	}
	if value, ok := (*ctx).Value(constants.DeviceKey).(string); ok {
		jwt.Device = value
	}
	if value, ok := (*ctx).Value(constants.AppKey).(string); ok {
		jwt.App = value
	}
	if value, ok := (*ctx).Value(constants.CodeKey).(string); ok {
		jwt.Code = value
	}
	if value, ok := (*ctx).Value(constants.IssuerKey).(string); ok {
		jwt.Issuer = value
	}

	// User
	user := &types.ContextUserData{}
	if value, ok := (*ctx).Value(constants.RoleIDKey).(int64); ok {
		user.RoleID = value
	}
	if value, ok := (*ctx).Value(constants.FeatureKey).(string); ok {
		user.Feature = value
	}
	if value, ok := (*ctx).Value(constants.DirectorIDKey).(int64); ok {
		user.DirectorID = value
	}
	if value, ok := (*ctx).Value(constants.TeacherIDKey).(int64); ok {
		user.TeacherID = value
	}
	if value, ok := (*ctx).Value(constants.StudentIDKey).(int64); ok {
		user.StudentID = value
	}
	if value, ok := (*ctx).Value(constants.ParentIDKey).(int64); ok {
		user.ParentID = value
	}

	result = &types.ContextData{
		Jwt:  jwt,
		User: user,
	}
	return
}

// GetSchoolContext Returns the school id
func GetSchoolContext(ctx *context.Context) (result int64) {
	if value, ok := (*ctx).Value(constants.SchoolIDKey).(int64); ok {
		result = value
	}
	return
}

// ExtractBearerContext Extracts Bearer token from standard context
func ExtractBearerContext(ctx *context.Context) string {
	if token, ok := (*ctx).Value(constants.TokenKey).(string); ok {
		return token
	}
	return ""
}
