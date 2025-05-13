package config

import (
	"api/common/helpers"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Establishes a connection to the database.
func ConnectDatabase() error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		Env.PostgresHost,
		Env.PostgresUsername,
		Env.PostgresPassword,
		Env.PostgresDatabase,
		Env.PostgresPort,
		Env.PostgresSslMode,
		Env.PostgresTimeZone,
	)
	helpers.Logger.Warn(
		"Database DSN: ",
		zap.String("Value: ", dsn),
	)
	var err error
	DB, err = gorm.Open(
		postgres.New(postgres.Config{
			DSN: dsn,
		}),
		&gorm.Config{},
	)
	return err
}
