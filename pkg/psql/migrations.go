package psql

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

const (
	DB = "postgres"
)

func MigrateLibUp(connString, migrationFile string) error {
	db, err := sql.Open(DB, connString)
	if err != nil {
		return err
	}
	defer db.Close()
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(migrationFile, DB, driver)
	if err != nil {
		return err
	}
	m.Up()
	return nil
}

// MakeMigrations - makes all migration via goose up
func MakeMigrationsUp(dsn, migrationFile string) error {
	mdb, _ := sql.Open(DB, dsn)
	err := mdb.Ping()
	if err != nil {
		return err
	}
	defer mdb.Close()

	err = goose.Up(mdb, migrationFile)
	if err != nil {
		return err
	}
	return nil
}

func MakeMigrationsDown(dsn, migrationFile string) error {
	mdb, _ := sql.Open(DB, dsn)
	err := mdb.Ping()
	if err != nil {
		return err
	}
	defer mdb.Close()

	err = goose.Down(mdb, migrationFile)
	if err != nil {
		return err
	}
	return nil
}
