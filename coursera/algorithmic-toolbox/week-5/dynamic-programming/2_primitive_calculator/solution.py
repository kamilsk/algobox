# python3
from collections import namedtuple
from sys import stdin
from typing import List
from unittest import TestCase

test = namedtuple('test', 'size expected')


def optimal_sequence(size: int) -> List[int]:
    optimal = []

    table = [0] * (size + 1)
    for i in range(1, size + 1):
        table[i] = table[i - 1] + 1
        if i % 2 == 0:
            table[i] = min(table[i], table[i // 2] + 1)
        if i % 3 == 0:
            table[i] = min(table[i], table[i // 3] + 1)

    while size >= 1:
        optimal.append(size)
        if size % 3 == 0 and table[size // 3] == table[size] - 1:
            size = size // 3
        elif size % 2 == 0 and table[size // 2] == table[size] - 1:
            size = size // 2
        else:
            size -= 1

    return list(reversed(optimal))


class Test(TestCase):
    def test_optimal_sequence(self):
        tests = [
            # samples
            test(1, [1]),
            test(5, [1, 2, 4, 5]),
            test(96234, [1, 3, 9, 10, 11, 22, 66, 198, 594, 1782, 5346, 16038, 16039, 32078, 96234]),
        ]
        for i, t in enumerate(tests):
            result = optimal_sequence(t.size)
            self.assertEqual(len(t.expected), len(result), msg='at {} position'.format(i))
            self.assertEqual(t.expected, result, msg='at {} position'.format(i))


if __name__ == '__main__':
    sequence = optimal_sequence(int(stdin.read()))
    print(len(sequence) - 1)
    for x in sequence:
        print(x, end=' ')
