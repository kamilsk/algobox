# python3
def gcd_naive(a: int, b: int):
    current_gcd = 1

    for d in range(2, min(a, b) + 1):
        if a % d == 0 and b % d == 0:
            if d > current_gcd:
                current_gcd = d

    return current_gcd


def gcd(a: int, b: int):
    if b == 0:
        return a

    if a < b:
        return gcd(a, b % a)
    return gcd(b, a % b)


if __name__ == '__main__':
    print(gcd(*map(int, input().split())))
