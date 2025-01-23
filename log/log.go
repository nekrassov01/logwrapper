package log

import (
	"fmt"
	"io"

	"github.com/aws/smithy-go/logging"
	"github.com/charmbracelet/log"
)

// baseLogger is a base logger for embedding.
type baseLogger struct {
	*log.Logger
}

// SetLevel sets the logging level. If the level is DebugLevel, it will also report the caller implicitly.
func (l *baseLogger) SetLevel(level Level) {
	l.Logger.SetLevel(level)
	l.Logger.SetReportCaller(level == DebugLevel)
}

// newBaseLogger creates a new base logger used by AppLogger and SDKLogger.
func newBaseLogger(w io.Writer, level Level, styles *Styles, prefix string) *baseLogger {
	l := log.New(w)
	l.SetLevel(level)
	l.SetStyles(styles)
	l.SetReportCaller(level == DebugLevel)
	if prefix != "" {
		l.SetPrefix(prefix)
	}
	return &baseLogger{l}
}

// AppLogger is a logger for the application.
type AppLogger struct {
	*baseLogger
}

// NewAppLogger creates a new logger for the application.
func NewAppLogger(w io.Writer, level Level, styles *Styles, prefix string) *AppLogger {
	l := newBaseLogger(w, level, styles, prefix)
	return &AppLogger{l}
}

// SDKLogger is a logger for the AWS SDK. This implemented the logging.Logger interface.
// See: https://github.com/aws/smithy-go/blob/main/logging/logger.go
type SDKLogger struct {
	*baseLogger
}

// Logf logs a message with formatting.
func (l *SDKLogger) Logf(c logging.Classification, format string, v ...any) {
	s := fmt.Sprintf(format, v...)
	switch c {
	case logging.Warn:
		l.Warn(s)
	case logging.Debug:
		l.Debug(s)
	default:
	}
}

// NewSDKLogger creates a new logger for AWS SDK.
func NewSDKLogger(w io.Writer, level Level, styles *Styles, prefix string) *SDKLogger {
	l := newBaseLogger(w, level, styles, prefix)
	return &SDKLogger{l}
}
