package logging

import (
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	zerolog.Logger
}

func NewLogger() *Logger {
	return &Logger{
		zerolog.New(os.Stdout).With().Timestamp().Logger(),
	}
}
