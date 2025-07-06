package main

import (
	"api/common/helpers"
	"api/common/utils/security"
	"api/config"
	configDeploy "api/config/deployment"
	"api/services/school/common/school/model"

	"go.uber.org/zap"
)

// Contains all errors during init() execution
var errInit error

func main() {
	// Check if there are any errors when initializing the app
	if errInit != nil {
		helpers.Logger.Warn(
			"There are some errors when initializing app!",
			zap.String("Error", "Please fix previous errors before."),
		)
		panic(errInit)
	}

	school := &model.School{
		Type:      "university",
		Favicon:   "https://www.google.com/favicon.ico",
		Logo:      "https://www.gstatic.com/marketing-cms/assets/images/c5/3a/200414104c669203c62270f7884f/google-wordmarks-2x.webp=n-w100-h32-fcrop64=1,00000000ffffffff-rw",
		LogoWhite: "https://www.gstatic.com/marketing-cms/assets/images/c5/3a/200414104c669203c62270f7884f/google-wordmarks-2x.webp=n-w100-h32-fcrop64=1,00000000ffffffff-rw",
		Config: &model.SchoolConfig{
			Protocol:            "https",
			DomainName:          "www.uy1.cm",
			DomainCert:          "CERT\nAAA",
			DomainKey:           "KEY\nBBB",
			WebsiteTitle:        "UY1",
			WebsiteDescription:  "School management app",
			ColorPrimary:        "#111111",
			ColorPrimaryBg:      "#F1F1F1",
			ColorPrimaryBgHover: "#D1D1D1",
		},
	}
	school.ID = 2
	configDeploy.DeploySchool(school)
	// configDeploy.DeleteSchoolDeployment(1)
	// configDeploy.DeleteSchoolDeployment(2)

	/*
		// Migrate
		err := migrate.Apply()
		if err != nil {
			panic(err)
		}
		// Load fixtures
		err = fixture.Load()
		if err != nil {
			panic(err)
		}

		di.InjectDependencies()
		api.Start()
	*/
}

// Called before the main entry point. It's useful for setting up
// configurations before starting the application.
func init() {
	helpers.EnableLogger()

	// Load env
	errEnv := config.LoadEnv()
	if errEnv != nil {
		errInit = errEnv
		helpers.Logger.Warn(
			"Failed to load env!",
			zap.String("Error", errEnv.Error()),
		)
	} else {
		helpers.Logger.Info("Env loaded!")
	}

	// Load jwt keys
	errKeys := config.LoadKeys()
	if errKeys != nil {
		errInit = errKeys
		helpers.Logger.Warn(
			"Failed to load keys!",
			zap.String("Error", errKeys.Error()),
		)
	} else {
		helpers.Logger.Info("Keys loaded!")
	}

	// Test Argon 2id with an empty password to ensure that everything works as expected
	_, errArgon2id := security.EncodeArgon2id("Testing")
	if errArgon2id != nil {
		errInit = errArgon2id
		helpers.Logger.Warn(
			"Failed to initialize argon2id!",
			zap.String("Error", errArgon2id.Error()),
		)
	} else {
		helpers.Logger.Info("Argon2id initialized ok!")
	}

	// Connect redis
	errRedis := config.ConnectRedis()
	if errRedis != nil {
		errInit = errRedis
		helpers.Logger.Warn(
			"Failed to connect to Redis!",
			zap.String("Error", errRedis.Error()),
		)
	} else {
		helpers.Logger.Info("Connected to Redis!")
	}

	// Connect database
	errDB := config.ConnectDatabase()
	if errDB != nil {
		errInit = errDB
		helpers.Logger.Warn(
			"Failed to connect to database!",
			zap.String("Error", errDB.Error()),
		)
	} else {
		helpers.Logger.Info("Connected to database!")
	}

	// Setup SMS
	errSms := config.SetupSMS()
	if errSms != nil {
		errInit = errSms
		helpers.Logger.Warn(
			"Failed to setup SMS!",
			zap.String("Error", errSms.Error()),
		)
	} else {
		helpers.Logger.Info("SMS configured!")
	}

	// Load OpenAPI templates
	errOpenAPITemplates := config.LoadOpenAPITemplates()
	if errOpenAPITemplates != nil {
		errInit = errOpenAPITemplates
		helpers.Logger.Warn(
			"Failed to load OpenAPI templates!",
			zap.String("Error", errOpenAPITemplates.Error()),
		)
	} else {
		helpers.Logger.Info("OpenAPI templates loaded!")
	}
}
