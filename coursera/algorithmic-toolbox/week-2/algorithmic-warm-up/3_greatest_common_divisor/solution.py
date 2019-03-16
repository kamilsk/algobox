# python3
def gcd_naive(x: int, y: int):
    current_gcd = 1

    for d in range(2, min(x, y) + 1):
        if x % d == 0 and y % d == 0:
            if d > current_gcd:
                current_gcd = d

    return current_gcd


def gcd(x: int, y: int):
    if y == 0:
        return x

    if x < y:
        return gcd(x, y % x)
    return gcd(y, x % y)


if __name__ == "__main__":
    a, b = map(int, input().split())
    print(gcd(a, b))
