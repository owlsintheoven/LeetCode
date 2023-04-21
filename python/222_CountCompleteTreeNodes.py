# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution(object):
    def countNodes(self, root):
        if root is None:
            return 0
        nodes = [root]
        res = 1
        while len(nodes) != 0:
            temp = []
            for node in nodes:
                if node.left is None:
                    return res
                elif node.right is None:
                    return res + 1
                else:
                    temp.extend([node.left, node.right])
                    res += 2
            nodes = temp
        return res


if __name__ == '__main__':
    solution = Solution()
    root = TreeNode(1, TreeNode(2, TreeNode(4), TreeNode(5)), TreeNode(3,  TreeNode(6)))
    ans = solution.countNodes(root)
    print(ans)

