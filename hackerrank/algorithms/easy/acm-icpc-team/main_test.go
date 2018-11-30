package main

import (
	"reflect"
	"testing"
)

func TestAcmTeam(t *testing.T) {
	tests := []struct {
		name     string
		topic    []string
		expected []int32
	}{
		{"Test case 0", []string{"10101", "11100", "11010", "00101"}, []int32{5, 2}},
		{"Test case 8", []string{"11101", "10101", "11001", "10111", "10000", "01110"}, []int32{5, 6}},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := acmTeam(tc.topic); !reflect.DeepEqual(tc.expected, obtained) {
				t.Errorf("expected %v, obtained %v", tc.expected, obtained)
			}
		})
	}
}
