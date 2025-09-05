package unit

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"api/common/constants"
	"api/common/types"
	"api/services/school/university/unit/data"
)

func RegisterEndpoints(
	humaApi *huma.API,
	controller *Controller,
) {
	var endpointConfig = types.ApiEndpointConfig{
		Group: "/schools/university/units",
		Tag:   []string{"University - Units"},
	}
	const tableName = constants.RESOURCE_TABLE_UNIVERSITY_UNIT

	// Create unit
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "post-unit",
			Summary:     "Create unit",
			Description: "Create new unit and return the created object.",
			Method:      http.MethodPost,
			Path:        endpointConfig.Group,
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecuritySchemeSchoolToken: {},
					constants.SecuritySchemeSchoolID:    {},
					constants.SecuritySchemeBearerToken: {
						fmt.Sprintf("%s,%s",
							constants.FeatureAdmin,
							constants.FeatureDirector,
						), // Feature
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
				Body data.UnitRequest
			},
		) (*struct{ Body data.UnitResponse }, error) {
			result, errCode, err := controller.Create(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.UnitResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Update unit with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-unit",
			Summary:     "Update unit",
			Description: "Update existing unit with matching id and return the updated object.",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/{id}", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecuritySchemeSchoolToken: {},
					constants.SecuritySchemeSchoolID:    {},
					constants.SecuritySchemeBearerToken: {
						fmt.Sprintf("%s,%s",
							constants.FeatureAdmin,
							constants.FeatureDirector,
						), // Feature
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
				data.UnitID
				Body data.UnitRequest
			},
		) (*struct{ Body data.UnitResponse }, error) {
			result, errCode, err := controller.Update(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			fmt.Print(result)
			return &struct{ Body data.UnitResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Delete unit with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-unit",
			Summary:     "Delete unit",
			Description: "Delete existing unit and return affected rows in database.",
			Method:      http.MethodDelete,
			Path:        fmt.Sprintf("%s/{id}", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecuritySchemeSchoolToken: {},
					constants.SecuritySchemeSchoolID:    {},
					constants.SecuritySchemeBearerToken: {
						fmt.Sprintf("%s,%s",
							constants.FeatureAdmin,
							constants.FeatureDirector,
						), // Feature
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
				data.UnitID
			},
		) (*struct{ Body types.DeletedResponse }, error) {
			result, errCode, err := controller.Delete(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DeletedResponse }{Body: types.DeletedResponse{AffectedRows: result}}, nil
		},
	)

	// Delete multiple unit
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-unit-multiple",
			Summary:     "Delete multiple unit",
			Description: "Delete multiple unit by providing a list of IDs and return affected rows in database.",
			Method:      http.MethodDelete,
			Path:        fmt.Sprintf("%s/multiple/delete", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecuritySchemeSchoolToken: {},
					constants.SecuritySchemeSchoolID:    {},
					constants.SecuritySchemeBearerToken: {
						fmt.Sprintf("%s,%s",
							constants.FeatureAdmin,
							constants.FeatureDirector,
						), // Feature
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

	// Get unit by id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-unit-id",
			Summary:     "Get unit by id",
			Description: "Return one unit with matching id",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/{id}", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecuritySchemeSchoolToken: {},
					constants.SecuritySchemeSchoolID:    {},
					constants.SecuritySchemeBearerToken: {
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
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				data.UnitID
			},
		) (*struct{ Body data.UnitResponse }, error) {
			result, errCode, err := controller.Get(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.UnitResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Get all unit
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-unit-list",
			Summary:     "Get all unit",
			Description: "Get all unit with support for search, filter and pagination",
			Method:      http.MethodGet,
			Path:        endpointConfig.Group,
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecuritySchemeSchoolToken: {},
					constants.SecuritySchemeSchoolID:    {},
					constants.SecuritySchemeBearerToken: {
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
			Body data.UnitResponseList
		}, error) {
			result, errCode, err := controller.GetAll(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}

			return &struct {
				Body data.UnitResponseList
			}{Body: *result}, nil
		},
	)
}
