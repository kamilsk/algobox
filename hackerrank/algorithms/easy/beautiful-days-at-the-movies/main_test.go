package main

import "testing"

func TestBeautifulDays(t *testing.T) {
	tests := []struct {
		name     string
		i, j, k  int32
		expected int32
	}{
		{"Test case 8", 20, 23, 6, 2},
		{"Test case 9", 13, 45, 3, 33},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := beautifulDays(tc.i, tc.j, tc.k); tc.expected != obtained {
				t.Errorf("expected %d, obtained %d", tc.expected, obtained)
			}
		})
	}
}
