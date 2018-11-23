package main

import "testing"

func TestJumpingOnClouds(t *testing.T) {
	tests := []struct {
		name     string
		c        []int32
		k        int32
		expected int32
	}{
		{"Test case 0", []int32{0, 0, 1, 0, 0, 1, 1, 0}, 2, 92},
		{"Test case 8", []int32{1, 1, 1, 0, 1, 1, 0, 0, 0, 0}, 3, 94},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := jumpingOnClouds(tc.c, tc.k); tc.expected != obtained {
				t.Errorf("expected %d, obtained %d", tc.expected, obtained)
			}
		})
	}
}
