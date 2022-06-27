package logger

import (
	"context"
	"fmt"
	"io"

	"github.com/sirupsen/logrus"
)

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

type TraceContextKey string

const LogKeyTraceID TraceContextKey = "trace_id"

const (
	logInfo = iota
	logWarn
	logError
)

type EnzoLogger struct {
	log *logrus.Logger
}

func NewEnzoLogger(output io.Writer) LoggerWithContext {
	log := logrus.New()
	log.SetOutput(output)
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.DebugLevel)
	return &EnzoLogger{log}
}

func (l *EnzoLogger) Info(ctx context.Context, args ...interface{}) {
	l.write(ctx, logInfo, args...)
}

func (l *EnzoLogger) Warn(ctx context.Context, args ...interface{}) {
	l.write(ctx, logWarn, args...)

}
func (l *EnzoLogger) Error(ctx context.Context, args ...interface{}) {
	l.write(ctx, logError, args...)
}

func (l *EnzoLogger) Infof(ctx context.Context, format string, args ...interface{}) {
	l.writef(ctx, logInfo, format, args...)
}

func (l *EnzoLogger) Warnf(ctx context.Context, format string, args ...interface{}) {
	l.writef(ctx, logWarn, format, args...)

}
func (l *EnzoLogger) Errorf(ctx context.Context, format string, args ...interface{}) {
	l.writef(ctx, logError, format, args...)
}

func (l *EnzoLogger) write(ctx context.Context, level int, args ...interface{}) {
	reqID := fmt.Sprintf("%v", ctx.Value(LogKeyTraceID))
	entry := l.log.WithField(string(LogKeyTraceID), fmt.Sprintf("%v", reqID))

	switch level {
	case logInfo:
		entry.Info(args...)
	case logWarn:
		entry.Warn(args...)
	case logError:
		entry.Error(args...)
	}
}

func (l *EnzoLogger) writef(ctx context.Context, level int, format string, args ...interface{}) {
	reqID := fmt.Sprintf("%v", ctx.Value(LogKeyTraceID))
	entry := l.log.WithField(string(LogKeyTraceID), fmt.Sprintf("%v", reqID))

	switch level {
	case logInfo:
		entry.Infof(format, args...)
	case logWarn:
		entry.Warnf(format, args...)
	case logError:
		entry.Errorf(format, args...)
	}
}
