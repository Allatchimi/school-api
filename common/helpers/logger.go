package helpers

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

// EnableLogger Enables the logger to print beautiful log messages.
func EnableLogger() {
	if Logger != nil {
		return
	}
	var err error
	Logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

// LogMigrationsWithoutFk Shows custom log message for migrations without foreign key.
func LogMigrationsWithoutFk(err error) {
	if err != nil {
		Logger.Error(
			"Failed to migrate some tables without foreign key!",
			zap.String("Error", err.Error()),
		)
		return
	}
	Logger.Info(
		"Migration done for all tables without foreign key!",
	)
}

// LogMigrations Shows custom log message for migrations with foreign key.
func LogMigrations(err error) {
	if err != nil {
		Logger.Error(
			"Failed to migrate some tables with foreign key!",
			zap.String("Error", err.Error()),
		)
		return
	}
	Logger.Info(
		"Migration done for all tables with foreign key!",
	)
}
