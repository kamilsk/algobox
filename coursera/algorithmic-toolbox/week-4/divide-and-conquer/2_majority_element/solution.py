# python3
from collections import namedtuple
from functools import reduce
from sys import stdin
from typing import List
from unittest import TestCase

test = namedtuple('test', 'votes expected')


def majority_element(votes: List[int], left: int, right: int) -> (int, int):
    if left is right:
        return -1, 0
    if left + 1 is right:
        return votes[left], 1

    middle = (left + right) // 2

    ml, cl = majority_element(votes, left, middle)
    if ml is not -1:
        cl += reduce(lambda count, x: count if x is not ml else count + 1, votes[middle:right], 0)
        if cl > (right - left) // 2:
            return ml, cl

    mr, cr = majority_element(votes, middle, right)
    if mr is not -1:
        cr += reduce(lambda count, x: count if x is not mr else count + 1, votes[left:middle], 0)
        if cr > (right - left) // 2:
            return mr, cr

    return -1, 0


class Test(TestCase):
    def test_majority_element(self):
        tests = [
            # samples
            test([2, 3, 9, 2, 2], (2, 3)),
            test([1, 2, 3, 4], (-1, 0)),
            test([1, 2, 3, 1], (-1, 0)),

            # acceptance
            test([512766168, 717383758, 5, 126144732, 5, 573799007, 5, 5, 5, 405079772], (-1, 0)),

            # additional
            test([1, 2, 2, 2, 2, 1], (2, 4)),
            test([1, 2, 2, 2, 1], (2, 3)),
            test([2, 2, 1, 1, 2, 2], (2, 4)),
            test([2, 2, 1, 2, 2], (2, 4)),
            test([1, 1, 1, 1, 1], (1, 5)),
            test([1, 2, 2, 1], (-1, 0)),
            test([1, 1, 1, 1], (1, 4)),
            test([1], (1, 1))
        ]
        for i, t in enumerate(tests):
            self.assertEqual(t.expected, majority_element(t.votes, 0, len(t.votes)), msg='at {} position'.format(i))


if __name__ == '__main__':
    n, *a = list(map(int, stdin.read().split()))
    print(0 if majority_element(a, 0, n)[0] is -1 else 1)
