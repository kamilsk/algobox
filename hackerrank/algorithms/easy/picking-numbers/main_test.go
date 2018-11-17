package main

import "testing"

func TestPickingNumbers(t *testing.T) {
	tests := []struct {
		name     string
		a        []int32
		expected int32
	}{
		{"Testcase 0", []int32{4, 6, 5, 3, 3, 1}, 3},
		{"Testcase 1", []int32{1, 2, 2, 3, 1, 2}, 5},
		{"Testcase 7", []int32{
			4, 97, 5, 97, 97, 4, 97, 4, 97, 97, 97, 97, 4, 4, 5, 5, 97, 5, 97, 99, 4, 97, 5, 97, 97, 97, 5, 5, 97, 4, 5,
			97, 97, 5, 97, 4, 97, 5, 4, 4, 97, 5, 5, 5, 4, 97, 97, 4, 97, 5, 4, 4, 97, 97, 97, 5, 5, 97, 4, 97, 97, 5,
			4, 97, 97, 4, 97, 97, 97, 5, 4, 4, 97, 4, 4, 97, 5, 97, 97, 97, 97, 4, 97, 5, 97, 5, 4, 97, 4, 5, 97, 97, 5,
			97, 5, 97, 5, 97, 97, 97,
		}, 50},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := pickingNumbers(tc.a); tc.expected != obtained {
				t.Errorf("expected %d, obtained %d", tc.expected, obtained)
			}
		})
	}
}
