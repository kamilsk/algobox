package main

import "testing"

func TestFindDigits(t *testing.T) {
	tests := []struct {
		name     string
		n        int32
		expected int32
	}{
		{"Test case 0#1", 12, 2},
		{"Test case 0#2", 1012, 3},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := findDigits(tc.n); tc.expected != obtained {
				t.Errorf("expected %d, obtained %d", tc.expected, obtained)
			}
		})
	}
}
