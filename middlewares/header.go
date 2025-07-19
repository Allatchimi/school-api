package middlewares

import (
	"fmt"

	"github.com/danielgtaylor/huma/v2"
)

// HeadersMiddleware Sets HTTP headers for responses.
func HeadersMiddleware(api huma.API) func(huma.Context, func(huma.Context)) {
	return func(humaCtx huma.Context, next func(huma.Context)) {

		humaCtx.SetHeader("X-Frame-Options", "DENY")
		humaCtx.SetHeader("X-Content-Type-Options", "nosniff")
		humaCtx.SetHeader("X-Xss-Protection", "1; mode=block")
		humaCtx.SetHeader("Content-Security-Policy", "default-src 'self'")
		humaCtx.SetHeader("Referrer-Policy", "strict-origin-when-cross-origin")
		humaCtx.SetHeader("Strict-Transport-Security", fmt.Sprintf("max-age=%d; %s", 31536000, "includeSubDomains"))

		next(humaCtx)
	}
}
