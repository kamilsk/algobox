package main

import (
	"reflect"
	"testing"
)

const format = `
expected %+v
obtained %+v`

func TestGradingStudents(t *testing.T) {
	tests := []struct {
		name     string
		grades   []int32
		expected []int32
	}{
		{"Testcase 0", []int32{73, 67, 38, 33}, []int32{75, 67, 40, 33}},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if obtained := gradingStudents(tc.grades); !reflect.DeepEqual(tc.expected, obtained) {
				t.Errorf(format, tc.expected, obtained)
			}
		})
	}
}
