package psql

import (
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	envDBURL      = "CONNSTR"
	envMigrations = "MIGRATES"
	envPort       = "PORT"
)

func Migrate(migrateDown bool) error {
	m, err := migrate.New(
		os.Getenv(envMigrations),
		os.Getenv(envDBURL),
	)
	if err != nil {
		return err
	}
	defer m.Close()
	if migrateDown {
		err = m.Down()
	} else {
		err = m.Up()
	}
	return err
}


