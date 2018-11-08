package main

import (
	"reflect"
	"testing"
)

const format = `
expected %+v
obtained %+v`

func TestSelection(t *testing.T) {
	tests := []struct {
		name     string
		array    []int
		expected []int
	}{
		{"case#01", []int{3, 1, 2}, []int{1, 2, 3}},
		{"case#02", []int{3, 2, 1}, []int{1, 2, 3}},
		{"case#03", []int{1, 2, 3}, []int{1, 2, 3}},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := selection(tc.array); !reflect.DeepEqual(tc.expected, obtained) {
				t.Errorf(format, tc.expected, obtained)
			}
		})
	}
}

func TestInsertion(t *testing.T) {
	tests := []struct {
		name     string
		array    []int
		expected []int
	}{
		{"case#01", []int{3, 1, 2}, []int{1, 2, 3}},
		{"case#02", []int{3, 2, 1}, []int{1, 2, 3}},
		{"case#03", []int{1, 2, 3}, []int{1, 2, 3}},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := insertion(tc.array); !reflect.DeepEqual(tc.expected, obtained) {
				t.Errorf(format, tc.expected, obtained)
			}
		})
	}
}

func TestBubble(t *testing.T) {
	tests := []struct {
		name     string
		array    []int
		expected []int
	}{
		{"case#01", []int{3, 1, 2}, []int{1, 2, 3}},
		{"case#02", []int{3, 2, 1}, []int{1, 2, 3}},
		{"case#03", []int{1, 2, 3}, []int{1, 2, 3}},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := bubble(tc.array); !reflect.DeepEqual(tc.expected, obtained) {
				t.Errorf(format, tc.expected, obtained)
			}
		})
	}
}
