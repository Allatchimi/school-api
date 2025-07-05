package profile

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"api/common/constants"
	"api/common/types"
	"api/services/user/profile/data"
	dataUser "api/services/user/user/data"
)

func RegisterEndpoints(
	humaApi *huma.API,
	controller *Controller,
) {
	var endpointConfig = types.ApiEndpointConfig{
		Group: "/profile",
		Tag:   []string{"Profile"},
	}

	// Update profile info
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-profile-info",
			Summary:     "Update profile info",
			Description: "Update profile information such as username, first name, last name, address, ...",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/info", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
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
				Body data.UpdateProfileInfoRequest
			},
		) (*struct{ Body dataUser.UserInfoResponse }, error) {
			result, errCode, err := controller.UpdateProfileInfo(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body dataUser.UserInfoResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Update password init
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-profile-password-init",
			Summary:     "Update profile password init",
			Description: "Update profile password init",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/password/init", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
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
			},
		) (*struct {
			Body data.UpdateProfilePasswordInitResponse
		}, error) {
			result, errCode, err := controller.UpdateProfilePasswordInit(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.UpdateProfilePasswordInitResponse
			}{Body: data.UpdateProfilePasswordInitResponse{Token: result}}, nil
		},
	)

	// Update password check code
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-profile-password-check-code",
			Summary:     "Update profile password check code",
			Description: "Update profile password check code",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/password/checkcode", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
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
				Body data.UpdateProfilePasswordCheckCodeRequest
			},
		) (*struct {
			Body data.UpdateProfilePasswordCheckCodeResponse
		}, error) {
			result, errCode, err := controller.UpdateProfilePasswordCheckCode(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.UpdateProfilePasswordCheckCodeResponse
			}{Body: data.UpdateProfilePasswordCheckCodeResponse{Token: result}}, nil
		},
	)

	// Update password new password
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-profile-password-new-password",
			Summary:     "Update profile password new password",
			Description: "Update profile password new password",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/password/new", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
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
				Body data.UpdateProfilePasswordNewPasswordRequest
			},
		) (*struct {
			Body types.DefaultResponse
		}, error) {
			errCode, err := controller.UpdateProfilePasswordNewPassword(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body types.DefaultResponse
			}{Body: types.DefaultResponse{Message: "Successfully updated profile password"}}, nil
		},
	)

	// Update phone number init
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-profile-phone-number-init",
			Summary:     "Update profile phone number init",
			Description: "Update profile phone number init",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/phonenumber/init", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
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
			},
		) (*struct {
			Body data.UpdateProfilePhoneNumberInitResponse
		}, error) {
			result, errCode, err := controller.UpdateProfilePhoneNumberInit(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.UpdateProfilePhoneNumberInitResponse
			}{Body: data.UpdateProfilePhoneNumberInitResponse{Token: result}}, nil
		},
	)

	// Update phone number check code
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-profile-phone-number-check-code",
			Summary:     "Update profile phone number check code",
			Description: "Update profile phone number check code",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/phonenumber/checkcode", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
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
				Body data.UpdateProfilePhoneNumberCheckCodeRequest
			},
		) (*struct {
			Body data.UpdateProfilePhoneNumberCheckCodeResponse
		}, error) {
			result, errCode, err := controller.UpdateProfilePhoneNumberCheckCode(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.UpdateProfilePhoneNumberCheckCodeResponse
			}{Body: data.UpdateProfilePhoneNumberCheckCodeResponse{Token: result}}, nil
		},
	)

	// Update phone number new phone number
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-profile-phone-number-new-phone-number",
			Summary:     "Update profile phone number new phone number",
			Description: "Update profile phone number new phone number",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/phonenumber/new", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
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
				Body data.UpdateProfilePhoneNumberNewPhoneNumberRequest
			},
		) (*struct {
			Body types.DefaultResponse
		}, error) {
			errCode, err := controller.UpdateProfilePhoneNumberNewPhoneNumber(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body types.DefaultResponse
			}{Body: types.DefaultResponse{Message: "Successfully updated phone number"}}, nil
		},
	)

	// Update mfa email init
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-profile-mfa-email-init",
			Summary:     "Update profile mfa email init",
			Description: "Update profile mfa email init",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/config/mfa/email/init", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
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
		) (*struct {
			Body data.UpdateProfileMfaEmailInitResponse
		}, error) {
			result, errCode, err := controller.UpdateProfileMfaEmailInit(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.UpdateProfileMfaEmailInitResponse
			}{Body: data.UpdateProfileMfaEmailInitResponse{Token: result}}, nil
		},
	)

	// Update mfa email check code
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-profile-mfa-email-check-code",
			Summary:     "Update profile mfa email check code",
			Description: "Update profile mfa email check code",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/config/mfa/email/checkcode", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
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
				Body data.UpdateProfileMfaEmailCheckCodeRequest
			},
		) (*struct {
			Body types.DefaultResponse
		}, error) {
			errCode, err := controller.UpdateProfileMfaEmailCheckCode(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body types.DefaultResponse
			}{Body: types.DefaultResponse{Message: "Successfully updated mfa email"}}, nil
		},
	)

	// Update notification
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-profile-notification",
			Summary:     "Update profile notification",
			Description: "Update profile notification",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/config/notification", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
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
				Body data.UpdateProfileSettingNotificationRequest
			},
		) (*struct {
			Body dataUser.UserResponse
		}, error) {
			result, errCode, err := controller.UpdateProfileNotification(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body dataUser.UserResponse
			}{Body: *result.ToResponse()}, nil
		},
	)

	// Update web push subscription
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-web-push-subscription",
			Summary:     "Update web push subscription",
			Description: "Update web push subscription",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/config/webpush/subscription", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
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
				Body data.UpdateProfileWebPushSubscriptionRequest
			},
		) (*struct{ Body types.DefaultResponse }, error) {
			errCode, err := controller.UpdateWebPushSubscription(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body types.DefaultResponse }{Body: types.DefaultResponse{Message: "Successfully updated web push subscription"}}, nil
		},
	)

	// Get profile
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-profile",
			Summary:     "Get profile info",
			Description: "Retrieve profile information for the current user",
			Method:      http.MethodGet,
			Path:        endpointConfig.Group,
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
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
		) (*struct{ Body dataUser.UserResponse }, error) {
			result, errCode, err := controller.GetProfile(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct{ Body dataUser.UserResponse }{Body: *result.ToResponse()}, nil
		},
	)

	// Get web push subscription public key
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "get-web-push-subscription-public-key",
			Summary:     "Get web push subscription public key",
			Description: "Retrieve web push subscription public key for the current user",
			Method:      http.MethodGet,
			Path:        fmt.Sprintf("%s/config/webpush/publickey", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{
					constants.SecurityAuthName: { // Authentication
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
		) (*struct {
			Body data.GetProfileWebPushSubscriptionResponse
		}, error) {
			result, errCode, err := controller.GetWebPushSubscriptionPublicKey(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error(), err)
			}
			return &struct {
				Body data.GetProfileWebPushSubscriptionResponse
			}{Body: data.GetProfileWebPushSubscriptionResponse{PublicKey: result}}, nil
		},
	)
}
