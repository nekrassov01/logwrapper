package log

import (
	"reflect"
	"testing"
)

func TestStyle_String(t *testing.T) {
	tests := []struct {
		name string
		tr   Style
		want string
	}{
		{
			name: "default",
			tr:   DefaultStyle,
			want: "default",
		},
		{
			name: "labeled",
			tr:   LabeledStyle,
			want: "labeled",
		},
		{
			name: "unknown",
			tr:   2,
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.String(); got != tt.want {
				t.Errorf("Style.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseStyles(t *testing.T) {
	type args struct {
		style string
	}
	tests := []struct {
		name    string
		args    args
		want    *Styles
		wantErr bool
	}{
		{
			name:    "default",
			args:    args{style: "default"},
			want:    DefaultStyles(),
			wantErr: false,
		},
		{
			name:    "labeled",
			args:    args{style: "labeled"},
			want:    LabeledStyles(),
			wantErr: false,
		},
		{
			name:    "unknown",
			args:    args{style: "unknown"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseStyles(tt.args.style)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseStyles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseStyles() = %v, want %v", got, tt.want)
			}
		})
	}
}
