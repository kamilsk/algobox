package main

import "testing"

func TestUtopianTree(t *testing.T) {
	tests := []struct {
		name     string
		n        int32
		expected int32
	}{

		{"Testcase 0#1", 0, 1},
		{"Testcase 0#1", 1, 2},
		{"Testcase 1#1", 4, 7},
		{"Testcase 2#2", 3, 6},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := utopianTree(tc.n); tc.expected != obtained {
				t.Errorf("expected %d, obtained %d", tc.expected, obtained)
			}
		})
	}
}
