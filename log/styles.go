package log

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

// Style is a style.
type Style int

const (
	// DefaultStyle is the default style.
	DefaultStyle Style = iota

	// LabeledStyle is the color labeled style.
	LabeledStyle
)

// String returns the string representation of the style.
func (t Style) String() string {
	switch t {
	case DefaultStyle:
		return "default"
	case LabeledStyle:
		return "labeled"
	default:
		return ""
	}
}

// Styles is a set of styles.
type Styles = log.Styles

// ParseStyles parses the styles.
func ParseStyles(styles string) (*Styles, error) {
	switch styles {
	case DefaultStyle.String():
		return DefaultStyles(), nil
	case LabeledStyle.String():
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
