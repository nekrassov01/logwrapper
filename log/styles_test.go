package log

import (
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
