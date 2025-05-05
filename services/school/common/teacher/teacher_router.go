package teacher

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"api/common/constants"
	"api/common/types"
	"api/services/school/common/teacher/data"
)

func RegisterEndpoints(
	humaApi *huma.API,
	controller *Controller,
) {
	var endpointConfig = types.ApiEndpointConfig{
		Group: "/schools/teachers",
		Tag:   []string{"Teachers"},
	}
	const tableName = "teachers"

	// Create teacher
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "post-teacher",
			Summary:     "Create teacher",
			Description: "Create new teacher by providing name and description and return created object. The name teacher should be unique.",
			Method:      http.MethodPost,
			Path:        endpointConfig.Group,
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						constants.FeatureAdmin,     // Feature scope
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
				Body data.TeacherRequest
			},
		) (*struct{ Body data.TeacherResponse }, error) {
			result, errCode, err := controller.Create(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.TeacherResponse }{Body: *result.ToTeacherResponse()}, nil
		},
	)

	// Create teacher tu/class
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "post-teacher-tu/class",
			Summary:     "Create teacher tu/class",
			Description: "Create new teacher teaching unit/class and return created object.",
			Method:      http.MethodPost,
			Path:        fmt.Sprintf("%s/tuclass", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						constants.FeatureAdmin,     // Feature scope
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
				Body data.TUSubjectRequest
			},
		) (*struct {
			Body data.TUSubjectResponse
		}, error) {
			result, errCode, err := controller.CreateTUSubject(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.TUSubjectResponse
			}{Body: *result.ToTUSubjectResponse()}, nil
		},
	)

	// Update teacher with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-teacher",
			Summary:     "Update teacher",
			Description: "Update existing teacher with matching id and return the new object.",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/{id}", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						constants.FeatureAdmin,     // Feature scope
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
				data.TeacherID
				Body data.TeacherRequest
			},
		) (*struct{ Body data.TeacherResponse }, error) {
			result, errCode, err := controller.Update(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.TeacherResponse }{Body: *result.ToTeacherResponse()}, nil
		},
	)

	// Update teacher tu/class with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-teacher-tu/class",
			Summary:     "Update teacher tu/class",
			Description: "Update existing teacher teaching unit/class with matching id and return the new object.",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/tuclass/{id}", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						constants.FeatureAdmin,     // Feature scope
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
				data.TUSubjectID
				Body data.TUSubjectRequest
			},
		) (*struct {
			Body data.TUSubjectResponse
		}, error) {
			result, errCode, err := controller.UpdateTUSubject(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.TUSubjectResponse
			}{Body: *result.ToTUSubjectResponse()}, nil
		},
	)

	// Delete teacher with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-teacher",
			Summary:     "Delete teacher",
			Description: "Delete existing teacher with matching id and return affected rows in database.",
			Method:      http.MethodDelete,
			Path:        fmt.Sprintf("%s/{id}", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						constants.FeatureAdmin,     // Feature scope
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
				data.TeacherID
			},
		) (*struct{ Body types.DeletedResponse }, error) {
			result, errCode, err := controller.Delete(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DeletedResponse }{Body: types.DeletedResponse{AffectedRows: result}}, nil
		},
	)

	// Delete teacher tu/class with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-teacher-tu/class",
			Summary:     "Delete teacher tu/class",
			Description: "Delete existing teacher teaching unit/class with matching id and return affected rows in database.",
			Method:      http.MethodDelete,
			Path:        fmt.Sprintf("%s/tuclass/{id}", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						constants.FeatureAdmin,     // Feature scope
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
				data.TUSubjectID
			},
		) (*struct{ Body types.DeletedResponse }, error) {
			result, errCode, err := controller.DeleteTUSubject(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DeletedResponse }{Body: types.DeletedResponse{AffectedRows: result}}, nil
		},
	)

	// Get teacher by id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-teacher-id",
			Summary:     "Get teacher by id",
			Description: "Return one teacher with matching id",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/{id}", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						constants.FeatureAdmin,   // Feature scope
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
				data.TeacherID
			},
		) (*struct{ Body data.TeacherResponse }, error) {
			result, errCode, err := controller.Get(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.TeacherResponse }{Body: *result.ToTeacherResponse()}, nil
		},
	)

	// Get teacher tu/class by id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-teacher-tu/class-id",
			Summary:     "Get teacher tu/class by id",
			Description: "Return one teacher teaching unit/class with matching id",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/tuclass/{id}", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						constants.FeatureAdmin,   // Feature scope
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
				data.TUSubjectID
			},
		) (*struct {
			Body data.TUSubjectResponse
		}, error) {
			result, errCode, err := controller.GetTUSubject(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.TUSubjectResponse
			}{Body: *result.ToTUSubjectResponse()}, nil
		},
	)

	// Get all teachers
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-teacher-list",
			Summary:     "Get all teachers",
			Description: "Get all teachers with support for search, filter and pagination",
			Method:      http.MethodGet,
			Path:        endpointConfig.Group,
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						constants.FeatureAdmin,   // Feature scope
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
			Body data.TeacherResponseList
		}, error) {
			result, errCode, err := controller.GetAll(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.TeacherResponseList
			}{Body: *result}, nil
		},
	)

	// Get all teachers tu/class
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-teacher-tu/class-list",
			Summary:     "Get all teachers tu/class",
			Description: "Get all teaching unit/class for specified teacher with support for search, filter and pagination",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/{id}/tuclass", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						constants.FeatureAdmin,   // Feature scope
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
				data.GetAllTUSubjectRequest
			},
		) (*struct {
			Body data.TUSubjectResponseList
		}, error) {
			result, errCode, err := controller.GetAllTUSubject(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.TUSubjectResponseList
			}{Body: *result}, nil
		},
	)
}
