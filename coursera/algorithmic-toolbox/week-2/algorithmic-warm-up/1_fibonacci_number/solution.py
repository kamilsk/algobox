# python3
def calc_fib_naive(n: int) -> int:
    if n <= 1:
        return n

    return calc_fib_naive(n - 1) + calc_fib_naive(n - 2)


def calc_fib(n: int) -> int:
    memo = list()
    memo.append(0)
    memo.append(1)

    for i in range(2, n + 1):
        memo.append(memo[i - 1] + memo[i - 2])

    return memo[n]


if __name__ == '__main__':
    print(calc_fib(int(input())))
