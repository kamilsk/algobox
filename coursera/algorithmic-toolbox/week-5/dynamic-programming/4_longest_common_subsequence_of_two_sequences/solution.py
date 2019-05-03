# python3
from collections import namedtuple
from sys import stdin
from typing import List
from unittest import TestCase

test = namedtuple('test', 'sequence_a sequence_b expected')


def lcs2(sequence_a: List[int], sequence_b: List[int]) -> int:
    m, n = len(sequence_a) + 1, len(sequence_b) + 1

    table = []
    for i in range(m):
        row = [0] * n
        row[0] = i
        table.append(row)
    for j in range(1, n):
        table[0][j] = j

    for i in range(1, m):
        for j in range(1, n):
            insertion = table[i][j - 1] + 1
            deletion = table[i - 1][j] + 1
            match = table[i - 1][j - 1]
            mismatch = match + 1

            table[i][j] = min(
                insertion,
                deletion,
                match if sequence_a[i - 1] == sequence_b[j - 1] else mismatch,
            )

    i, j = len(sequence_a), len(sequence_b)
    matches = 0
    while i > 0 and j > 0:
        if i > 0 and table[i][j] == table[i - 1][j] + 1:
            i -= 1
        elif j > 0 and table[i][j] == table[i][j - 1] + 1:
            j -= 1
        else:
            i, j = i - 1, j - 1
            matches += 1 if sequence_a[i] == sequence_b[j] else 0

    return matches


class Test(TestCase):
    def test_change(self):
        tests = [
            # samples
            test([2, 7, 5], [2, 5], 2),
            test([7], [1, 2, 3, 4], 0),
            test([2, 7, 8, 3], [5, 2, 8, 7], 2),
        ]
        for i, t in enumerate(tests):
            self.assertEqual(t.expected, lcs2(t.sequence_a, t.sequence_b), msg='at {} position'.format(i))


if __name__ == '__main__':
    data = list(map(int, stdin.read().split()))

    len_a = data[0]
    data = data[1:]
    a = data[:len_a]

    data = data[len_a:]
    len_b = data[0]
    data = data[1:]
    b = data[:len_b]

    print(lcs2(a, b))
