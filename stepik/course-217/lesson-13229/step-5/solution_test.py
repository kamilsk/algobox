from unittest import TestCase
from solution import gcd


class Test(TestCase):
    def test_gcd(self):
        tests = [
            {'name': 'case#01', 'a': 18, 'b': 35, 'expected': 1},
            {'name': 'case#02', 'a': 14159572, 'b': 63967072, 'expected': 4},
        ]
        for test in tests:
            self.assertEqual(gcd(test['a'], test['b']), test['expected'])
