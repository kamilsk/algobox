# python3

from collections import namedtuple
from unittest import TestCase

from random import randint
from typing import Any, Callable, List, Optional

contact = namedtuple('contact', 'phone name')
test = namedtuple('test', 'commands expected')


class Node:
    def __init__(self, value):
        self.next = self.prev = None
        self.value = value


class LinkedList:
    def __init__(self, compare: Callable[[Any, Any], bool]):
        self.head = self.tail = None
        self.compare = compare

    def add(self, value):
        if self.tail is None:
            self.head = self.tail = Node(value)
            return
        found = self.find(value)
        if found is None:
            n, t = Node(value), self.tail
            t.next, n.prev = n, t
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
            if self.compare(found.value, value):
                break
            found = found.next
        return found


class Query:
    def __init__(self, query: List[str]):
        self.type, self.phone = query[0], int(query[1])
        if self.type == 'add':
            self.name = query[2]


class PhoneBook:
    @staticmethod
    def __universal_family(cardinality: int) -> Callable[[int], int]:
        primary = 10 ** 7 + randint(10, 100)
        a, b = randint(1, primary - 1), randint(0, primary - 1)
        return lambda x: ((a * x + b) % primary) % cardinality

    def __init__(self):
        cardinality = 10 ** 7 // 10 ** 4
        self.__hash = self.__universal_family(cardinality)
        self.__storage: List[Optional[LinkedList]] = [None] * cardinality

    def process(self, queries: List[str]) -> List[str]:
        result = []

        for query in [Query(query.split()) for query in queries]:
            h = self.__hash(query.phone)
            if self.__storage[h] is None:
                self.__storage[h] = LinkedList(lambda x, y: x.phone == y.phone)
            ll = self.__storage[h]

            if query.type == 'add':
                ll.add(contact(query.phone, query.name))
                continue

            if query.type == 'del':
                ll.remove(contact(query.phone, None))
                continue

            response = 'not found'
            c = ll.find(contact(query.phone, None))
            if c is not None:
                response = c.value.name
            result.append(response)

        return result


class Test(TestCase):
    def test_phone_book(self):
        tests = [
            # samples
            test([
                'add 911 police',
                'add 76213 Mom',
                'add 17239 Bob',
                'find 76213',
                'find 910',
                'find 911',
                'del 910',
                'del 911',
                'find 911',
                'find 76213',
                'add 76213 daddy',
                'find 76213',
            ], [
                'Mom',
                'not found',
                'police',
                'not found',
                'Mom',
                'daddy',
            ]),
            test([
                'find 3839442',
                'add 123456 me',
                'add 0 granny',
                'find 0',
                'find 123456',
                'del 0',
                'del 0',
                'find 0',
            ], [
                'not found',
                'granny',
                'me',
                'not found',
            ]),
        ]
        for i, t in enumerate(tests):
            self.assertEqual(t.expected, PhoneBook().process(t.commands), msg='at {} position'.format(i))


if __name__ == '__main__':
    print('\n'.join(PhoneBook().process([input() for _ in range(int(input()))])))
