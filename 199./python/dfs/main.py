# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def rightSideView(self, root: Optional[TreeNode]) -> List[int]:
        self.result = []
        self.dfs(root, 0)
        return self.result

    def dfs(self, node, depth: int):
        if not node:
            return
        
        if depth == len(self.result):
            self.result.append(node.val)
        
        self.dfs(node.right, depth+1)
        self.dfs(node.left, depth+1)
