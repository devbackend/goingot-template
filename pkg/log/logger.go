package log

import (
	"io"

	"github.com/rs/zerolog"
)

type Logger struct {
	zerolog.Logger
}

// New return instance of log
func New(w io.Writer) Logger {
	return Logger{zerolog.New(w).With().Caller().Timestamp().Logger()}
}
