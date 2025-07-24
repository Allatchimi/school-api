package middlewares

import (
	"net/http"
	"strings"

	"github.com/danielgtaylor/huma/v2"

	"api/common/constants"
	httpHelper "api/common/helpers/http"
	"api/common/types"
	securityUtil "api/common/utils/security"
)

// ExtractBearerTokenHeader Retrieves the bearer token from the current request context.
func ExtractBearerTokenHeader(humaCtx *huma.Context) string {
	return strings.TrimPrefix((*humaCtx).Header("Authorization"), "Bearer ")
}

// SetAuthContext Adds information such as JWT token and bearer token to context in order
// to pass information to middleware, operation and handler func
func SetAuthContext(humaCtx *huma.Context, token string, jwtToken *types.JwtToken) *huma.Context {
	ctxToken := huma.WithValue(*humaCtx, constants.TokenKey, token)
	ctxUserID := huma.WithValue(ctxToken, constants.UserIDKey, jwtToken.UserID)
	ctxSchoolID := huma.WithValue(ctxUserID, constants.SchoolIDKey, jwtToken.SchoolID)
	ctxIssuer := huma.WithValue(ctxSchoolID, constants.IssuerKey, jwtToken.Issuer)
	ctxPlatform := huma.WithValue(ctxIssuer, constants.PlatformKey, jwtToken.Platform)
	ctxDevice := huma.WithValue(ctxPlatform, constants.DeviceKey, jwtToken.Device)
	ctxApp := huma.WithValue(ctxDevice, constants.AppKey, jwtToken.App)
	ctxCode := huma.WithValue(ctxApp, constants.CodeKey, jwtToken.Code)
	return &ctxCode
}

// AuthMiddleware Handles authentication for API requests.
func AuthMiddleware(api huma.API) func(huma.Context, func(huma.Context)) {
	return func(humaCtx huma.Context, next func(huma.Context)) {
		// Check if current endpoint requires authorization
		isAuthorizationBearerRequired := false
		for _, opScheme := range humaCtx.Operation().Security {
			if _, ok := opScheme[constants.SecuritySchemeBearerToken]; ok {
				isAuthorizationBearerRequired = true
				break
			}
		}
		if !isAuthorizationBearerRequired {
			next(humaCtx)
			return
		}

		// Now check bearer authorization
		// Parse and decode the token
		token := ExtractBearerTokenHeader(&humaCtx)
		jwtToken, errCode, err := securityUtil.ValidateAuthToken(token)
		if err != nil {
			_ = huma.WriteErr(
				api,
				humaCtx,
				errCode,
				err.Error(),
				err,
			)
			return
		}
		if jwtToken != nil {
			// Get school context and compare school id form server with jwt token
			ctx := humaCtx.Context()
			schoolID := httpHelper.GetSchoolContext(&ctx)
			if jwtToken.SchoolID == schoolID {
				authCtx := SetAuthContext(&humaCtx, token, jwtToken)
				next(*authCtx)
				return
			}
		}

		tempErr := constants.Http401InvalidTokenErrorMessage()
		_ = huma.WriteErr(api, humaCtx, http.StatusUnauthorized, tempErr.Error(), tempErr)
	}
}
