# python3
from collections import namedtuple
from unittest import TestCase

from sys import stdin
from typing import List

testCase = namedtuple('test', 'votes expected')


def majority_element(votes: List[int], left: int, right: int) -> int:
    if left == right:
        return -1
    if left + 1 == right:
        return votes[left]
    return -1


class Test(TestCase):
    def test_majority_element(self):
        tests = [
            testCase([2, 3, 9, 2, 2], 2),
            testCase([1, 2, 3, 4], -1),
            testCase([1, 2, 3, 1], -1),
        ]
        for test in tests:
            self.assertEqual(test.expected, majority_element(test.votes, 0, len(test.votes)))


if __name__ == '__main__':
    n, *a = list(map(int, stdin.read().split()))
    print(1 if majority_element(a, 0, n) != -1 else 0)
