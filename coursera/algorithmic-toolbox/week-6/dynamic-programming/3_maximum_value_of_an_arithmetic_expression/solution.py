# python3
from collections import namedtuple
from unittest import TestCase

test = namedtuple('test', 'expression expected')


def calculate(a: int, b: int, op: chr) -> int:
    if op == '+':
        return a + b
    elif op == '-':
        return a - b
    elif op == '*':
        return a * b
    else:
        assert False


def maximum_value(expression: str) -> int:
    n = 1 + len(expression) // 2
    mins, maxs = [], []
    for i in range(n):
        mins.append([0] * n)
        maxs.append([0] * n)
        mins[i][i] = maxs[i][i] = int(expression[2 * i])

    for s in range(1, n):
        for i in range(n - s):
            j = i + s
            cur_min, cur_max = 9 ** 14, -9 ** 14
            for k in range(i, j):
                a = calculate(maxs[i][k], maxs[k + 1][j], expression[2 * k + 1])
                b = calculate(maxs[i][k], mins[k + 1][j], expression[2 * k + 1])
                c = calculate(mins[i][k], maxs[k + 1][j], expression[2 * k + 1])
                d = calculate(mins[i][k], mins[k + 1][j], expression[2 * k + 1])
                cur_min = min(cur_min, a, b, c, d)
                cur_max = max(cur_max, a, b, c, d)
            mins[i][j], maxs[i][j] = cur_min, cur_max

    return maxs[0].pop()


class Test(TestCase):
    def test_maximum_value(self):
        tests = [
            # samples
            test('1+5', 6),
            test('5-8+7*4-8+9', 200),
        ]
        for i, t in enumerate(tests):
            self.assertEqual(t.expected, maximum_value(t.expression), msg='at {} position'.format(i))


if __name__ == '__main__':
    print(maximum_value(input()))
