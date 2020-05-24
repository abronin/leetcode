from typing import List

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def removeLeafNodes(self, root: TreeNode, target: int) -> TreeNode:
        collected = self.collect(root, target, [])
        for node in reversed(collected):
            if not self.isLeaf(node) and node.left is not None and node.left.val == target and self.isLeaf(node.left):
                node.left = None
            if not self.isLeaf(node) and node.right is not None and node.right.val == target and self.isLeaf(node.right):
                node.right = None
        if root is not None and self.isLeaf(root) and root.val == target:
            return None
        return root
    
    def collect(self, root: TreeNode, target: int, collected: List[TreeNode]) -> List[TreeNode]:
        if root == None:
            return collected
        if (root.left is not None and root.left.val == target) or (root.right is not None and root.right.val == target):
            collected.append(root)
        self.collect(root.left, target, collected)
        self.collect(root.right, target, collected)
        return collected

    def isLeaf(self, node: TreeNode) -> bool:
        return node is None or (node.left is None and node.right is None)

def test_second():
    s = Solution()
    root = deserialize('[1,2,3]')
    target = 1
    expected = deserialize('[1,2,3]')
    given = s.removeLeafNodes(root, target)
    assert given == expected

def deserialize(string):
    if string == '{}':
        return None
    nodes = [None if val == 'null' else TreeNode(int(val))
             for val in string.strip('[]{}').split(',')]
    kids = nodes[::-1]
    root = kids.pop()
    for node in nodes:
        if node:
            if kids: node.left  = kids.pop()
            if kids: node.right = kids.pop()
    return root
