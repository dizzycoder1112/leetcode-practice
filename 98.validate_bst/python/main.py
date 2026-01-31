from typing import Optional

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
class Solution:
    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        return self.validate(root, float('-inf'), float('inf'))
        
    
    def validate(self, node: Optional[TreeNode], minimum: int, maximum: int) -> bool:
        if not node:
            return True
        
        if node.val<=minimum or node.val>=maximum:
            return False
        

        left = self.validate(node.left, minimum, node.val)
        right = self.validate(node.right, node.val, maximum)

        return left and right


def build_tree(values: list[int | None]) -> Optional[TreeNode]:
    """从 LeetCode 格式的 list 创建二叉树"""
    if not values or values[0] is None:
        return None

    root = TreeNode(values[0])
    queue = [root]
    i = 1

    while queue and i < len(values):
        node = queue.pop(0)

        if i < len(values) and values[i] is not None:
            node.left = TreeNode(values[i])
            queue.append(node.left)
        i += 1

        if i < len(values) and values[i] is not None:
            node.right = TreeNode(values[i])
            queue.append(node.right)
        i += 1

    return root


if __name__ == "__main__":
    sol = Solution()

    # Test 1: [2,1,3] -> Valid BST
    root1 = build_tree([2, 1, 3])
    print(f"Test 1: {sol.isValidBST(root1)}")  # Expected: True

    # Test 2: [5,1,4,null,null,3,6] -> Invalid BST (4 < 5 but on right)
    root2 = build_tree([5, 1, 4, None, None, 3, 6])
    print(f"Test 2: {sol.isValidBST(root2)}")  # Expected: False

    # Test 3: [5,4,6,null,null,3,7] -> Invalid BST (3 < 5 but in right subtree)
    root3 = build_tree([5, 4, 6, None, None, 3, 7])
    print(f"Test 3: {sol.isValidBST(root3)}")  # Expected: False

    # Test 4: Single node [1] -> Valid BST
    root4 = build_tree([1])
    print(f"Test 4: {sol.isValidBST(root4)}")  # Expected: True