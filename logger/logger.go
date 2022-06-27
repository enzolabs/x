package logger

import "context"

type LoggerWithContext interface {
	// Debug(ctx context.Context, args ...interface{})
	Info(ctx context.Context, args ...interface{})
	Warn(ctx context.Context, args ...interface{})
	Error(ctx context.Context, args ...interface{})
	// Fatal(ctx context.Context, args ...interface{})
	// Panic(ctx context.Context, args ...interface{})

	// Debugf(ctx context.Context, format string, args ...interface{})
	Infof(ctx context.Context, format string, args ...interface{})
	Warnf(ctx context.Context, format string, args ...interface{})
	Errorf(ctx context.Context, format string, args ...interface{})
	// Fatalf(ctx context.Context, format string, args ...interface{})
	// Panicf(ctx context.Context, format string, args ...interface{})
}
