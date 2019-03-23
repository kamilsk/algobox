# python3
from typing import List


def get_optimal_value(cap: int, loots: List[tuple]):
    optimal = 0.0

    loots = sorted(loots, key=lambda x: x[0] / x[1], reverse=True)
    while cap > 0 and len(loots) > 0:
        best, loots = loots[0], loots[1:]
        v, w = best
        score = v / w
        w = min(w, cap)
        optimal += score * w
        cap -= w

    return optimal


if __name__ == "__main__":
    n, capacity = list(map(int, input().split()))
    data = []
    for _ in range(n):
        value, weight = map(int, input().split())
        data.append((value, weight))
    opt_value = get_optimal_value(capacity, data)
    print("{:.4f}".format(opt_value))
