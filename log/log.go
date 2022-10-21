// package log is a wrapper around github.com/uber-go/zap
package log

import (
	"context"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// log levels
const (
	DebugLevel = "debug"
	InfoLevel  = "info"
	WarnLevel  = "warn"
	ErrorLevel = "error"
	PanicLevel = "panic"
	FatalLevel = "fatal"
)

type (
	Logger struct {
		level zap.AtomicLevel
		*zap.Logger
	}

	contextKey struct{}
)

// SetLevel sets the log level on this Logger instance.
// It defaults to INFO if an invalid level is given.
func (l *Logger) SetLevel(level string) {
	l.level.SetLevel(zapLevel(level))
}

// New returns a new Logger instance which logs at the given level
func New(level string) *Logger {
	zapConfig := zap.NewProductionConfig()

	zapConfig.Level = zap.NewAtomicLevelAt(zapLevel(level))
	zapConfig.EncoderConfig.MessageKey = "message"
	zapConfig.EncoderConfig.TimeKey = "@timestamp"
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	zapLogger, err := zapConfig.Build()
	if err != nil {
		panic(err)
	}

	return &Logger{
		level:  zapConfig.Level,
		Logger: zapLogger,
	}
}

// WithContext embeds the given Logger instance into the given context.
// If context is nil it returns nil.
func WithContext(ctx context.Context, logger *Logger) context.Context {
	if ctx == nil {
		return ctx
	}
	return context.WithValue(ctx, contextKey{}, logger)
}

// FromContext returns a Logger instance from the given context.
// If no Logger instance is found, it returns a new Logger instance.
func FromContext(ctx context.Context) *Logger {
	if ctx == nil {
		return New(InfoLevel)
	}
	if logger, ok := ctx.Value(contextKey{}).(*Logger); ok {
		return logger
	}
	return New(InfoLevel)
}

// zapLevel resolves a given level string to its zapcore.Level equivalent.
// It defaults to INFO if an invalid level is given.
func zapLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case DebugLevel:
		return zap.DebugLevel
	case InfoLevel:
		return zap.InfoLevel
	case WarnLevel:
		return zap.WarnLevel
	case ErrorLevel:
		return zap.ErrorLevel
	case PanicLevel:
		return zap.PanicLevel
	case FatalLevel:
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}
