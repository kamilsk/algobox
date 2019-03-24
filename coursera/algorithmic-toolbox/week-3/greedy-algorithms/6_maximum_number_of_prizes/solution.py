# python3
from sys import stdin
from typing import List
from unittest import TestCase


def optimal_summands(n: int) -> List[int]:
    optimal = []

    last_prize = 0
    while n:
        current_prize = min(n, last_prize + 1)
        if n - current_prize <= current_prize:
            current_prize = n
        n -= current_prize
        last_prize = current_prize
        optimal.append(last_prize)

    return optimal


class Test(TestCase):
    def test_optimal_summands(self):
        tests = [
            {'n': 6, 'expected': [1, 2, 3]},
            {'n': 8, 'expected': [1, 2, 5]},
            {'n': 2, 'expected': [2]},
            {'n': 1, 'expected': [1]},
            {'n': 3, 'expected': [1, 2]},
        ]
        for test in tests:
            expected, obtained = test['expected'], optimal_summands(test['n'])
            self.assertEqual(expected, obtained)


if __name__ == '__main__':
    summands = optimal_summands(int(stdin.read()))
    print(len(summands))
    for x in summands:
        print(x, end=' ')
