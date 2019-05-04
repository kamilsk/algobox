# python3
def naive_calc_fib(n: int) -> int:
    if n <= 1:
        return n

    return naive_calc_fib(n - 1) + naive_calc_fib(n - 2)


def fast_calc_fib(n: int) -> int:
    memo = list()
    memo.append(0)
    memo.append(1)

    for i in range(2, n + 1):
        memo.append(memo[i - 1] + memo[i - 2])

    return memo[n]


if __name__ == '__main__':
    print(fast_calc_fib(int(input())))
