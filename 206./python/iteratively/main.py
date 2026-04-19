from typing import Optional

# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if not head:
            return head

        current = head
        node = None

        while current:
            nextNode = current.next
            current.next = node
            node = current
            current = nextNode
        
        return node