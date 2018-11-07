package main

import "testing"

func TestGetTotalX(t *testing.T) {
	tests := []struct {
		name     string
		a, b     []int32
		expected int32
	}{
		{"case#01", []int32{}, []int32{}, 0},
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
