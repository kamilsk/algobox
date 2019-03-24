# python3
from collections import namedtuple
from sys import stdin
from typing import List
from unittest import TestCase

Interval = namedtuple('Interval', 'start end')


def optimal_points(timeline: List[Interval]) -> List[int]:
    optimal = []

    timeline = sorted(timeline, key=lambda x: x.start)
    left, timeline = timeline[0], timeline[1:]
    right = left.end
    while timeline:
        current, timeline = timeline[0], timeline[1:]
        if current.start > right:
            optimal.append(right)
            left, right = current, current.end
            continue
        right = min(right, current.end)
    optimal.append(right)

    return optimal


class Test(TestCase):
    def test_optimal_points(self):
        tests = [
            {'timeline': [Interval(1, 3), Interval(2, 5), Interval(3, 6)], 'expected': [3]},
            {'timeline': [Interval(4, 7), Interval(1, 3), Interval(2, 5), Interval(5, 6)], 'expected': [3, 6]},
        ]
        for test in tests:
            expected, obtained = test['expected'], optimal_points(test['timeline'])
            self.assertEqual(expected, obtained)


if __name__ == '__main__':
    n, *data = map(int, stdin.read().split())
    segments = list(map(lambda x: Interval(x[0], x[1]), zip(data[::2], data[1::2])))
    points = optimal_points(segments)
    print(len(points))
    for p in points:
        print(p, end=' ')
