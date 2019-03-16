# python3
def lcm_naive(x: int, y: int):
    for l in range(1, x * y + 1):
        if l % x == 0 and l % y == 0:
            return l

    return x * y


def gcd(x: int, y: int):
    if y == 0:
        return x

    if x < y:
        return gcd(x, y % x)
    return gcd(y, x % y)


def lcm(x: int, y: int):
    return (x * y) // gcd(x, y)


if __name__ == '__main__':
    a, b = map(int, input().split())
    print(lcm(a, b))
