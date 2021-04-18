package main

import "testing"

// Solution for https://leetcode.com/problems/contains-duplicate-iii/.
func TestContainsNearbyAlmostDuplicate(t *testing.T) {
	tests := []struct {
		nums     []int
		k, t     int
		expected bool
	}{
		{[]int{1, 2, 3, 1}, 3, 0, true},
		{[]int{1, 0, 1, 1}, 1, 2, true},
		{[]int{1, 5, 9, 1, 5, 9}, 2, 3, false},

		{[]int{1, 2, 1, 1}, 1, 0, true},
		{[]int{1, 2, 5, 6, 7, 2, 4}, 4, 0, true},
		{[]int{7, 1, 3}, 2, 3, true},
		{[]int{2147483647, -1, 2147483647}, 1, 2147483647, false}, // bucket fails
	}

	for i, test := range tests {
		obtained := containsNearbyAlmostDuplicate(test.nums, test.k, test.t)
		if obtained != test.expected {
			t.Errorf("expected: %v, obtained: %v (case %d)",
				test.expected, obtained, i)
		}
	}
}

// naive - n*k, bucket - n
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	abs := func(num int) int {
		if num < 0 {
			return -num
		}
		return num
	}

	naive := func(nums []int, k int, t int) bool {
		for i, n := range nums {
			window := i + k + 1
			if window > len(nums) {
				window = len(nums)
			}
			for _, m := range nums[i+1 : window] {
				if abs(n-m) <= t {
					return true
				}
			}
		}
		return false
	}

	bucket := func(nums []int, k int, t int) bool {
		buckets := make(map[int]int)
		for i, num := range nums {
			b := num / (t + 1)
			if _, has := buckets[b]; has {
				return true
			}
			buckets[b] = num
			if _, has := buckets[b-1]; has && num-buckets[b-1] <= t {
				return true
			}
			if _, has := buckets[b+1]; has && buckets[b+1]-num <= t {
				return true
			}
			if i >= k {
				delete(buckets, nums[i-k]/(t+1))
			}
		}
		return false
	}

	_ = bucket
	return naive(nums, k, t)
}
