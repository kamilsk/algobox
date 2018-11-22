class Calculator:
    def __init__(self):
        pass

    def power(self, n: int, p: int) -> float:
        if n < 0 or p < 0:
            raise ValueError('n and p should be non-negative')
        return pow(n, p)
