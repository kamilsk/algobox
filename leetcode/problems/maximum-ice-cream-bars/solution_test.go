package main

import (
	"sort"
	"testing"
)

// Solution for https://leetcode.com/problems/maximum-ice-cream-bars/.
func TestMaxIceCream(t *testing.T) {
	tests := []struct {
		costs    []int
		coins    int
		expected int
	}{
		{[]int{1, 3, 2, 4, 1}, 7, 4},
		{[]int{10, 6, 8, 7, 7, 8}, 5, 0},
		{[]int{1, 6, 3, 1, 2, 5}, 20, 6},
	}

	for i, test := range tests {
		if obtained := maxIceCream(test.costs, test.coins); obtained != test.expected {
			t.Errorf("expected: %d, obtained: %d (case %d)", test.expected, obtained, i)
		}
	}
}

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)

	var total int
	for _, cost := range costs {
		if coins >= cost {
			coins -= cost
			total++
		}
		if coins == 0 {
			break
		}
	}
	return total
}
