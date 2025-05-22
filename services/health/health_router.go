package health

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"api/common/constants"
	"api/common/types"
	"api/services/health/data"
)

func RegisterEndpoints(
	humaApi *huma.API,
	controller *Controller,
) {
	var endpointConfig = types.ApiEndpointConfig{
		Group: "/healthz",
		Tag:   []string{"Healthz"},
	}

	// Get postgres health
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "get-health-postgres",
			Summary:       "Get postgres health",
			Description:   "Checks if the API can communicate with postgres",
			Method:        http.MethodGet,
			Path:          fmt.Sprintf("%s/postgres", endpointConfig.Group),
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
			isOk, errCode, err := controller.GetPostgresHealth(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			if !isOk {
				return nil, huma.NewError(errCode, "can't join postgres", err)
			}
			return &struct{ Body data.HealthResponse }{Body: data.HealthResponse{
				Message: "OK",
			}}, nil
		},
	)

	// Get redis health
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "get-health-redis",
			Summary:       "Get redis health",
			Description:   "Checks if the API can communicate with redis",
			Method:        http.MethodGet,
			Path:          fmt.Sprintf("%s/redis", endpointConfig.Group),
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
			result, errCode, err := controller.GetRedisHealth(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			if !result {
				return nil, huma.NewError(errCode, "can't join redis", err)
			}
			return &struct{ Body data.HealthResponse }{Body: data.HealthResponse{
				Message: "OK",
			}}, nil
		},
	)
}
