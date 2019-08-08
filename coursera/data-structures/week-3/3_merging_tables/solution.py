# python3

from collections import namedtuple
from sys import stdin
from typing import List
from unittest import TestCase

merge = namedtuple('merge', 'dst src expected')
symlink = namedtuple('symlink', 'dst')
test = namedtuple('test', 'tables merges expected')


class Storage:
    def __init__(self, tables: List[int]):
        self._dsu = [symlink(0)]
        self._dsu.extend([i for i in tables])
        self._max = max(tables)

    def __len__(self) -> int:
        return self._max

    def find(self, table: int) -> int:
        update = []
        while isinstance(self._dsu[table], symlink):
            update.append(table)
            table = self._dsu[table].dst
        if len(update) > 1:
            for child in update[:len(update) - 1]:
                self._dsu[child] = symlink(table)
        return table

    def merge(self, dst: int, src: int) -> bool:
        real_dst, real_src = self.find(dst), self.find(src)

        if real_dst == real_src:
            return False

        self._dsu[real_dst] = self._dsu[real_dst] + self._dsu[real_src]
        self._dsu[real_src] = symlink(real_dst)
        self._max = max(self._max, self._dsu[real_dst])

        return True


class Test(TestCase):
    def test_merging_tables(self):
        tests = [
            # samples
            test([1, 1, 1, 1, 1], [
                merge(3, 5, True),
                merge(2, 4, True),
                merge(1, 4, True),
                merge(5, 4, True),
                merge(5, 3, False),
            ], [2, 2, 3, 5, 5]),
            test([10, 0, 5, 0, 3, 3], [
                merge(6, 6, False),
                merge(6, 5, True),
                merge(5, 4, True),
                merge(4, 3, True),
            ], [10, 10, 10, 11]),
        ]
        for i, t in enumerate(tests):
            s = Storage(t.tables)
            for j, mrg in enumerate(t.merges):
                self.assertEqual(mrg.expected, s.merge(mrg.dst, mrg.src), msg='merge at {}:{} position'.format(i, j))
                self.assertEqual(t.expected[j], len(s), msg='max size at {}:{} position'.format(i, j))


if __name__ == '__main__':
    n, m = map(int, stdin.readline().split())
    storage = Storage(list(map(int, stdin.readline().split())))
    for _ in range(m):
        destination, source = map(int, stdin.readline().split())
        storage.merge(destination, source)
        print(len(storage))
