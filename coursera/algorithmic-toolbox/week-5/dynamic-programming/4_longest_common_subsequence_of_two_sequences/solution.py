# python3
from collections import namedtuple
from sys import stdin
from typing import List
from unittest import TestCase

test = namedtuple('test', 'seq_a seq_b expected')


def lcs2(seq_a: List[int], seq_b: List[int]) -> int:
    m, n = len(seq_a) + 1, len(seq_b) + 1

    table = [[0] * n for _ in range(m)]
    for i in range(1, m):
        for j in range(1, n):
            if seq_a[i - 1] == seq_b[j - 1]:
                table[i][j] = table[i - 1][j - 1] + 1
            else:
                table[i][j] = max(table[i - 1][j], table[i][j - 1])

    return table.pop().pop()


class Test(TestCase):
    def test_lcs2(self):
        tests = [
            # samples
            test([2, 7, 5], [2, 5], 2),
            test([7], [1, 2, 3, 4], 0),
            test([2, 7, 8, 3], [5, 2, 8, 7], 2),

            # additional
            test([6, 3, 6, 3, 6, 5, 8], [10, 10, 5, 1, 4, 3], 1)
        ]
        for i, t in enumerate(tests):
            self.assertEqual(t.expected, lcs2(t.seq_a, t.seq_b), msg='at {} position'.format(i))


if __name__ == '__main__':
    data = list(map(int, stdin.read().split()))
    an = data[0]
    data = data[1:]
    a = data[:an]

    data = data[an:]
    bn = data[0]
    data = data[1:]
    b = data[:bn]

    print(lcs2(a, b))
