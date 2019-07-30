# python3

from collections import namedtuple
from unittest import TestCase

from typing import List, Optional

test = namedtuple('test', 'buckets commands expected')


class Node:
    def __init__(self, value):
        self.next = self.prev = None
        self.value = value


class LinkedList:
    def __init__(self):
        self.head = self.tail = None

    def add(self, value):
        if self.tail is None:
            self.head = self.tail = Node(value)
            return
        found = self.find(value)
        if found is None:
            n, h = Node(value), self.tail
            h.next, n.prev = n, h
            self.tail = n
            return
        found.value = value

    def remove(self, value):
        found = self.find(value)
        if found is None:
            return
        p, n = found.prev, found.next
        if p is None and n is None:
            self.head = self.tail = None
            return
        if p is None:
            self.head = n
            n.prev = None
            return
        if n is None:
            self.tail = p
            p.next = None
            return
        p.next, n.prev = n, p

    def find(self, value) -> Optional[Node]:
        if self.head is None:
            return None
        found = self.head
        while found is not None:
            if found.value == value:
                break
            found = found.next
        return found


class Query:
    def __init__(self, query: List[str]):
        self.type = query[0]
        if self.type == 'check':
            self.idx = int(query[1])
        else:
            self.str = query[1]


class QueryProcessor:
    _multiplier = 263
    _prime = 1000000007

    def __init__(self, buckets: int):
        self.__storage: List[Optional[LinkedList]] = [None] * buckets

    def __hash(self, s: str) -> int:
        bucket = 0
        for c in reversed(s):
            bucket = (bucket * self._multiplier + ord(c)) % self._prime
        return bucket % len(self.__storage)

    def process(self, queries: List[str]) -> List[str]:
        result = []

        for query in [Query(query.split()) for query in queries]:
            if query.type == 'check':
                ll = self.__storage[query.idx]
                dump = ''
                if ll is not None:
                    buf, ptr = [], ll.tail
                    while ptr is not None:
                        buf.append(ptr.value)
                        ptr = ptr.prev
                    dump = ' '.join(buf)
                result.append(dump)
                continue

            h = self.__hash(query.str)
            if self.__storage[h] is None:
                self.__storage[h] = LinkedList()
            ll = self.__storage[h]

            if query.type == 'add':
                ll.add(query.str)
                continue

            if query.type == 'del':
                ll.remove(query.str)
                continue

            result.append('yes' if ll.find(query.str) is not None else 'no')

        return result


class Test(TestCase):
    def test_query_processor(self):
        tests = [
            # samples
            test(5, [
                'add world',
                'add HellO',
                'check 4',
                'find World',
                'find world',
                'del world',
                'check 4',
                'del HellO',
                'add luck',
                'add GooD',
                'check 2',
                'del good',
            ], ['HellO world', 'no', 'yes', 'HellO', 'GooD luck']),
            test(4, [
                'add test',
                'add test',
                'find test',
                'del test',
                'find test',
                'find Test',
                'add Test',
                'find Test',
            ], ['yes', 'no', 'no', 'yes']),
            test(3, [
                'check 0',
                'find help',
                'add help',
                'add del',
                'add add',
                'find add',
                'find del',
                'del del',
                'find del',
                'check 0',
                'check 1',
                'check 2',
            ], ['', 'no', 'yes', 'yes', 'no', '', 'add help', '']),
        ]
        for i, t in enumerate(tests):
            self.assertEqual(t.expected, QueryProcessor(t.buckets).process(t.commands),
                             msg='at {} position'.format(i))


if __name__ == '__main__':
    print('\n'.join(QueryProcessor(int(input())).process([input() for _ in range(int(input()))])))
