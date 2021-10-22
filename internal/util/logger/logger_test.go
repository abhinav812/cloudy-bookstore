// +build !binary_log

package logger_test

import (
	"flag"
	"testing"
	"time"

	"github.com/abhinav812/cloudy-bookstore/internal/util/logger"
	"github.com/rs/zerolog"
)

func setup() *logger.Logger {
	// setting empty string will write with UNIX time
	zerolog.TimeFieldFormat = ""

	// In order to always return a static time to pass the tests.
	zerolog.TimestampFunc = func() time.Time {
		return time.Date(2021, 1, 1, 01, 01, 01, 0, time.UTC)
	}

	return logger.NewConsole(true)
}

// Simple logging example using the Printf function in the log package
func TestPrintf(_ *testing.T) {
	l := setup()
	l.Printf("hello %s", "world")

	// Output: {"level":"debug","time":1609462861,"message":"hello world"}
}

// Simple logging example using the Print function in the log package
func TestPrint(_ *testing.T) {
	l := setup()
	l.Print("hello world")

	// Output: {"level":"debug","time":1609462861,"message":"hello world"}
}

// Test of a log with no particular "level"
func TestLog(_ *testing.T) {
	l := setup()
	l.Log().Msg("hello world")

	// Output: {"time":1609462861,"message":"hello world"}
}

// Test of a log at a particular "level" (in this case, "debug")
func TestDebug(_ *testing.T) {
	l := setup()
	l.Debug().Msg("hello world")

	// Output: {"level":"debug","time":1609462861,"message":"hello world"}
}

// Test of a log at a particular "level" (in this case, "info")
func TestInfo(_ *testing.T) {
	l := setup()
	l.Info().Msg("hello world")

	// Output: {"level":"info","time":1609462861,"message":"hello world"}
}

// Test of a log at a particular "level" (in this case, "warn")
func TestWarn(_ *testing.T) {
	l := setup()
	l.Warn().Msg("hello world")

	// Output: {"level":"warn","time":1609462861,"message":"hello world"}
}

// Test of a log at a particular "level" (in this case, "error")
func TestError(_ *testing.T) {
	l := setup()
	l.Error().Msg("hello world")

	// Output: {"level":"error","time":1609462861,"message":"hello world"}
}

// This test uses command-line flags to demonstrate various outputs
// depending on the chosen log level.
func TestLogger_Level(_ *testing.T) {
	l := setup()
	debug := flag.Bool("debug", false, "sets log level to debug")

	flag.Parse()

	// Default level for this test is info, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	l.Debug().Msg("This message appears only when log level set to Debug")
	l.Info().Msg("This message appears when log level set to Debug or Info")

	if e := l.Debug(); e.Enabled() {
		// Compute log output only if enabled.
		value := "bar"
		e.Str("foo", value).Msg("some debug message")
	}

	// Output: {"level":"info","time":1609462861,"message":"This message appears when log level set to Debug or Info"}
}
