# python3
from collections import namedtuple
from sys import stdin
from typing import List
from unittest import TestCase

test = namedtuple('test', 'where what expected')


def binary_search(where: List[int], what: int) -> int:
    left, right = 0, len(where) - 1
    while left <= right:
        middle = (left + right) // 2
        if what == where[middle]:
            return middle
        if what > where[middle]:
            left = middle + 1
            continue
        right = middle - 1
    return -1


class Test(TestCase):
    def test_binary_search(self):
        tests = [
            # samples
            test([1, 5, 8, 12, 13], [8, 1, 23, 1, 11], [2, 0, -1, 0, -1]),

            # acceptance
            test([1, 2, 3, 4, 5], [1, 2, 3, 4, 5], [0, 1, 2, 3, 4]),
        ]
        for i, t in enumerate(tests):
            obtained = []
            for what in t.what:
                obtained.append(binary_search(t.where, what))
            self.assertEqual(t.expected, obtained, msg='at {} position'.format(i))


if __name__ == '__main__':
    data = list(map(int, stdin.read().split()))
    n = data[0]
    m, a = data[n + 1], data[1: n + 1]
    for x in data[n + 2:]:
        print(binary_search(a, x), end=' ')
