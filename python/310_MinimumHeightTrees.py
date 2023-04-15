class Solution(object):
    def findMinHeightTrees(self, n, edges):
        if n <= 2:
            return [i for i in range(n)]
        adj = [set() for _ in range(n)]
        for edge in edges:
            adj[edge[0]].add(edge[1])
            adj[edge[1]].add(edge[0])
        leaves, left = [], set(i for i in range(n))
        for i in range(n):
            if len(adj[i]) == 1:
                leaves.append(i)
                left.remove(i)
        while len(left) > 2:
            new_leaves = []
            for leaf in leaves:
                parent = adj[leaf].pop()
                adj[parent].remove(leaf)
                if len(adj[parent]) == 1:
                    new_leaves.append(parent)
                    left.remove(parent)
                elif len(adj[parent]) == 0:
                    return [parent]
            leaves = new_leaves
        return list(left)


if __name__ == '__main__':
    solution = Solution()
#    n = 4
#    edges = [[1, 0], [1, 2], [1, 3]]
#    n = 6
#    edges = [[3, 0], [3, 1], [3, 2], [3, 4], [5, 4]]
#    n = 3
#    edges = [[0,1],[0,2]]
    n = 6
    edges = [[0,1],[0,2],[0,3],[3,4],[4,5]]
    ans = solution.findMinHeightTrees(n, edges)
    print(ans)