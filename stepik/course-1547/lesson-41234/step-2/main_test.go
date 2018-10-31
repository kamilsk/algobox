package main

import "testing"

func TestHeight(t *testing.T) {
	tests := []struct {
		name     string
		parents  []int
		expected int
	}{
		{"case#01", []int{4, -1, 4, 1, 1}, 3},
		{"case#02", []int{-1, 0, 4, 0, 3}, 4},
		{"case#03", []int{9, 7, 5, 5, 2, 9, 9, 9, 2, -1}, 4},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := height(tc.parents); obtained != tc.expected {
				t.Errorf("expected %d, obtained %d", tc.expected, obtained)
			}
		})
	}
}
