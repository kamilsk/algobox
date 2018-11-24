class AdvancedArithmetic(object):
    def divisorSum(self, n: int) -> int:
        raise NotImplementedError


class Calculator(AdvancedArithmetic):
    def divisorSum(self, n: int) -> int:
        s = 1
        for i in range(2, n + 1):
            if n % i == 0:
                s += i
        return s
