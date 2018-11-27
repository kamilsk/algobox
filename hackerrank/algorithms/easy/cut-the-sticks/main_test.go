package main

import (
	"reflect"
	"testing"
)

func TestCutTheSticks(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int32
		expected []int32
	}{
		{"Test case 8", []int32{5, 4, 4, 2, 2, 8}, []int32{6, 4, 2, 1}},
		{"Test case 9", []int32{1, 2, 3, 4, 3, 3, 2, 1}, []int32{8, 6, 4, 1}},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := cutTheSticks(tc.arr); !reflect.DeepEqual(tc.expected, obtained) {
				t.Errorf("expected %v, obtained %v", tc.expected, obtained)
			}
		})
	}
}
