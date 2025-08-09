package initialize

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"api/common/constants"
	"api/common/types"
	"api/services/others/initialize/data"
)

func RegisterEndpoints(
	humaApi *huma.API,
	controller *Controller,
) {
	var endpointConfig = types.ApiEndpointConfig{
		Group: fmt.Sprintf("/%s", constants.RESOURCE_TABLE_INITIALIZE),
		Tag:   []string{"Initialize"},
	}
	const tableName = constants.RESOURCE_TABLE_INITIALIZE

	// Get initial school by id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-initialize",
			Summary:     "Get initial",
			Description: "Return one initial configuration",
			Method:      http.MethodGet,
			Path:        endpointConfig.Group,
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecuritySchemeSchoolToken: {},
					constants.SecuritySchemeSchoolID:    {},
				},
			},
			MaxBodyBytes:  constants.DefaultBodySize,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct{},
		) (*struct{ Body data.InitializeResponse }, error) {
			result, errCode, err := controller.GetInitial(&ctx)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.InitializeResponse }{Body: *result}, nil
		},
	)
}
