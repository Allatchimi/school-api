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
	return func(humaCtx huma.Context, next func(huma.Context)) {
		humaCtx.SetHeader("Access-Control-Allow-Origin", config.Env.AllowedHosts)

		// Check for allowed hosts
		if !isOriginKnown(humaCtx.Host()) {
			errMessage := "CORS error. Our system detected your request as malicious! Please fix that before."
			_ = huma.WriteErr(api, humaCtx, http.StatusForbidden, errMessage, fmt.Errorf("%s", errMessage))
			return
		}

		next(humaCtx)
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
