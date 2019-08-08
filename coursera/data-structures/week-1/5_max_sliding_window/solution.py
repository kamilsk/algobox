# python3

from collections import namedtuple, deque
from typing import List
from unittest import TestCase

test = namedtuple('test', 'size data expected')


class SlidingWindow:
    @staticmethod
    def max(data: List[int], size: int) -> List[int]:
        result = []

        q = deque()
        for i in range(size):
            while q and data[i] >= data[q[-1]]:
                q.pop()
            q.append(i)

        for i in range(size, len(data)):
            result.append(data[q[0]])

            while q and q[0] <= i - size:
                q.popleft()

            while q and data[i] >= data[q[-1]]:
                q.pop()

            q.append(i)

        result.append(data[q[0]])
        return result


class Test(TestCase):
    def test_max_sliding_window(self):
        tests = [
            # samples
            test(4, [2, 7, 3, 1, 5, 2, 6, 2], [7, 7, 5, 6, 6]),
        ]
        for i, t in enumerate(tests):
            self.assertEqual(t.expected, SlidingWindow.max(t.data, t.size), msg='at {} position'.format(i))


if __name__ == '__main__':
    n = int(input())
    for max_value in SlidingWindow.max(data=[int(i) for i in input().split()], size=int(input())):
        print(max_value, end=' ')
