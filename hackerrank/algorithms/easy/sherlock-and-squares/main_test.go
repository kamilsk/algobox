package main

import "testing"

func TestSquares(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int32
		expected int32
	}{
		{"Test case 8#1", 3, 9, 2},
		{"Test case 8#2", 17, 24, 0},
		{"Test case 9#1", 35, 70, 3},
		{"Test case 9#2", 100, 1000, 22},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := squares(tc.a, tc.b); tc.expected != obtained {
				t.Errorf("expected %d, obtained %d", tc.expected, obtained)
			}
		})
	}
}
