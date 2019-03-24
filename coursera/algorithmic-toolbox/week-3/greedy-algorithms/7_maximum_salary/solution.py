# python3
from sys import stdin
from typing import List
from unittest import TestCase


def largest_number(a: List[str]):
    res = ''

    m = 0
    while a:
        for i, v in enumerate(a):
            if greater_or_equal(v, a[m]):
                m = i
        res += a[m]
        a[0], a[m] = a[m], a[0]
        m, a = 0, a[1:]

    return res


def greater_or_equal(a: str, b: str) -> bool:
    la, lb = len(a), len(b)

    if la < lb and b.startswith(a):
        b = b[b.rindex(a) + la:]

    if lb < la and a.startswith(b):
        a = a[a.rindex(b) + lb:]

    return a >= b


class Test(TestCase):
    def test_largest_number(self):
        tests = [
            {'a': ['21', '2'], 'expected': '221'},
            {'a': ['9', '4', '6', '1', '9'], 'expected': '99641'},
            {'a': ['23', '39', '92'], 'expected': '923923'},
            # Failed case #10/11: Wrong answer
        ]
        for test in tests:
            expected, obtained = test['expected'], largest_number(test['a'])
            self.assertEqual(expected, obtained)

    def test_greater_or_equal(self):
        tests = [
            {'a': '21', 'b': '2', 'expected': False},
            {'a': '2', 'b': '21', 'expected': True},
            {'a': '21213', 'b': '21', 'expected': True},
            {'a': '21211', 'b': '21', 'expected': False},
        ]
        for test in tests:
            expected, obtained = test['expected'], greater_or_equal(test['a'], test['b'])
            self.assertEqual(expected, obtained)


if __name__ == '__main__':
    print(largest_number(stdin.read().split()[1:]))
