package main

import "testing"

func TestRepeatedString(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		n        int64
		expected int64
	}{
		{"Test case 0", "aba", 10, 7},
		{"Test case 1", "a", 1000000000000, 1000000000000},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := repeatedString(tc.s, tc.n); tc.expected != obtained {
				t.Errorf("expected %d, obtained %d", tc.expected, obtained)
			}
		})
	}
}
