package main

import (
	"reflect"
	"testing"
)

const format = `
expected %+v
obtained %+v`

func TestSearch(t *testing.T) {
	tests := []struct {
		name     string
		sorted   []int
		find     []int
		expected []string
	}{
		{"case#01", []int{10, 20, 30}, []int{9, 15, 35}, []string{"0", "0", "2"}},
		{"case#02", []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90}, []int{-1, 11, 40, 41, 92},
			[]string{"0", "1", "4", "4", "9"}},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := search(tc.sorted, tc.find); !reflect.DeepEqual(tc.expected, obtained) {
				t.Errorf(format, tc.expected, obtained)
			}
		})
	}
}
