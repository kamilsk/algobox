def gcd(a: int, b: int) -> int:
    if b == 0:
        return a
    while a % b != 0:
        a, b = b, (a % b)
    return b


def main():
    a, b = map(int, input().split())
    print(gcd(a, b))


if __name__ == "__main__":
    main()
