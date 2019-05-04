# python3
def naive_lcm(a: int, b: int) -> int:
    for l in range(1, a * b + 1):
        if l % a == 0 and l % b == 0:
            return l

    return a * b


def fast_gcd(a: int, b: int) -> int:
    if b == 0:
        return a

    if a < b:
        return fast_gcd(a, b % a)
    return fast_gcd(b, a % b)


def fast_lcm(a: int, b: int) -> int:
    return (a * b) // fast_gcd(a, b)


if __name__ == '__main__':
    print(fast_lcm(*map(int, input().split())))
