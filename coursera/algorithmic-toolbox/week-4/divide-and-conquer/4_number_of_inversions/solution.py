# python3
from collections import namedtuple
from sys import stdin
from typing import List
from unittest import TestCase

test = namedtuple('test', 'seq expected')


def number_of_inversions(seq: List[int], buf: List[int], left: int, right: int) -> int:
    count = 0
    if right - left <= 1:
        return count
    ave = (left + right) // 2
    count += number_of_inversions(seq, buf, left, ave)
    count += number_of_inversions(seq, buf, ave, right)

    i, j = left, ave
    while i < ave or j < right:
        if i == ave:
            buf[i + j - ave] = seq[j]
            j += 1
            continue
        if j == right:
            buf[i + j - ave] = seq[i]
            i += 1
            count += 1
            continue
        if seq[i] <= seq[j]:
            buf[i + j - ave] = seq[i]
            i += 1
        else:
            buf[i + j - ave] = seq[j]
            j += 1
            count += 1

    seq[left:right] = buf[left:right]
    return count


class Test(TestCase):
    def test_number_of_inversions(self):
        tests = [
            # samples
            test([2, 3, 9, 2, 9], 2),

            # acceptance
            test([9, 8, 7, 3, 2, 1], 15),

            # additional
            test([1, 2, 3, 4, 5], 0),
            test([5, 4, 3, 2, 1], 11),
        ]
        for i, t in enumerate(tests):
            self.assertEqual(t.expected, number_of_inversions(t.seq, [0] * len(t.seq), 0, len(t.seq)),
                             msg='at {} position'.format(i))


if __name__ == '__main__':
    n, *a = list(map(int, stdin.read().split()))
    print(number_of_inversions(a, n * [0], 0, len(a)))
