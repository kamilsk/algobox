package main

import "testing"

func TestJumpingOnClouds(t *testing.T) {
	tests := []struct {
		name     string
		c        []int32
		expected int32
	}{
		{"Test case 0", []int32{0, 0, 1, 0, 0, 1, 0}, 4},
		{"Test case 1", []int32{0, 0, 0, 1, 0, 0}, 3},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := jumpingOnClouds(tc.c); tc.expected != obtained {
				t.Errorf("expected %d, obtained %d", tc.expected, obtained)
			}
		})
	}
}
