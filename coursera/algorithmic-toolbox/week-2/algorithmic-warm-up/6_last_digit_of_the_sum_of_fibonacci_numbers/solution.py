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


def fibonacci_sum_naive(n: int):
    if n <= 1:
        return n

    previous, current = 0, 1
    total = previous + current
    for _ in range(n - 1):
        previous, current = current, previous + current
        total += current

    return total % 10


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


def fibonacci_sum(n: int):
    # f0 = f2 - f1
    # f1 = f3 - f2
    # ...
    # fn = fn+2 - fn+1
    # sum(fn) = fn+2 - f1
    last = get_fibonacci_huge(n + 2, 10)
    return last - 1 if last > 0 else 9


if __name__ == '__main__':
    print(fibonacci_sum(int(input())))
