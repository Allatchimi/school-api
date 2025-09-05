package course

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"api/common/constants"
	"api/common/types"
	"api/services/school/common/course/data"
)

func RegisterEndpoints(
	humaApi *huma.API,
	controller *Controller,
) {
	var endpointConfig = types.ApiEndpointConfig{
		Group: "/schools/courses",
		Tag:   []string{"Courses"},
	}
	const tableName = constants.RESOURCE_TABLE_COURSE

	// Create course
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "post-course",
			Summary:     "Create course",
			Description: "Create new course and return created object.",
			Method:      http.MethodPost,
			Path:        endpointConfig.Group,
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecuritySchemeSchoolToken: {},
					constants.SecuritySchemeSchoolID:    {},
					constants.SecuritySchemeBearerToken: {
						fmt.Sprintf("%s,%s,%s",
							constants.FeatureAdmin,
							constants.FeatureDirector,
							constants.FeatureTeacher,
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
				Body data.CourseRequest
			},
		) (*struct{ Body data.CourseResponse }, error) {
			result, errCode, err := controller.Create(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.CourseResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Create course comment
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "post-course-comment",
			Summary:     "Create course comment",
			Description: "Create new course comment and return created object.",
			Method:      http.MethodPost,
			Path:        fmt.Sprintf("%s/{id}/comments", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecuritySchemeSchoolToken: {},
					constants.SecuritySchemeSchoolID:    {},
					constants.SecuritySchemeBearerToken: {
						fmt.Sprintf("%s,%s,%s,%s",
							constants.FeatureAdmin,
							constants.FeatureDirector,
							constants.FeatureTeacher,
							constants.FeatureStudent,
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
				data.CourseID
				Body data.CourseCommentRequest
			},
		) (*struct{ Body data.CourseCommentResponse }, error) {
			result, errCode, err := controller.CreateCourseComment(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.CourseCommentResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Update course with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-course",
			Summary:     "Update course",
			Description: "Update existing course with matching id and return the new course object.",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/{id}", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecuritySchemeSchoolToken: {},
					constants.SecuritySchemeSchoolID:    {},
					constants.SecuritySchemeBearerToken: {
						fmt.Sprintf("%s,%s,%s",
							constants.FeatureAdmin,
							constants.FeatureDirector,
							constants.FeatureTeacher,
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
				data.CourseID
				Body data.CourseRequest
			},
		) (*struct{ Body data.CourseResponse }, error) {
			result, errCode, err := controller.Update(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.CourseResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Update course comment with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-course-comment",
			Summary:     "Update course comment",
			Description: "Update existing course comment with matching id and return the new course object.",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/comments/{id}", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecuritySchemeSchoolToken: {},
					constants.SecuritySchemeSchoolID:    {},
					constants.SecuritySchemeBearerToken: {
						fmt.Sprintf("%s,%s,%s,%s",
							constants.FeatureAdmin,
							constants.FeatureDirector,
							constants.FeatureTeacher,
							constants.FeatureStudent,
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
				data.CourseID
				Body data.CourseCommentRequest
			},
		) (*struct{ Body data.CourseCommentResponse }, error) {
			result, errCode, err := controller.UpdateCourseComment(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.CourseCommentResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Delete course with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-course",
			Summary:     "Delete course",
			Description: "Delete existing course with matching id and return affected rows in database.",
			Method:      http.MethodDelete,
			Path:        fmt.Sprintf("%s/{id}", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecuritySchemeSchoolToken: {},
					constants.SecuritySchemeSchoolID:    {},
					constants.SecuritySchemeBearerToken: {
						fmt.Sprintf("%s,%s,%s",
							constants.FeatureAdmin,
							constants.FeatureDirector,
							constants.FeatureTeacher,
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
				data.CourseID
			},
		) (*struct{ Body types.DeletedResponse }, error) {
			result, errCode, err := controller.Delete(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DeletedResponse }{Body: types.DeletedResponse{AffectedRows: result}}, nil
		},
	)

	// Delete course comment with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-course-comment",
			Summary:     "Delete course comment",
			Description: "Delete existing course comment with matching id and return affected rows in database.",
			Method:      http.MethodDelete,
			Path:        fmt.Sprintf("%s/{id}/comments/{commentID}", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecuritySchemeSchoolToken: {},
					constants.SecuritySchemeSchoolID:    {},
					constants.SecuritySchemeBearerToken: {
						fmt.Sprintf("%s,%s,%s,%s",
							constants.FeatureAdmin,
							constants.FeatureDirector,
							constants.FeatureTeacher,
							constants.FeatureStudent,
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
				data.CourseID
				data.CourseCommentID
			},
		) (*struct{ Body types.DeletedResponse }, error) {
			result, errCode, err := controller.DeleteCourseComment(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DeletedResponse }{Body: types.DeletedResponse{AffectedRows: result}}, nil
		},
	)

	// Delete multiple course
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-course-multiple",
			Summary:     "Delete multiple course",
			Description: "Delete multiple course by providing a list of IDs and return affected rows in database.",
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

	// Get course by id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-course-id",
			Summary:     "Get course by id",
			Description: "Return one course with matching id",
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
				data.CourseID
			},
		) (*struct{ Body data.CourseResponse }, error) {
			result, errCode, err := controller.Get(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.CourseResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Get all course
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-course-list",
			Summary:     "Get all course",
			Description: "Get all course with support for search, filter and pagination",
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
			Body data.CourseResponseList
		}, error) {
			result, errCode, err := controller.GetAll(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}

			return &struct {
				Body data.CourseResponseList
			}{Body: *result}, nil
		},
	)

	// Get all course comment
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-course-comment-list",
			Summary:     "Get all course comment",
			Description: "Get all course comment with support for search, filter and pagination",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/{id}/comments", endpointConfig.Group),
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
				data.CourseID
				data.GetAllCourseCommentRequest
			},
		) (*struct {
			Body data.CourseCommentResponseList
		}, error) {
			result, errCode, err := controller.GetAllCourseComment(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}

			return &struct {
				Body data.CourseCommentResponseList
			}{Body: *result}, nil
		},
	)
}
