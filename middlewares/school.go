package middlewares

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"api/common/constants"
	securityUtil "api/common/utils/security"
	"api/config"
)

// ExtractSchoolTokenHeader Retrieves the school token from the current request context.
func ExtractSchoolTokenHeaders(humaCtx *huma.Context) (string, string) {
	return (*humaCtx).Header("X-School-Api-Key"), (*humaCtx).Header("X-School-Id")
}

// SetAuthContext Adds information such as JWT token and bearer token to context in order
// to pass information to middleware, operation and handler func
func SetSchoolContext(humaCtx *huma.Context, schoolIDStr string) *huma.Context {
	ctxSchoolID := huma.WithValue(*humaCtx, constants.SchoolIDKey, schoolIDStr)
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

		schoolToken, schoolIDStr := ExtractSchoolTokenHeaders(&humaCtx)
		if len(schoolToken) < 1 && len(schoolIDStr) < 1 {
			next(humaCtx)
			return
		}
		ok, errSecretProof := securityUtil.VerifyHMAC_SHA256_Base64URL(
			schoolIDStr,
			schoolToken,
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

		next(*SetSchoolContext(&humaCtx, schoolIDStr))
	}
}
