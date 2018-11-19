class Difference:
    maximumDifference = None

    def __init__(self, a):
        self.__elements = a

    def computeDifference(self):
        min_el, max_el = min(self.__elements), max(self.__elements)
        self.maximumDifference = max_el - min_el


_ = input()
a = [int(e) for e in input().split(' ')]

d = Difference(a)
d.computeDifference()

print(d.maximumDifference)
