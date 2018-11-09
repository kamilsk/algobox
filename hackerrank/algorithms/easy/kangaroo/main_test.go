package main

import "testing"

func TestKangaroo(t *testing.T) {
	tests := []struct {
		name     string
		x1, v1   int32
		x2, v2   int32
		expected string
	}{
		{"Testcase 0", 0, 3, 4, 2, "YES"},
		{"Testcase 1", 0, 2, 5, 3, "NO"},
		{"custom#01", 0, 2, 4, 2, "NO"},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := kangaroo(tc.x1, tc.v1, tc.x2, tc.v2); tc.expected != obtained {
				t.Errorf("expected %s, obtained %s", tc.expected, obtained)
			}
		})
	}
}
