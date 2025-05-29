// Package xraylogr provides a means to pass a [logr.Logger] to the AWS X-Ray
// tracing SDK as an [xraylog.Logger].
//
//	l := stdr.New(nil)
//
//	nl, err := xraylogr.New(l)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	xray.SetLogger(nl)
package xraylogr

import (
	"fmt"

	"github.com/aws/aws-xray-sdk-go/xraylog"
	"github.com/go-logr/logr"
)

type logger struct {
	l logr.Logger
	f func(xraylog.LogLevel) int
}

func (l *logger) Log(level xraylog.LogLevel, msg fmt.Stringer) {
	nl := l.l.V(l.f(level))

	if level == xraylog.LogLevelError {
		nl.Error(nil, msg.String())
	} else if nl.Enabled() {
		nl.Info(msg.String())
	}
}

var _ xraylog.Logger = new(logger)

// WithLogLevel allows setting a custom mapping function from
// [xraylog.LogLevel] to logr verbosity levels.
func WithLogLevel(f func(xraylog.LogLevel) int) func(*logger) error {
	return func(l *logger) error {
		l.f = f

		return nil
	}
}

// New returns a new [xraylog.Logger] that by default maps
// [xraylog.LogLevelDebug] to verbosity level 4, [xraylog.LogLevelInfo] to
// verbosity level 1, otherwise level 0 is used.
func New(l logr.Logger, options ...func(*logger) error) (xraylog.Logger, error) {
	nl := &logger{
		l: l,
		f: func(level xraylog.LogLevel) int {
			switch level { //nolint:exhaustive
			case xraylog.LogLevelDebug:
				return 4
			case xraylog.LogLevelInfo:
				return 1
			}

			return 0
		},
	}

	for _, option := range options {
		if err := option(nl); err != nil {
			return nil, err
		}
	}

	return nl, nil
}
