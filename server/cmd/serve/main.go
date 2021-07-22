package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"regexp"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	log "github.com/sirupsen/logrus"

	"github.com/bensaufley/graphql-preact-starter/internal/graphiql"
	"github.com/bensaufley/graphql-preact-starter/internal/resolver"
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
	s, err := schema.String()
	if err != nil {
		log.WithError(err).Fatal("could not build schema string")
	}
	schm, err := graphql.ParseSchema(s, resolver.NewRoot(nil))
	if err != nil {
		log.WithError(err).Fatal("could not parse schema")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/static/", handleStatic)
	mux.Handle("/graphiql", graphiql.Serve)
	mux.Handle("/graphql", &relay.Handler{Schema: schm})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "/public/index.html")
	})

	log.Info("server listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.WithError(err).Fatal("could not start server")
	}
}
