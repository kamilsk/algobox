package main

import "testing"

// Solution for https://leetcode.com/problems/non-decreasing-array/.
func TestCheckPossibility(t *testing.T) {
	tests := []struct {
		nums     []int
		expected bool
	}{
		{[]int{4, 2, 3}, true},
		{[]int{4, 2, 1}, false},

		{[]int{3, 4, 2, 3}, false},
	}

	for i, test := range tests {
		if obtained := checkPossibility(test.nums); obtained != test.expected {
			t.Errorf("expected: %v, obtained: %v (case %d)",
				test.expected, obtained, i)
		}
	}
}

// n
func checkPossibility(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	failFast := func(num int, sub []int) bool {
		for _, next := range sub {
			if next < num {
				return false
			}
			num = next
		}
		return true
	}

	for i, num := range nums[1:] {
		if num >= nums[i] {
			continue
		}
		if failFast(nums[i], nums[i+2:]) {
			return true
		}
		if i > 0 {
			return failFast(nums[i-1], nums[i+1:])
		}
		return failFast(num, nums[i+2:])
	}
	return true
}
