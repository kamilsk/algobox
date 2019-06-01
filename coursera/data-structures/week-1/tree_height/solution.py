# python3
# from __future__ import annotations

from collections import namedtuple, deque
from typing import List
from unittest import TestCase

test = namedtuple('test', 'parents expected')


class Node(object):
    def __init__(self):
        self.children = []

    def add_child(self, node):
        self.children.append(node)


def compute_height(parents: List[int]) -> int:
    root, nodes = None, [Node() for _ in parents]

    for i in range(len(parents)):
        parent = parents[i]
        if parent == -1:
            root = nodes[i]
            continue
        nodes[parent].add_child(nodes[i])

    height = 0
    if root:
        q = deque()
        q.append(height + 1)
        q.append(root)
        while q:
            node = some = q.popleft()
            if type(some) is int:
                height = some
                node = q.popleft()
            if node.children:
                q.append(height + 1)
                for child in node.children:
                    q.append(child)

    return height


class Test(TestCase):
    def test_compute_height(self):
        tests = [
            # samples
            test([4, -1, 4, 1, 1], 3),
            test([-1, 0, 4, 0, 3], 4)
        ]
        for i, t in enumerate(tests):
            self.assertEqual(t.expected, compute_height(t.parents), msg='at {} position'.format(i))


if __name__ == '__main__':
    n = int(input())
    print(compute_height(list(map(int, input().split()))))
