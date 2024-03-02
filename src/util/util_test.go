package util

import (
	"reflect"
	"testing"

	"github.com/nsf/termbox-go"
)

func TestFlagColor(t *testing.T) {
	type args struct {
		ColorString string
	}
	tests := []struct {
		name string
		args args
		want termbox.Attribute
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagColor(tt.args.ColorString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlagColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
