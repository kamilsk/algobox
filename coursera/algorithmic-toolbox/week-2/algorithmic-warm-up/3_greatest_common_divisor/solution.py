# python3
def naive_gcd(a: int, b: int) -> int:
    current_gcd = 1

    for d in range(2, min(a, b) + 1):
        if a % d == 0 and b % d == 0:
            if d > current_gcd:
                current_gcd = d

    return current_gcd


def fast_gcd(a: int, b: int) -> int:
    if b == 0:
        return a

    if a < b:
        return fast_gcd(a, b % a)
    return fast_gcd(b, a % b)


if __name__ == '__main__':
    print(fast_gcd(*map(int, input().split())))
