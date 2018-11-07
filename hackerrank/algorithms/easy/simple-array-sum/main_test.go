package main

import (
	"reflect"
	"testing"
)

const format = `
expected %+v
obtained %+v`

func TestSimpleArraySum(t *testing.T) {
	tests := []struct {
		name     string
		ar       []int32
		expected int32
	}{
		{"case#01", []int32{1, 2, 3, 4, 10, 11}, 31},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := simpleArraySum(tc.ar); !reflect.DeepEqual(tc.expected, obtained) {
				t.Errorf(format, tc.expected, obtained)
			}
		})
	}
}
