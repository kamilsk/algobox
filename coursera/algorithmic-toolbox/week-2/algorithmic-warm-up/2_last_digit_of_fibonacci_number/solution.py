# python3
def naive_fibonacci_last_digit(n: int) -> int:
    if n <= 1:
        return n

    previous, current = 0, 1
    for _ in range(n - 1):
        previous, current = current, previous + current

    return current % 10


def fast_fibonacci_last_digit(n: int) -> int:
    if n <= 1:
        return n

    previous, current = 0, 1
    for _ in range(n - 1):
        previous, current = current, previous + current % 10

    return current % 10


if __name__ == '__main__':
    print(fast_fibonacci_last_digit(int(input())))
