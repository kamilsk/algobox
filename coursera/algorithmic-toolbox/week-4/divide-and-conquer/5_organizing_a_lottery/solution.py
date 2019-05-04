# python3
from collections import namedtuple
from sys import stdin
from typing import List
from unittest import TestCase

test = namedtuple('test', 'starts ends points expected')


def fast_count_segments(starts: List[int], ends: List[int], points: List[int]) -> List[int]:
    cnt = {}
    line = [(x, 'left') for x in starts]
    line += [(x, 'point') for x in points]
    line += [(x, 'right') for x in ends]
    line.sort()

    segments = 0
    for pos, kind in line:
        if kind == 'left':
            segments += 1
        elif kind == 'right':
            segments -= 1
        else:
            cnt[pos] = segments

    return [cnt[x] for x in points]


def naive_count_segments(starts: List[int], ends: List[int], points: List[int]) -> List[int]:
    result = [0] * len(points)
    for i in range(len(points)):
        for j in range(len(starts)):
            if starts[j] <= points[i] <= ends[j]:
                result[i] += 1
    return result


class Test(TestCase):
    def test_count_segments(self):
        tests = [
            # samples
            test([0, 7], [5, 10], [1, 6, 11], [1, 0, 0]),
            test([-10], [10], [-100, 100, 0], [0, 0, 1]),
            test([0, -3, 7], [5, 2, 10], [1, 6], [2, 0]),
        ]
        for i, t in enumerate(tests):
            self.assertEqual(t.expected, fast_count_segments(t.starts, t.ends, t.points),
                             msg='at {} position'.format(i))


if __name__ == '__main__':
    data = list(map(int, stdin.read().split()))
    n = data[0]
    m = data[1]
    print(' '.join(map(str, fast_count_segments(data[2:2 * n + 2:2], data[3:2 * n + 2:2], data[2 * n + 2:]))))
