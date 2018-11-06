package main

import (
	"reflect"
	"testing"
)

var format = `
expected %+v
obtained %+v`

func TestBirthdayCakeCandles(t *testing.T) {
	tests := []struct {
		name     string
		ar       []int32
		expected int32
	}{
		{"case#01", []int32{3, 2, 1, 3}, 2},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := birthdayCakeCandles(tc.ar); !reflect.DeepEqual(tc.expected, obtained) {
				t.Errorf(format, tc.expected, obtained)
			}
		})
	}
}
