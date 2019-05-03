# python3
from collections import namedtuple
from sys import stdin
from typing import List
from unittest import TestCase

test = namedtuple('test', 'seq_a seq_b seq_c expected')


def lcs3(seq_a: List[int], seq_b: List[int], seq_c: List[int]) -> int:
    m, n, o = len(seq_a) + 1, len(seq_b) + 1, len(seq_c) + 1

    table = [[[0 for _ in range(o)] for _ in range(n)] for _ in range(m)]
    for i in range(1, m):
        for j in range(1, n):
            for k in range(1, o):
                if seq_a[i - 1] == seq_b[j - 1] == seq_c[k - 1]:
                    table[i][j][k] = table[i - 1][j - 1][k - 1] + 1
                else:
                    table[i][j][k] = max(table[i - 1][j][k], table[i][j - 1][k], table[i][j][k - 1])

    return table.pop().pop().pop()


class Test(TestCase):
    def test_lcs3(self):
        tests = [
            # samples
            test([1, 2, 3], [2, 1, 3], [1, 3, 5], 2),
            test([8, 3, 2, 1, 7], [8, 2, 1, 3, 8, 10, 7], [6, 8, 3, 1, 4, 7], 3),
        ]
        for i, t in enumerate(tests):
            self.assertEqual(t.expected, lcs3(t.seq_a, t.seq_b, t.seq_c), msg='at {} position'.format(i))


if __name__ == '__main__':
    data = list(map(int, stdin.read().split()))
    an = data[0]
    data = data[1:]
    a = data[:an]

    data = data[an:]
    bn = data[0]
    data = data[1:]
    b = data[:bn]

    data = data[bn:]
    cn = data[0]
    data = data[1:]
    c = data[:cn]

    print(lcs3(a, b, c))
