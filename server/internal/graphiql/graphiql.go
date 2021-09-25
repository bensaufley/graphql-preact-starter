// Package graphiql is for dev only and is used to serve the interactive Graphiql
// UI in your local environment.
package graphiql

import (
	_ "embed"
	"net/http"

	log "github.com/sirupsen/logrus"
)

//go:embed index.html
var html []byte

var Serve = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write(html); err != nil {
		log.WithContext(r.Context()).WithError(err).Error("encountered error serving graphiql")
	}
})
