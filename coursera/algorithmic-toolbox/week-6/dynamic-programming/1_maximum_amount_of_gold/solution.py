# python3
from collections import namedtuple
from sys import stdin
from typing import List
from unittest import TestCase

test = namedtuple('test', 'capacity bars expected')


def optimal_weight(capacity: int, bars: List[int]) -> int:
    m, n = len(bars) + 1, capacity + 1

    table = []
    for i in range(m):
        row = [0] * n
        table.append(row)

    for i in range(1, m):
        wt = bars[i - 1]
        for capacity in range(1, n):
            table[i][capacity] = table[i - 1][capacity]
            if wt <= capacity:
                value = table[i - 1][capacity - wt] + wt
                if table[i][capacity] < value:
                    table[i][capacity] = value

    return table.pop().pop()


class Test(TestCase):
    def test_optimal_weight(self):
        tests = [
            # samples
            test(10, [1, 4, 8], 9),

            # additional
            test(1, [8, 2, 10], 0),
            test(11, [6, 11, 1, 3], 11),
            test(11, [6, 12, 1, 3], 10),
        ]
        for i, t in enumerate(tests):
            self.assertEqual(t.expected, optimal_weight(t.capacity, t.bars), msg='at {} position'.format(i))


if __name__ == '__main__':
    W, _, *w = list(map(int, stdin.read().split()))
    print(optimal_weight(W, w))
