package profile

import (
	"context"

	httpHelper "api/common/helpers/http"
	"api/services/user/profile/data"
	"api/services/user/user/model"
)

type Controller struct {
	Service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{Service: service}
}

// ----------------- Info -----------------
func (controller *Controller) UpdateProfileInfo(
	ctx *context.Context,
	input *struct {
		Body data.UpdateProfileInfoRequest
	},
) (result *model.UserInfo, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateProfileInfo(
		httpHelper.GetJwtContext(ctx),
		&input.Body,
	)
	return
}

// ----------------- Config -----------------
func (controller *Controller) UpdateProfileConfigMessage(
	ctx *context.Context,
	input *struct {
		Body data.UpdateProfileMessageRequest
	},
) (result *model.UserConfig, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateProfileConfigMessage(
		httpHelper.GetJwtContext(ctx),
		&input.Body,
	)
	return
}

// ----------------- Password -----------------
func (controller *Controller) UpdateProfilePasswordInit(
	ctx *context.Context,
	input *struct {
	},
) (result string, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateProfilePasswordInit(httpHelper.GetJwtContext(ctx))
	return
}
func (controller *Controller) UpdateProfilePasswordCheckCode(
	ctx *context.Context,
	input *struct {
		Body data.UpdateProfilePasswordCheckCodeRequest
	},
) (result string, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateProfilePasswordCheckCode(httpHelper.GetJwtContext(ctx), input.Body.Token, input.Body.Code)
	return
}
func (controller *Controller) UpdateProfilePasswordNewPassword(
	ctx *context.Context,
	input *struct {
		Body data.UpdateProfilePasswordNewPasswordRequest
	},
) (errCode int, err error) {
	errCode, err = controller.Service.UpdateProfilePasswordNewPassword(httpHelper.GetJwtContext(ctx), input.Body.Token, input.Body.CurrentPassword, input.Body.NewPassword)
	return
}

// ----------------- Phone number -----------------
func (controller *Controller) UpdateProfilePhoneNumberInit(
	ctx *context.Context,
	input *struct {
	},
) (result string, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateProfilePhoneNumberInit(httpHelper.GetJwtContext(ctx))
	return
}
func (controller *Controller) UpdateProfilePhoneNumberCheckCode(
	ctx *context.Context,
	input *struct {
		Body data.UpdateProfilePhoneNumberCheckCodeRequest
	},
) (result string, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateProfilePhoneNumberCheckCode(httpHelper.GetJwtContext(ctx), input.Body.Token, input.Body.Code)
	return
}
func (controller *Controller) UpdateProfilePhoneNumberNewPhoneNumber(
	ctx *context.Context,
	input *struct {
		Body data.UpdateProfilePhoneNumberNewPhoneNumberRequest
	},
) (errCode int, err error) {
	errCode, err = controller.Service.UpdateProfilePhoneNumberNewPhoneNumber(httpHelper.GetJwtContext(ctx), input.Body.Token, input.Body.PhoneNumber)
	return
}

// ----------------- Mfa Email -----------------
func (controller *Controller) UpdateProfileMfaEmailInit(
	ctx *context.Context,
	input *struct{},
) (result string, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateProfileMfaEmailInit(httpHelper.GetJwtContext(ctx))
	return
}
func (controller *Controller) UpdateProfileMfaEmailCheckCode(
	ctx *context.Context,
	input *struct {
		Body data.UpdateProfileMfaEmailCheckCodeRequest
	},
) (errCode int, err error) {
	errCode, err = controller.Service.UpdateProfileMfaEmailCheckCode(httpHelper.GetJwtContext(ctx), input.Body.Token, input.Body.Code)
	return
}

// ----------------- Notification -----------------
func (controller *Controller) UpdateProfileNotification(
	ctx *context.Context,
	input *struct {
		Body data.UpdateProfileSettingNotificationRequest
	},
) (result *model.User, errCode int, err error) {
	result, errCode, err = controller.Service.UpdateProfileNotification(httpHelper.GetJwtContext(ctx), input.Body.IsEnabled)
	return
}

// ----------------- Web push subscription -----------------
func (controller *Controller) UpdateWebPushSubscription(
	ctx *context.Context,
	input *struct {
		Body data.UpdateProfileWebPushSubscriptionRequest
	},
) (errCode int, err error) {
	errCode, err = controller.Service.UpdateProfileWebPushSubscription(httpHelper.GetJwtContext(ctx), &input.Body)
	if err != nil {
		return
	}
	return
}

// ----------------- Get -----------------
func (controller *Controller) GetProfile(
	ctx *context.Context,
	input *struct{},
) (result *model.User, errCode int, err error) {
	result, errCode, err = controller.Service.GetProfile(httpHelper.GetJwtContext(ctx))
	if err != nil {
		return
	}
	return
}
func (controller *Controller) GetWebPushSubscriptionPublicKey(
	ctx *context.Context,
	input *struct{},
) (result string, errCode int, err error) {
	result, errCode, err = controller.Service.GetWebPushSubscriptionPublicKey(httpHelper.GetJwtContext(ctx))
	if err != nil {
		return
	}
	return
}
