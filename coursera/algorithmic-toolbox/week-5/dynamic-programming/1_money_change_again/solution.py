# python3
from collections import namedtuple
from sys import stdin
from unittest import TestCase

test = namedtuple('test', 'money expected')
denominations = [1, 3, 4]


def change(money: int) -> int:
    table = [0] * (money + 1)

    for m in range(1, money + 1):
        table[m] = -1
        for coin in denominations:
            if m >= coin:
                coins = table[m - coin] + 1
                if table[m] == -1 or coins < table[m]:
                    table[m] = coins

    return table[money]


class Test(TestCase):
    def test_change(self):
        tests = [
            # samples
            test(2, 2),
            test(34, 9),
        ]
        for i, t in enumerate(tests):
            self.assertEqual(t.expected, change(t.money), msg='at {} position'.format(i))


if __name__ == '__main__':
    print(change(int(stdin.read())))
