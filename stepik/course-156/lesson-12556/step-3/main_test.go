package main

import (
	"reflect"
	"testing"
)

const format = `
expected %+v
obtained %+v`

func TestCompare(t *testing.T) {
	tests := []struct {
		name     string
		first    []int
		second   []int
		expected [2]int
	}{
		{"case#01", []int{4, -8, 6, 0}, []int{-10, 3, 1, 1}, [2]int{0, 1}},
		{"case#02", []int{1, 1, 39, 40, 1, 1}, []int{1, 30, 29, 1, 1, 1}, [2]int{2, 2}},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if i0, j0 := compare(tc.first, tc.second); !reflect.DeepEqual(tc.expected, [2]int{i0, j0}) {
				t.Errorf(format, tc.expected, [2]int{i0, j0})
			}
		})
	}
}
