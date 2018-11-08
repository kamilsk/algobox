package main

import (
	"reflect"
	"testing"
)

const format = `
expected %+v
obtained %+v`

func TestHeap(t *testing.T) {
	tests := []struct {
		name     string
		points   []point
		expected []point
	}{
		{"case#01", []point{
			{0, 0},
			{1, 1},
			{1, 0},
			{0, 1},
		}, []point{
			{0, 0},
			{0, 1},
			{1, 0},
			{1, 1},
		}},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := new(heap).sort(tc.points); !reflect.DeepEqual(tc.expected, obtained) {
				t.Errorf(format, tc.expected, obtained)
			}
		})
	}
}
