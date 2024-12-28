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
// If level is not valid, it will use InfoLevel.
// If styles is not valid, it will use DefaultStyles.
func NewAppLogger(w io.Writer, level, styles, prefix string) *AppLogger {
	lv, err := ParseLevel(level)
	if err != nil {
		lv = InfoLevel
	}
	st, err := ParseStyles(styles)
	if err != nil {
		st = DefaultStyles()
	}
	l := log.New(w)
	l.SetLevel(lv)
	l.SetStyles(st)
	if level == DebugLevel.String() {
		l.SetReportCaller(true)
	}
	if prefix != "" {
		l.SetPrefix(prefix)
	}
	return l
}

// SDKLogger is a logger for the AWS SDK. This implemented the logging.Logger interface.
// https://github.com/aws/smithy-go/blob/main/logging/logger.go
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
func NewSDKLogger(w io.Writer, level, styles, prefix string) *SDKLogger {
	return &SDKLogger{
		Logger: NewAppLogger(w, level, styles, prefix),
	}
}
