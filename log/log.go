package log

import (
	"fmt"
	"io"

	"github.com/aws/smithy-go/logging"
	"github.com/charmbracelet/log"
)

// AppLogger is a logger for the application.
type AppLogger = log.Logger

// NewAppLogger creates a new logger for the application.
func NewAppLogger(w io.Writer, level Level, styles *Styles, prefix string) *AppLogger {
	l := log.New(w)
	l.SetLevel(level)
	l.SetStyles(styles)
	l.SetReportCaller(level == DebugLevel)
	if prefix != "" {
		l.SetPrefix(prefix)
	}
	return l
}

// SDKLogger is a logger for the AWS SDK. This implemented the logging.Logger interface.
// See: https://github.com/aws/smithy-go/blob/main/logging/logger.go
type SDKLogger struct {
	*log.Logger
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
	return &SDKLogger{
		Logger: NewAppLogger(w, level, styles, prefix),
	}
}
