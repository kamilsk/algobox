# python3
from collections import namedtuple
from typing import List
from unittest import TestCase

Loot = namedtuple('Loot', 'value weight')


def get_optimal_value(cap: int, loots: List[Loot]) -> float:
    optimal = 0.0

    loots = sorted(loots, key=lambda x: x.value / x.weight, reverse=True)
    while cap > 0 and loots:
        best, loots = loots[0], loots[1:]
        score = best.value / best.weight
        w = min(best.weight, cap)
        optimal += score * w
        cap -= w

    return optimal


class Test(TestCase):
    def test_get_optimal_value(self):
        tests = [
            {'cap': 50, 'loots': [Loot(60, 20), Loot(100, 50), Loot(120, 30)], 'expected': 180.0000},
            {'cap': 10, 'loots': [Loot(500, 30)], 'expected': 166.6667},
        ]
        for test in tests:
            expected, obtained = test['expected'], get_optimal_value(test['cap'], test['loots'])
            self.assertEqual('{:.4f}'.format(expected), '{:.4f}'.format(obtained))


if __name__ == '__main__':
    n, capacity = list(map(int, input().split()))
    data = []
    for _ in range(n):
        value, weight = map(int, input().split())
        data.append(Loot(value, weight))
    opt_value = get_optimal_value(capacity, data)
    print('{:.4f}'.format(opt_value))
