package middlewares

import (
	"api/common/helpers"
	"api/services/user/permission"
	"api/services/user/user"
	"net/http"
	"strings"

	"github.com/danielgtaylor/huma/v2"

	"api/common/constants"
)

// PermissionMiddleware Checks resource permissions
func PermissionMiddleware(api huma.API, userRepo *user.Repository, permissionRepo *permission.Repository) func(huma.Context, func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		// Retrieve jwtToken
		ctxContext := ctx.Context()
		jwtToken := helpers.GetJwtContext(&ctxContext)
		if jwtToken == nil || jwtToken.UserID <= 0 {
			next(ctx)
			return
		}

		// Retrieve feature permissions
		var featuresScope, tableName, tableOperation string
		for _, opScheme := range ctx.Operation().Security {
			if securityScheme, ok := opScheme[constants.SecurityAuthName]; ok {
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
			// Retrieve user data
			foundUser, errFound := userRepo.GetByID(jwtToken.UserID)
			if errFound != nil {
				tempErr := constants.Http500ErrorMessage("interact with role model")
				_ = huma.WriteErr(api, ctx, http.StatusInternalServerError, tempErr.Error(), tempErr)
				return
			}
			if foundUser == nil {
				tempErr := constants.Http403InvalidPermissionErrorMessage()
				_ = huma.WriteErr(api, ctx, http.StatusForbidden, tempErr.Error(), tempErr)
				return
			}

			// Check if the use have feature
			var haveFeature = false
			for _, feat := range strings.Split(featuresScope, ",") {
				if feat == foundUser.Role.Feature {
					haveFeature = true
				}
			}
			if !haveFeature {
				tempErr := constants.Http403InvalidPermissionErrorMessage()
				_ = huma.WriteErr(api, ctx, http.StatusForbidden, tempErr.Error(), tempErr)
				return
			}

			// Check if the user have table name permission
			if len(tableName) > 0 {
				// Retrieve permission
				userPermission, errPerm := permissionRepo.GetByRoleIDTableNameMultiple(foundUser.RoleID, tableName, "*")
				if errPerm != nil {
					tempErr := constants.Http500ErrorMessage("interact with permission model")
					_ = huma.WriteErr(api, ctx, http.StatusInternalServerError, tempErr.Error(), tempErr)
					return
				}
				tempErr := constants.Http403InvalidPermissionErrorMessage()
				if !(userPermission != nil && (userPermission.TableName == tableName || userPermission.TableName == "*")) {
					_ = huma.WriteErr(api, ctx, http.StatusForbidden, tempErr.Error(), tempErr)
					return
				}

				if len(tableOperation) >= 1 {
					if tableOperation == constants.PermissionCreate && !userPermission.Create {
						_ = huma.WriteErr(api, ctx, http.StatusForbidden, tempErr.Error(), tempErr)
						return
					} else if tableOperation == constants.PermissionRead && !userPermission.Read {
						_ = huma.WriteErr(api, ctx, http.StatusForbidden, tempErr.Error(), tempErr)
						return
					} else if tableOperation == constants.PermissionUpdate && !userPermission.Update {
						_ = huma.WriteErr(api, ctx, http.StatusForbidden, tempErr.Error(), tempErr)
						return
					} else if tableOperation == constants.PermissionDelete && !userPermission.Delete {
						_ = huma.WriteErr(api, ctx, http.StatusForbidden, tempErr.Error(), tempErr)
						return
					}
				}
			}
		}

		next(ctx)

	}
}
