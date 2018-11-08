from unittest import TestCase
from solution import gcd


class Test(TestCase):
    def test_gcd(self):
        self.assertEqual(gcd(18, 35), 1)
        self.assertEqual(gcd(14159572, 63967072), 4)
