package main

import "testing"

func TestEqualizeArray(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int32
		expected int32
	}{
		{"Test case 0", []int32{3, 3, 2, 1, 3}, 2},
		{"Test case 17", []int32{1, 2, 3, 1, 2, 3, 3, 3}, 4},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := equalizeArray(tc.arr); tc.expected != obtained {
				t.Errorf("expected %d, obtained %d", tc.expected, obtained)
			}
		})
	}
}
