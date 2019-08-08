# python3

from collections import namedtuple
from typing import Optional
from unittest import TestCase

bracket = namedtuple('bracket', 'char position')
test = namedtuple('test', 'text expected')


def are_matching(left: str, right: str) -> bool:
    return (left + right) in ['()', '[]', '{}']


def find_mismatch(text: str) -> Optional[int]:
    opening_brackets_stack = []

    for i, char in enumerate(text):
        if char in '([{':
            opening_brackets_stack.append(bracket(char, i + 1))
            continue

        if char in ')]}':
            if not opening_brackets_stack:
                return i + 1
            prev = opening_brackets_stack.pop()
            if not are_matching(prev.char, char):
                return i + 1

    return None if not opening_brackets_stack else opening_brackets_stack[0].position


class Test(TestCase):
    def test_brackets_in_code(self):
        tests = [
            # samples
            test('[]', None),
            test('{}[]', None),
            test('[()]', None),
            test('(())', None),
            test('{[]}()', None),
            test('{', 1),
            test('{[}', 3),
            test('foo(bar);', None),
            test('foo(bar[i);', 10),

            # acceptance
            test('}', 1),

            # additional
            test('foo() {}; bar({', 14),
        ]
        for i, t in enumerate(tests):
            self.assertEqual(t.expected, find_mismatch(t.text), msg='at {} position'.format(i))


if __name__ == '__main__':
    mismatch = find_mismatch(input())
    print('Success' if not mismatch else mismatch)
