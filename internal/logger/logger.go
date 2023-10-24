package logger

import (
	"errors"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

type LogConfig struct {
	Environment string
	LogLevel    Level `envconfig:"LOG_LEVEL" default:"DEBUG"`
}

type Level struct {
	slug string
}

func (r Level) String() string {
	return r.slug
}

var (
	Unknown = Level{""}
	TRACE   = Level{"TRACE"}
	DEBUG   = Level{"DEBUG"}
	INFO    = Level{"INFO"}
	WARN    = Level{"WARN"}
	ERROR   = Level{"ERROR"}
	PANIC   = Level{"PANIC"}
)

func NewLevel(s string) (Level, error) {
	switch s {
	case TRACE.slug:
		return TRACE, nil
	case DEBUG.slug:
		return DEBUG, nil
	case INFO.slug:
		return INFO, nil
	case WARN.slug:
		return WARN, nil
	case ERROR.slug:
		return ERROR, nil
	case PANIC.slug:
		return PANIC, nil
	}

	return Unknown, errors.New("Unknown Level: " + s)
}

func New(cfg LogConfig, env string) zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	switch cfg.Environment {
	case "production":
		return zerolog.New(os.Stdout).
			Level(logLevelToZero(cfg.LogLevel)).
			With().
			Timestamp().
			Caller().
			Logger()
	default:
		return zerolog.New(zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
			w.TimeFormat = "03:04:05.000PM"
		})).
			Level(logLevelToZero(cfg.LogLevel)).
			With().
			Timestamp().
			Caller().
			Logger()
	}
}

func logLevelToZero(level Level) zerolog.Level {
	switch level {
	case PANIC:
		return zerolog.PanicLevel
	case ERROR:
		return zerolog.ErrorLevel
	case WARN:
		return zerolog.WarnLevel
	case INFO:
		return zerolog.InfoLevel
	case DEBUG:
		return zerolog.DebugLevel
	case TRACE:
		return zerolog.TraceLevel
	default:
		return zerolog.InfoLevel
	}
}
