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
		{"case#01", 3, 2, []weight{1, 2, 2}, 4},
		{"discussion=149459#2", 3, 6, []weight{4, 3, 5}, 5},
		{"discussion=149459#3", 7, 3, []weight{1, 1, 1, 1, 1, 1, 1}, 3},
		{"discussion=149459#4", 2, 6, []weight{4, 1}, 3},
		{"discussion=149459#5", 10, 10, []weight{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}, 11},
		{"discussion=149459#6", 4, 12, []weight{8, 3, 2, 1}, 4},
		{"discussion=148843#1", 12, 100, []weight{1, 4, 8, 9, 10, 40, 45, 50, 60, 70, 80, 100}, 13},
		{"discussion=148843#2", 7, 25, []weight{1, 3, 5, 7, 9, 11, 13}, 5},
		{"discussion=148843#3", 10, 11, []weight{1, 1, 1, 1, 1, 1, 1, 1, 1, 10}, 4},
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
