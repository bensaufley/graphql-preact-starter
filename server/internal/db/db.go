package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Get() (*sql.DB, error) {
	return sql.Open("sqlite3", "file:/db/data.db?mode=rw")
}
