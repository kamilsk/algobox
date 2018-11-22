package main

import (
	"reflect"
	"testing"
)

func TestCircularArrayRotation(t *testing.T) {
	tests := []struct {
		name     string
		a        []int32
		k        int32
		queries  []int32
		expected []int32
	}{
		{"Test case 15", []int32{1, 2, 3}, 2, []int32{0, 1, 2}, []int32{2, 3, 1}},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := circularArrayRotation(tc.a, tc.k, tc.queries); !reflect.DeepEqual(tc.expected, obtained) {
				t.Errorf("expected %v, obtained %bv", tc.expected, obtained)
			}
		})
	}
}
