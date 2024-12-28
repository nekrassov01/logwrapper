package log

import (
	"bytes"
	"regexp"
	"testing"

	"github.com/aws/smithy-go/logging"
	"github.com/charmbracelet/log"
)

type test struct {
	name    string
	level   string
	styles  string
	prefix  string
	msg     string
	want    string
	wantErr bool
}

var (
	re    = regexp.MustCompile(`log/log_test.go:\d*`)
	tests = []test{
		{
			name:    "no prefix",
			level:   "info",
			styles:  "default",
			prefix:  "",
			msg:     "This is a debug message",
			want:    "INF This is a debug message\n",
			wantErr: false,
		},
		{
			name:    "default debug",
			level:   "debug",
			styles:  "default",
			prefix:  "MyApp",
			msg:     "This is a debug message",
			want:    "DBG <> MyApp: This is a debug message\n",
			wantErr: false,
		},
		{
			name:    "default info",
			level:   "info",
			styles:  "default",
			prefix:  "MyApp",
			msg:     "This is an infomational message",
			want:    "INF MyApp: This is an infomational message\n",
			wantErr: false,
		},
		{
			name:    "default warn",
			level:   "warn",
			styles:  "default",
			prefix:  "MyApp",
			msg:     "This is a warning message",
			want:    "WRN MyApp: This is a warning message\n",
			wantErr: false,
		},
		{
			name:    "default error",
			level:   "error",
			styles:  "default",
			prefix:  "MyApp",
			msg:     "This is an error message",
			want:    "ERR MyApp: This is an error message\n",
			wantErr: false,
		},
		{
			name:    "labeled debug",
			level:   "debug",
			styles:  "labeled",
			prefix:  "MyApp",
			msg:     "This is a debug message",
			want:    " DBG  <> MyApp: This is a debug message\n",
			wantErr: false,
		},
		{
			name:    "labeled info",
			level:   "info",
			styles:  "labeled",
			prefix:  "MyApp",
			msg:     "This is an infomational message",
			want:    " INF  MyApp: This is an infomational message\n",
			wantErr: false,
		},
		{
			name:    "labeled warn",
			level:   "warn",
			styles:  "labeled",
			prefix:  "MyApp",
			msg:     "This is a warning message",
			want:    " WRN  MyApp: This is a warning message\n",
			wantErr: false,
		},
		{
			name:    "labeled error",
			level:   "error",
			styles:  "labeled",
			prefix:  "MyApp",
			msg:     "This is an error message",
			want:    " ERR  MyApp: This is an error message\n",
			wantErr: false,
		},
		{
			name:    "invalid level",
			level:   "invalid level",
			styles:  "default",
			prefix:  "MyApp",
			msg:     "",
			want:    "",
			wantErr: true,
		},
		{
			name:    "invalid styles",
			level:   "info",
			styles:  "invalid styles",
			prefix:  "MyApp",
			msg:     "",
			want:    "",
			wantErr: true,
		},
	}
)

func TestNewAppLogger(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			l, err := NewAppLogger(w, tt.level, tt.styles, tt.prefix)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAppLogger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			switch tt.level {
			case "debug":
				l.Debug(tt.msg)
			case "info":
				l.Info(tt.msg)
			case "warn":
				l.Warn(tt.msg)
			case "error":
				l.Error(tt.msg)
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
			l, err := NewSDKLogger(w, tt.level, tt.styles, tt.prefix)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSDKLogger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			switch tt.level {
			case "debug":
				l.Debug(tt.msg)
			case "info":
				l.Info(tt.msg)
			case "warn":
				l.Warn(tt.msg)
			case "error":
				l.Error(tt.msg)
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
			l, err := NewSDKLogger(w, tt.level, tt.styles, tt.prefix)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSDKLogger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			switch tt.level {
			case "debug":
				l.Debug(tt.msg)
			case "info":
				l.Info(tt.msg)
			case "warn":
				l.Warn(tt.msg)
			case "error":
				l.Error(tt.msg)
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
