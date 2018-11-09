package main

import (
	"reflect"
	"testing"
)

const format = `
expected %+v
obtained %+v`

func TestBreakingRecords(t *testing.T) {
	tests := []struct {
		name     string
		scores   []int32
		expected []int32
	}{
		{"Testcase 0", []int32{10, 5, 20, 20, 4, 5, 2, 25, 1}, []int32{2, 4}},
		{"Testcase 1", []int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42}, []int32{4, 0}},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := breakingRecords(tc.scores); !reflect.DeepEqual(tc.expected, obtained) {
				t.Errorf(format, tc.expected, obtained)
			}
		})
	}
}
