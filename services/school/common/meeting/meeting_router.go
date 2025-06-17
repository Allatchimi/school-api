package meeting

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/gin-gonic/gin"

	"api/common/constants"
	"api/common/types"
	"api/config"
	"api/middlewares"
	"api/services/school/common/meeting/data"
)

func RegisterEndpoints(
	humaApi *huma.API,
	controller *Controller,
) {
	var endpointConfig = types.ApiEndpointConfig{
		Group: "/schools/meetings",
		Tag:   []string{"Meetings"},
	}
	const tableName = "meetings"

	// Create room
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "post-meeting-room",
			Summary:     "Create a room",
			Description: "Create new room and return created object.",
			Method:      http.MethodPost,
			Path:        endpointConfig.Group,
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						fmt.Sprintf("%s,%s",
							constants.FeatureAdmin,
							constants.FeatureDirector,
						), // Features scope
						tableName,                  // Table name
						constants.PermissionCreate, // Operation
					},
				},
			},
			MaxBodyBytes:  constants.DefaultBodySize,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusFound},
		},
		func(
			ctx context.Context,
			input *struct {
				Body data.MeetingRoomRequest
			},
		) (*struct{ Body data.MeetingRoomResponse }, error) {
			result, errCode, err := controller.Create(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.MeetingRoomResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Delete room with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-meeting-room",
			Summary:     "Delete room",
			Description: "Delete existing room with matching id and return affected rows in database.",
			Method:      http.MethodDelete,
			Path:        fmt.Sprintf("%s/{id}", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						fmt.Sprintf("%s,%s",
							constants.FeatureAdmin,
							constants.FeatureDirector,
						), // Features scope
						tableName,                  // Table name
						constants.PermissionDelete, // Operation
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
				data.MeetingRoomID
			},
		) (*struct{ Body types.DeletedResponse }, error) {
			result, errCode, err := controller.Delete(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DeletedResponse }{Body: types.DeletedResponse{AffectedRows: result}}, nil
		},
	)

	// Delete multiple room
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-meeting-room-multiple",
			Summary:     "Delete multiple room",
			Description: "Delete multiple room by providing a lis of IDs and return affected rows in database.",
			Method:      http.MethodDelete,
			Path:        fmt.Sprintf("%s/multiple/delete", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						fmt.Sprintf("%s,%s",
							constants.FeatureAdmin,
							constants.FeatureDirector,
						), // Features scope
						tableName,                  // Table name
						constants.PermissionDelete, // Operation
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
				Body types.DeleteMultipleRequest
			},
		) (*struct{ Body types.DeletedResponse }, error) {
			result, errCode, err := controller.DeleteMultiple(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DeletedResponse }{Body: types.DeletedResponse{AffectedRows: result}}, nil
		},
	)

	// Get room by id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-meeting-room-id",
			Summary:     "Get room by id",
			Description: "Return one room with matching id",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/{id}", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						fmt.Sprintf("%s,%s,%s,%s,%s",
							constants.FeatureAdmin,
							constants.FeatureDirector,
							constants.FeatureTeacher,
							constants.FeatureStudent,
							constants.FeatureParent,
						), // Features scope
						tableName,                // Table name
						constants.PermissionRead, // Operation
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
				data.MeetingRoomID
			},
		) (*struct{ Body data.MeetingRoomResponse }, error) {
			result, errCode, err := controller.Get(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.MeetingRoomResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Get all rooms
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-meeting-room-list",
			Summary:     "Get all rooms",
			Description: "Get all rooms with support for search, filter and pagination",
			Method:      http.MethodGet,
			Path:        endpointConfig.Group,
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						fmt.Sprintf("%s,%s,%s,%s,%s",
							constants.FeatureAdmin,
							constants.FeatureDirector,
							constants.FeatureTeacher,
							constants.FeatureStudent,
							constants.FeatureParent,
						), // Features scope
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
			Body data.MeetingRoomResponseList
		}, error) {
			result, errCode, err := controller.GetAll(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.MeetingRoomResponseList
			}{Body: *result}, nil
		},
	)

	// Join room
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-meeting-room-join",
			Summary:     "Join room",
			Description: "Join the provided room",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/join/{id}", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						fmt.Sprintf("%s,%s,%s,%s,%s",
							constants.FeatureAdmin,
							constants.FeatureDirector,
							constants.FeatureTeacher,
							constants.FeatureStudent,
							constants.FeatureParent,
						), // Features scope
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
				data.MeetingRoomID
			},
		) (
			*huma.StreamResponse,
			error,
		) {
			result, errCode, err := controller.Join(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error())
			}
			if len(result) < 1 {
				return nil, huma.NewError(http.StatusNotFound, "Invalid url!")
			}

			ginCtx := ctx.Value(middlewares.GIN_CONTEXT_KEY).(*gin.Context)
			if ginCtx != nil {
				ginCtx.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s/?access_token=%s", config.Env.MeetingApiUrl, result))
			}
			return nil, nil
		},
	)
}
