package main

import "testing"

func TestTaumBday(t *testing.T) {
	tests := []struct {
		name            string
		b, w, bc, wc, z int32
		expected        int32
	}{
		{"0#1", 10, 10, 1, 1, 1, 20},
		{"0#2", 5, 9, 2, 3, 4, 37},
		{"0#3", 3, 6, 9, 1, 1, 12},
		{"0#4", 7, 7, 4, 2, 1, 35},
		{"0#5", 3, 3, 1, 9, 2, 12},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := taumBday(tc.b, tc.w, tc.bc, tc.wc, tc.z); tc.expected != obtained {
				t.Errorf("expected %d, obtained %d", tc.expected, obtained)
			}
		})
	}
}
