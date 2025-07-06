package auth

import (
	"context"
	"fmt"
	"net/http"

	"api/common/constants"
	"api/common/types"
	"api/services/user/auth/data"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterEndpoints(
	humaApi *huma.API,
	controller *Controller,
) {
	endpointConfig := types.ApiEndpointConfig{
		Group: "/auth",
		Tag:   []string{"Authentication"},
	}

	// Login with email
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "login-email",
			Summary:       "Login with email",
			Description:   "Login user with email and password. Account need to be activated to retrieve OK response.",
			Method:        http.MethodPost,
			Path:          fmt.Sprintf("%s/login/email", endpointConfig.Group),
			Tags:          endpointConfig.Tag,
			MaxBodyBytes:  constants.DefaultBodySize,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				data.LoginDevice
				Body data.LoginWithEmailRequest
			},
		) (*struct{ Body data.LoginResponse }, error) {
			result, errCode, err := controller.LoginWithEmail(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.LoginResponse }{Body: *result}, nil
		},
	)

	// Login with provider
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "login-provider",
			Summary:       "Login with provider",
			Description:   "Login user with a provider(Google, Facebook, ...) and token.",
			Method:        http.MethodPost,
			Path:          fmt.Sprintf("%s/login/provider", endpointConfig.Group),
			Tags:          endpointConfig.Tag,
			MaxBodyBytes:  constants.DefaultBodySize,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				data.LoginDevice
				Body data.LoginWithProviderRequest
			},
		) (*struct{ Body data.LoginResponse }, error) {
			result, errCode, err := controller.LoginWithProvider(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.LoginResponse }{Body: *result}, nil
		},
	)

	// Register with email
	huma.Register(
		*humaApi,
		huma.Operation{
			Hidden:        true,
			OperationID:   "register-email",
			Summary:       "Register with email",
			Description:   "Register new user with email and password.",
			Method:        http.MethodPost,
			Path:          fmt.Sprintf("%s/register/email", endpointConfig.Group),
			Tags:          endpointConfig.Tag,
			MaxBodyBytes:  constants.DefaultBodySize,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				Body data.RegisterWithEmailRequest
			},
		) (*struct{ Body data.RegisterResponse }, error) {
			result, errCode, err := controller.RegisterWithEmail(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.RegisterResponse }{Body: *result}, nil
		},
	)

	// Activate account
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "activate",
			Summary:       "Activate",
			Description:   "Activate user account.",
			Method:        http.MethodPost,
			Path:          fmt.Sprintf("%s/activate", endpointConfig.Group),
			Tags:          endpointConfig.Tag,
			MaxBodyBytes:  constants.DefaultBodySize,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				Body data.ActivateAccountRequest
			},
		) (*struct{ Body data.ActivateAccountResponse }, error) {
			result, errCode, err := controller.ActivateAccount(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.ActivateAccountResponse }{Body: *result}, nil
		},
	)

	// Forgot password step 1 with email
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "forgot-password-init-email",
			Summary:       "Forgot step 1 - email",
			Description:   "Forgot password step 1 initialize request with email.",
			Method:        http.MethodPost,
			Path:          fmt.Sprintf("%s/forgot/initemail", endpointConfig.Group),
			Tags:          endpointConfig.Tag,
			MaxBodyBytes:  constants.DefaultBodySize,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				Body data.ForgotPasswordWithEmailInitRequest
			},
		) (*struct {
			Body data.ForgotPasswordInitResponse
		}, error) {
			result, errCode, err := controller.ForgotPasswordEmailInit(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.ForgotPasswordInitResponse
			}{Body: *result}, nil
		},
	)

	// Forgot password step 2
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "forgot-password-code",
			Summary:       "Forgot step 2",
			Description:   "Forgot password step 2 validate your request with your received(email) code and token from step 1.",
			Method:        http.MethodPost,
			Path:          fmt.Sprintf("%s/forgot/checkcode", endpointConfig.Group),
			Tags:          endpointConfig.Tag,
			MaxBodyBytes:  constants.DefaultBodySize,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				Body data.ForgotPasswordCodeRequest
			},
		) (*struct {
			Body data.ForgotPasswordCodeResponse
		}, error) {
			result, errCode, err := controller.ForgotPasswordCode(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.ForgotPasswordCodeResponse
			}{Body: *result}, nil
		},
	)

	// Forgot password step 3
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "forgot-password-new-password",
			Summary:       "Forgot step 3",
			Description:   "Forgot password step 3 set your new password by providing a token received from step 2.",
			Method:        http.MethodPost,
			Path:          fmt.Sprintf("%s/forgot/newpassword", endpointConfig.Group),
			Tags:          endpointConfig.Tag,
			MaxBodyBytes:  constants.DefaultBodySize,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				Body data.ForgotPasswordNewPasswordRequest
			},
		) (*struct {
			Body data.ForgotPasswordNewPasswordResponse
		}, error) {
			result, errCode, err := controller.ForgotPasswordNewPassword(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.ForgotPasswordNewPasswordResponse
			}{Body: *result}, nil
		},
	)

	// Logout
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "logout",
			Summary:     "Logout",
			Description: "Logout user with provided token.",
			Method:      http.MethodPost,
			Path:        fmt.Sprintf("%s/logout", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{constants.SecuritySchemeBearerToken: {}}, // Used to require authentication
			},
			MaxBodyBytes:  constants.DefaultBodySize,
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct{},
		) (*struct{ Body data.LogoutResponse }, error) {
			result, errCode, err := controller.Logout(&ctx)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body data.LogoutResponse }{Body: *result}, nil
		},
	)
}
