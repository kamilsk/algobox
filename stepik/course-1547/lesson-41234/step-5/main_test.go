package main

import (
	"reflect"
	"testing"
)

func TestMaximumsInSlidingWindow(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		window   int
		expected []string
	}{
		{"case#01", []int{2, 7, 3, 1, 5, 2, 6, 2}, 4, []string{"7", "7", "5", "6", "6"}},
		{"case#02", []int{2, 1, 5}, 1, []string{"2", "1", "5"}},
		{"case#03", []int{2, 3, 9}, 3, []string{"9"}},
		{"custom#01", []int{2, 3}, 5, []string{"3"}},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := New(tc.window).calculate(tc.numbers); !reflect.DeepEqual(tc.expected, obtained) {
				t.Errorf("expected %+v, obtained %+v", tc.expected, obtained)
			}
		})
	}
}
