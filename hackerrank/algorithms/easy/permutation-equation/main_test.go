package main

import (
	"reflect"
	"testing"
)

func TestPermutationEquation(t *testing.T) {
	tests := []struct {
		name     string
		p        []int32
		expected []int32
	}{
		{"Test case 0", []int32{2, 3, 1}, []int32{2, 3, 1}},
		{"Test case 11", []int32{4, 3, 5, 1, 2}, []int32{1, 3, 5, 4, 2}},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := permutationEquation(tc.p); !reflect.DeepEqual(tc.expected, obtained) {
				t.Errorf("expected %v, obtained %v", tc.expected, obtained)
			}
		})
	}
}
