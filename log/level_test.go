package log

import (
	"reflect"
	"testing"
)

func TestParseLevel(t *testing.T) {
	type args struct {
		level string
	}
	tests := []struct {
		name    string
		args    args
		want    Level
		wantErr bool
	}{
		{
			name:    "debug",
			args:    args{level: "debug"},
			want:    DebugLevel,
			wantErr: false,
		},
		{
			name:    "default",
			args:    args{level: "info"},
			want:    InfoLevel,
			wantErr: false,
		},
		{
			name:    "warn",
			args:    args{level: "warn"},
			want:    WarnLevel,
			wantErr: false,
		},
		{
			name:    "error",
			args:    args{level: "error"},
			want:    ErrorLevel,
			wantErr: false,
		},
		{
			name:    "fatal",
			args:    args{level: "fatal"},
			want:    FatalLevel,
			wantErr: false,
		},
		{
			name:    "unknown",
			args:    args{level: "unknown"},
			want:    InfoLevel,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseLevel(tt.args.level)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseLevel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
