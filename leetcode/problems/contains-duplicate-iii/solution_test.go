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
		{[]int{2147483647, -1, 2147483647}, 1, 2147483647, false},
		{[]int{-2147483648, 2147483647}, 1, 1, false},
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
	if len(nums) < 2 || k < 0 || t < 0 {
		return false
	}

	abs := func(num int) int {
		if num < 0 {
			return -num
		}
		return num
	}

	naive := func(nums []int, k int, t int) bool {
		for i, num := range nums {
			window := i + k + 1
			if window > len(nums) {
				window = len(nums)
			}
			for _, cmp := range nums[i+1 : window] {
				if abs(num-cmp) <= t {
					return true
				}
			}
		}
		return false
	}

	bucket := func(nums []int, k int, t int) bool {
		buckets, width := make(map[int]int), t+1
		for i, num := range nums {
			bucket := num / width
			if num < 0 {
				bucket--
			}

			if _, has := buckets[bucket]; has {
				return true
			}
			buckets[bucket] = num

			if cmp, has := buckets[bucket-1]; has && abs(num-cmp) <= t {
				return true
			}
			if cmp, has := buckets[bucket+1]; has && abs(num-cmp) <= t {
				return true
			}
			if i >= k {
				delete(buckets, nums[i-k]/width)
			}
		}
		return false
	}

	_, _ = naive, bucket
	return bucket(nums, k, t)
}
