# python3
def lcm_naive(a: int, b: int):
    for l in range(1, a * b + 1):
        if l % a == 0 and l % b == 0:
            return l

    return a * b


def gcd(a: int, b: int):
    if b == 0:
        return a

    if a < b:
        return gcd(a, b % a)
    return gcd(b, a % b)


def lcm(a: int, b: int):
    return (a * b) // gcd(a, b)


if __name__ == '__main__':
    n, m = map(int, input().split())
    print(lcm(n, m))
