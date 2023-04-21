# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution(object):
    def widthOfBinaryTree(self, root):
        # the leftmost of the current level is the left child of the leftmost of the previous level
        # if the previous one has a left child
        widest = 1
        nodes = [(0, root)]
        while True:
            temp = []
            for i in range(len(nodes)):
                idx, node = nodes[i]
                if node.left is not None:
                    temp.append((2 * idx, node.left))
                if node.right is not None:
                    temp.append((2 * idx + 1, node.right))
            nodes = temp
            if len(nodes) != 0:
                width = nodes[-1][0] - nodes[0][0] + 1
                widest = width if width > widest else widest
            else:
                break
        return widest


#    def widthOfBinaryTree(self, root):
#
#            widest = 1
#        children = [root]
#        while True:
#            if children.count(None) == len(children):
#                break
#            width = len(children)
#            widest = width if width > widest else widest
#            for _ in range(width):
#                node = children.pop(0)
#                if node is None:
#                    children.extend([None, None])
#                else:
#                    children.extend([node.left, node.right])
#            while len(children) != 0 and children[-1] == None:
#                children.pop()
#            while len(children) != 0 and children[0] == None:
#                children.pop(0)
#        return widest


if __name__ == '__main__':
    solution = Solution()
    root = TreeNode(1, TreeNode(3, TreeNode(5), TreeNode(3)), TreeNode(2, None, TreeNode(9)))
    ans = solution.widthOfBinaryTree(root)
    print(ans)

