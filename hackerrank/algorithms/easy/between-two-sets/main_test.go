package main

import "testing"

func TestGetTotalX(t *testing.T) {
	tests := []struct {
		name     string
		a, b     []int32
		expected int32
	}{
		{"Testcase 0", []int32{2, 4}, []int32{16, 32, 96}, 3},
		{"Testcase 8", []int32{3, 4}, []int32{24, 48}, 2},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := getTotalX(tc.a, tc.b); tc.expected != obtained {
				t.Errorf("expected %d, obtained %d", tc.expected, obtained)
			}
		})
	}
}
