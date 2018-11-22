class Node:
    def __init__(self, data: int):
        self.data = data
        self.next = None


class Solution:
    @staticmethod
    def display(head: Node):
        current = head
        while current:
            print(current.data, end=' ')
            current = current.next

    @staticmethod
    def insert(head: Node, data: int) -> Node:
        if not head:
            return Node(data)
        current = head
        while current.next:
            current = current.next
        current.next = Node(data)
        return head
