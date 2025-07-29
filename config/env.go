package config

import (
	"api/common/constants"

	"github.com/spf13/viper"
)

type Environment struct {
	// Application config
	AppPort int    `mapstructure:"APP_PORT"`
	AppName string `mapstructure:"APP_NAME"`

	// API config
	ApiGroup       string `mapstructure:"API_GROUP"`
	WebsiteBaseURL string `mapstructure:"WEBSITE_BASE_URL"`
	ApiBaseURL     string `mapstructure:"API_BASE_URL"`
	GinMode        string `mapstructure:"GIN_MODE"`
	AllowedHosts   string `mapstructure:"ALLOWED_HOSTS"`

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

	// Web Push
	WebPushSubscriber      string `mapstructure:"WEB_PUSH_SUBSCRIBER"`
	WebPushVapidPublicKey  string `mapstructure:"WEB_PUSH_VAPID_PUBLIC_KEY"`
	WebPushVapidPrivateKey string `mapstructure:"WEB_PUSH_VAPID_PRIVATE_KEY"`

	// SMTP
	SmtpHost        string `mapstructure:"SMTP_HOST"`
	SmtpPort        int    `mapstructure:"SMTP_PORT"`
	SmtpUsername    string `mapstructure:"SMTP_USERNAME"`
	SmtpPassword    string `mapstructure:"SMTP_PASSWORD"`
	SmtpDomainName  string `mapstructure:"SMTP_DOMAIN_NAME"`
	SmtpUserNoReply string `mapstructure:"SMTP_USER_NO_REPLY"`
	SmtpUserSupport string `mapstructure:"SMTP_USER_SUPPORT"`

	// SMS
	SmsAfrikaTalkingApiKey   string `mapstructure:"SMS_AFRIKA_TALKING_API_KEY"`
	SmsAfrikaTalkingUserName string `mapstructure:"SMS_AFRIKA_TALKING_USER_NAME"`
	SmsAfrikaTalkingSenderId string `mapstructure:"SMS_AFRIKA_TALKING_SENDER_ID"`

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
	MeetingJoinUrl   string `mapstructure:"MEETING_JOIN_URL"`
	MeetingApiKey    string `mapstructure:"MEETING_API_KEY"`
	MeetingApiSecret string `mapstructure:"MEETING_API_SECRET"`

	// School api secret
	SchoolApiSecret  string `mapstructure:"SCHOOL_API_SECRET"`
	SchoolApiBaseURL string `mapstructure:"SCHOOL_API_BASE_URL"`
	SchoolCdnUrl     string `mapstructure:"SCHOOL_CDN_URL"`
	SchoolCdnKey     string `mapstructure:"SCHOOL_CDN_KEY"`

	// GIT
	GitRepoSshUrl               string `mapstructure:"GIT_DEPLOY_REPO_SSH_URL"`
	GitRepoBranch               string `mapstructure:"GIT_DEPLOY_REPO_BRANCH"`
	GitRepoSshEd25519PrivateKey string `mapstructure:"GIT_DEPLOY_REPO_SSH_ED25519_PRIVATE_KEY"`

	// Database fixtures
	FixtureRoleDefault  string `mapstructure:"FIXTURE_ROLE_DEFAULT"`
	FixtureRoleAdmin    string `mapstructure:"FIXTURE_ROLE_ADMIN"`
	FixtureRoleDirector string `mapstructure:"FIXTURE_ROLE_DIRECTOR"`
	FixtureRoleTeacher  string `mapstructure:"FIXTURE_ROLE_TEACHER"`
	FixtureRoleStudent  string `mapstructure:"FIXTURE_ROLE_STUDENT"`
	FixtureRoleParent   string `mapstructure:"FIXTURE_ROLE_PARENT"`
	// Fixture user admin
	FixtureUserAdminEmail    string `mapstructure:"FIXTURE_USER_ADMIN_EMAIL"`
	FixtureUserAdminPassword string `mapstructure:"FIXTURE_USER_ADMIN_PASSWORD"`
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
