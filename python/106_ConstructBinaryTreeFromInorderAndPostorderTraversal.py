# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution(object):
    def buildTree(self, inorder, postorder):
        if len(inorder) == 0:
            return None
        elif len(inorder) == 1:
            return TreeNode(inorder[0])
        # decide left group / right group
        node_val = postorder.pop()
        index_node_inorder = inorder.index(node_val)
        # no left tree
        if index_node_inorder == 0:
            return TreeNode(node_val, None, self.buildTree(inorder[1:], postorder))
        # no right tree
        elif index_node_inorder == len(inorder) - 1:
            return TreeNode(node_val, self.buildTree(inorder[: index_node_inorder], postorder))
        else:
            index_right_post = len(postorder) + 1
            for val in inorder[index_node_inorder + 1:]:
                index_right_post = postorder.index(val) if postorder.index(val) < index_right_post else index_right_post
                # l -> r -> val
            return TreeNode(node_val, self.buildTree(inorder[: index_node_inorder], postorder[: index_right_post]), self.buildTree(inorder[index_node_inorder + 1:], postorder[index_right_post:]))

    def inorder(self, node):
        if node != None:
            self.inorder(node.left)
            print(node.val)
            self.inorder(node.right)

    def postorder(self, node):
        if node != None:
            self.postorder(node.left)
            self.postorder(node.right)
            print(node.val)

    def preorder(self, node):
        if node != None:
            print(node.val)
            self.preorder(node.left)
            self.preorder(node.right)


if __name__ == '__main__':
    solution = Solution()
#    root = TreeNode(3, TreeNode(9), TreeNode(20, TreeNode(15), TreeNode(7)))
#    inorder = [9, 3, 15, 20, 7]
#    postorder = [9, 15, 7, 20, 3]
#    postorder = [9, 15, 3, 7, 20]
#    solution.postorder(root)
#    inorder = [3, 2, 1]
#    postorder = [3, 2, 1]
#    inorder = [1, 2, 3]
#    postorder = [3, 2, 1]
#    inorder = [1]
#    postorder = [1]
    inorder = [1, 2, 3, 4, 5]
    postorder = [1, 5, 4, 3, 2]
    ans = solution.buildTree(inorder, postorder)
    solution.preorder(ans)
#    root = TreeNode(1, None, TreeNode(2, None, TreeNode(3)))
#    solution.inorder(root)
#    print()
#    solution.postorder(root)
