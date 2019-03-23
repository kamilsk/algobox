# python3
def sum_of_two_digits(first: int, second: int) -> int:
    return first + second


if __name__ == '__main__':
    a, b = map(int, input().split())
    print(sum_of_two_digits(a, b))
