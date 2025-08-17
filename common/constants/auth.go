package constants

// API security name
const SecuritySchemeBearerToken = "BearerToken"
const SecuritySchemeSchoolID = "SchoolID"
const SecuritySchemeSchoolToken = "SchoolToken"

// Authentication - login methods
const AuthLoginMethodDefault = "default"
const AuthLoginMethodProvider = "provider"

// Authentication - providers
const AuthProviderGoogle = "google"
const AuthProviderFacebook = "facebook"

var AuthProviders = []string{
	AuthProviderGoogle,
	AuthProviderFacebook,
}
var AuthLoginWithFacebookRequiredScopes = []string{
	"email",
	"public_profile",
}

// Multiple Factor Authentication
const AuthMfaMethodEmail = "email"
const AuthMfaMethodAuthenticator = "authenticator"

var AuthMfaMethods = []string{
	AuthMfaMethodEmail,
	AuthMfaMethodAuthenticator,
}

// JWT issuers for user session
var JwtIssuerSession string
var JwtIssuerSessionApiKey string

// JWT issuers for authentication
var JwtIssuerAuthActivate string
var JwtIssuerAuthForgotPasswordCode string
var JwtIssuerAuthForgotPasswordNewPassword string
var JwtIssuerAuthList []string

// JWT issuers for user profile
var JwtIssuerProfileUpdatePasswordCode string
var JwtIssuerProfileUpdatePasswordNewPassword string
var JwtIssuerProfileUpdatePhoneNumberCode string
var JwtIssuerProfileUpdatePhoneNumberNewPhoneNumber string
var JwtIssuerProfileUpdateMfaEmailCode string
var JwtIssuerProfile []string

// InitializeJwtIssuerConst Initializes the JWT issuer with the provided passphrase.
// Needs to be called inside the "init" function in the "main.go" file.
func InitializeJwtIssuerConst(
	sessionPassPhrase string,
	sessionApiKeyPassPhrase string,
	sessionAuthPassPhrase string,
	profileUpdatePasswordPassPhrase string,
	profileUpdateEmailPassPhrase string,
	profileUpdatePhoneNumberPassPhrase string,
) {
	// Auth issuers
	JwtIssuerSession = sessionPassPhrase + "issuer_session"
	JwtIssuerSessionApiKey = sessionApiKeyPassPhrase + "issuer_session_api_key"
	JwtIssuerAuthActivate = sessionAuthPassPhrase + "issuer_auth_activate"
	JwtIssuerAuthForgotPasswordCode = sessionAuthPassPhrase + "issuer_auth_forgot_password_code"
	JwtIssuerAuthForgotPasswordNewPassword = sessionAuthPassPhrase + "issuer_auth_forgot_password_new_password"
	JwtIssuerAuthList = []string{
		JwtIssuerAuthActivate,
		JwtIssuerAuthForgotPasswordCode,
		JwtIssuerAuthForgotPasswordNewPassword,
	}

	// Profile issuers
	JwtIssuerProfileUpdatePasswordCode = profileUpdatePasswordPassPhrase + "issuer_profile_update_password_code"
	JwtIssuerProfileUpdatePasswordNewPassword = profileUpdatePasswordPassPhrase + "issuer_profile_update_password_code"
	JwtIssuerProfileUpdatePhoneNumberCode = profileUpdatePhoneNumberPassPhrase + "issuer_profile_update_phone_number_code"
	JwtIssuerProfileUpdatePhoneNumberNewPhoneNumber = profileUpdatePhoneNumberPassPhrase + "issuer_profile_update_phone_number_new_phone_number"
	JwtIssuerProfileUpdateMfaEmailCode = profileUpdateEmailPassPhrase + "issuer_profile_update_mfa_email_code"
	JwtIssuerProfile = []string{
		JwtIssuerProfileUpdatePasswordCode,
		JwtIssuerProfileUpdatePasswordNewPassword,
		JwtIssuerProfileUpdatePhoneNumberCode,
		JwtIssuerProfileUpdatePhoneNumberNewPhoneNumber,
		JwtIssuerProfileUpdateMfaEmailCode,
	}
}
