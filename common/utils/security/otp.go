package securityUtil

import (
	"fmt"

	"api/common/constants"
	"api/config"

	"github.com/pquerna/otp/totp"
)

func GenerateOTP(userId int64, username string, issuer string) (string, error) {
	// Generate OTP secret if not already generated
	secret, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: username,
	})
	if err != nil {
		return "", constants.Http500ErrorMessage("generate TOTP code")
	}

	// Save secret on redis
	err = config.SetRedisString(GetJWTCachedKey(userId, issuer), secret.Secret())
	if err != nil {
		return "", err
	}

	// Construct the OTP URL for generating QR code
	otpURL := fmt.Sprintf("otpauth://totp/%s:%s?secret=%s&issuer=%s",
		config.Env.AppName,
		username,
		secret.Secret(),
		issuer,
	)

	// Generate QR code
	// TODO generate QR code with otpURL

	return otpURL, nil
}

// ValidateOTP Validates the OTP code
func ValidateOTP(otpCode int, userId int64, issuer string) bool {
	userSecret, err := config.GetRedisString(GetJWTCachedKey(userId, issuer))
	if err != nil {
		return false
	}
	return totp.Validate(fmt.Sprintf("%d", otpCode), userSecret)
}
