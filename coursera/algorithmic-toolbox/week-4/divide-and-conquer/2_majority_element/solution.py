# python3
from collections import namedtuple
from sys import stdin
from typing import List
from unittest import TestCase

test = namedtuple('test', 'votes expected')


def majority_element(votes: List[int], left: int, right: int) -> int:
    if left is right:
        return -1
    if left + 1 is right:
        return votes[left] if votes[left] is votes[right] else -1

    ml = majority_element(votes, left, (left + right) // 2)
    mr = majority_element(votes, (left + right) // 2, right)

    if ml is -1 and mr is not -1:
        return mr
    if mr is -1 and ml is not -1:
        return ml
    return -1


class Test(TestCase):
    def test_majority_element(self):
        tests = [
            # samples
            test([2, 3, 9, 2, 2], 2),
            test([1, 2, 3, 4], -1),
            test([1, 2, 3, 1], -1),

            # acceptance
            test([512766168, 717383758, 5, 126144732, 5, 573799007, 5, 5, 5, 405079772], -1),

            # additional
            test([1, 2, 2, 2, 1], 2),
            test([2, 2, 1, 2, 2], 2),
            test([1, 2, 2, 1], -1),
        ]
        for t in tests:
            self.assertEqual(t.expected, majority_element(t.votes, 0, len(t.votes) - 1))


if __name__ == '__main__':
    n, *a = list(map(int, stdin.read().split()))
    print(0 if majority_element(a, 0, n - 1) is -1 else 1)
