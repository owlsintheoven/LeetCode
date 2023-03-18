# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution(object):
    def isCompleteTree(self, root):
        found_last_level = False
        nodes = [root]
        while len(nodes) != 0:
            next_nodes = []
            for i in range(len(nodes)):
                node = nodes[i]
                if node.left == None and node.right == None:
                    found_last_level = True
                    next_nodes.extend(nodes[i + 1:])
                    break
                elif node.left != None and node.right == None:
                    found_last_level = True
                    next_nodes.append((node.left))
                    next_nodes.extend(nodes[i + 1:])
                    break
                elif node.left == None and node.right != None:
                    return False
                else:
                    next_nodes.append(node.left)
                    next_nodes.append(node.right)
            nodes = next_nodes
            if found_last_level:
                break
        for node in nodes:
            if node.left != None or node.right != None:
                return False
        return True


if __name__ == '__main__':
#    node2_0 = TreeNode(4)
#    node2_1 = TreeNode(5)
#    node2_3 = TreeNode(6)
#    node1_0 = TreeNode(2, node2_0, node2_1)
#    node1_1 = TreeNode(3, None, node2_3)
#    root = TreeNode(1, node1_0, node1_1)
#    root = TreeNode(1, TreeNode(2), TreeNode(3))
#    root = TreeNode(1, TreeNode(2, TreeNode(3)))
    root = TreeNode(1, TreeNode(2, TreeNode(5)), TreeNode(3))
    solution = Solution()
    ans = solution.isCompleteTree(root)
    print(ans)

