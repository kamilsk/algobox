# python3
def get_change(m: int):
    coins, pars = 0, [10, 5, 1]
    for par in pars:
        coins += m // par
        m %= par
    return coins


if __name__ == '__main__':
    print(get_change(int(input())))
