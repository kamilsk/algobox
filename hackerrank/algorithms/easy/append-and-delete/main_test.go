package main

import "testing"

func TestAppendAndDelete(t *testing.T) {
	tests := []struct {
		name     string
		s, t     string
		k        int32
		expected string
	}{
		{"Test case 0", "hackerhappy", "hackerrank", 9, "Yes"},
		{"Test case 1", "aba", "aba", 7, "Yes"},
		{"Test case 4", "qwerasdf", "qwerbsdf", 6, "No"},
		{"Test case 5", "y", "yu", 2, "No"},
		{"Test case 9", "abcdef", "fedcba", 15, "Yes"},
		{"Test case 13", "ashley", "ash", 2, "No"},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := appendAndDelete(tc.s, tc.t, tc.k); tc.expected != obtained {
				t.Errorf("expected %s, obtained %s", tc.expected, obtained)
			}
		})
	}
}
