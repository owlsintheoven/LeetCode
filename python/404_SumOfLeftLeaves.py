# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution(object):
    def sumOfLeftLeaves(self, root):
        if root.left == None and root.right == None:
            return 0
        left_values = []
        self.search(root, left_values)
        return sum(left_values)

    def search(self, node, left_values):
        if node != None:
            self.search(node.right, left_values)
            if node.left != None and node.left.left == None and node.left.right == None:
                left_values.append(node.left.val)
                return
            self.search(node.left, left_values)


if __name__ == '__main__':
    solution = Solution()
    root = TreeNode(3, TreeNode(9), TreeNode(20, TreeNode(15), TreeNode(7)))
    ans = solution.sumOfLeftLeaves(root)
    print(ans)