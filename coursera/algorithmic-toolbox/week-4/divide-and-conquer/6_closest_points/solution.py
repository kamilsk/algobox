# python3
from collections import namedtuple
from math import sqrt
from sys import stdin
from typing import List
from unittest import TestCase

point = namedtuple('point', 'x y')
test = namedtuple('test', 'x y expected')


def closest_pair(points: List[point]) -> (point, point):
    return None, None


def closes_split_pair(points: List[point], delta: float) -> (point, point):
    return None, None


def distance(p1: point, p2: point) -> float:
    return sqrt((p1.x - p2.x) ** 2 + (p1.y - p2.y) ** 2)


def naive_minimum_distance(x: List[int], y: List[int]) -> float:
    points = [point(x[i], y[i]) for i in range(len(x))]

    result = distance(point(-10 ** 9, -10 ** 9), point(10 ** 9, 10 ** 9))
    for i in range(len(points)):
        for j in range(i + 1, len(points)):
            result = min(result, distance(points[i], points[j]))

    return result


def fast_minimum_distance(x: List[int], y: List[int]) -> float:
    pass


class Test(TestCase):
    def test_minimum_distance(self):
        tests = [
            # samples
            test([0, 3], [0, 4], 5.0),
            test([7, 1, 4, 7], [7, 100, 8, 7], 0.0),
            test([4, -2, -3, -1, 2, -4, 1, -1, 3, -4, -2], [4, -2, -4, 3, 3, 0, 1, -1, -1, 2, 4], 1.414213),
        ]
        for i, t in enumerate(tests):
            self.assertEqual(t.expected, int(10 ** 6 * naive_minimum_distance(t.x, t.y)) / 10 ** 6,
                             msg='at {} position'.format(i))


if __name__ == '__main__':
    data = list(map(int, stdin.read().split()))
    n = data[0]
    print('{0:.9f}'.format(fast_minimum_distance(data[1::2], data[2::2])))
