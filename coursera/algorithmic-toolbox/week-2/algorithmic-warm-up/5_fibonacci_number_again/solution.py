# python3
from time import time


def timeit(method):
    def timed(*args, **kw):
        ts = time()
        result = method(*args, **kw)
        te = time()

        print('%2.2f ms' % ((te - ts) * 1000))
        return result

    return timed


def get_fibonacci_huge_naive(x: int, y: int):
    if x <= 1:
        return x

    previous, current = 0, 1
    for _ in range(x - 1):
        previous, current = current, previous + current

    return current % y


def get_fibonacci_huge(x: int, y: int):
    pisano = list()
    pisano.append(0)
    pisano.append(1)

    previous, current, last = 0, 1, 2
    for i in range(x - 1):
        previous, current = current, previous + current
        pisano.append(current % y)
        for step in range(last, len(pisano) // 2):
            if pisano[:step] == pisano[step:2 * step]:
                return pisano[x % step]
            last = step

    return current % y


if __name__ == '__main__':
    n, m = map(int, input().split())
    print(get_fibonacci_huge(n, m))
