package main

import "testing"

func TestSockMerchant(t *testing.T) {
	tests := []struct {
		name     string
		n        int32
		ar       []int32
		expected int32
	}{
		{"Testcase 0", 9, []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}, 3},
		{"Testcase 8", 9, []int32{1, 1, 3, 1, 2, 1, 3, 3, 3, 3}, 4},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := sockMerchant(tc.n, tc.ar); tc.expected != obtained {
				t.Errorf("expected %d, obtained %d", tc.expected, obtained)
			}
		})
	}
}
