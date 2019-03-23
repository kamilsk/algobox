# python3
from sys import stdin
from typing import List
from unittest import TestCase


def max_dot_product(profits: List[int], clicks: List[int]) -> int:
    profits = sorted(profits, reverse=True)
    clicks = sorted(clicks, reverse=True)
    res = 0
    for i in range(len(profits)):
        res += profits[i] * clicks[i]
    return res


class Test(TestCase):
    def test_max_dot_product(self):
        tests = [
            {'profits': [23], 'clicks': [39], 'expected': 897},
            {'profits': [1, 3, -5], 'clicks': [-2, 4, 1], 'expected': 23},
        ]
        for test in tests:
            expected, obtained = test['expected'], max_dot_product(test['profits'], test['clicks'])
            self.assertEqual(expected, obtained)


if __name__ == '__main__':
    data = list(map(int, stdin.read().split()))
    n = data[0]
    a = data[1:n + 1]
    b = data[n + 1:]
    print(max_dot_product(a, b))
