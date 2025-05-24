package config

import (
	"api/common/constants"

	"github.com/spf13/viper"
)

type Environment struct {
	// Application config
	AppPort  int    `mapstructure:"APP_PORT"`
	AppName  string `mapstructure:"APP_NAME"`
	Hostname string `mapstructure:"HOST_NAME"`

	// API config
	ApiGroup     string `mapstructure:"API_GROUP"`
	GinMode      string `mapstructure:"GIN_MODE"`
	AllowedHosts string `mapstructure:"ALLOWED_HOSTS"`

	// Redis for fast memory key-value storage
	RedisHost     string `mapstructure:"SERVER_REDIS_HOST"`
	RedisPort     int    `mapstructure:"SERVER_REDIS_PORT"`
	RedisUsername string `mapstructure:"SERVER_REDIS_USER"`
	RedisPassword string `mapstructure:"SERVER_REDIS_PASSWORD"`
	RedisDatabase int    `mapstructure:"SERVER_REDIS_DB"`

	// Postgres database
	PostgresHost     string `mapstructure:"SERVER_POSTGRES_HOST"`
	PostgresPort     int    `mapstructure:"SERVER_POSTGRES_PORT"`
	PostgresUsername string `mapstructure:"SERVER_POSTGRES_USER"`
	PostgresPassword string `mapstructure:"SERVER_POSTGRES_PASSWORD"`
	PostgresDatabase string `mapstructure:"SERVER_POSTGRES_DB"`
	PostgresSslMode  string `mapstructure:"SERVER_POSTGRES_SSL_MODE"`
	PostgresTimeZone string `mapstructure:"SERVER_POSTGRES_TIME_ZONE"`

	// Argon 2id to hash password
	ArgonMemoryLeft  int `mapstructure:"ARGON_PARAM_MEMORY_L"`
	ArgonMemoryRight int `mapstructure:"ARGON_PARAM_MEMORY_R"`
	ArgonIterations  int `mapstructure:"ARGON_PARAM_ITERATIONS"`
	ArgonSaltLength  int `mapstructure:"ARGON_PARAM_SALT_LENGTH"`
	ArgonKeyLength   int `mapstructure:"ARGON_PARAM_KEY_LENGTH"`

	// Jwt issuers auth
	JwtIssuerAuthPassphrase string `mapstructure:"JWT_ISSUER_AUTH_PASSPHRASE"`
	// Jwt issuers session
	JwtIssuerSessionPassphrase       string `mapstructure:"JWT_ISSUER_SESSION_PASSPHRASE"`
	JwtIssuerSessionApiKeyPassphrase string `mapstructure:"JWT_ISSUER_SESSION_API_KEY_PASSPHRASE"`
	// Jwt issuers profile
	JwtIssuerProfileUpdatePasswordPassphrase    string `mapstructure:"JWT_ISSUER_PROFILE_UPDATE_PASSWORD_PASSPHRASE"`
	JwtIssuerProfileUpdateEmailPassphrase       string `mapstructure:"JWT_ISSUER_PROFILE_UPDATE_EMAIL_PASSPHRASE"`
	JwtIssuerProfileUpdatePhoneNumberPassphrase string `mapstructure:"JWT_ISSUER_PROFILE_UPDATE_PHONE_NUMBER_PASSPHRASE"`

	// GOOGLE reCAPTCHA
	GoogleReCAPTCHASiteKey string  `mapstructure:"GOOGLE_RECAPTCHA_SITE_KEY"`
	GoogleReCAPTCHAScore   float32 `mapstructure:"GOOGLE_RECAPTCHA_SCORE"`

	// SMTP
	SmtpHost     string `mapstructure:"SMTP_HOST"`
	SmtpPort     int    `mapstructure:"SMTP_PORT"`
	SmtpUsername string `mapstructure:"SMTP_USERNAME"`
	SmtpPassword string `mapstructure:"SMTP_PASSWORD"`
	SmtpSender   string `mapstructure:"SMTP_SENDER"`

	// Twilio SMS
	TwilioAccountSid   string `mapstructure:"TWILIO_ACCOUNT_SID"`
	TwilioApiKey       string `mapstructure:"TWILIO_API_KEY"`
	TwilioApiSecret    string `mapstructure:"TWILIO_API_SECRET"`
	TwilioSenderNumber string `mapstructure:"TWILIO_SENDER_NUMBER"`

	// Login with Google
	GooglePlusClientID string `mapstructure:"GOOGLE_PLUS_CLIENT_ID"`

	// Login with Facebook
	FacebookAppName       string `mapstructure:"FACEBOOK_APP_NAME"`
	FacebookAppID         string `mapstructure:"FACEBOOK_APP_ID"`
	FacebookClientSecret  string `mapstructure:"FACEBOOK_CLIENT_SECRET"`
	FacebookDebugTokenUrl string `mapstructure:"FACEBOOK_DEBUG_TOKEN_URL"`
	FacebookProfileUrl    string `mapstructure:"FACEBOOK_PROFILE_URL"`

	// Meetings
	MeetingApiUrl    string `mapstructure:"MEETING_API_URL"`
	MeetingApiKey    string `mapstructure:"MEETING_API_KEY"`
	MeetingApiSecret string `mapstructure:"MEETING_API_SECRET"`

	// Database fixtures
	RoleDefault  string `mapstructure:"ROLE_DEFAULT"`
	RoleAdmin    string `mapstructure:"ROLE_ADMIN"`
	RoleDirector string `mapstructure:"ROLE_DIRECTOR"`
	RoleTeacher  string `mapstructure:"ROLE_TEACHER"`
	RoleStudent  string `mapstructure:"ROLE_STUDENT"`
	RoleParent   string `mapstructure:"ROLE_PARENT"`

	UserAdminEmail    string `mapstructure:"USER_ADMIN_EMAIL"`
	UserAdminPassword string `mapstructure:"USER_ADMIN_PASSWORD"`
}

var Env = &Environment{}

// LoadEnv Loads environment variables.
func LoadEnv() error {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err == nil {
		err = viper.Unmarshal(Env)
		if err == nil {
			// Initialize the JWT issuer passphrase after the environment file is loaded
			constants.InitializeJwtIssuerConst(
				Env.JwtIssuerSessionPassphrase,
				Env.JwtIssuerSessionApiKeyPassphrase,
				Env.JwtIssuerAuthPassphrase,
				Env.JwtIssuerProfileUpdatePasswordPassphrase,
				Env.JwtIssuerProfileUpdateEmailPassphrase,
				Env.JwtIssuerProfileUpdatePhoneNumberPassphrase,
			)
		}
	}
	return err
}
