// Package entities contains all entities or models that the
// API may expose.
package entities

import (
	"time"

	"gorm.io/gorm"

	"github.com/bensaufley/graphql-preact-starter/internal/ulid"
)

// Base is the core GORM model that all GORM entities should embed.
type Base struct {
	ULID      string `gorm:"column:ulid;size:32;primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (b *Base) BeforeCreate(*gorm.DB) error {
	gen := ulid.NewGenerator()
	id := gen.String()
	b.ULID = id

	return nil
}
