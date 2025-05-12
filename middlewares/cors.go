package middlewares

import (
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/danielgtaylor/huma/v2"

	"api/config"
)

func CorsMiddleware(api huma.API) func(huma.Context, func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		ctx.SetHeader("Access-Control-Allow-Origin", config.Env.AllowedHosts)

		// Check for allowed hosts
		if !isOriginKnown(ctx.Host()) {
			errMessage := "CORS error. Our system detected your request as malicious! Please fix that before."
			_ = huma.WriteErr(api, ctx, http.StatusForbidden, errMessage, fmt.Errorf("%s", errMessage))
			return
		}

		next(ctx)
	}
}

// Utility function for CORS
func isOriginKnown(host string) bool {
	if config.Env.AllowedHosts == "*" {
		return true
	}
	hosts := strings.Split(config.Env.AllowedHosts, ",")
	return slices.Contains(hosts, host)
}
