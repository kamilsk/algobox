# python3
from collections import namedtuple
from unittest import TestCase

test = namedtuple('test', 'src target expected')


def edit_distance(src: str, target: str) -> int:
    m, n = len(target) + 1, len(src) + 1

    table = []
    for i in range(n):
        row = [0] * m
        row[0] = i
        table.append(row)
    for j in range(1, m):
        table[0][j] = j

    for j in range(1, m):
        for i in range(1, n):
            insertion = table[i][j - 1] + 1
            deletion = table[i - 1][j] + 1
            match = table[i - 1][j - 1]
            mismatch = match + 1

            table[i][j] = min(
                insertion,
                deletion,
                match if src[i - 1] == target[j - 1] else mismatch,
            )

    return table[n - 1][m - 1]


class Test(TestCase):
    def test_edit_distance(self):
        tests = [
            # samples
            test('ab', 'ab', 0),
            test('short', 'ports', 3),
            test('editing', 'distance', 5),
        ]
        for i, t in enumerate(tests):
            self.assertEqual(t.expected, edit_distance(t.src, t.target), msg='at {} position'.format(i))


if __name__ == '__main__':
    print(edit_distance(input(), input()))
