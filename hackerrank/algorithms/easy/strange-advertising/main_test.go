package main

import "testing"

func TestViralAdvertising(t *testing.T) {
	tests := []struct {
		name string
		n int32
		expected int32
	}{
		{"Test case 0", 3, 9},
		{"Test case 1", 4, 15},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := viralAdvertising(tc.n); tc.expected != obtained {
				t.Errorf("expected %d, obtained %d", tc.expected, obtained)
			}
		})
	}
}
