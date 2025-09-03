package migrate

import "api/cmd/migrate/migrations"

func ApplyManually() error {
	err := migrations.M1()
	return err
}
