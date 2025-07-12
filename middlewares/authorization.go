package middlewares

import (
	"net/http"
	"strings"

	"github.com/danielgtaylor/huma/v2"

	"api/common/constants"
	"api/common/types"
	securityUtil "api/common/utils/security"
	"api/config"
)

// ExtractBearerTokenHeader Retrieves the bearer token from the current request context.
func ExtractBearerTokenHeader(ctx *huma.Context) string {
	return strings.TrimPrefix((*ctx).Header("Authorization"), "Bearer ")
}

// ExtractSchoolTokenHeader Retrieves the school token from the current request context.
func ExtractSchoolTokenHeaders(ctx *huma.Context) (string, string) {
	return (*ctx).Header("X-School-Api-Key"), (*ctx).Header("X-School-Id")
}

// SetAuthContext Adds information such as JWT token and bearer token to context in order
// to pass information to middleware, operation and handler func
func SetAuthContext(ctx *huma.Context, token string, jwtToken *types.JwtToken) *huma.Context {
	ctxToken := huma.WithValue(*ctx, constants.TokenKey, token)
	ctxUserID := huma.WithValue(ctxToken, constants.UserIDKey, jwtToken.UserID)
	ctxIssuer := huma.WithValue(ctxUserID, constants.IssuerKey, jwtToken.Issuer)
	ctxPlatform := huma.WithValue(ctxIssuer, constants.PlatformKey, jwtToken.Platform)
	ctxDevice := huma.WithValue(ctxPlatform, constants.DeviceKey, jwtToken.Device)
	ctxApp := huma.WithValue(ctxDevice, constants.AppKey, jwtToken.App)
	ctxCode := huma.WithValue(ctxApp, constants.CodeKey, jwtToken.Code)
	return &ctxCode
}

// AuthMiddleware Handles authentication for API requests.
func AuthMiddleware(api huma.API) func(huma.Context, func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		// Check if current endpoint requires authorization
		isAuthorizationBearerRequired := false
		isAuthorizationSchoolTokenRequired := false
		for _, opScheme := range ctx.Operation().Security {
			if _, ok := opScheme[constants.SecuritySchemeBearerToken]; ok {
				isAuthorizationBearerRequired = true
				break
			}
			if _, ok := opScheme[constants.SecuritySchemeSchoolToken]; ok {
				isAuthorizationSchoolTokenRequired = true
				break
			}
		}
		if !isAuthorizationBearerRequired && !isAuthorizationSchoolTokenRequired {
			next(ctx)
			return
		}

		// Check school authorization
		if isAuthorizationSchoolTokenRequired {
			schoolToken, schoolIDStr := ExtractSchoolTokenHeaders(&ctx)
			ok, errSecretProof := securityUtil.VerifyHMAC_SHA256_Base64URL(
				schoolIDStr,
				schoolToken,
				config.Env.SchoolApiSecret,
			)
			if errSecretProof != nil {
				_ = huma.WriteErr(
					api,
					ctx,
					http.StatusInternalServerError,
					constants.Http500ErrorMessage("verify HMAC SHA256 Base64URL").Error(),
					constants.Http500ErrorMessage("verify HMAC SHA256 Base64URL"),
				)
				return
			}
			if !ok {
				_ = huma.WriteErr(
					api,
					ctx,
					http.StatusUnauthorized,
					constants.Http401InvalidTokenErrorMessage().Error(),
					constants.Http401InvalidTokenErrorMessage(),
				)
				return
			}
		}

		// Now check bearer authorization
		// Parse and decode the token
		token := ExtractBearerTokenHeader(&ctx)
		jwtToken, errCode, err := securityUtil.ValidateAuthToken(token)
		if err != nil {
			_ = huma.WriteErr(
				api,
				ctx,
				errCode,
				err.Error(),
				err,
			)
			return
		}
		if jwtToken != nil {
			next(*SetAuthContext(&ctx, token, jwtToken))
			return
		}

		tempErr := constants.Http401InvalidTokenErrorMessage()
		_ = huma.WriteErr(api, ctx, http.StatusUnauthorized, tempErr.Error(), tempErr)
	}
}
