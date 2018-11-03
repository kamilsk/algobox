package main

import (
	"reflect"
	"testing"
)

func TestStack(t *testing.T) {
	tests := []struct {
		name     string
		commands []string
		expected []int
	}{
		{"case#01", []string{"push 1", "push 7", "pop"}, []int{}},
		{"case#02", []string{"push 2", "push 1", "max", "pop", "max"}, []int{2, 2}},
		{"case#03", []string{"push 7", "push 1", "push 7", "max", "pop", "max"}, []int{7, 7}},
		{"case#04", []string{"push 1", "push 2", "max", "pop", "max"}, []int{2, 1}},
		{"case#05", []string{"push 2", "push 3", "push 9", "push 7", "push 2", "max", "max", "max", "pop", "max"},
			[]int{9, 9, 9, 9}},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := New(len(tc.commands)).process(tc.commands); !reflect.DeepEqual(tc.expected, obtained) {
				t.Errorf("expected %d, obtained %d", tc.expected, obtained)
			}
		})
	}
}
