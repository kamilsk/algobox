package main

import "testing"

func TestSolveMeFirst(t *testing.T) {
	tests := []struct {
		name     string
		a, b     uint32
		expected uint32
	}{
		{"Testcase 0", 2, 3, 5},
		{"Testcase 1", 100, 1000, 1100},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := solveMeFirst(tc.a, tc.b); tc.expected != obtained {
				t.Errorf("expected %d, obtained %d", tc.expected, obtained)
			}
		})
	}
}
