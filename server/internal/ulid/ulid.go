// Package ulid provides convenience functions around oklog/ulid
// to easily generate ULIDs and stringify them
package ulid

import (
	"crypto/rand"
	"time"

	oklog "github.com/oklog/ulid/v2"
)

type Generator struct {
	entropy *oklog.MonotonicEntropy
}

func (g *Generator) New() oklog.ULID {
	return oklog.MustNew(oklog.Timestamp(time.Now()), g.entropy)
}

func (g *Generator) String() string {
	return g.New().String()
}

func NewGenerator() *Generator {
	entropy := oklog.Monotonic(rand.Reader, 0)
	return &Generator{entropy: entropy}
}
