# python3

from abc import ABC
from collections import namedtuple
from sys import stdin
from typing import AnyStr, IO, List
from unittest import TestCase

test = namedtuple('test', 'input expected')


class StackWithMax:
    def __init__(self):
        self.__stack = []
        self.__maxes = []

    def push(self, a: int):
        self.__maxes.append(max(a, self.__maxes[len(self.__maxes) - 1]) if self.__maxes else a)
        self.__stack.append(a)

    def pop(self) -> int:
        self.__maxes.pop()
        return self.__stack.pop()

    def max(self) -> int:
        return self.__maxes[len(self.__maxes) - 1]


class Processor:
    @staticmethod
    def process(src: IO) -> List[str]:
        buffer, stack = [], StackWithMax()

        for _ in range(int(src.readline())):
            query = src.readline().split()
            if query[0] == 'push':
                stack.push(int(query[1]))
            elif query[0] == 'pop':
                stack.pop()
            elif query[0] == 'max':
                buffer.append(stack.max())

        return buffer


class Fake(IO, ABC):
    def __init__(self, rows: List[str]):
        self.__i = -1
        self.__rows = [str(len(rows))] + rows

    def readline(self, limit: int = -1) -> AnyStr:
        self.__i += 1
        return self.__rows[self.__i]


class Test(TestCase):
    def test_stack_with_max(self):
        tests = [
            # samples
            test([
                'push 2',
                'push 1',
                'max',
                'pop',
                'max',
            ], [
                2,
                2,
            ]),
            test([
                'push 1',
                'push 2',
                'max',
                'pop',
                'max',
            ], [
                2,
                1,
            ]),
            test([
                'push 2',
                'push 3',
                'push 9',
                'push 7',
                'push 2',
                'max',
                'max',
                'max',
                'pop',
                'max',
            ], [
                9,
                9,
                9,
                9,
            ]),
            test([
                'push 1',
                'push 7',
                'pop',
            ], []),
            test([
                'push 7',
                'push 1',
                'push 7',
                'max',
                'pop',
                'max',
            ], [
                7,
                7,
            ]),
        ]
        for i, t in enumerate(tests):
            self.assertEqual(t.expected, Processor.process(Fake(t.input)), msg='at {} position'.format(i))


if __name__ == '__main__':
    for entry in Processor.process(stdin):
        print(entry)
