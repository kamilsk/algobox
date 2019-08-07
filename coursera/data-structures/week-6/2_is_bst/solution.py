# python3

from abc import ABC
from collections import namedtuple
from sys import setrecursionlimit, stdin
from threading import stack_size, Thread
from typing import AnyStr, IO, List
from unittest import TestCase

setrecursionlimit(10 ** 6)
stack_size(2 ** 27)

test = namedtuple('test', 'input expected')


class TreeChecker:
    def __init__(self):
        self.n = 0
        self.key = self.left = self.right = None

    def read(self, src: IO):
        self.n = int(src.readline())
        self.key = [0 for _ in range(self.n)]
        self.left = [0 for _ in range(self.n)]
        self.right = [0 for _ in range(self.n)]
        for i in range(self.n):
            [self.key[i], self.left[i], self.right[i]] = map(int, src.readline().split())
        return self

    def check(self) -> bool:
        if self.n == 0:
            return True

        buf = [0 for _ in range(self.n)]
        self.in_order(buf)

        for i in range(len(buf) - 1):
            if buf[i] >= buf[i + 1]:
                return False
        return True

    def in_order(self, buf: List[int], node=0, position=0) -> int:
        if self.left[node] != -1:
            position = self.in_order(buf, self.left[node], position)

        buf[position] = self.key[node]
        position += 1

        if self.right[node] != -1:
            position = self.in_order(buf, self.right[node], position)

        return position


class Fake(IO, ABC):
    def __init__(self, rows: List[str]):
        self.__i = -1
        self.__rows = [str(len(rows))] + rows

    def readline(self, limit: int = -1) -> AnyStr:
        self.__i += 1
        return self.__rows[self.__i]


class Test(TestCase):
    def test_tree_checker(self):
        tests = [
            # samples
            test([
                '2 1 2',
                '1 -1 -1',
                '3 -1 -1',
            ], True),
            test([
                '1 1 2',
                '2 -1 -1',
                '3 -1 -1',
            ], False),
            test([], True),
            test([
                '1 -1 1',
                '2 -1 2',
                '3 -1 3',
                '4 -1 4',
                '5 -1 -1',
            ], True),
            test([
                '4 1 2',
                '2 3 4',
                '6 5 6',
                '1 -1 -1',
                '3 -1 -1',
                '5 -1 -1',
                '7 -1 -1',
            ], True),
            test([
                '4 1 -1',
                '2 2 3',
                '1 -1 -1',
                '5 -1 -1',
            ], False),
        ]
        for i, t in enumerate(tests):
            src = Fake(t.input)
            self.assertEqual(t.expected, TreeChecker().read(src).check(), msg='at {} position'.format(i))


def main():
    print('CORRECT') if TreeChecker().read(stdin).check() else print('INCORRECT')


if __name__ == '__main__':
    Thread(target=main).start()
