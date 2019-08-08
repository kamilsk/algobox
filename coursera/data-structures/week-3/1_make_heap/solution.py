# python3

from collections import namedtuple
from typing import List
from unittest import TestCase

swap = namedtuple('swap', 'i j')
test = namedtuple('test', 'data expected')


def build_heap(data: List[int]) -> List[swap]:
    swaps = []
    for i in range(len(data) // 2, -1, -1):
        swaps.extend(sift_down(i, data))
    return swaps


def sift_down(i: int, data: List[int]) -> List[swap]:
    swaps = []
    min_idx = i
    left = 2 * i + 1
    if left < len(data) and data[left] < data[min_idx]:
        min_idx = left
    right = 2 * (i + 1)
    if right < len(data) and data[right] < data[min_idx]:
        min_idx = right
    if i != min_idx:
        data[i], data[min_idx] = data[min_idx], data[i]
        swaps.append(swap(i, min_idx))
        swaps.extend(sift_down(min_idx, data))
    return swaps


class Test(TestCase):
    def test_make_heap(self):
        tests = [
            # samples
            test([5, 4, 3, 2, 1], [swap(1, 4), swap(0, 1), swap(1, 3)]),
            test([1, 2, 3, 4, 5], []),
        ]
        for i, t in enumerate(tests):
            self.assertEqual(t.expected, build_heap(t.data), msg='at {} position'.format(i))


if __name__ == '__main__':
    n = int(input())
    result = build_heap([int(s) for s in input().split()])
    print(len(result))
    for swap in result:
        print(swap.i, swap.j)
