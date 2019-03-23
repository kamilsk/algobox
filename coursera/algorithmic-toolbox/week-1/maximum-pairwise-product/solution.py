# python3
from typing import List


def max_pairwise_product(numbers: List[int]) -> int:
    local, idx1 = 0, -1
    for k, v in enumerate(numbers):
        if v >= local:
            local, idx1 = v, k
    local, idx2 = 0, -1
    for k, v in enumerate(numbers):
        if k != idx1 and v >= local:
            local, idx2 = v, k
    return numbers[idx1] * numbers[idx2]


if __name__ == '__main__':
    input_n = int(input())
    input_numbers = [int(x) for x in input().split()]
    print(max_pairwise_product(input_numbers))
