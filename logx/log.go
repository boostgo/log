package logx

import (
	"context"
	"os"
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	_ctx       context.Context
	_ctxCancel context.CancelFunc

	_pretty bool

	_logger zerolog.Logger
	_once   sync.Once
)

// Pretty enables logging mode.
//
// This mode could be activated by "PRETTY_LOGGER=true" env
func Pretty() {
	_pretty = true
}

func IsPretty() bool {
	return _pretty
}

func Cancel() {
	if _ctxCancel == nil {
		return
	}

	_ctxCancel()
}

func SetCancel(cancel context.CancelFunc) {
	_ctxCancel = cancel
}

func InitLogger() {
	switch os.Getenv("PRETTY_LOGGER") {
	case "true", "TRUE":
		Pretty()
	}

	switch os.Getenv("PRETTY_LOG") {
	case "true", "TRUE":
		Pretty()
	}

	if IsPretty() {
		_logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		return
	}

	_logger = zerolog.
		New(os.Stdout).
		With().
		Timestamp().
		Logger()
}

func Logger() zerolog.Logger {
	_once.Do(func() {
		InitLogger()
	})
	return _logger
}
