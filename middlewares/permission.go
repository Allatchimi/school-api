package middlewares

import (
	httpHelper "api/common/helpers/http"
	"api/services/school/common/director"
	"api/services/school/common/parent"
	"api/services/school/common/student"
	"api/services/school/common/teacher"
	"api/services/user/permission"
	"api/services/user/user"
	"fmt"
	"net/http"
	"strings"

	"github.com/danielgtaylor/huma/v2"

	"api/common/constants"
)

// SetUserContext Adds user information such role, feature,...
func SetUserContext(humaCtx *huma.Context, roleID int64, feature string, userFeatureID int64) *huma.Context {
	ctxRoleID := huma.WithValue(*humaCtx, constants.RoleIDKey, roleID)
	ctxFeature := huma.WithValue(ctxRoleID, constants.FeatureKey, feature)
	var ctxResult *huma.Context = &ctxFeature
	switch feature {
	case constants.FeatureDirector:
		tmpCtx := huma.WithValue(*ctxResult, constants.DirectorIDKey, userFeatureID)
		ctxResult = &tmpCtx
	case constants.FeatureTeacher:
		tmpCtx := huma.WithValue(*ctxResult, constants.TeacherIDKey, userFeatureID)
		ctxResult = &tmpCtx
	case constants.FeatureStudent:
		tmpCtx := huma.WithValue(*ctxResult, constants.StudentIDKey, userFeatureID)
		ctxResult = &tmpCtx
	case constants.FeatureParent:
		tmpCtx := huma.WithValue(*ctxResult, constants.ParentIDKey, userFeatureID)
		ctxResult = &tmpCtx
	}
	return ctxResult
}

// PermissionMiddleware Checks resource permissions
func PermissionMiddleware(
	api huma.API,
	userRepo *user.Repository,
	permissionRepo *permission.Repository,
	directorRepo *director.Repository,
	teacherRepo *teacher.Repository,
	studentRepo *student.Repository,
	parentRepo *parent.Repository,
) func(huma.Context, func(huma.Context)) {
	return func(humaCtx huma.Context, next func(huma.Context)) {
		// Retrieve context data
		ctx := humaCtx.Context()
		ctxData := httpHelper.GetContextData(&ctx)
		if ctxData == nil || ctxData.Jwt == nil || ctxData.Jwt.UserID < 1 {
			next(humaCtx)
			return
		}
		// Find user
		foundUser, errFound := userRepo.GetByID(ctxData.Jwt.UserID)
		if errFound != nil {
			tempErr := constants.Http500ErrorMessage("interact with user model")
			_ = huma.WriteErr(api, humaCtx, http.StatusInternalServerError, tempErr.Error(), tempErr)
			return
		}
		if foundUser == nil {
			tempErr := constants.Http403InvalidPermissionErrorMessage()
			_ = huma.WriteErr(api, humaCtx, http.StatusForbidden, tempErr.Error(), tempErr)
			return
		}
		// Check if the status is enabled
		if foundUser.Status != constants.USER_STATUS_ENABLED {
			tempErr := fmt.Errorf("%s", "User account is disabled! Please contact support.")
			_ = huma.WriteErr(api, humaCtx, http.StatusUnavailableForLegalReasons, tempErr.Error(), tempErr)
			return
		}

		// Retrieve feature permissions
		var featuresScope, tableName, tableOperation string
		for _, opScheme := range humaCtx.Operation().Security {
			if securityScheme, ok := opScheme[constants.SecuritySchemeBearerToken]; ok {
				if len(securityScheme) > 0 {
					featuresScope = securityScheme[0]
					if len(securityScheme) > 1 {
						tableName = securityScheme[1]
						if len(securityScheme) > 2 {
							tableOperation = securityScheme[2]
						}
					}
				}
				break
			}
		}

		// Check for required permissions
		if len(featuresScope) > 0 {
			// Check if the user has feature
			var haveFeature bool = false
			for _, feat := range strings.Split(featuresScope, ",") {
				if feat == foundUser.Role.Feature {
					haveFeature = true
				}
			}
			if !haveFeature {
				tempErr := constants.Http403InvalidPermissionErrorMessage()
				_ = huma.WriteErr(api, humaCtx, http.StatusForbidden, tempErr.Error(), tempErr)
				return
			}

			// Check if the user has table name permission
			if len(tableName) > 0 {
				// Retrieve permission
				userPermission, errPerm := permissionRepo.GetByRoleIDTableNameMultiple(foundUser.RoleID, tableName, "*")
				if errPerm != nil {
					tempErr := constants.Http500ErrorMessage("interact with permission model")
					_ = huma.WriteErr(api, humaCtx, http.StatusInternalServerError, tempErr.Error(), tempErr)
					return
				}
				tempErr := constants.Http403InvalidPermissionErrorMessage()
				if !(userPermission != nil && (userPermission.TableName == tableName || userPermission.TableName == "*")) {
					_ = huma.WriteErr(api, humaCtx, http.StatusForbidden, tempErr.Error(), tempErr)
					return
				}

				if len(tableOperation) >= 1 {
					if tableOperation == constants.PermissionCreate && !userPermission.Create {
						_ = huma.WriteErr(api, humaCtx, http.StatusForbidden, tempErr.Error(), tempErr)
						return
					} else if tableOperation == constants.PermissionRead && !userPermission.Read {
						_ = huma.WriteErr(api, humaCtx, http.StatusForbidden, tempErr.Error(), tempErr)
						return
					} else if tableOperation == constants.PermissionUpdate && !userPermission.Update {
						_ = huma.WriteErr(api, humaCtx, http.StatusForbidden, tempErr.Error(), tempErr)
						return
					} else if tableOperation == constants.PermissionDelete && !userPermission.Delete {
						_ = huma.WriteErr(api, humaCtx, http.StatusForbidden, tempErr.Error(), tempErr)
						return
					}
				}
			}
		}

		// Set user context data
		var userFeatureID int64
		switch foundUser.Role.Feature {
		case constants.FeatureDirector:
			foundFeat, errFoundFeat := directorRepo.GetByUserID(foundUser.ID)
			if errFoundFeat != nil {
				tempErr := constants.Http500ErrorMessage("interact with user model")
				_ = huma.WriteErr(api, humaCtx, http.StatusInternalServerError, tempErr.Error(), tempErr)
				return
			}
			if foundFeat == nil || foundFeat.ID < 1 {
				tempErr := constants.Http403InvalidPermissionErrorMessage()
				_ = huma.WriteErr(api, humaCtx, http.StatusForbidden, tempErr.Error(), tempErr)
				return
			}
			userFeatureID = foundFeat.ID
		case constants.FeatureTeacher:
			foundFeat, errFoundFeat := teacherRepo.GetByUserID(foundUser.ID)
			if errFoundFeat != nil {
				tempErr := constants.Http500ErrorMessage("interact with user model")
				_ = huma.WriteErr(api, humaCtx, http.StatusInternalServerError, tempErr.Error(), tempErr)
				return
			}
			if foundFeat == nil || foundFeat.ID < 1 {
				tempErr := constants.Http403InvalidPermissionErrorMessage()
				_ = huma.WriteErr(api, humaCtx, http.StatusForbidden, tempErr.Error(), tempErr)
				return
			}
			userFeatureID = foundFeat.ID
		case constants.FeatureStudent:
			foundFeat, errFoundFeat := studentRepo.GetByUserID(foundUser.ID)
			if errFoundFeat != nil {
				tempErr := constants.Http500ErrorMessage("interact with user model")
				_ = huma.WriteErr(api, humaCtx, http.StatusInternalServerError, tempErr.Error(), tempErr)
				return
			}
			if foundFeat == nil || foundFeat.ID < 1 {
				tempErr := constants.Http403InvalidPermissionErrorMessage()
				_ = huma.WriteErr(api, humaCtx, http.StatusForbidden, tempErr.Error(), tempErr)
				return
			}
			userFeatureID = foundFeat.ID
		case constants.FeatureParent:
			foundFeat, errFoundFeat := parentRepo.GetByUserID(foundUser.ID)
			if errFoundFeat != nil {
				tempErr := constants.Http500ErrorMessage("interact with user model")
				_ = huma.WriteErr(api, humaCtx, http.StatusInternalServerError, tempErr.Error(), tempErr)
				return
			}
			if foundFeat == nil || foundFeat.ID < 1 {
				tempErr := constants.Http403InvalidPermissionErrorMessage()
				_ = huma.WriteErr(api, humaCtx, http.StatusForbidden, tempErr.Error(), tempErr)
				return
			}
			userFeatureID = foundFeat.ID
		}
		userCtx := SetUserContext(&humaCtx, foundUser.RoleID, foundUser.Role.Feature, userFeatureID)

		// Next
		next(*userCtx)

	}
}
