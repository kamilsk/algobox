package main

import "testing"

func TestSaveThePrisoner(t *testing.T) {
	tests := []struct {
		name     string
		n, m, s  int32
		expected int32
	}{
		{"Test case 4#1", 5, 2, 1, 2},
		{"Test case 4#2", 5, 2, 2, 3},
		{"Test case 11#1", 7, 19, 2, 6},
		{"Test case 11#2", 3, 7, 3, 3},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := saveThePrisoner(tc.n, tc.m, tc.s); tc.expected != obtained {
				t.Errorf("expected %d, obtained %d", tc.expected, obtained)
			}
		})
	}
}
