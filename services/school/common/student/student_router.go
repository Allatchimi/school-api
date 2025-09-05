package student

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"api/common/constants"
	"api/common/types"
	"api/services/school/common/student/data"
)

func RegisterEndpoints(
	humaApi *huma.API,
	controller *Controller,
) {
	var endpointConfig = types.ApiEndpointConfig{
		Group: "/schools/students",
		Tag:   []string{"Students"},
	}
	const tableName = constants.RESOURCE_TABLE_STUDENT

	// Create student
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "post-student",
			Summary:     "Create student",
			Description: "Create new student and return created object.",
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
				Body data.StudentRequest
			},
		) (*struct{ Body data.StudentResponse }, error) {
			result, errCode, err := controller.Create(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.StudentResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Create student enroll
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "post-student-enroll",
			Summary:     "Create student enroll",
			Description: "Create new student enroll and return created object.",
			Method:      http.MethodPost,
			Path:        fmt.Sprintf("%s/enrolls", endpointConfig.Group),
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
				Body data.StudentEnrollRequest
			},
		) (*struct{ Body data.StudentEnrollResponse }, error) {
			result, errCode, err := controller.CreateStudentEnroll(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.StudentEnrollResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Create student pre enroll
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "post-student--pre-enroll",
			Summary:     "Create student pre enroll",
			Description: "Create new student pre enroll and return created object.",
			Method:      http.MethodPost,
			Path:        fmt.Sprintf("%s/enrolls/pre", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecuritySchemeSchoolToken: {},
					constants.SecuritySchemeSchoolID:    {},
					constants.SecuritySchemeBearerToken: {
						fmt.Sprintf("%s",
							constants.FeatureDefault,
						), // Feature
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
				Body data.StudentPreEnrollRequest
			},
		) (*struct{ Body data.StudentPreEnrollResponse }, error) {
			result, errCode, err := controller.CreateStudentPreEnroll(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.StudentPreEnrollResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Update student with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-student",
			Summary:     "Update student",
			Description: "Update existing student with matching id and return the new object.",
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
				data.StudentID
				Body data.StudentRequest
			},
		) (*struct{ Body data.StudentResponse }, error) {
			result, errCode, err := controller.Update(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.StudentResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Update student enroll with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-student-enroll",
			Summary:     "Update student enroll",
			Description: "Update existing student enroll with matching id and return the new object.",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/enrolls/{id}", endpointConfig.Group),
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
				data.StudentEnrollID
				Body data.StudentEnrollRequest
			},
		) (*struct {
			Body data.StudentEnrollResponse
		}, error) {
			result, errCode, err := controller.UpdateStudentEnroll(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.StudentEnrollResponse
			}{Body: *result.ToResponse()}, nil
		},
	)

	// Update student pre enroll with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-student-pre-enroll",
			Summary:     "Update student pre enroll",
			Description: "Update existing student pre enroll with matching id and return the new object.",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/enrolls/pre/{id}", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecuritySchemeSchoolToken: {},
					constants.SecuritySchemeSchoolID:    {},
					constants.SecuritySchemeBearerToken: {
						fmt.Sprintf("%s",
							constants.FeatureDefault,
						), // Feature
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
				data.StudentPreEnrollID
				Body data.StudentPreEnrollRequest
			},
		) (*struct {
			Body data.StudentPreEnrollResponse
		}, error) {
			result, errCode, err := controller.UpdateStudentPreEnroll(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.StudentPreEnrollResponse
			}{Body: *result.ToResponse()}, nil
		},
	)

	// Update student pre enroll status with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-student-pre-enroll-status",
			Summary:     "Update student pre enroll status",
			Description: "Update existing student pre enroll status with matching id and return the new object.",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/enrolls/pre/{id}/status", endpointConfig.Group),
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
				data.StudentPreEnrollID
				Body data.StudentPreEnrollStatusRequest
			},
		) (*struct {
			Body data.StudentPreEnrollResponse
		}, error) {
			result, errCode, err := controller.UpdateStudentPreEnrollStatus(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.StudentPreEnrollResponse
			}{Body: *result.ToResponse()}, nil
		},
	)

	// Delete student with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-student",
			Summary:     "Delete student",
			Description: "Delete existing student with matching id and return affected rows in database.",
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
				data.StudentID
			},
		) (*struct{ Body types.DeletedResponse }, error) {
			result, errCode, err := controller.Delete(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DeletedResponse }{Body: types.DeletedResponse{AffectedRows: result}}, nil
		},
	)

	// Delete student enroll with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-student-enroll",
			Summary:     "Delete student enroll",
			Description: "Delete existing student enroll with matching id and return affected rows in database.",
			Method:      http.MethodDelete,
			Path:        fmt.Sprintf("%s/enrolls/{id}", endpointConfig.Group),
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
				data.StudentEnrollID
			},
		) (*struct{ Body types.DeletedResponse }, error) {
			result, errCode, err := controller.DeleteStudentEnroll(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DeletedResponse }{Body: types.DeletedResponse{AffectedRows: result}}, nil
		},
	)

	// Delete student pre enroll with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-student-pre-enroll",
			Summary:     "Delete student pre enroll",
			Description: "Delete existing student pre enroll with matching id and return affected rows in database.",
			Method:      http.MethodDelete,
			Path:        fmt.Sprintf("%s/enrolls/pre/{id}", endpointConfig.Group),
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
				data.StudentPreEnrollID
			},
		) (*struct{ Body types.DeletedResponse }, error) {
			result, errCode, err := controller.DeleteStudentPreEnroll(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DeletedResponse }{Body: types.DeletedResponse{AffectedRows: result}}, nil
		},
	)

	// Delete multiple student
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-student-multiple",
			Summary:     "Delete multiple student",
			Description: "Delete multiple student by providing a list of IDs and return affected rows in database.",
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

	// Delete multiple student enroll
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-student-enroll-multiple",
			Summary:     "Delete multiple student enroll",
			Description: "Delete multiple student enroll by providing a list of IDs and return affected rows in database.",
			Method:      http.MethodDelete,
			Path:        fmt.Sprintf("%s/enrolls/multiple/delete", endpointConfig.Group),
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
			result, errCode, err := controller.DeleteMultipleStudentEnroll(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DeletedResponse }{Body: types.DeletedResponse{AffectedRows: result}}, nil
		},
	)

	// Delete multiple student pre enroll
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-student-pre-enroll-multiple",
			Summary:     "Delete multiple student pre enroll",
			Description: "Delete multiple student pre enroll by providing a list of IDs and return affected rows in database.",
			Method:      http.MethodDelete,
			Path:        fmt.Sprintf("%s/enrolls/pre/multiple/delete", endpointConfig.Group),
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
			result, errCode, err := controller.DeleteMultipleStudentPreEnroll(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DeletedResponse }{Body: types.DeletedResponse{AffectedRows: result}}, nil
		},
	)

	// Get student by id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-student-id",
			Summary:     "Get student by id",
			Description: "Return one student with matching id",
			Method:      http.MethodGet,
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
				data.StudentID
			},
		) (*struct{ Body data.StudentResponse }, error) {
			result, errCode, err := controller.Get(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.StudentResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Get student enroll by id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-student-enroll-id",
			Summary:     "Get student enroll by id",
			Description: "Return one student enroll with matching id",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/enrolls/{id}", endpointConfig.Group),
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
				data.StudentEnrollID
			},
		) (*struct {
			Body data.StudentEnrollResponse
		}, error) {
			result, errCode, err := controller.GetStudentEnroll(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.StudentEnrollResponse
			}{Body: *result.ToResponse()}, nil
		},
	)

	// Get student pre enroll by id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-student-pre-enroll-id",
			Summary:     "Get student pre enroll by id",
			Description: "Return one student pre enroll with matching id",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/enrolls/pre/{id}", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecuritySchemeSchoolToken: {},
					constants.SecuritySchemeSchoolID:    {},
					constants.SecuritySchemeBearerToken: {
						fmt.Sprintf("%s,%s,%s",
							constants.FeatureAdmin,
							constants.FeatureDirector,
							constants.FeatureDefault,
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
				data.StudentPreEnrollID
			},
		) (*struct {
			Body data.StudentPreEnrollResponse
		}, error) {
			result, errCode, err := controller.GetStudentPreEnroll(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.StudentPreEnrollResponse
			}{Body: *result.ToResponse()}, nil
		},
	)

	// Get all student
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-student-list",
			Summary:     "Get all student",
			Description: "Get all student with support for search, filter and pagination",
			Method:      http.MethodGet,
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
			Body data.StudentResponseList
		}, error) {
			result, errCode, err := controller.GetAll(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}

			return &struct {
				Body data.StudentResponseList
			}{Body: *result}, nil
		},
	)

	// Get all student enroll
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-student-enroll-list",
			Summary:     "Get all student enroll",
			Description: "Get all student enroll with support for search, filter and pagination",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/enrolls", endpointConfig.Group),
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
				data.GetAllStudentEnrollRequest
			},
		) (*struct {
			Body data.StudentEnrollResponseList
		}, error) {
			result, errCode, err := controller.GetAllStudentEnroll(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}

			return &struct {
				Body data.StudentEnrollResponseList
			}{Body: *result}, nil
		},
	)

	// Get all student pre enroll
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-student-pre-enroll-list",
			Summary:     "Get all student pre enroll",
			Description: "Get all student pre enroll with support for search, filter and pagination",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/enrolls/pre", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecuritySchemeSchoolToken: {},
					constants.SecuritySchemeSchoolID:    {},
					constants.SecuritySchemeBearerToken: {
						fmt.Sprintf("%s,%s,%s",
							constants.FeatureAdmin,
							constants.FeatureDirector,
							constants.FeatureDefault,
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
				data.GetAllStudentPreEnrollRequest
			},
		) (*struct {
			Body data.StudentPreEnrollResponseList
		}, error) {
			result, errCode, err := controller.GetAllStudentPreEnroll(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}

			return &struct {
				Body data.StudentPreEnrollResponseList
			}{Body: *result}, nil
		},
	)

	// Get all student public
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "get-student-list-public",
			Summary:       "Get all student public",
			Description:   "Get all student public with support for search, filter and pagination",
			Method:        http.MethodGet,
			Path:          fmt.Sprintf("%s/public", endpointConfig.Group),
			Tags:          endpointConfig.Tag,
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
			Body data.StudentResponseList
		}, error) {
			result, errCode, err := controller.GetAll(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}

			return &struct {
				Body data.StudentResponseList
			}{Body: *result}, nil
		},
	)
}
