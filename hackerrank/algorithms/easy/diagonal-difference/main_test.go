package main

import "testing"

func TestDiagonalDifference(t *testing.T) {
	tests := []struct {
		name     string
		arr      [][]int32
		expected int32
	}{
		{"case#01", [][]int32{
			{11, 2, 4},
			{4, 5, 6},
			{10, 8, -12},
		}, 15},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := diagonalDifference(tc.arr); tc.expected != obtained {
				t.Errorf("expected %d, obtained %d", tc.expected, obtained)
			}
		})
	}
}
