package db

import (
	"database/sql"
	_ "embed"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

func Get() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "file:/db/data.db?mode=rwc")
	if err != nil {
		return nil, errors.Wrap(err, "could not open database")
	}
	dbDriver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "could not initialize migration driver")
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"sqlite3",
		dbDriver,
	)
	if err != nil {
		return nil, errors.Wrap(err, "could not initialize migrator")
	}
	err = m.Up()
	if err != nil {
		return nil, errors.Wrap(err, "could not run migrations")
	}
	return db, nil
}
