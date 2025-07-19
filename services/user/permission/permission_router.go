package permission

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"api/common/constants"
	"api/common/types"
	"api/services/user/permission/data"
)

func RegisterEndpoints(
	humaApi *huma.API,
	controller *Controller,
) {
	var endpointConfig = types.ApiEndpointConfig{
		Group: "/permissions",
		Tag:   []string{"Permissions"},
	}
	const tableName = "permissions"

	// Update permission
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-permission",
			Summary:     "Update permission",
			Description: "Update permission with matching role id, table name with actions(CRUD)",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/role/{roleID}", endpointConfig.Group),
			Tags:        endpointConfig.Tag,

			Security: []map[string][]string{
				{
					constants.SecuritySchemeBearerToken: {
						fmt.Sprintf("%s", constants.FeatureAdmin), // Feature
						tableName,                  // Table name
						constants.PermissionUpdate, // Operation
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
				data.PermissionRoleID
				Body data.UpdatePermissionRequest
			},
		) (*struct {
			Body data.PermissionResponse
		}, error) {
			result, errCode, err := controller.Update(
				&ctx, input,
			)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.PermissionResponse
			}{Body: *result}, nil
		},
	)

	// Delete permission with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-permission-id",
			Summary:     "Delete permission",
			Description: "Delete existing permission with matching id and return affected rows in database.",
			Method:      http.MethodDelete,
			Path:        fmt.Sprintf("%s/{id}", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecuritySchemeBearerToken: {
						fmt.Sprintf("%s", constants.FeatureAdmin), // Feature
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
				data.PermissionID
			},
		) (*struct{ Body types.DeletedResponse }, error) {
			result, errCode, err := controller.Delete(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DeletedResponse }{Body: types.DeletedResponse{AffectedRows: result}}, nil
		},
	)

	// Delete multiple permission
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-permission-multiple",
			Summary:     "Delete multiple permission",
			Description: "Delete multiple permission by providing a lis of IDs and return affected rows in database.",
			Method:      http.MethodDelete,
			Path:        fmt.Sprintf("%s/multiple/delete", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecuritySchemeBearerToken: {
						fmt.Sprintf("%s", constants.FeatureAdmin), // Feature
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

	// Get all permission
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-permission-list",
			Summary:     "Get all permission",
			Description: "Get all permission and support for search, filter and pagination",
			Method:      http.MethodGet,
			Path:        endpointConfig.Group,
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecuritySchemeBearerToken: {
						fmt.Sprintf("%s", constants.FeatureAdmin), // Feature
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
			Body data.PermissionListResponse
		}, error) {
			result, errCode, err := controller.GetAll(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}

			// Generate items
			tempResult := make([]data.PermissionResponse, 10)
			for i := range result.Data {
				tempModel := data.PermissionResponse{}
				tempModel.ID = int64(i)

				tempResult[i] = tempModel
			}
			result.Data = tempResult

			return &struct {
				Body data.PermissionListResponse
			}{Body: *result}, nil
		},
	)
}
