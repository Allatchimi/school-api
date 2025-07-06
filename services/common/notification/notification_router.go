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

	// Update notification seen status
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-notification-seen",
			Summary:     "Update notification seen status",
			Description: "Update existing notification seen status with matching id and return the new notification object.",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/{id}/seen", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecuritySchemeBearerToken: { // Authentication
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
				Body data.NotificationSeenRequest
			},
		) (*struct{ Body data.NotificationResponse }, error) {
			result, errCode, err := controller.UpdateSeen(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.NotificationResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Update notification seen all status
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-notification-seen-all",
			Summary:     "Update notification seen status for all notifications",
			Description: "Update existing notification seen status for all notifications and return the new notification object list.",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/seen/all", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecuritySchemeBearerToken: { // Authentication
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
				Body data.NotificationSeenAllRequest
			},
		) (*struct{ Body types.DefaultResponse }, error) {
			errCode, err := controller.UpdateSeenAll(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DefaultResponse }{Body: types.DefaultResponse{
				Message: "Successful updated!",
			}}, nil
		},
	)

	// Delete notification with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-notification",
			Summary:     "Delete notification",
			Description: "Delete existing notification and return affected rows in database.",
			Method:      http.MethodDelete,
			Path:        fmt.Sprintf("%s/{id}", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecuritySchemeBearerToken: { // Authentication
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
		) (*struct{ Body types.DeletedResponse }, error) {
			result, errCode, err := controller.Delete(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DeletedResponse }{Body: types.DeletedResponse{AffectedRows: result}}, nil
		},
	)

	// Delete all notification
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-notification-all-user",
			Summary:     "Delete all notification for user",
			Description: "Delete all existing notification for user and return affected rows in database.",
			Method:      http.MethodDelete,
			Path:        fmt.Sprintf("%s/all", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecuritySchemeBearerToken: { // Authentication
					},
				},
			},
			MaxBodyBytes:  constants.DefaultBodySize,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct{},
		) (*struct{ Body types.DeletedResponse }, error) {
			result, errCode, err := controller.DeleteAll(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DeletedResponse }{Body: types.DeletedResponse{AffectedRows: result}}, nil
		},
	)

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
					constants.SecuritySchemeBearerToken: { // Authentication
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

	// Get notification count not seen
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-notification-count-not-seen",
			Summary:     "Get notification count not seen",
			Description: "Return notification count not seen",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/notseen/count", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecuritySchemeBearerToken: { // Authentication
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
				data.NotificationNotSeenCountRequest
			},
		) (*struct {
			Body data.NotificationNotSeenResponse
		}, error) {
			result, errCode, err := controller.GetNotSeenCount(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.NotificationNotSeenResponse
			}{Body: data.NotificationNotSeenResponse{Count: result}}, nil
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
					constants.SecuritySchemeBearerToken: { // Authentication
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
