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


def naive_fibonacci_sum(n: int) -> int:
    if n <= 1:
        return n

    previous, current = 0, 1
    total = previous + current
    for _ in range(n - 1):
        previous, current = current, previous + current
        total += current

    return total % 10


def fast_fibonacci_huge(n: int, m: int) -> int:
    pisano = list()
    pisano.append(0)
    pisano.append(1)

    previous, current, last = 0, 1, 2
    for i in range(n - 1):
        previous, current = current, previous + current
        pisano.append(current % m)
        for step in range(last, len(pisano) // 2):
            if pisano[:step] == pisano[step:2 * step]:
                return pisano[n % step]
            last = step + 1

    return current % m


def fast_fibonacci_sum(n: int) -> int:
    # f0 = f2 - f1
    # f1 = f3 - f2
    # ...
    # fn = fn+2 - fn+1
    # sum(fn) = fn+2 - f1
    last = fast_fibonacci_huge(n + 2, 10)
    return last - 1 if last > 0 else 9


if __name__ == '__main__':
    print(fast_fibonacci_sum(int(input())))
