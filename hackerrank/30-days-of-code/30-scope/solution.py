class Difference:
    maximumDifference = None

    def __init__(self, a: list):
        self.__elements = a

    def computeDifference(self):
        min_el, max_el = min(self.__elements), max(self.__elements)
        self.maximumDifference = max_el - min_el
