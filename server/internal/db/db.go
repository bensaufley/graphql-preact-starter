// Package db provides database-related structs and methods
package db

import (
	_ "embed"
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/bensaufley/graphql-preact-starter/internal/entities"
)

type Config struct {
	DBPath string
}

func (cfg *Config) Get() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("file:%s?mode=rwc", cfg.DBPath)), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("could not open database at %s: %w", cfg.DBPath, err)
	}
	if err := db.AutoMigrate(
		&entities.Todo{},
	); err != nil {
		return nil, err
	}
	return db, nil
}
