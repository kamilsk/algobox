package main

import "testing"

func TestPageCount(t *testing.T) {
	tests := []struct {
		name     string
		n, p     int32
		expected int32
	}{
		{"Testcase 0", 6, 2, 1},
		{"Testcase 1", 5, 4, 0},
		{"Testcase 26", 6, 5, 1},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := pageCount(tc.n, tc.p); tc.expected != obtained {
				t.Errorf("expected %d, obtained %d", tc.expected, obtained)
			}
		})
	}
}
