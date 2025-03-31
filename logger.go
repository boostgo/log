package log

import (
	"context"
	"github.com/boostgo/log/logx"
)

// Debug print log on debug level.
//
// Provided context use trace id
func Debug(ctx ...context.Context) Event {
	l := logx.Logger()
	return newEvent(l.Debug(), ctx...)
}

// Info print log on info level.
//
// Provided context use trace id
func Info(ctx ...context.Context) Event {
	l := logx.Logger()
	return newEvent(l.Info(), ctx...)
}

// Warn print log on warning level.
//
// Provided context use trace id
func Warn(ctx ...context.Context) Event {
	l := logx.Logger()
	return newEvent(l.Warn(), ctx...)
}

// Error print log on error level.
//
// Provided context use trace id
func Error(ctx ...context.Context) Event {
	l := logx.Logger()
	return newEvent(l.Error(), ctx...)
}

// Fatal print log on error level but with bool fatal=true.
// Provided context use trace id.
//
// Call AppCancel function
func Fatal(ctx ...context.Context) Event {
	defer logx.Cancel()
	l := logx.Logger()
	return newEvent(l.Error().Bool("fatal", true), ctx...)
}

// Logger is wrap over zerolog logger
type Logger interface {
	// Debug print log on debug level.
	// Provided context use trace id
	Debug() Event
	// Info print log on info level.
	// Provided context use trace id
	Info() Event
	// Warn print log on warning level.
	// Provided context use trace id
	Warn() Event
	// Error print log on error level.
	// Provided context use trace id
	Error() Event
	// Fatal print log on error level but with bool fatal=true.
	// Provided context use trace id.
	//
	// Call life.Cancel() method which call graceful shutdown
	Fatal() Event
}

type wrapper struct {
	namespace string
	ctx       context.Context
}

// Namespace creates Logger implementation with namespace
func Namespace(namespace string) Logger {
	return Context(context.Background(), namespace)
}

// Context creates Logger implementation with context & namespace
func Context(ctx context.Context, namespace string) Logger {
	return &wrapper{
		ctx:       ctx,
		namespace: namespace,
	}
}

func (logger *wrapper) Debug() Event {
	return Debug().
		Ctx(logger.ctx).
		Namespace(logger.namespace)
}

func (logger *wrapper) Info() Event {
	return Info().
		Ctx(logger.ctx).
		Namespace(logger.namespace)
}

func (logger *wrapper) Warn() Event {
	return Warn().
		Ctx(logger.ctx).
		Namespace(logger.namespace)
}

func (logger *wrapper) Error() Event {
	return Error().
		Ctx(logger.ctx).
		Namespace(logger.namespace)
}

func (logger *wrapper) Fatal() Event {
	return Fatal().
		Ctx(logger.ctx).
		Namespace(logger.namespace)
}
