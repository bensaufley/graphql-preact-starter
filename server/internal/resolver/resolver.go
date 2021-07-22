package resolver

import (
	"context"
	"database/sql"
)

type QueryResolver struct {
	Db *sql.DB
}

func NewRoot(db *sql.DB) *QueryResolver {
	return &QueryResolver{Db: db}
}

type Bar struct {
	name string
}

func (r QueryResolver) Foo(ctx context.Context) (*Bar, error) {
	return &Bar{}, nil
}

func (b *Bar) Name() string {
	if b == nil {
		return ""
	}
	return b.name
}
