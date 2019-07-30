# python3

from collections import namedtuple
from unittest import TestCase

from typing import List, Optional

thread = namedtuple('thread', 'id start_at')
test = namedtuple('test', 'threads jobs expected')


class Scheduler(object):
    def __init__(self, threads: int):
        self._threads = [thread(i, 0) for i in range(threads)]

    def plan(self, jobs: List[int]) -> List[thread]:
        log = []
        for job in jobs:
            t = self.pop()
            log.append(thread(t.id, t.start_at))
            self.schedule(thread(t.id, t.start_at + job))
        return log

    def pop(self) -> Optional[thread]:
        last = len(self._threads) - 1
        if last < 0:
            return None
        if last == 0:
            return self._threads.pop()
        self._threads[0], self._threads[last] = self._threads[last], self._threads[0]
        nearest = self._threads.pop()
        self.sift_down(0)
        return nearest

    def schedule(self, t: thread):
        last = len(self._threads)
        self._threads.append(t)
        self.sift_up(last)

    def sift_down(self, i: int):
        min_idx = i

        left = 2 * i + 1
        if left < len(self._threads) and self._threads[left].start_at <= self._threads[min_idx].start_at:
            if self._threads[left].start_at < self._threads[min_idx].start_at:
                min_idx = left
            else:
                min_idx = left if self._threads[left].id < self._threads[min_idx].id else min_idx

        right = 2 * (i + 1)
        if right < len(self._threads) and self._threads[right].start_at <= self._threads[min_idx].start_at:
            if self._threads[right].start_at < self._threads[min_idx].start_at:
                min_idx = right
            else:
                min_idx = right if self._threads[right].id < self._threads[min_idx].id else min_idx

        if i != min_idx:
            self._threads[i], self._threads[min_idx] = self._threads[min_idx], self._threads[i]
            self.sift_down(min_idx)

    def sift_up(self, i: int):
        while i > 0 and self._threads[self.parent(i)].start_at >= self._threads[i].start_at:
            parent = self.parent(i)
            swap = False
            swap = swap or self._threads[parent].start_at > self._threads[i].start_at
            swap = swap or self._threads[parent].id > self._threads[i].id
            if not swap:
                break
            self._threads[parent], self._threads[i] = self._threads[i], self._threads[parent]
            i = parent

    @staticmethod
    def parent(i) -> int:
        return (i - 1) // 2


class Test(TestCase):
    def test_scheduler(self):
        tests = [
            # samples
            test(2, [1, 2, 3, 4, 5], [thread(0, 0), thread(1, 0), thread(0, 1), thread(1, 2), thread(0, 4)]),
            test(4, [1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1], [
                thread(0, 0),
                thread(1, 0),
                thread(2, 0),
                thread(3, 0),
                thread(0, 1),
                thread(1, 1),
                thread(2, 1),
                thread(3, 1),
                thread(0, 2),
                thread(1, 2),
                thread(2, 2),
                thread(3, 2),
                thread(0, 3),
                thread(1, 3),
                thread(2, 3),
                thread(3, 3),
                thread(0, 4),
                thread(1, 4),
                thread(2, 4),
                thread(3, 4),
            ]),
        ]
        for i, t in enumerate(tests):
            self.assertEqual(t.expected, Scheduler(t.threads).plan(t.jobs), msg='at {} position'.format(i))


if __name__ == '__main__':
    n, m = map(int, input().split())
    for plan in Scheduler(n).plan(list(map(int, input().split()))):
        print(plan.id, plan.start_at)
