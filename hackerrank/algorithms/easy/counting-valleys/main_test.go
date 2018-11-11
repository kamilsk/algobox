package main

import "testing"

func TestCountingValleys(t *testing.T) {
	tests := []struct {
		name     string
		n        int32
		s        string
		expected int32
	}{
		{"Testcase 0", 8, "UDDDUDUU", 1},
		{"Testcase 1", 12, "DDUUDDUDUUUD", 2},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := countingValleys(tc.n, tc.s); tc.expected != obtained {
				t.Errorf("expected %d, obtained %d", tc.expected, obtained)
			}
		})
	}
}
