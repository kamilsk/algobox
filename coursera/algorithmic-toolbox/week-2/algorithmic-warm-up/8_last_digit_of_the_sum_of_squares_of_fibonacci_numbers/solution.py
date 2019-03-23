# python3
def fibonacci_sum_squares_naive(n: int) -> int:
    if n <= 1:
        return n

    previous, current = 0, 1
    total = previous + current
    for _ in range(n - 1):
        previous, current = current, previous + current
        total += current * current

    return total % 10


def get_fibonacci_huge(n: int, m: int) -> int:
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


def fibonacci_sum_squares(n: int) -> int:
    if n == 0:
        return 0

    # sum(fn^2) = fn*fn+1 (see the picture at the notes.pdf)
    return get_fibonacci_huge(n, 10) * get_fibonacci_huge(n + 1, 10) % 10


if __name__ == '__main__':
    print(fibonacci_sum_squares(int(input())))
