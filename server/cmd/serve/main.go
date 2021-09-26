package main

import (
	"net/http"

	"github.com/bensaufley/graphql-preact-starter/internal/db"
	"github.com/bensaufley/graphql-preact-starter/internal/graphql"
	"github.com/bensaufley/graphql-preact-starter/internal/log"
	"github.com/bensaufley/graphql-preact-starter/internal/schema"
	"github.com/bensaufley/graphql-preact-starter/internal/server"
)

func main() {
	cfg := &graphql.Config{
		DB: &db.Config{
			DBPath: "/storage/data.db",
		},
		SchemaString: schema.String,
	}

	mux, err := server.New(cfg)
	if err != nil {
		log.Logger.WithError(err).Fatal("could noot initialize server")
	}

	log.Logger.Info("server listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Logger.WithError(err).Fatal("could not start server")
	}
}
