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
	const tableName = "students"

	// Create student
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "post-student",
			Summary:     "Create student",
			Description: "Create new student by providing name and description and return created object. The name student should be unique.",
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
				Body data.StudentRequest
			},
		) (*struct{ Body data.StudentResponse }, error) {
			result, errCode, err := controller.Create(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.StudentResponse }{Body: *result.ToStudentResponse()}, nil
		},
	)

	// Create student level domain/class
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "post-student-level-domain/class",
			Summary:     "Create student level domain/class",
			Description: "Create new student level domain/class and return created object.",
			Method:      http.MethodPost,
			Path:        fmt.Sprintf("%s/leveldomainclasses", endpointConfig.Group),
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
				Body data.StudentEnrollRequest
			},
		) (*struct {
			Body data.StudentEnrollResponse
		}, error) {
			result, errCode, err := controller.CreateStudentEnroll(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.StudentEnrollResponse
			}{Body: *result.ToStudentEnrollResponse()}, nil
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
					constants.SecurityAuthName: { // Authentication
						fmt.Sprintf("%s,%s",
							constants.FeatureAdmin,
							constants.FeatureDirector,
						), // Features scope
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
			return &struct{ Body data.StudentResponse }{Body: *result.ToStudentResponse()}, nil
		},
	)

	// Update student level domain/class with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-student-level-domain/class",
			Summary:     "Update student level domain/class",
			Description: "Update existing student level domain/class with matching id and return the new object.",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/leveldomainclasses/{id}", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
						fmt.Sprintf("%s,%s",
							constants.FeatureAdmin,
							constants.FeatureDirector,
						), // Features scope
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
			}{Body: *result.ToStudentEnrollResponse()}, nil
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

	// Delete student level domain/class with id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-student-level-domain/class",
			Summary:     "Delete student level domain/class",
			Description: "Delete existing student level domain/class with matching id and return affected rows in database.",
			Method:      http.MethodDelete,
			Path:        fmt.Sprintf("%s/leveldomainclasses/{id}", endpointConfig.Group),
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
				data.StudentID
			},
		) (*struct{ Body data.StudentResponse }, error) {
			result, errCode, err := controller.Get(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.StudentResponse }{Body: *result.ToStudentResponse()}, nil
		},
	)

	// Get student level domain/class by id
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-student-level-domain/class-id",
			Summary:     "Get student level domain/class by id",
			Description: "Return one student level domain/class with matching id",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/leveldomainclasses/{id}", endpointConfig.Group),
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
			}{Body: *result.ToStudentEnrollResponse()}, nil
		},
	)

	// Get all students
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-student-list",
			Summary:     "Get all students",
			Description: "Get all students with support for search, filter and pagination",
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

	// Get all students enrolls
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-student-enrolls-list",
			Summary:     "Get all students enrolls",
			Description: "Get all students enrolls with support for search, filter and pagination",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/enrolls", endpointConfig.Group),
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

	// Get all students level domain/class
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-student-level-domain/class-list",
			Summary:     "Get all students level domain/class",
			Description: "Get all level domain/class for specified student with support for search, filter and pagination",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/leveldomainclasses", endpointConfig.Group),
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
}
