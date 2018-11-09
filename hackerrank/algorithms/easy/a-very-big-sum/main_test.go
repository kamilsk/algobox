package main

import "testing"

func TestAVeryBigSum(t *testing.T) {
	tests := []struct {
		name     string
		ar       []int64
		expected int64
	}{
		{"Testcase 0", []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}, 5000000015},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := aVeryBigSum(tc.ar); tc.expected != obtained {
				t.Errorf("expected %d, obtained %d", tc.expected, obtained)
			}
		})
	}
}
