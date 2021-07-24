package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"regexp"

	log "github.com/sirupsen/logrus"

	"github.com/bensaufley/graphql-preact-starter/internal/db"
	"github.com/bensaufley/graphql-preact-starter/internal/graphiql"
	"github.com/bensaufley/graphql-preact-starter/internal/graphql"
	"github.com/bensaufley/graphql-preact-starter/internal/schema"
)

var fs = http.Dir("/public")
var publicRegExp = regexp.MustCompile(`^/public/`)

func handleStatic(w http.ResponseWriter, r *http.Request) {
	p, err := filepath.Rel("/static", r.URL.Path)
	if err != nil {
		log.WithError(err).Warn("could not get relative path from static")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	pp := filepath.Clean(fmt.Sprintf("/public/%s", p))
	if !publicRegExp.MatchString(pp) {
		log.WithField("path", pp).Warn("clean public filepath is not within public")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http.ServeFile(w, r, pp)
}

func main() {
	cfg := &graphql.GraphQLConfig{
		DB: &db.Config{
			DBPath:         "/db/data.db",
			MigrationsPath: "migrations",
		},
		SchemaString: schema.String,
	}
	gqlhandler, err := cfg.NewHandler()
	if err != nil {
		log.WithError(err).Fatal("could not initialize GraphQL handler")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/static/", handleStatic)
	mux.Handle("/graphiql", graphiql.Serve)
	mux.Handle("/graphql", gqlhandler)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "/public/index.html")
	})

	log.Info("server listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.WithError(err).Fatal("could not start server")
	}
}
