package log

import (
	"fmt"
	"io"

	"github.com/aws/smithy-go/logging"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

// AppLogger is a logger for the application.
type AppLogger = log.Logger

// NewAppLogger creates a new logger for the application.
func NewAppLogger(w io.Writer, level, styles, prefix string) (*AppLogger, error) {
	l, err := newLogger(w, level, styles, prefix)
	if err != nil {
		return nil, err
	}
	return l, nil
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
func NewSDKLogger(w io.Writer, level, styles, prefix string) (*SDKLogger, error) {
	logger, err := newLogger(w, level, styles, prefix)
	if err != nil {
		return nil, err
	}
	l := &SDKLogger{
		Logger: logger,
	}
	return l, nil
}

// newLogger creates a new logger.
func newLogger(w io.Writer, level, styles, prefix string) (*log.Logger, error) {
	lv, err := log.ParseLevel(level)
	if err != nil {
		return nil, err
	}
	st, err := ParseStyles(styles)
	if err != nil {
		return nil, err
	}
	l := log.New(w)
	l.SetLevel(lv)
	l.SetStyles(st)
	if level == "debug" {
		l.SetReportCaller(true)
	}
	if prefix != "" {
		l.SetPrefix(prefix)
	}
	return l, nil
}

// Level is a logging level.
type Level = log.Level

var InfoLevel = log.InfoLevel

// Styles is a set of styles.
type Styles = log.Styles

// ParseStyles parses the styles.
func ParseStyles(styles string) (*Styles, error) {
	switch styles {
	case "default":
		return DefaultStyles(), nil
	case "labeled":
		return LabeledStyles(), nil
	default:
		return nil, fmt.Errorf("unsupported styles: %s", styles)
	}
}

// DefaultStyles returns the default styles.
func DefaultStyles() *Styles {
	styles := log.DefaultStyles()
	styles.Levels[log.DebugLevel] = lipgloss.NewStyle().
		SetString("DBG").
		Bold(true).
		MaxWidth(3).
		Foreground(lipgloss.Color("63"))
	styles.Levels[log.InfoLevel] = lipgloss.NewStyle().
		SetString("INF").
		Bold(true).
		MaxWidth(3).
		Foreground(lipgloss.Color("86"))
	styles.Levels[log.WarnLevel] = lipgloss.NewStyle().
		SetString("WRN").
		Bold(true).
		MaxWidth(3).
		Foreground(lipgloss.Color("192"))
	styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
		SetString("ERR").
		Bold(true).
		MaxWidth(3).
		Foreground(lipgloss.Color("204"))
	styles.Levels[log.FatalLevel] = lipgloss.NewStyle().
		SetString("FTL").
		Bold(true).
		MaxWidth(3).
		Foreground(lipgloss.Color("134"))
	return styles
}

// LabeledStyles returns the color labeled styles.
func LabeledStyles() *Styles {
	styles := log.DefaultStyles()
	styles.Levels[log.DebugLevel] = lipgloss.NewStyle().
		SetString("DBG").
		Bold(true).
		Padding(0, 1, 0, 1).
		MaxWidth(5).
		Background(lipgloss.Color("63")).
		Foreground(lipgloss.Color("0"))
	styles.Levels[log.InfoLevel] = lipgloss.NewStyle().
		SetString("INF").
		Bold(true).
		Padding(0, 1, 0, 1).
		MaxWidth(5).
		Background(lipgloss.Color("86")).
		Foreground(lipgloss.Color("0"))
	styles.Levels[log.WarnLevel] = lipgloss.NewStyle().
		SetString("WRN").
		Bold(true).
		Padding(0, 1, 0, 1).
		MaxWidth(5).
		Background(lipgloss.Color("192")).
		Foreground(lipgloss.Color("0"))
	styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
		SetString("ERR").
		Bold(true).
		Padding(0, 1, 0, 1).
		MaxWidth(5).
		Background(lipgloss.Color("204")).
		Foreground(lipgloss.Color("0"))
	styles.Levels[log.FatalLevel] = lipgloss.NewStyle().
		SetString("FTL").
		Bold(true).
		Padding(0, 1, 0, 1).
		MaxWidth(5).
		Background(lipgloss.Color("134")).
		Foreground(lipgloss.Color("0"))
	return styles
}
