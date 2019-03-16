# python3
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

    previous, current = 0, 1
    for i in range(x - 1):
        previous, current = current, previous + current
        pisano.append(current % y)
        for step in range(2, len(pisano) // 2):
            if pisano[:step] == pisano[step:2 * step]:
                return pisano[x % step]

    return current % y


if __name__ == '__main__':
    n, m = map(int, input().split())
    print(get_fibonacci_huge(n, m))
