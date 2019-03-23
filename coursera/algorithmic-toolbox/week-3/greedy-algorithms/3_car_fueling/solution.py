# python3
from sys import stdin
from typing import List
from unittest import TestCase


def compute_min_refills(distance: int, tank: int, points: List[int]) -> int:
    points.insert(0, 0)
    if distance != points[-1]:
        points.append(distance)

    num_refills = current_refill = 0
    destination = len(points) - 1

    while current_refill < destination:
        last_refill = current_refill
        while current_refill < destination and points[current_refill + 1] - points[last_refill] <= tank:
            current_refill += 1
        if last_refill is current_refill:
            return -1
        if current_refill < destination:
            num_refills += 1

    return num_refills


class Test(TestCase):
    def test_compute_min_refills(self):
        tests = [
            {'distance': 950, 'tank': 400, 'points': [200, 375, 550, 750], 'expected': 2},
            {'distance': 10, 'tank': 3, 'points': [1, 2, 5, 9], 'expected': -1},
            {'distance': 200, 'tank': 250, 'points': [100, 150], 'expected': 0},
        ]
        for test in tests:
            expected, obtained = test['expected'], compute_min_refills(test['distance'], test['tank'], test['points'])
            self.assertEqual(expected, obtained)


if __name__ == '__main__':
    d, m, _, *stops = map(int, stdin.read().split())
    print(compute_min_refills(d, m, stops))
