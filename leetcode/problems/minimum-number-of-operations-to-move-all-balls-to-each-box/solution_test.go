package main

import (
	"reflect"
	"testing"
)

// Solution for https://leetcode.com/problems/minimum-number-of-operations-to-move-all-balls-to-each-box/.
func TestMinOperations(t *testing.T) {
	tests := []struct {
		boxes    string
		expected []int
	}{
		{"110", []int{1, 1, 3}},
		{"001011", []int{11, 8, 5, 4, 3, 4}},
	}

	for i, test := range tests {
		obtained := minOperations(test.boxes)
		if !reflect.DeepEqual(obtained, test.expected) {
			t.Errorf("expected: %v, obtained: %v (case %d)",
				test.expected, obtained, i)
		}
	}
}

// n^2
func minOperations(boxes string) []int {
	result := make([]int, len(boxes))
	for i := range boxes {
		for j, r := range boxes {
			if i == j {
				continue
			}
			if r == '1' {
				if i > j {
					result[i] += i - j
					continue
				}
				result[i] += j - i
			}
		}
	}
	return result
}
