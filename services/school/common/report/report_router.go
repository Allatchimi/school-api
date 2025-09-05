package report

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"api/common/constants"
	"api/common/types"
	"api/services/school/common/report/data"
)

func RegisterEndpoints(
	humaApi *huma.API,
	controller *Controller,
) {
	var endpointConfig = types.ApiEndpointConfig{
		Group: "/schools/reports",
		Tag:   []string{"Reports"},
	}
	const tableName = constants.RESOURCE_TABLE_REPORT

	// Create report entry
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "post-report-entry",
			Summary:     "Create report entry",
			Description: "Create new report entry and return created object.",
			Method:      http.MethodPost,
			Path:        fmt.Sprintf("%s/entries", endpointConfig.Group),
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
				Body data.ReportEntryRequest
			},
		) (*struct{ Body data.ReportEntryResponse }, error) {
			result, errCode, err := controller.CreateEntry(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.ReportEntryResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Create report grade
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "post-report-grade",
			Summary:     "Create report grade",
			Description: "Create new report grade and return created object.",
			Method:      http.MethodPost,
			Path:        fmt.Sprintf("%s/grades", endpointConfig.Group),
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
				Body data.ReportGradeRequest
			},
		) (*struct{ Body data.ReportGradeResponse }, error) {
			result, errCode, err := controller.CreateGrade(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.ReportGradeResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Create report correspondence
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "post-report-correspondence",
			Summary:     "Create report correspondence",
			Description: "Create new report correspondence and return created object.",
			Method:      http.MethodPost,
			Path:        fmt.Sprintf("%s/correspondences", endpointConfig.Group),
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
				Body data.ReportCorrespondenceRequest
			},
		) (*struct {
			Body data.ReportCorrespondenceResponse
		}, error) {
			result, errCode, err := controller.CreateCorrespondence(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.ReportCorrespondenceResponse
			}{Body: *result.ToResponse()}, nil
		},
	)

	// Create report config
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "post-report-config",
			Summary:     "Create report config",
			Description: "Create new report config and return created object.",
			Method:      http.MethodPost,
			Path:        fmt.Sprintf("%s/configs", endpointConfig.Group),
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
				Body data.ReportConfigRequest
			},
		) (*struct {
			Body data.ReportConfigResponse
		}, error) {
			result, errCode, err := controller.CreateConfig(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.ReportConfigResponse
			}{Body: *result.ToResponse()}, nil
		},
	)

	// Update report grade with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-report-grade",
			Summary:     "Update report grade",
			Description: "Update existing report grade with matching id and return the new report grade object.",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/grades/{id}", endpointConfig.Group),
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
				data.ReportGradeID
				Body data.ReportGradeRequest
			},
		) (*struct{ Body data.ReportGradeResponse }, error) {
			result, errCode, err := controller.UpdateGrade(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.ReportGradeResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Update report correspondence with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-report-correspondence",
			Summary:     "Update report correspondence",
			Description: "Update existing report correspondence with matching id and return the new report grade object.",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/correspondences/{id}", endpointConfig.Group),
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
				data.ReportCorrespondenceID
				Body data.ReportCorrespondenceRequest
			},
		) (*struct {
			Body data.ReportCorrespondenceResponse
		}, error) {
			result, errCode, err := controller.UpdateCorrespondence(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.ReportCorrespondenceResponse
			}{Body: *result.ToResponse()}, nil
		},
	)

	// Update report config with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-report-config",
			Summary:     "Update report config",
			Description: "Update existing report config with matching id and return the new report config object.",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/configs/{id}", endpointConfig.Group),
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
				data.ReportConfigID
				Body data.ReportConfigRequest
			},
		) (*struct {
			Body data.ReportConfigResponse
		}, error) {
			result, errCode, err := controller.UpdateConfig(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.ReportConfigResponse
			}{Body: *result.ToResponse()}, nil
		},
	)

	// Update report table status with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-report-table-status",
			Summary:     "Update report table status",
			Description: "Update existing report table status with matching id and return the new report config object.",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/tables/{id}/status", endpointConfig.Group),
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
				data.ReportTableID
				Body data.ReportTableStatusRequest
			},
		) (*struct {
			Body data.ReportTableResponse
		}, error) {
			result, errCode, err := controller.UpdateTableStatus(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.ReportTableResponse
			}{Body: *result.ToResponse()}, nil
		},
	)

	// Delete report grade with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-report-grade",
			Summary:     "Delete report grade",
			Description: "Delete existing report grade with matching id and return affected rows in database.",
			Method:      http.MethodDelete,
			Path:        fmt.Sprintf("%s/grades/{id}", endpointConfig.Group),
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
				data.ReportGradeID
			},
		) (*struct{ Body types.DeletedResponse }, error) {
			result, errCode, err := controller.DeleteGrade(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DeletedResponse }{Body: types.DeletedResponse{AffectedRows: result}}, nil
		},
	)

	// Delete report correspondence with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-report-correspondence",
			Summary:     "Delete report correspondence",
			Description: "Delete existing report correspondence with matching id and return affected rows in database.",
			Method:      http.MethodDelete,
			Path:        fmt.Sprintf("%s/correspondences/{id}", endpointConfig.Group),
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
				data.ReportCorrespondenceID
			},
		) (*struct{ Body types.DeletedResponse }, error) {
			result, errCode, err := controller.DeleteCorrespondence(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DeletedResponse }{Body: types.DeletedResponse{AffectedRows: result}}, nil
		},
	)

	// Delete report config with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-report-config",
			Summary:     "Delete report config",
			Description: "Delete existing report config with matching id and return affected rows in database.",
			Method:      http.MethodDelete,
			Path:        fmt.Sprintf("%s/configs/{id}", endpointConfig.Group),
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
				data.ReportConfigID
			},
		) (*struct{ Body types.DeletedResponse }, error) {
			result, errCode, err := controller.DeleteConfig(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DeletedResponse }{Body: types.DeletedResponse{AffectedRows: result}}, nil
		},
	)

	// Delete multiple report grade
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-report-grade-multiple",
			Summary:     "Delete multiple report grade",
			Description: "Delete multiple report grade by providing a list of IDs and return affected rows in database.",
			Method:      http.MethodDelete,
			Path:        fmt.Sprintf("%s/grades/multiple/delete", endpointConfig.Group),
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
			result, errCode, err := controller.DeleteMultipleGrade(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DeletedResponse }{Body: types.DeletedResponse{AffectedRows: result}}, nil
		},
	)

	// Delete multiple report correspondence
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-report-grade-correspondence",
			Summary:     "Delete multiple report correspondence",
			Description: "Delete multiple report correspondence by providing a list of IDs and return affected rows in database.",
			Method:      http.MethodDelete,
			Path:        fmt.Sprintf("%s/correspondences/multiple/delete", endpointConfig.Group),
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
			result, errCode, err := controller.DeleteMultipleCorrespondence(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DeletedResponse }{Body: types.DeletedResponse{AffectedRows: result}}, nil
		},
	)

	// Delete multiple report config
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-report-config-multiple",
			Summary:     "Delete multiple report config",
			Description: "Delete multiple report config by providing a list of IDs and return affected rows in database.",
			Method:      http.MethodDelete,
			Path:        fmt.Sprintf("%s/configs/multiple/delete", endpointConfig.Group),
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
			result, errCode, err := controller.DeleteMultipleConfig(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DeletedResponse }{Body: types.DeletedResponse{AffectedRows: result}}, nil
		},
	)

	// Get report entry by id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-report-entry-id",
			Summary:     "Get report entry by id",
			Description: "Return one report entry with matching id",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/entries/{id}", endpointConfig.Group),
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
				data.ReportEntryID
			},
		) (*struct{ Body data.ReportEntryResponse }, error) {
			result, errCode, err := controller.GetEntry(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.ReportEntryResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Get report grade by id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-report-grade-id",
			Summary:     "Get report grade by id",
			Description: "Return one report grade with matching id",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/grades/{id}", endpointConfig.Group),
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
				data.ReportGradeID
			},
		) (*struct{ Body data.ReportGradeResponse }, error) {
			result, errCode, err := controller.GetGrade(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.ReportGradeResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Get report correspondence by id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-report-correspondence-id",
			Summary:     "Get report correspondence by id",
			Description: "Return one report correspondence with matching id",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/correspondences/{id}", endpointConfig.Group),
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
				data.ReportCorrespondenceID
			},
		) (*struct {
			Body data.ReportCorrespondenceResponse
		}, error) {
			result, errCode, err := controller.GetCorrespondence(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.ReportCorrespondenceResponse
			}{Body: *result.ToResponse()}, nil
		},
	)

	// Get report config by id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-report-config-id",
			Summary:     "Get report config by id",
			Description: "Return one report config with matching id",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/configs/{id}", endpointConfig.Group),
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
				data.ReportConfigID
			},
		) (*struct {
			Body data.ReportConfigResponse
		}, error) {
			result, errCode, err := controller.GetConfig(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.ReportConfigResponse
			}{Body: *result.ToResponse()}, nil
		},
	)

	// Get all report entry
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-report-entry-list",
			Summary:     "Get all report entry",
			Description: "Get all report entry with support for search, filter and pagination",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/entries", endpointConfig.Group),
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
				data.GetAllReportEntryRequest
			},
		) (*struct {
			Body data.ReportEntryResponseList
		}, error) {
			result, errCode, err := controller.GetAllEntry(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}

			return &struct {
				Body data.ReportEntryResponseList
			}{Body: *result}, nil
		},
	)

	// Get all report grade
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-report-grade-list",
			Summary:     "Get all report grade",
			Description: "Get all report grade with support for search, filter and pagination",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/grades", endpointConfig.Group),
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
				data.GetAllReportGradeRequest
			},
		) (*struct {
			Body data.ReportGradeResponseList
		}, error) {
			result, errCode, err := controller.GetAllGrade(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}

			return &struct {
				Body data.ReportGradeResponseList
			}{Body: *result}, nil
		},
	)

	// Get all report correspondence
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-report-correspondence-list",
			Summary:     "Get all report correspondence",
			Description: "Get all report correspondence with support for search, filter and pagination",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/correspondences", endpointConfig.Group),
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
				data.GetAllReportCorrespondenceRequest
			},
		) (*struct {
			Body data.ReportCorrespondenceResponseList
		}, error) {
			result, errCode, err := controller.GetAllCorrespondence(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}

			return &struct {
				Body data.ReportCorrespondenceResponseList
			}{Body: *result}, nil
		},
	)

	// Get all report config
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-report-config-list",
			Summary:     "Get all report config",
			Description: "Get all report config with support for search, filter and pagination",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/configs", endpointConfig.Group),
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
				data.GetAllReportConfigRequest
			},
		) (*struct {
			Body data.ReportConfigResponseList
		}, error) {
			result, errCode, err := controller.GetAllConfig(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}

			return &struct {
				Body data.ReportConfigResponseList
			}{Body: *result}, nil
		},
	)

	// Get all report average
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-report-average-list",
			Summary:     "Get all report average",
			Description: "Get all report average with support for search, filter and pagination",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/averages", endpointConfig.Group),
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
				data.GetAllReportAverageRequest
			},
		) (*struct {
			Body data.ReportAverageResponseList
		}, error) {
			result, errCode, err := controller.GetAllAverage(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}

			return &struct {
				Body data.ReportAverageResponseList
			}{Body: *result}, nil
		},
	)

	// Get all report table
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-report-table-list",
			Summary:     "Get all report table",
			Description: "Get all report table with support for search, filter and pagination",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/tables", endpointConfig.Group),
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
				data.GetAllReportTableRequest
			},
		) (*struct {
			Body data.ReportTableResponseList
		}, error) {
			result, errCode, err := controller.GetAllTable(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}

			return &struct {
				Body data.ReportTableResponseList
			}{Body: *result}, nil
		},
	)
}
