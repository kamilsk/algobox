# python3
from collections import namedtuple
from sys import stdin
from typing import List
from unittest import TestCase

test = namedtuple('test', 'x y expected')


def minimum_distance(x: List[int], y: List[int]) -> int:
    return 10 ** 18


class Test(TestCase):
    def test_minimum_distance(self):
        tests = [
            # samples
            test([0, 3], [0, 4], 5.0),
            test([7, 1, 4, 7], [7, 100, 8, 7], 0.0),
            test([4, -2, -3, -1, 2, -4, 1, -1, 3, -4, -2], [4, -2, -4, 3, 3, 0, 1, -1, -1, 2, 4], 1.414213),
        ]
        for i, t in enumerate(tests):
            self.assertEqual(t.expected, minimum_distance(t.x, t.y), msg='at {} position'.format(i))


if __name__ == '__main__':
    data = list(map(int, stdin.read().split()))
    n = data[0]
    print("{0:.9f}".format(minimum_distance(data[1::2], data[2::2])))
