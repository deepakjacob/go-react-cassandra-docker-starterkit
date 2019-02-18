package logger

import (
	"context"
	"os"
	"path"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type correlationIdType int

const (
	requestIdKey correlationIdType = iota
	sessionIdKey
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

// WithRqId returns a context which knows its request ID
func WithRqId(ctx context.Context, rqId uuid.UUID) context.Context {
	return context.WithValue(ctx, requestIdKey, rqId)
}

// WithSessionId returns a context which knows its session ID
func WithSessionId(ctx context.Context, sessionId uuid.UUID) context.Context {
	return context.WithValue(ctx, sessionIdKey, sessionId)
}

// Logger returns a zap logger with as much context as possible
func Logger(ctx context.Context) *zap.Logger {
	newLogger := logger
	pid := os.Getpid()
	exec := path.Base(os.Args[0])
	if ctx != nil {
		if ctxRqId, ok := ctx.Value(requestIdKey).(uuid.UUID); ok {
			newLogger = newLogger.With(
				zap.String("rqId", ctxRqId.String()),
				zap.Int("pid", pid),
				zap.String("exec", exec),
			)
		}
		if ctxSessionId, ok := ctx.Value(sessionIdKey).(uuid.UUID); ok {
			newLogger = newLogger.With(
				zap.String("sessionId", ctxSessionId.String()),
				zap.Int("pid", pid),
				zap.String("exec", exec),
			)
		}
	}
	return newLogger
}
