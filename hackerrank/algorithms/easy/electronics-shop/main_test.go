package main

import "testing"

func TestGetMoneySpent(t *testing.T) {
	tests := []struct {
		name              string
		keyboards, drives []int32
		b                 int32
		expected          int32
	}{
		{"Testcase 0", []int32{3, 1}, []int32{5, 2, 8}, 10, 9},
		{"Testcase 1", []int32{4}, []int32{5}, 5, -1},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := getMoneySpent(tc.keyboards, tc.drives, tc.b); tc.expected != obtained {
				t.Errorf("expected %d, obtained %d", tc.expected, obtained)
			}
		})
	}
}
