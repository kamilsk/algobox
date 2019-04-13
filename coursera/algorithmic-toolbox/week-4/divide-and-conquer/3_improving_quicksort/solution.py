# python3
from collections import namedtuple
from random import randint
from sys import stdin
from typing import List
from unittest import TestCase

test = namedtuple('test', 'nums expected')


def partition3(nums: List[int], left: int, right: int) -> (int, int):
    pivot = nums[left]
    lt, gt = left, right
    i = left
    while i <= gt:
        if nums[i] > pivot:
            nums[gt], nums[i] = nums[i], nums[gt]
            gt -= 1
            continue
        if nums[i] < pivot:
            nums[lt], nums[i] = nums[i], nums[lt]
            lt += 1
            i += 1
            continue
        i += 1
    return lt, gt


def partition2(nums: List[int], left: int, right: int) -> int:
    pivot = nums[left]
    j = left
    for i in range(left + 1, right + 1):
        if nums[i] <= pivot:
            j += 1
            nums[i], nums[j] = nums[j], nums[i]
    nums[left], nums[j] = nums[j], nums[left]
    return j


def randomized_quick_sort(nums: List[int], lo: int, hi: int):
    if lo >= hi:
        return
    k = randint(lo, hi)
    nums[lo], nums[k] = nums[k], nums[lo]
    # left = right = partition2(nums, lo, hi)
    left, right = partition3(nums, lo, hi)
    randomized_quick_sort(nums, lo, left - 1)
    randomized_quick_sort(nums, right + 1, hi)


class Test(TestCase):
    def test_randomized_quick_sort(self):
        tests = [
            # samples
            test([2, 3, 9, 2, 2], [2, 2, 2, 3, 9]),

            # acceptance
            test([2, 3, 9, 2, 9], [2, 2, 3, 9, 9]),
        ]
        for i, t in enumerate(tests):
            randomized_quick_sort(t.nums, 0, len(t.nums) - 1)
            self.assertEqual(t.expected, t.nums, msg='at {} position'.format(i))


if __name__ == '__main__':
    n, *a = list(map(int, stdin.read().split()))
    randomized_quick_sort(a, 0, n - 1)
    for x in a:
        print(x, end=' ')
