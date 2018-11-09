package main

import "testing"

func TestBirthday(t *testing.T) {
	tests := []struct {
		name     string
		s        []int32
		d, m     int32
		expected int32
	}{
		{"Testcase 0", []int32{1, 2, 1, 3, 2}, 3, 2, 2},
		{"Testcase 1", []int32{1, 1, 1, 1, 1, 1}, 3, 2, 0},
		{"Testcase 2", []int32{4}, 4, 1, 1},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := birthday(tc.s, tc.d, tc.m); tc.expected != obtained {
				t.Errorf("expected %d, obtained %d", tc.expected, obtained)
			}
		})
	}
}
