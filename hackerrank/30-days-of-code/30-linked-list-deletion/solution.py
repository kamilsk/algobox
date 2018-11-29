class Node:
    def __init__(self, data):
        self.data = data
        self.next = None


class Solution:
    @staticmethod
    def insert(head, data):
        p = Node(data)
        if head is None:
            head = p
        elif head.next is None:
            head.next = p
        else:
            start = head
            while start.next is not None:
                start = start.next
            start.next = p
        return head

    @staticmethod
    def display(head):
        current = head
        while current:
            print(current.data, end=' ')
            current = current.next

    @staticmethod
    def removeDuplicates(head):
        if head is None:
            return head
        compare = head
        while compare is not None:
            prev, current = compare, compare.next
            while current is not None:
                if compare.data == current.data:
                    prev.next = current.next
                else:
                    prev = current
                current = current.next
            compare = compare.next
        return head
