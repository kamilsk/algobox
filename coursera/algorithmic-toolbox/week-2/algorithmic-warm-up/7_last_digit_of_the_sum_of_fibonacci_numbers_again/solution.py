# python3
def naive_fibonacci_partial_sum(m: int, n: int) -> int:
    total = 0

    previous, current = 0, 1
    for i in range(n + 1):
        if i >= m:
            total += previous
        previous, current = current, previous + current

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


def fast_fibonacci_partial_sum(m: int, n: int) -> int:
    # sum(fm-1) = fm+1 - f1
    # sum(fn) = fn+2 - f1 (see the previous solution)
    # partial sum(fm...fn) = sum(fn) - sum(fm-1) = fn+2 - fm+1
    subset, full = fast_fibonacci_huge(m + 1, 10), fast_fibonacci_huge(n + 2, 10)
    return full - subset if full > subset else (10 - subset + full) % 10


if __name__ == '__main__':
    print(fast_fibonacci_partial_sum(*map(int, input().split())))
