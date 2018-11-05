package main

import "testing"

func TestDEQ(t *testing.T) {
	tests := []struct {
		name       string
		operations [][2]int
		expected   bool
	}{
		{"case#01", [][2]int{
			{1, 44},
			{3, 50},
			{2, 44},
			{2, 50},
			{2, -1},
		}, true},
		{"case#02", [][2]int{
			{4, 44},
			{4, -1},
			{3, 22},
			{4, 22},
			{2, -1},
			{1, 11},
			{4, 11},
		}, false},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := new(deq).process(tc.operations); tc.expected != obtained {
				t.Errorf("expected %v, obtained %v", tc.expected, obtained)
			}
		})
	}
}
