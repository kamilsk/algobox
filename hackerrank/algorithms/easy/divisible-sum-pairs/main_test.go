package main

import "testing"

func TestDivisibleSumPairs(t *testing.T) {
	tests := []struct {
		name     string
		n, k     int32
		ar       []int32
		expected int32
	}{
		{"Testcase 0", 6, 3, []int32{1, 3, 2, 6, 1, 2}, 5},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := divisibleSumPairs(tc.n, tc.k, tc.ar); tc.expected != obtained {
				t.Errorf("expected %d, obtained %d", tc.expected, obtained)
			}
		})
	}
}
