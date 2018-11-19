package main

import "testing"

func TestAngryProfessor(t *testing.T) {
	tests := []struct {
		name     string
		k        int32
		a        []int32
		expected string
	}{
		{"Testcase 0#1", 3, []int32{-1, -3, 4, 2}, "YES"},
		{"Testcase 0#2", 2, []int32{0, -1, 2, 1}, "NO"},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := angryProfessor(tc.k, tc.a); tc.expected != obtained {
				t.Errorf("expected %s, obtained %s", tc.expected, obtained)
			}
		})
	}
}
