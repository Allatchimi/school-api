package middlewares

import (
	"net/http"
	"strconv"

	"github.com/danielgtaylor/huma/v2"

	"api/common/constants"
	securityUtil "api/common/utils/security"
	"api/config"
)

// ExtractSchoolTokenHeader Retrieves the school token from the current request context.
func ExtractSchoolTokenHeaders(humaCtx *huma.Context) (apiKey string, schoolID string) {
	apiKey = (*humaCtx).Header("X-School-Api-Key")
	schoolID = (*humaCtx).Header("X-School-Id")
	return
}

// SetAuthContext Adds information such as JWT token and bearer token to context in order
// to pass information to middleware, operation and handler func
func SetSchoolContext(humaCtx *huma.Context, schoolID int64) *huma.Context {
	ctxSchoolID := huma.WithValue(*humaCtx, constants.SchoolIDKey, schoolID)
	return &ctxSchoolID
}

// AuthMiddleware Handles authentication for API requests.
func SchoolMiddleware(api huma.API) func(huma.Context, func(huma.Context)) {
	return func(humaCtx huma.Context, next func(huma.Context)) {
		// Check if current endpoint requires authorization
		isAuthorizationSchoolTokenRequired := false
		for _, opScheme := range humaCtx.Operation().Security {
			if _, ok := opScheme[constants.SecuritySchemeSchoolToken]; ok {
				isAuthorizationSchoolTokenRequired = true
				break
			}
		}
		if !isAuthorizationSchoolTokenRequired {
			next(humaCtx)
			return
		}

		schoolApiKey, schoolIDStr := ExtractSchoolTokenHeaders(&humaCtx)
		if len(schoolApiKey) > 0 && len(schoolIDStr) > 0 {
			ok, errSecretProof := securityUtil.VerifyHMAC_SHA256_Base64URL(
				schoolIDStr,
				schoolApiKey,
				config.Env.SchoolApiSecret,
			)
			if errSecretProof != nil {
				_ = huma.WriteErr(
					api,
					humaCtx,
					http.StatusInternalServerError,
					constants.Http500ErrorMessage("verify HMAC SHA256 Base64URL").Error(),
					constants.Http500ErrorMessage("verify HMAC SHA256 Base64URL"),
				)
				return
			}
			if !ok {
				_ = huma.WriteErr(
					api,
					humaCtx,
					http.StatusUnauthorized,
					constants.Http401InvalidTokenErrorMessage().Error(),
					constants.Http401InvalidTokenErrorMessage(),
				)
				return
			}

			schoolID, err := strconv.ParseInt(schoolIDStr, 10, 64)
			if err != nil {
				_ = huma.WriteErr(
					api,
					humaCtx,
					http.StatusInternalServerError,
					constants.Http500ErrorMessage("").Error(),
					constants.Http500ErrorMessage(""),
				)
				return
			}
			next(*SetSchoolContext(&humaCtx, schoolID))
			return
		}

		next(humaCtx)
	}
}
