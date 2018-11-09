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
		{"discussion=376300", [][2]int{{0, 7}, {0, 0}, {2, 0}, {3, 3}, {4, 0}, {5, 0}}, 3, []int{0, 7, 7, -1, -1, -1}},
		{"discussion=376375&reply=376392", [][2]int{
			{6, 23}, {15, 44}, {24, 28}, {25, 15}, {33, 7}, {47, 41}, {58, 25}, {65, 5}, {70, 14}, {79, 8}, {93, 43},
			{103, 11}, {110, 25}, {123, 27}, {138, 40}, {144, 19}, {159, 2}, {167, 23}, {179, 43}, {182, 31}, {186, 7},
			{198, 16}, {208, 41}, {222, 23}, {235, 26},
		}, 11, []int{
			6, 29, 73, 101, 116, 123, 164, 189, 194, 208, 216,
			259, 270, 295, 322, 362, -1, 381, -1, -1, -1,
			404, 420, 461, 484,
		}},
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
