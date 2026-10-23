# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def tree2str(self, root):
        # top-down DFS with global state
        self.result = []
        def dfs(node):
            if not node: 
                return
            self.result.append(str(node.val))
            if not node.left and not node.right: 
                return # no children 
            # alwyas process left (even if empty, when right exists)
            self.result.append("(")
            dfs(node.left)
            self.result.append(")")
            # only process right if it exists
            if node.right:
                self.result.append("(")
                dfs(node.right)
                self.result.append(")")
        dfs(root)
        return "".join(self.result)

