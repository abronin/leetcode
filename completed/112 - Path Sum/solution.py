class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
    
    def __string__(self):
        return "" + self.val + " - " + self.left + " - " + self.right

class Solution:
    def hasPathSum(self, root: TreeNode, sum: int) -> bool:
        return self.check(root, 0, sum)
    
    def check(self, node: TreeNode, sum: int, target: int) -> bool:
        if node is None: return False
        
        newsum = sum+node.val
        if self.isLeaf(node):
            return newsum == target
        
        return self.check(node.left, newsum, target) or self.check(node.right, newsum, target)

    def isLeaf(self, node: TreeNode) -> bool:
        return node.left is None and node.right is None

# left = TreeNode(2, None, None)
right = TreeNode(-3, None, None)
root = TreeNode(-2, None, right)

res = Solution().hasPathSum(root, -5)
print(res)