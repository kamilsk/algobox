package main

import (
	"reflect"
	"testing"
)

func TestProcess(t *testing.T) {
	tests := []struct {
		name     string
		traffic  [][2]int
		capacity int
		expected []int
	}{
		{"case#01", [][2]int{}, 1, nil},
		{"case#02", [][2]int{{0, 0}}, 1, []int{0}},
		{"case#03", [][2]int{{0, 1}, {0, 1}}, 1, []int{0, -1}},
		{"case#04", [][2]int{{0, 1}, {1, 1}}, 1, []int{0, 1}},
		{"case#05", [][2]int{{0, 1}}, 1, []int{0}},
		{"discussion#376300", [][2]int{{0, 7}, {0, 0}, {2, 0}, {3, 3}, {4, 0}, {5, 0}}, 3, []int{0, 7, 7, -1, -1, -1}},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := process(tc.traffic, tc.capacity); !reflect.DeepEqual(tc.expected, obtained) {
				t.Errorf("expected %d, obtained %d", tc.expected, obtained)
			}
		})
	}
}
