package main

import (
	"reflect"
	"testing"
)

const format = `
expected %+v
obtained %+v`

func TestHeap(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected [][]string
	}{
		{"case#01", []int{5, 4, 3, 2, 1}, [][]string{{"1", "4"}, {"0", "1"}, {"1", "3"}}},
		{"case#02", []int{1, 2, 3, 4, 5}, [][]string{}},
		{"case#03", []int{0, 1, 2, 3, 4, 5}, [][]string{}},
		{"case#04", []int{7, 6, 5, 4, 3, 2}, [][]string{{"2", "5"}, {"1", "4"}, {"0", "2"}, {"2", "5"}}},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := new(heap).build(tc.numbers); !reflect.DeepEqual(tc.expected, obtained) {
				t.Errorf(format, tc.expected, obtained)
			}
		})
	}
}
