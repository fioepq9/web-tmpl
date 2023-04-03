package logger

import (
	"io"

	"os"
	"time"
	"webtmpl/internal/config"

	"github.com/rs/zerolog"
)

var L zerolog.Logger

func Init(c config.Config) {
	var output io.Writer
	if c.App.Mode.Dev() {
		output = zerolog.NewConsoleWriter(
			func(w *zerolog.ConsoleWriter) {
				w.Out = os.Stdout
				w.TimeFormat = time.RFC3339
			},
		)
	} else {
		output = os.Stdout
	}
	L = zerolog.New(output).With().Caller().Timestamp().Logger()
	if c.App.Mode.Dev() {
		L = L.Level(zerolog.TraceLevel)
	} else {
		L = L.Level(zerolog.InfoLevel)
	}
}
