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


class TreeOrders:
    def __init__(self):
        self.n = self.key = self.left = self.right = None

    def read(self, src: IO):
        self.n = int(src.readline())
        self.key = [0 for _ in range(self.n)]
        self.left = [0 for _ in range(self.n)]
        self.right = [0 for _ in range(self.n)]
        for i in range(self.n):
            [self.key[i], self.left[i], self.right[i]] = map(int, src.readline().split())
        return self

    def walk(self) -> List[str]:
        return [
            ' '.join(str(x) for x in self.in_order()),
            ' '.join(str(x) for x in self.pre_order()),
            ' '.join(str(x) for x in self.post_order()),
        ]

    def in_order(self, node=0) -> List[int]:
        result = []

        if self.left[node] != -1:
            result += self.in_order(self.left[node])
        result.append(self.key[node])
        if self.right[node] != -1:
            result += self.in_order(self.right[node])

        return result

    def pre_order(self, node=0) -> List[int]:
        result = [self.key[node]]

        if self.left[node] != -1:
            result += self.pre_order(self.left[node])
        if self.right[node] != -1:
            result += self.pre_order(self.right[node])

        return result

    def post_order(self, node=0) -> List[int]:
        result = []

        if self.left[node] != -1:
            result += self.post_order(self.left[node])
        if self.right[node] != -1:
            result += self.post_order(self.right[node])
        result.append(self.key[node])

        return result


class Fake(IO, ABC):
    def __init__(self, rows: List[str]):
        self.__i = -1
        self.__rows = [str(len(rows))] + rows

    def readline(self, limit: int = -1) -> AnyStr:
        self.__i += 1
        return self.__rows[self.__i]


class Printer:
    def __init__(self, tree: TreeOrders):
        self.__tree = tree

    def walk(self):
        self.in_order()
        self.pre_order()
        self.post_order()

    def in_order(self, node=0):
        if self.__tree.left[node] != -1:
            self.in_order(self.__tree.left[node])

        print(self.__tree.key[node])

        if self.__tree.right[node] != -1:
            self.in_order(self.__tree.right[node])

    def pre_order(self, node=0):
        print(self.__tree.key[node])

        if self.__tree.left[node] != -1:
            self.pre_order(self.__tree.left[node])

        if self.__tree.right[node] != -1:
            self.pre_order(self.__tree.right[node])

    def post_order(self, node=0):
        if self.__tree.left[node] != -1:
            self.post_order(self.__tree.left[node])

        if self.__tree.right[node] != -1:
            self.post_order(self.__tree.right[node])

        print(self.__tree.key[node])


class Test(TestCase):
    def test_tree_orders(self):
        tests = [
            # samples
            test([
                '4 1 2',
                '2 3 4',
                '5 -1 -1',
                '1 -1 -1',
                '3 -1 -1',
            ], [
                '1 2 3 4 5',
                '4 2 1 3 5',
                '1 3 2 5 4',
            ]),
            test([
                '0 7 2',
                '10 -1 -1',
                '20 -1 6',
                '30 8 9',
                '40 3 -1',
                '50 -1 -1',
                '60 1 -1',
                '70 5 4',
                '80 -1 -1',
                '90 -1 -1',
            ], [
                '50 70 80 30 90 40 0 20 10 60',
                '0 70 50 40 30 80 90 20 60 10',
                '50 80 90 30 40 70 10 60 20 0',
            ]),
        ]
        for i, t in enumerate(tests):
            self.assertEqual(t.expected, TreeOrders().read(Fake(t.input)).walk(), msg='at {} position'.format(i))


def main():
    Printer(TreeOrders().read(stdin)).walk()
    # for way in TreeOrders().read(stdin).walk():
    #     print(way)


if __name__ == '__main__':
    Thread(target=main).start()
