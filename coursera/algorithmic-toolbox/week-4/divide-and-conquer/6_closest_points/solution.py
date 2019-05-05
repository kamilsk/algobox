# python3
from collections import namedtuple
from math import sqrt
from random import randint
from sys import stdin
from typing import List, Optional
from unittest import TestCase

pair = namedtuple('pair', 'p q')
point = namedtuple('point', 'x y')
test = namedtuple('test', 'x y expected')


def naive_minimum_distance(x: List[int], y: List[int]) -> float:
    points = [point(x[i], y[i]) for i in range(len(x))]

    p1, p2 = None, None
    count, best = len(points), distance(point(-10 ** 9, -10 ** 9), point(10 ** 9, 10 ** 9))
    for i in range(count):
        for j in range(i + 1, count):
            d = distance(points[i], points[j])
            if best > d:
                best = d
                p1, p2 = points[i], points[j]

    return distance(p1, p2)


def fast_minimum_distance(x: List[int], y: List[int]) -> float:
    points = [point(x[i], y[i]) for i in range(len(x))]
    points.sort()

    p1, p2 = closest_pair(points)
    return distance(p1, p2)


def distance(p1: point, p2: point) -> Optional[float]:
    if p1 is None or p2 is None:
        return None
    return sqrt((p1.x - p2.x) ** 2 + (p1.y - p2.y) ** 2)


def closest_pair(points: List[point]) -> pair:
    if len(points) < 2:
        return pair(None, None)

    mid = len(points) // 2
    p1 = closest_pair(points[:mid])
    d1 = distance(*p1)
    p2 = closest_pair(points[mid:])
    d2 = distance(*p2)
    best = min(d1, d2) if d1 is not None and d2 is not None else (d1 if d2 is None else d2)

    strip = []
    for i in range(len(points)):
        if best is None or abs(points[i].x - points[mid].x) < best:
            strip.append(points[i])

    return choose([p1, p2, closest_split_pair(strip, best)])


def choose(points: List[pair]) -> pair:
    points = list(filter(lambda x: x.p is not None and x.q is not None, points))
    points.sort(key=lambda x: distance(*x), reverse=True)
    return points.pop()


def closest_split_pair(points: List[point], delta: float) -> pair:
    points.sort(key=lambda x: x.y)

    p1, p2 = None, None
    count, best = len(points), delta
    for i in range(count):
        for j in range(i + 1, count):
            d = distance(points[i], points[j])
            if best is None:
                best = d
                p1, p2 = points[i], points[j]
                continue
            if points[j].y - points[i].y >= best:
                break
            if best > d:
                best = d
                p1, p2 = points[i], points[j]

    return pair(p1, p2)


class Test(TestCase):
    def test_minimum_distance(self):
        tests = [
            # samples
            test([0, 3], [0, 4], 5.0),
            test([7, 1, 4, 7], [7, 100, 8, 7], 0.0),
            test([4, -2, -3, -1, 2, -4, 1, -1, 3, -4, -2], [4, -2, -4, 3, 3, 0, 1, -1, -1, 2, 4], 1.414213),

            # stress
            test([9, 9, 9, 0, 10, -3], [8, 0, 5, -4, 0, -8], 1.0),
            test([-127678177, -331970333, -800285983, -314341947, 233375965, -812839577, 322317687, -813575754,
                  304407788, 271791828, 391444914, 57930264, -547266271, -100947032, 980647496],
                 [-598319026, -950101333, -928393983, 522743622, 430981688, 436461075, -424578224, -396917943,
                  637314539, -383018865, -284424151, -319904913, 843840201, 855896685, 374328039],
                 65422035.646858),
        ]
        for i, t in enumerate(tests):
            self.assertEqual(t.expected, int(10 ** 6 * fast_minimum_distance(t.x, t.y)) / 10 ** 6,
                             msg='at {} position'.format(i))

    def test_stress(self):
        while True:
            points = []
            for i in range(2 * randint(2, 100)):
                points.append(randint(-pow(10, 2), pow(10, 2)))
            x, y = points[0::2], points[1::2]
            try:
                expected = '{0:.9f}'.format(naive_minimum_distance(x, y))
                obtained = '{0:.9f}'.format(fast_minimum_distance(x, y))
                if expected != obtained:
                    self.fail(msg='failed on input ({}, {})'.format(x, y))
            except (RuntimeError, TypeError) as e:
                self.fail(msg='failed on input ({}, {})'.format(x, y))


if __name__ == '__main__':
    data = list(map(int, stdin.read().split()))
    n = data[0]
    print('{0:.9f}'.format(fast_minimum_distance(data[1::2], data[2::2])))
