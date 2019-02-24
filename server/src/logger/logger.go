package logger

import (
	"context"
	"os"
	"path"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type correlationIDType int

const (
	requestIDKey correlationIDType = iota
	sessionIDKey
)

var logger *zap.Logger

func init() {
	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	logger, _ = cfg.Build()
}

// WithRqID returns a context which knows its request ID
func WithRqID(ctx context.Context, rqID uuid.UUID) context.Context {
	return context.WithValue(ctx, requestIDKey, rqID)
}

// WithSessionID returns a context which knows its session ID
func WithSessionID(ctx context.Context, sessionID uuid.UUID) context.Context {
	return context.WithValue(ctx, sessionIDKey, sessionID)
}

// Logger returns a zap logger with as much context as possible
func Logger(ctx context.Context) *zap.Logger {
	newLogger := logger
	pid := os.Getpid()
	exec := path.Base(os.Args[0])
	if ctx != nil {
		if ctxRqID, ok := ctx.Value(requestIDKey).(uuid.UUID); ok {
			newLogger = newLogger.With(
				zap.String("rqId", ctxRqID.String()),
				zap.Int("pid", pid),
				zap.String("exec", exec),
			)
		}
		if ctxSessionID, ok := ctx.Value(sessionIDKey).(uuid.UUID); ok {
			newLogger = newLogger.With(
				zap.String("sessionId", ctxSessionID.String()),
				zap.Int("pid", pid),
				zap.String("exec", exec),
			)
		}
	}
	return newLogger
}
