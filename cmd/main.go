package main

import (
	"api/cmd/api"
	"api/cmd/di"
	"api/cmd/fixture"
	"api/cmd/migrate"
	"api/cmd/test"
	"api/common/helpers"
	securityUtil "api/common/utils/security"
	"api/config"

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

	test.Testssssss()

	di.InjectDependencies()
	api.StartGin()
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
	_, errArgon2id := securityUtil.EncodeArgon2id("Testing")
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
