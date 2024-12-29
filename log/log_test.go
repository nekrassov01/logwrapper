package log

import (
	"bytes"
	"regexp"
	"testing"

	"github.com/aws/smithy-go/logging"
	"github.com/charmbracelet/log"
)

type test struct {
	name   string
	level  Level
	styles *Styles
	prefix string
	msg    string
	want   string
}

var (
	re    = regexp.MustCompile(`log/log_test.go:\d*`)
	tests = []test{
		{
			name:   "no prefix",
			level:  InfoLevel,
			styles: DefaultStyles(),
			prefix: "",
			msg:    "This is a debug message",
			want:   "INF This is a debug message\n",
		},
		{
			name:   "default debug",
			level:  DebugLevel,
			styles: DefaultStyles(),
			prefix: "MyApp",
			msg:    "This is a debug message",
			want:   "DBG <> MyApp: This is a debug message\n",
		},
		{
			name:   "default info",
			level:  InfoLevel,
			styles: DefaultStyles(),
			prefix: "MyApp",
			msg:    "This is an infomational message",
			want:   "INF MyApp: This is an infomational message\n",
		},
		{
			name:   "default warn",
			level:  WarnLevel,
			styles: DefaultStyles(),
			prefix: "MyApp",
			msg:    "This is a warning message",
			want:   "WRN MyApp: This is a warning message\n",
		},
		{
			name:   "default error",
			level:  ErrorLevel,
			styles: DefaultStyles(),
			prefix: "MyApp",
			msg:    "This is an error message",
			want:   "ERR MyApp: This is an error message\n",
		},
		{
			name:   "labeled debug",
			level:  DebugLevel,
			styles: LabeledStyles(),
			prefix: "MyApp",
			msg:    "This is a debug message",
			want:   " DBG  <> MyApp: This is a debug message\n",
		},
		{
			name:   "labeled info",
			level:  InfoLevel,
			styles: LabeledStyles(),
			prefix: "MyApp",
			msg:    "This is an infomational message",
			want:   " INF  MyApp: This is an infomational message\n",
		},
		{
			name:   "labeled warn",
			level:  WarnLevel,
			styles: LabeledStyles(),
			prefix: "MyApp",
			msg:    "This is a warning message",
			want:   " WRN  MyApp: This is a warning message\n",
		},
		{
			name:   "labeled error",
			level:  ErrorLevel,
			styles: LabeledStyles(),
			prefix: "MyApp",
			msg:    "This is an error message",
			want:   " ERR  MyApp: This is an error message\n",
		},
	}
)

func TestNewAppLogger(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			l := NewAppLogger(w, tt.level, tt.styles, tt.prefix)
			switch tt.level {
			case DebugLevel:
				l.Debug(tt.msg)
			case InfoLevel:
				l.Info(tt.msg)
			case WarnLevel:
				l.Warn(tt.msg)
			case ErrorLevel:
				l.Error(tt.msg)
			default:
				l.Info(tt.msg)
			}
			got := re.ReplaceAllString(w.String(), "")
			if got != tt.want {
				t.Errorf("NewAppLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSDKLogger(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			l := NewSDKLogger(w, tt.level, tt.styles, tt.prefix)
			switch tt.level {
			case DebugLevel:
				l.Debug(tt.msg)
			case InfoLevel:
				l.Info(tt.msg)
			case WarnLevel:
				l.Warn(tt.msg)
			case ErrorLevel:
				l.Error(tt.msg)
			default:
				l.Info(tt.msg)
			}
			got := re.ReplaceAllString(w.String(), "")
			if got != tt.want {
				t.Errorf("NewSDKLogger() = %v, want %v", got, tt.want)
			}
			l.Logf(logging.Warn, "This is a warning message")
		})
	}
}

func Test_newLogger(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			l := NewSDKLogger(w, tt.level, tt.styles, tt.prefix)
			switch tt.level {
			case DebugLevel:
				l.Debug(tt.msg)
			case InfoLevel:
				l.Info(tt.msg)
			case WarnLevel:
				l.Warn(tt.msg)
			case ErrorLevel:
				l.Error(tt.msg)
			default:
				l.Info(tt.msg)
			}
			got := re.ReplaceAllString(w.String(), "")
			if got != tt.want {
				t.Errorf("NewSDKLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSDKLogger_Logf(t *testing.T) {
	type fields struct {
		Logger *log.Logger
	}
	type args struct {
		c      logging.Classification
		format string
		v      []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "warn",
			fields: fields{
				Logger: log.NewWithOptions(&bytes.Buffer{}, log.Options{
					Level: log.WarnLevel,
				}),
			},
			args: args{
				c:      logging.Warn,
				format: "Warning: %d issues found",
				v:      []any{3},
			},
			want: "WRN Warning: 3 issues found\n",
		},
		{
			name: "debug",
			fields: fields{
				Logger: log.NewWithOptions(&bytes.Buffer{}, log.Options{
					Level: log.DebugLevel,
				}),
			},
			args: args{
				c:      logging.Debug,
				format: "Debug message: %s",
				v:      []any{"debugging"},
			},
			want: "DBG Debug message: debugging\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			logger := func() *log.Logger {
				l := log.New(w)
				l.SetLevel(tt.fields.Logger.GetLevel())
				l.SetStyles(DefaultStyles())
				return l
			}
			l := &SDKLogger{Logger: logger()}
			l.Logf(tt.args.c, tt.args.format, tt.args.v...)
			got := re.ReplaceAllString(w.String(), "")
			if got != tt.want {
				t.Errorf("SDKLogger.Logf() = %v, want %v", got, tt.want)
			}
		})
	}
}
