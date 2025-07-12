package constants

type MailTemplate struct {
	Subject string
	Title   string
	Message string
}

var (
	// Verify email
	MailVerifyEmailCheckCode = MailTemplate{
		Subject: "Verify your email address",
		Title:   "Verify your email address",
		Message: "Thanks for starting the new account creation process. We want to make sure it's really you. Please enter the following verification code when prompted. If you don't want to create an account, you can ignore this message.",
	}
	MailWelcomeVerifiedEmail = MailTemplate{
		Subject: "Successfully activated your account",
		Title:   "Welcome to our platform",
		Message: "Your account has been successfully activated! You can now start using your account and enjoy all the features we have to offer. If you have any questions or need assistance, please don't hesitate to contact our support team. Thank you!",
	}
	// Forgot password
	MailForgotPasswordCheckCode = MailTemplate{
		Subject: "Forgot your password",
		Title:   "Forgot your password",
		Message: "You have requested to reset your password. Please enter the following verification code when prompted. If you did not request a password reset, please contact our support team.",
	}
	// Update password
	MailUpdatePasswordCheckCode = MailTemplate{
		Subject: "Update your password",
		Title:   "Update your password",
		Message: "You have requested to update your password. Please enter the following verification code when prompted. If you did not request a password update, please contact our support team.",
	}
	MailUpdatePasswordSuccess = MailTemplate{
		Subject: "Your password has been successfully updated",
		Title:   "Your password has been successfully updated",
		Message: "Your password has been recently updated. If you did not request a password update, please contact our support team.",
	}
	// Update phone nmber
	MailUpdatePhoneNumberCheckCode = MailTemplate{
		Subject: "Update your phone number",
		Title:   "Update your phone number",
		Message: "You have requested to update your phone number. Please enter the following verification code when prompted. If you did not request a phone number update, please contact our support team.",
	}
	// Update mfa email
	MailUpdateMfaEmailCheckCode = MailTemplate{
		Subject: "Update your Mfa email",
		Title:   "Update your Mfa email",
		Message: "You have requested to update your Mfa email. Please enter the following verification code when prompted. If you did not request a Mfa email update, please contact our support team.",
	}
)
