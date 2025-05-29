package xraylogr_test

import (
	"bytes"
	"fmt"
	"log"

	"github.com/aws/aws-xray-sdk-go/xraylog"
	"github.com/bodgit/xraylogr"
	"github.com/tonglil/buflogr"
)

type stringer string

func (s stringer) String() string { return string(s) }

func ExampleLogger() {
	buf := new(bytes.Buffer)

	l, err := xraylogr.New(buflogr.NewWithBuffer(buf), xraylogr.WithLogLevel(func(l xraylog.LogLevel) int {
		return 0
	}))
	if err != nil {
		log.Fatal(err)
	}

	l.Log(xraylog.LogLevelWarn, stringer("a warning"))

	l, err = xraylogr.New(buflogr.NewWithBuffer(buf))
	if err != nil {
		log.Fatal(err)
	}

	l.Log(xraylog.LogLevelDebug, stringer("some debug"))
	l.Log(xraylog.LogLevelInfo, stringer("some info"))
	l.Log(xraylog.LogLevelWarn, stringer("another warning"))
	l.Log(xraylog.LogLevelError, stringer("an error"))

	fmt.Print(buf.String())

	// Output:
	// INFO a warning
	// V[4] some debug
	// V[1] some info
	// INFO another warning
	// ERROR <nil> an error
}
