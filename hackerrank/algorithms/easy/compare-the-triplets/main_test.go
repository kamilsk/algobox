package main

import (
	"reflect"
	"testing"
)

const format = `
expected %+v
obtained %+v`

func TestCompareTriplets(t *testing.T) {
	tests := []struct {
		name     string
		a, b     []int32
		expected []int32
	}{
		{"case#01", []int32{5, 6, 7}, []int32{3, 6, 10}, []int32{1, 1}},
		{"case#02", []int32{17, 28, 30}, []int32{99, 16, 8}, []int32{2, 1}},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := compareTriplets(tc.a, tc.b); !reflect.DeepEqual(tc.expected, obtained) {
				t.Errorf(format, tc.expected, obtained)
			}
		})
	}
}
