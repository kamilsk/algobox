package main

import (
	"reflect"
	"testing"
)

var format = `
expected %+v
obtained %+v`

func TestSplit(t *testing.T) {
	tests := []struct {
		name     string
		number   int
		expected []string
	}{
		{"case#01", 75, []string{"3", "5", "5"}},
		{"case#02", 2, []string{"2"}},
		{"case#03", 109, []string{"109"}},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := split(tc.number); !reflect.DeepEqual(tc.expected, obtained) {
				t.Errorf(format, tc.expected, obtained)
			}
		})
	}
}
