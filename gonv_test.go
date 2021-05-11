package gonv

import (
	"reflect"
	"testing"
)

func TestAtob(t *testing.T) {
	tests := []struct {
		s    string
		want []byte
	}{
		{"gonv", []byte("gonv")},
		{"", []byte(nil)},
		{"\n", []byte("\n")},
		{"#%&", []byte("#%&")},
	}

	for _, tt := range tests {
		got := Atob(tt.s)
		if !reflect.DeepEqual(tt.want, got) {
			t.Errorf("not equals. want: %#v, but got: %#v", tt.want, got)
		}
	}
}
