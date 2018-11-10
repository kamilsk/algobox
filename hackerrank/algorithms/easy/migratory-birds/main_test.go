package main

import "testing"

func TestMigratoryBirds(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int32
		expected int32
	}{
		{"Testcase 0", []int32{1, 4, 4, 4, 5, 3}, 4},
		{"Testcase 5", []int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4}, 3},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := migratoryBirds(tc.arr); tc.expected != obtained {
				t.Errorf("expected %d, obtained %d", tc.expected, obtained)
			}
		})
	}
}
