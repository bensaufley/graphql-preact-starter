// Package graphql provides the core logic for the GraphQL server itself
package graphql

import (
	"net/http"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/graph-gophers/graphql-transport-ws/graphqlws"

	"github.com/bensaufley/graphql-preact-starter/internal/db"
	"github.com/bensaufley/graphql-preact-starter/internal/log"
	"github.com/bensaufley/graphql-preact-starter/internal/schema"
)

type Config struct {
	DB           *db.Config
	SchemaString func() (string, error)
}

func (cfg *Config) NewHandler() (http.HandlerFunc, error) {
	s, err := cfg.SchemaString()
	if err != nil {
		log.Logger.WithError(err).Fatal("could not build schema string")
	}
	sqlite, err := cfg.DB.Get()
	if err != nil {
		log.Logger.WithError(err).Fatal("error initializing database")
	}
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers(), graphql.Logger(log.Logger)}
	schm, err := graphql.ParseSchema(s, schema.NewRoot(sqlite), opts...)
	if err != nil {
		log.Logger.WithError(err).Fatal("could not parse schema")
	}
	return graphqlws.NewHandlerFunc(schm, &relay.Handler{Schema: schm}), nil
}
