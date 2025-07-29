package health

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"api/common/constants"
	"api/common/types"
	"api/services/others/health/data"
)

func RegisterEndpoints(
	humaApi *huma.API,
	controller *Controller,
) {
	var endpointConfig = types.ApiEndpointConfig{
		Group: "/healthz",
		Tag:   []string{"Healthz"},
	}

	// Get live health
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "get-health-live",
			Summary:       "Get health live",
			Description:   "Checks if the API is alive.",
			Method:        http.MethodGet,
			Path:          fmt.Sprintf("%s/live", endpointConfig.Group),
			Tags:          endpointConfig.Tag,
			MaxBodyBytes:  constants.DefaultBodySize,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				data.HealthRequest
			},
		) (*struct{ Body data.HealthResponse }, error) {
			isOk, errCode, err := controller.HealthLive(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			if !isOk {
				return nil, huma.NewError(errCode, "can't join http server", err)
			}
			return &struct{ Body data.HealthResponse }{Body: data.HealthResponse{
				Message: "OK",
			}}, nil
		},
	)

	// Get dependencies health
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "get-health-deps",
			Summary:       "Get dependencies health",
			Description:   "Checks if the API can communicate with dependencies: postgres, redis",
			Method:        http.MethodGet,
			Path:          fmt.Sprintf("%s/dependencies", endpointConfig.Group),
			Tags:          endpointConfig.Tag,
			MaxBodyBytes:  constants.DefaultBodySize,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				data.HealthRequest
			},
		) (*struct{ Body data.HealthResponse }, error) {
			result, errCode, err := controller.HealthDepencencies(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			if !result {
				return nil, huma.NewError(errCode, "can't join dependencies", err)
			}
			return &struct{ Body data.HealthResponse }{Body: data.HealthResponse{
				Message: "OK",
			}}, nil
		},
	)
}
