package logger

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

type output struct {
	RequestID string    `json:"trace_id"`
	Level     string    `json:"level"`
	Message   string    `json:"msg"`
	Time      time.Time `jsom:"time"`
}

const (
	tMessage   = "message"
	tRequestID = "request-id"
)

const (
	levelInfo  = "info"
	levelWarn  = "warning"
	levelError = "error"
)

func TestAppLogger(t *testing.T) {
	buf := new(bytes.Buffer)

	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{})
	l.SetOutput(buf)

	log := &EnzoLogger{
		log: l,
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, LogKeyTraceID, tRequestID)

	// Test Info()
	t.Run("EnzoLogger.Info()", func(t *testing.T) {
		log.Info(ctx, tMessage)
		if err := isValidLogOutput(buf, levelInfo, tMessage); err != nil {
			t.Error(err.Error())
		}
	})

	buf.Reset()

	// Test Warn()
	t.Run("EnzoLogger.Warn()", func(t *testing.T) {

		log.Warn(ctx, tMessage)
		if err := isValidLogOutput(buf, levelWarn, tMessage); err != nil {
			t.Error(err.Error())
		}
	})

	buf.Reset()

	// Test Error()
	t.Run("EnzoLogger.Warn()", func(t *testing.T) {
		log.Error(ctx, tMessage)
		if err := isValidLogOutput(buf, levelError, tMessage); err != nil {
			t.Error(err.Error())
		}
	})

	buf.Reset()

	t.Run("EnzoLogger.Infof()", func(t *testing.T) {
		log.Infof(ctx, "x:%v", tMessage)
		if err := isValidLogOutput(buf, levelInfo, fmt.Sprintf("x:%s", tMessage)); err != nil {
			t.Error(err.Error())
		}
	})

	buf.Reset()

	t.Run("EnzoLogger.Warnf()", func(t *testing.T) {
		log.Warnf(ctx, "x:%v", tMessage)
		if err := isValidLogOutput(buf, levelWarn, fmt.Sprintf("x:%s", tMessage)); err != nil {
			t.Error(err.Error())
		}
	})

	buf.Reset()

	t.Run("EnzoLogger.Errorf()", func(t *testing.T) {
		log.Errorf(ctx, "x:%v", tMessage)
		if err := isValidLogOutput(buf, levelError, fmt.Sprintf("x:%s", tMessage)); err != nil {
			t.Error(err.Error())
		}
	})

}

func isValidLogOutput(buf *bytes.Buffer, level string, message string) error {
	o := new(output)
	if err := json.Unmarshal(buf.Bytes(), o); err != nil {
		return fmt.Errorf("expected a JSON output, got= %v", buf.String())
	}
	if o.Level != level {
		return fmt.Errorf("expected log level to be '%s', got=%v", level, o.Level)
	}
	if o.RequestID != tRequestID {
		return fmt.Errorf("expected requestID to be '%s', got=%v", tRequestID, o.RequestID)
	}
	if o.Message != message {
		return fmt.Errorf("expected msg to be %s, got=%v", message, o.Message)
	}
	return nil
}
