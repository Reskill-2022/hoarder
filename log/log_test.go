package log

import (
	"context"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestWithFromContext(t *testing.T) {
	logger := New("info")
	ctx := WithContext(context.Background(), logger)

	got := FromContext(ctx)
	if got != logger {
		t.Errorf("got %v, want %v", got, logger)
	}
}

func TestLoggerSetLevel(t *testing.T) {
	logger := New("info")

	logger.SetLevel("debug")
	if logger.level.Level() != zap.DebugLevel {
		t.Errorf("expected debug level, got %s", logger.level.Level())
	}
}

func TestZapLevel(t *testing.T) {
	testCases := []struct {
		level string
		want  zapcore.Level
	}{
		{"debug", zap.DebugLevel},
		{InfoLevel, zap.InfoLevel},
		{"warn", zap.WarnLevel},
		{"error", zap.ErrorLevel},
		{"panic", zap.PanicLevel},
		{FatalLevel, zap.FatalLevel},
		{"", zap.InfoLevel},
		{"unknown", zap.InfoLevel},
	}

	for _, tc := range testCases {
		t.Run(tc.level, func(t *testing.T) {
			got := zapLevel(tc.level)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
