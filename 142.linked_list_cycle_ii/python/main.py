from typing import Optional

# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def detectCycle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        slow = fast = head
        found = False

        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next
            if slow == fast:
                found = True
                break

        if not found:
            return None
        
        slow = head

        while slow != fast:
            slow = slow.next
            fast = fast.next
        
        return slow


def create_linked_list_with_cycle(values: list[int], pos: int) -> Optional[ListNode]:
    """创建带环的链表，pos 是环起始位置的索引，-1 表示无环"""
    if not values:
        return None

    nodes = [ListNode(v) for v in values]
    for i in range(len(nodes) - 1):
        nodes[i].next = nodes[i + 1]

    if pos >= 0:
        nodes[-1].next = nodes[pos]

    return nodes[0]


if __name__ == "__main__":
    sol = Solution()

    # Test 1: [3,2,0,-4], pos=1 -> 环起点是 node(2)
    head1 = create_linked_list_with_cycle([3, 2, 0, -4], 1)
    result1 = sol.detectCycle(head1)
    print(f"Test 1: {result1.val if result1 else None}")  # Expected: 2

    # Test 2: [1,2], pos=0 -> 环起点是 node(1)
    head2 = create_linked_list_with_cycle([1, 2], 0)
    result2 = sol.detectCycle(head2)
    print(f"Test 2: {result2.val if result2 else None}")  # Expected: 1

    # Test 3: [1], pos=-1 -> 无环
    head3 = create_linked_list_with_cycle([1], -1)
    result3 = sol.detectCycle(head3)
    print(f"Test 3: {result3.val if result3 else None}")  # Expected: None