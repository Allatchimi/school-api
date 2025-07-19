package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	httpHelper "api/common/helpers/http"
	"api/services/user/auth/data"
)

type Controller struct {
	Service *Service
}

func NewAuthController(service *Service) *Controller {
	return &Controller{Service: service}
}

func (controller *Controller) LoginWithEmail(
	ctx *context.Context,
	input *struct {
		data.LoginDevice
		Body data.LoginWithEmailRequest
	},
) (result *data.LoginResponse, errCode int, err error) {
	accessToken, accessExpires, activateAccountToken, errCode, err := controller.Service.Login(
		httpHelper.GetSchoolContext(ctx),
		&data.LoginRequest{
			Email:         input.Body.Email,
			Password:      input.Body.Password,
			StayConnected: input.Body.StayConnected,
		},
		&input.LoginDevice,
	)
	if err != nil && (len(activateAccountToken) < 1 || errCode != http.StatusForbidden) {
		return
	}
	err = nil
	result = &data.LoginResponse{
		AccessToken:          accessToken,
		Expires:              accessExpires,
		ActivateAccountToken: activateAccountToken,
	}
	return
}

func (controller *Controller) LoginWithProvider(
	ctx *context.Context,
	input *struct {
		data.LoginDevice
		Body data.LoginWithProviderRequest
	},
) (result *data.LoginResponse, errCode int, err error) {
	var accessToken string
	var accessExpires *time.Time
	accessToken, accessExpires, errCode, err = controller.Service.LoginWithProvider(
		httpHelper.GetSchoolContext(ctx),
		&input.Body,
		&input.LoginDevice,
	)
	if err != nil {
		return
	}
	result = &data.LoginResponse{
		AccessToken: accessToken,
		Expires:     accessExpires,
	}
	return
}

func (controller *Controller) RegisterWithEmail(
	ctx *context.Context,
	input *struct {
		Body data.RegisterWithEmailRequest
	},
) (result *data.RegisterResponse, errCode int, err error) {
	var activateAccountToken string
	activateAccountToken, errCode, err = controller.Service.Register(
		httpHelper.GetSchoolContext(ctx),
		&data.RegisterRequest{
			Email:    input.Body.Email,
			Password: input.Body.Password,
		},
	)
	if err != nil {
		return
	}
	result = &data.RegisterResponse{
		ActivateAccountToken: activateAccountToken,
		Message:              "Account created! Please activate your account to start using your services.",
	}
	return
}

func (controller *Controller) ActivateAccount(
	ctx *context.Context,
	input *struct {
		Body data.ActivateAccountRequest
	},
) (result *data.ActivateAccountResponse, errCode int, err error) {
	activatedAt, errCode, err := controller.Service.ActivateAccount(&input.Body)
	if err != nil {
		return
	}
	result = &data.ActivateAccountResponse{
		ActivatedAt: activatedAt,
	}
	return
}

func (controller *Controller) ForgotPasswordEmailInit(
	ctx *context.Context,
	input *struct {
		Body data.ForgotPasswordWithEmailInitRequest
	},
) (result *data.ForgotPasswordInitResponse, errCode int, err error) {
	token, errCode, err := controller.Service.ForgotPasswordInit(
		&data.ForgotPasswordInitRequest{
			Email: input.Body.Email,
		},
	)
	if err != nil {
		return
	}
	if len(token) <= 0 {
		errCode = http.StatusInternalServerError
		err = fmt.Errorf("%s", "Failed to start the process! Please try again later.")
		return
	}
	result = &data.ForgotPasswordInitResponse{
		Token: token,
	}
	return
}

func (controller *Controller) ForgotPasswordCode(
	ctx *context.Context,
	input *struct {
		Body data.ForgotPasswordCodeRequest
	},
) (result *data.ForgotPasswordCodeResponse, errCode int, err error) {
	token, errCode, err := controller.Service.ForgotPasswordCode(&input.Body)
	if err != nil {
		return
	}
	result = &data.ForgotPasswordCodeResponse{
		Token: token,
	}
	return
}

func (controller *Controller) ForgotPasswordNewPassword(
	ctx *context.Context,
	input *struct {
		Body data.ForgotPasswordNewPasswordRequest
	},
) (result *data.ForgotPasswordNewPasswordResponse, errCode int, err error) {
	errCode, err = controller.Service.ForgotPasswordNewPassword(&input.Body)
	if err != nil {
		return
	}
	result = &data.ForgotPasswordNewPasswordResponse{
		Message: "Password successful changed! Please sign in to start using our services.",
	}
	return
}

func (controller *Controller) Logout(
	ctx *context.Context,
) (result *data.LogoutResponse, errCode int, err error) {
	errCode, err = controller.Service.Logout(httpHelper.GetJwtContext(ctx), httpHelper.ExtractBearerContext(ctx))
	if err != nil {
		return
	}
	result = &data.LogoutResponse{
		Message: "Successful signed out! See you soon bye.",
	}
	return
}
