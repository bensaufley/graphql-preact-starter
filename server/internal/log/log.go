// Package log wraps sirupsen/logrus to provide default settings for
// a logger and also adds a wrapper struct that can be passed to
// graphql as a logger that matches the expected graphql.Logger interface
package log

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
)

type logger struct {
	*logrus.Logger
}

var Logger = &logger{
	Logger: &logrus.Logger{
		Out:       os.Stdout,
		Formatter: &logrus.JSONFormatter{},
		Level:     logrus.DebugLevel,
	},
}

type Entry = logrus.Entry

type Fields = logrus.Fields

func (l *logger) LogPanic(ctx context.Context, value interface{}) {
	l.Logger.WithContext(ctx).Error(value)
}
