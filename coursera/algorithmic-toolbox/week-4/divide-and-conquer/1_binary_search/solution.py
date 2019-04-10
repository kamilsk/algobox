# python3
from unittest import TestCase

from sys import stdin
from typing import List


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
            {'where': [1, 5, 8, 12, 13], 'what': [8, 1, 23, 1, 11], 'expected': [2, 0, -1, 0, -1]},
            {'where': [1, 2, 3, 4, 5], 'what': [1, 2, 3, 4, 5], 'expected': [0, 1, 2, 3, 4]},
        ]
        for test in tests:
            obtained = []
            for what in test['what']:
                obtained.append(binary_search(test['where'], what))
            self.assertEqual(test['expected'], obtained)


if __name__ == '__main__':
    data = list(map(int, stdin.read().split()))
    n = data[0]
    m = data[n + 1]
    a = data[1: n + 1]
    for x in data[n + 2:]:
        print(binary_search(a, x), end=' ')
