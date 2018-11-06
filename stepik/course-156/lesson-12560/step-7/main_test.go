package main

import "testing"

func TestGreedy(t *testing.T) {
	tests := []struct {
		name     string
		cap      int
		f        weight
		fruit    []weight
		expected int
	}{
		//{"case#01", 3, 2, []weight{1, 2, 2}, 4},
		//{"case#02", 3, 6, []weight{4, 3, 5}, 5},
		//{"case#03", 7, 3, []weight{1, 1, 1, 1, 1, 1, 1}, 3},
		//{"case#04", 2, 6, []weight{4, 1}, 3},
		{"case#05", 10, 10, []weight{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}, 12},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := (greedy{tc.f, new(basket).with(tc.cap).insert(tc.fruit...)}).eat(); tc.expected != obtained {
				t.Errorf("expected %d, obtained %d", tc.expected, obtained)
			}
		})
	}
}
