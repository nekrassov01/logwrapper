package log

import (
	"github.com/charmbracelet/log"
)

// Level is a logging level.
type Level = log.Level

var (
	DebugLevel = log.DebugLevel // DebugLevel is the debug level.
	InfoLevel  = log.InfoLevel  // InfoLevel is the info level.
	WarnLevel  = log.WarnLevel  // WarnLevel is the warn level.
	ErrorLevel = log.ErrorLevel // ErrorLevel is the error level.
	FatalLevel = log.FatalLevel // FatalLevel is the fatal level.
)

// ParseLevel converts level in string to Level type. Default level is InfoLevel.
func ParseLevel(level string) (Level, error) {
	return log.ParseLevel(level)
}

// SetLevel sets the logging level. If level is DebugLevel, it will also report caller implicitly.
func SetLevel(level Level) {
	log.SetLevel(level)
	log.SetReportCaller(level == DebugLevel)
}
