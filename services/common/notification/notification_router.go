package notification

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"api/common/constants"
	"api/common/types"
	"api/services/common/notification/data"
)

func RegisterEndpoints(
	humaApi *huma.API,
	controller *Controller,
) {
	var endpointConfig = types.ApiEndpointConfig{
		Group: "/notifications",
		Tag:   []string{"Notifications"},
	}

	// Get notification by id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-notification-id",
			Summary:     "Get notification by id",
			Description: "Return one notification with matching id",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/{id}", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
					},
				},
			},
			MaxBodyBytes:  constants.DefaultBodySize,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				data.NotificationID
			},
		) (*struct{ Body data.NotificationResponse }, error) {
			result, errCode, err := controller.Get(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.NotificationResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Get all notifications
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-notification-list",
			Summary:     "Get all notifications",
			Description: "Get all notifications with support for search, filter and pagination",
			Method:      http.MethodGet,
			Path:        endpointConfig.Group,
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
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
			Body data.NotificationResponseList
		}, error) {
			result, errCode, err := controller.GetAll(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}

			// Generate items
			tempResult := make([]data.NotificationResponse, 10)
			for i := range result.Data {
				tempModel := data.NotificationResponse{}
				tempModel.ID = int64(i)

				tempResult[i] = tempModel
			}
			result.Data = tempResult

			return &struct {
				Body data.NotificationResponseList
			}{Body: *result}, nil
		},
	)
}
