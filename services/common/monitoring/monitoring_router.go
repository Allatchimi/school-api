package monitoring

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"api/common/constants"
	"api/common/types"
	"api/services/common/monitoring/data"
)

func RegisterEndpoints(
	humaApi *huma.API,
	controller *Controller,
) {
	var endpointConfig = types.ApiEndpointConfig{
		Group: "/monitorings",
		Tag:   []string{"Monitorings"},
	}
	const tableName = "monitorings"

	// Get all monitoring
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-monitoring-list",
			Summary:     "Get all monitoring",
			Description: "Get all monitoring with support for search, filter and pagination",
			Method:      http.MethodGet,
			Path:        endpointConfig.Group,
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecuritySchemeBearerToken: { // Authentication
						fmt.Sprintf("%s,%s,%s,%s,%s",
							constants.FeatureAdmin,
							constants.FeatureDirector,
							constants.FeatureTeacher,
							constants.FeatureStudent,
							constants.FeatureParent,
						), // Feature
						tableName,                // Table name
						constants.PermissionRead, // Operation
					},
				},
			},
			MaxBodyBytes:  constants.DefaultBodySize,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden},
		},
		func(
			ctx context.Context,
			input *struct {
				types.Filter
				types.PaginationRequest
				data.GetAllRequest
			},
		) (*struct {
			Body data.MonitoringResponseList
		}, error) {
			result, errCode, err := controller.GetAll(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}

			// Generate items
			tempResult := &data.MonitoringResponse{}
			tempResult.Count = &data.CountResponse{
				Schools:   2,
				Directors: 4,
				Teachers:  29,
				Students:  3842,
				Parents:   123,
			}
			tempResult.UsersByYear = make([]data.UsersByYearResponse, 5)
			tempResult.SuccessBySchool = make([]data.SuccessBySchoolResponse, 5)
			tempResult.SuccessBySchoolGender = make([]data.SuccessBySchoolGenderResponse, 5)
			result.Data = tempResult

			return &struct {
				Body data.MonitoringResponseList
			}{Body: *result}, nil
		},
	)
}
