package main

import "testing"

func TestHurdleRace(t *testing.T) {
	tests := []struct {
		name     string
		k        int32
		height   []int32
		expected int32
	}{
		{"Testcase 0", 4, []int32{1, 6, 3, 5, 2}, 2},
		{"Testcase 1", 7, []int32{2, 5, 4, 5, 2}, 0},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := hurdleRace(tc.k, tc.height); tc.expected != obtained {
				t.Errorf("expected %d, obtained %d", tc.expected, obtained)
			}
		})
	}
}
